package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	//adds csv as a flag
	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
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

	//timer sends message through a channel once the provided duration is completed.
	//This will help us keep a track of the quiz time.
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctAnswers := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.Q)

		answerCh := make(chan string)
		//Anonymous function running which scans the answer.
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.A {
				correctAnswers++
			}
		default:
		}
	}
	fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(problems))
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
