package main

import (
	"fmt"
	"log"
	"math-skills/calculations"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide one txt file.")
		return
	}

	if !strings.HasSuffix(args[0], ".txt") {
		fmt.Println("Please provide txt files.")
		return
	}
	data := ReadFile(args[0])

	checkSpaces := strings.ReplaceAll(data, "\n", "")
	checkSpaces = strings.ReplaceAll(checkSpaces, "\t", "")
	checkSpaces = strings.ReplaceAll(checkSpaces, " ", "")

	if data == "" || checkSpaces == "" {
		fmt.Println("The data.txt is empty.")
		return
	}

	average, median, variance, stdDev, err := calculations.Calculate(data)
	if err != nil {
		fmt.Println("The data in data.txt is not a statistical population.")
		return
	}
	fmt.Printf(`Average: %d
Median: %d
Variance: %d
Standard Deviation: %d
`, average, median, variance, stdDev)
}

func ReadFile(sampleFile string) string {
	data, err := os.ReadFile(sampleFile)
	if err != nil {
		log.Print(err)
	}
	dataStr := string(data)
	return dataStr
}
