package calculations_test

import (
	"fmt"
	"math-skills/calculations"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	for i := 0; i <= 9; i++ {
		num := strconv.Itoa(i)
		sampleFile := "tests/testCases/case" + num + ".txt"
		initialText := ReadFile(sampleFile)

		targetFile := "tests/testX/case" + num + "X.txt"
		targetText := ReadFile(targetFile)

		average, median, variance, stdDev, err := calculations.Calculate(initialText)
		foundText := fmt.Sprintf("Average: %d\nMedian: %d\nVariance: %d\nStandard Deviation: %d", average, median, variance, stdDev)

		if err != nil {
			foundText = ""
		}
		modifiedFile := "tests/testOut/case" + num + "Out.txt"
		WriteFile(modifiedFile, foundText)

		result, textFound, textTarget := deepCompare(foundText, targetText)

		if !result {
			t.Errorf(`Content mismatch at test case#: %v`, num)
			for i := 0; i < len(textFound); i++ {
				t.Errorf(`
				_________________________________
				Text after modification
				%v
				----------------------
				Target text
				%v
				________________________________
				`, textFound[i], textTarget[i])
			}
		}
	}
}

func deepCompare(text1, text2 string) (bool, []string, []string) {
	if text1 == text2 {
		return true, nil, nil
	}
	lengthShort := 0
	lines1 := strings.Split(text1, "\n")
	lines2 := strings.Split(text2, "\n")
	if len(lines1) < len(lines2) {
		lengthShort = len(lines1)
	} else {
		lengthShort = len(lines2)
	}
	var textModified, textTarget []string
	for i := 0; i < lengthShort; i++ {
		if lines1[i] != lines2[i] {
			num := strconv.Itoa(i + 1)
			textModified = append(textModified, "line #"+num+": "+lines1[i])
			textTarget = append(textTarget, "line #"+num+": "+lines2[i])
		}
	}

	for i := lengthShort; i < len(lines1); i++ {
		num := strconv.Itoa(i + 1)
		textModified = append(textModified, "line #"+num+": "+lines1[i])
		textTarget = append(textTarget, "line #"+num+": "+"***<empty>***")
	}
	for i := lengthShort; i < len(lines2); i++ {
		num := strconv.Itoa(i + 1)
		textModified = append(textModified, "line #"+num+": "+"***<empty>***")
		textTarget = append(textTarget, "line #"+num+": "+lines2[i])
	}

	return false, textModified, textTarget
}

func ReadFile(sampleFile string) string {
	data, err := os.ReadFile(sampleFile)
	if err != nil {
		fmt.Printf("Error reading file %s", sampleFile)
		return ""
	}

	originalText := string(data)

	return originalText
}

func WriteFile(result, finalText string) {
	resultFile, err := os.Create(result)
	if err != nil {
		fmt.Println("Error in creating file: ", err)
	}

	// Write the text to result.txt
	openResultFile, _ := os.OpenFile(result, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	_, err3 := openResultFile.WriteString(finalText)
	if err3 != nil {
		fmt.Println("Error in writing file: ", err)
	}
	defer resultFile.Close()
}
