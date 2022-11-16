package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to quiz game..")
	var csvPath *string = flag.String("csv", "problems.csv", "This is the flag you use to read a csv file that contains the questions and answers in the form 5+5,10")
	flag.Parse()
	f, err := os.Open(*csvPath)
	if err != nil {
		handleError(fmt.Sprintf("%s of name %s", "Could not open the file", *csvPath), err)
	}
	problems, err := readFile(f)
	if err != nil {
		handleError("Failed to parse the file", errors.New("Improper formatting"))
	}
	correctAnswers := startQuiz(problems)
	fmt.Println("You have answered", correctAnswers, "correctly out of", len(problems))
}

type problem struct {
	question string
	answer   string
}

func handleError(message string, err error) {
	fmt.Printf("%s: %v", message, err)
	os.Exit(1)
}

func readFile(file *os.File) ([]problem, error) {
	problems := make([]problem, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		q := strings.Split(scanner.Text(), ",")[0]
		a := strings.TrimSpace(strings.Split(scanner.Text(), ",")[1])
		problems = append(problems, problem{question: q, answer: a})
	}
	return problems, nil
}

func startQuiz(problems []problem) int {
	var correctAnswers int
	for index, problem := range problems {
		fmt.Printf("Question #%d: %s = ?\n", index+1, problem.question)
		var answer string
		a, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		answer = strings.TrimSpace(string(a))
		if answer == problem.answer {
			correctAnswers++
		}
	}
	return correctAnswers
}
