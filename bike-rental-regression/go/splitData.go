package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strconv"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// Open the advertising dataset file.
	f, err := os.Open("day.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	df := dataframe.ReadCSV(f)

	// Calculate the number of elements in each set.
	trainCount := (4 * df.Nrow()) / 5
	testCount := df.Nrow() / 5

	// Create the subset indices.
	trainingIdx := make([]int, trainCount)
	testIdx := make([]int, testCount)

	// Enumerate the training indices.
	for i := 0; i < trainCount; i++ {
		trainingIdx[i] = i
	}
	// Enumerate the test indices.
	for i := 0; i < testCount; i++ {
		testIdx[i] = trainCount + i
	}

	// Create the subset dataframes.
	trainingDF := df.Subset(trainingIdx)
	testDF := df.Subset(testIdx)

	// Create a map that will be used in writing the data
	// to files.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Create the respective files.
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// Save the filtered dataset file.
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// Create a buffered writer.
		w := bufio.NewWriter(f)

		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Created training.csv with "+strconv.Itoa(trainCount)+" rows")
	fmt.Println("Created test.csv with "+strconv.Itoa(testCount)+" rows")
}