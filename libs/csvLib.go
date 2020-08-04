package libs

import (
	"encoding/csv"
	"fmt"
	"go-projects/quiz/models"
	"os"
)

//ReadCsvFile reads a csv file and returns a constructed slice of Problem struct
func ReadCsvFile(csvFileName *string) []models.Problem {

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv file.")
	}

	return parseLines(lines)
}

func parseLines(lines [][]string) []models.Problem {
	parsedLines := make([]models.Problem, len(lines))
	for i, line := range lines {
		parsedLines[i] = models.Problem{
			Q: line[0],
			A: line[1],
		}
	}
	return parsedLines
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
