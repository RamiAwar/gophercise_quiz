package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	for i := 0; i < len(records); i++ {
		record := records[i]
		question := record[0]

		// Display question for X seconds
		fmt.Print(question + " = ")

		var user_answer string
		fmt.Scanln(&user_answer)

		answer, err := strconv.ParseInt(record[1], 10, 0)
		if err != nil {
			log.Fatal("Unable to parse record as integer: ", record[1])
		}

		user_answer_int, err := strconv.ParseInt(user_answer, 10, 0)
		if err != nil {
			fmt.Println("Please input a valid number.")
			i--
		}
		if answer == user_answer_int {
			correct += 1
		}
	}

	var result string = strconv.Itoa(correct) + "/" + strconv.Itoa(total)
	fmt.Println("You have answered " + result + " correctly")
}
