package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide a path to a CSV file.")
	}
	csvFilePath := args[1]
	file, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	uniqueValues := make(map[string]bool)
	for _, record := range records {
		for _, cell := range record {
			uniqueValues[cell] = true
		}
	}

	file, err = os.Create("csvuniq.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for value := range uniqueValues {
		err = writer.Write([]string{value})
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
