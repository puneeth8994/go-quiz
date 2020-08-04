package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	//adds csv as a flag
	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv file.")
	}

	problems := parseLines(lines)

	correctAnswers := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.Q)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.A {
			correctAnswers++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(problems))
}

//Problem represents a struct of question and answer
type Problem struct {
	Q string
	A string
}

func parseLines(lines [][]string) []Problem {
	parsedLines := make([]Problem, len(lines))
	for i, line := range lines {
		parsedLines[i] = Problem{
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
