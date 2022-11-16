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
	fmt.Println(problems)

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
