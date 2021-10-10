package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	// Parse inputs
	var csvFilePath *string = flag.String("csv", "quiz.csv", "name of csv file to read")
	flag.Parse()

	records := readCsvFile(*csvFilePath)

	// Process records
	var correct, total int = 0, len(records)
	timer := time.NewTimer(10 * time.Second)

outerloop:
	for i := 0; i < len(records); i++ {
		record := records[i]
		fmt.Printf("Problem #%d: %s = ", i+1, record[0])

		answer := make(chan string)

		go func() {
			var y string
			fmt.Scanf("%s\n", &y)
			answer <- y
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTimeout!")
			break outerloop
		case x := <-answer:
			if record[1] == x {
				correct += 1
			}
		}
	}

	fmt.Printf("\nYou have answered %d/%d correctly\n", correct, total)
}
