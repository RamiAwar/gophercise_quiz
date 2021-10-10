package main

import (
	"errors"
	"io"
)

type CancelableReader struct {
	ch   chan bool
	data chan []byte
	err  error
	r    io.Reader
}

func (c *CancelableReader) begin() {
	buf := make([]byte, 1024)
	for {
		n, err := c.r.Read(buf)
		if n > 0 {
			tmp := make([]byte, n)
			copy(tmp, buf[:n])
			c.data <- tmp
		}
		if err != nil {
			c.err = err
			close(c.data)
			return
		}
	}
}

func (c *CancelableReader) Read(p []byte) (int, error) {
	select {
	case <-c.ch:
		return 0, errors.New("cancelled read")
	case d, ok := <-c.data:
		if !ok {
			return 0, c.err
		}
		copy(p, d)
		return len(d), nil
	}
}

func New(ch chan bool, r io.Reader) *CancelableReader {
	c := &CancelableReader{
		r:    r,
		ch:   ch,
		data: make(chan []byte),
	}
	go c.begin()
	return c
}
