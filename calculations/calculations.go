package calculations

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Calculate(originalText string) (average, median, variance, stdDev int, err error) {
	slice := []float64{}
	sliceStr := strings.Split(originalText, "\n")
	for _, word := range sliceStr {
		if word != "" {

			num, isSmall, err := StringToFloat(word)
			if !isSmall || err != nil {
				return 0, 0, 0, 0, err
			}
			slice = append(slice, num)
		}
	}

	averageFloat := math.Round(CalculateAverage(slice))
	average = int(averageFloat)
	medianFloat := math.Round(CalculateMedian(slice))
	median = int(medianFloat)
	varianceFloat := math.Round(CalculateVariance(slice))
	variance = int(varianceFloat)
	stdDevFloat := math.Round(CalculateStdDev(slice))
	stdDev = int(stdDevFloat)
	return average, median, variance, stdDev, nil
}

func StringToFloat(word string) (float64, bool, error) {
	isSmall := true
	word = strings.TrimSpace(word)
	word = strings.ReplaceAll(word, ",", ".")
	wholeNumber := word
	if strings.Contains(word, ".") {
		wholeNumber = word[:strings.IndexByte(word, '.')]
	}
	if len(wholeNumber) > 19 {
		isSmall = false
	}
	num, err := strconv.ParseFloat(word, 64)

	return num, isSmall, err
}

func CalculateAverage(slice []float64) float64 {
	var average float64
	for _, num := range slice {
		average += num
	}
	lengthOfSlice := float64(len(slice))
	average = average / lengthOfSlice
	return average
}

func CalculateMedian(slice []float64) float64 {
	var median float64
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	if len(slice)%2 != 0 {
		median = slice[len(slice)/2]
	} else {
		i := len(slice)/2 - 1
		median = (slice[i] + slice[i+1]) / 2
	}
	return median
}

func CalculateVariance(slice []float64) float64 {
	var variance float64
	average := CalculateAverage(slice)
	for _, num := range slice {
		variance += (num - average) * (num - average)
	}
	lengthOfSlice := float64(len(slice))
	variance = variance / (lengthOfSlice)
	return variance
}

func CalculateStdDev(slice []float64) float64 {
	var stdDev float64
	variance := CalculateVariance(slice)
	stdDev = math.Sqrt(variance)
	return stdDev
}
