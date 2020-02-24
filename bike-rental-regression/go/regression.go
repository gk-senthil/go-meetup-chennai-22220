package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {

	// Open the training dataset file.
	f, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 16
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(trainingData[0])
	// let's create
	// the struct needed to train a model using github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("Number of bike rentals per day")
	r.SetVar(0, "Temperature")
	r.SetVar(1,"Season")
	//r.SetVar(2,"Month")
	r.SetVar(2,"Weather")

	// Loop of records in the CSV, adding the training data to the regression value.
	for i, record := range trainingData {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the "Rental count" measure, or "y".
		yVal, err := strconv.ParseFloat(record[15], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the attributes of significance for predicting daily rentals.
		temp, err := strconv.ParseFloat(record[9], 64)
		if err != nil {
			log.Fatal(err)
		}

		season, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		/* month, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatal(err)
		} */

		weather, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Add these points to the regression value.
		r.Train(regression.DataPoint(yVal, []float64{temp,season,weather}))
	}

	// Train/fit the regression model.
	r.Run()

	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)

	// Open the test dataset file.
	f, err = os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a CSV reader reading from the opened file.
	reader = csv.NewReader(f)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 16
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the test data predicting y and evaluating the prediction
	// with the mean absolute error.
	var mAE float64
	for i, record := range testData {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the observed diabetes progression measure, or "y".
		yObserved, err := strconv.ParseFloat(record[15], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the attributes of significance for predicting daily rentals.
		temp, err := strconv.ParseFloat(record[9], 64)
		if err != nil {
			log.Fatal(err)
		}

		season, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		/* month, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatal(err)
		} */

		weather, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model.
		yPredicted, err := r.Predict([]float64{temp,season,weather})

		// Add the to the mean absolute error.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// Output the MAE to standard out.
	fmt.Printf("MAE = %0.2f\n\n", mAE)
}