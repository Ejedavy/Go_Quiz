package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to quiz game..")
	var csvPath *string = flag.String("csv", "problems.csv", "This is the flag you use to read a csv file that contains the questions and answers in the form 5+5,10")
	var quizDuration *int = flag.Int("duration", 5, "This is the duration of the quiz and by default it is 5 seconds")
	flag.Parse()
	f, err := os.Open(*csvPath)
	if err != nil {
		handleError(fmt.Sprintf("%s of name %s", "Could not open the file", *csvPath), err)
	}
	problems, err := readFile(f)
	if err != nil {
		handleError("Failed to parse the file", errors.New("Improper formatting"))
	}
	correctAnswers, timeUp := startQuiz(problems, time.Duration(*quizDuration))
	if timeUp {
		fmt.Println("TIME UP")
	}
	fmt.Println("You have answered", correctAnswers, "correctly out of", len(problems))
	os.Exit(0)
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

func startQuiz(problems []problem, duration time.Duration) (int, bool) {
	timer := time.NewTimer(duration * time.Second)

	var correctAnswers int
	for index, problem := range problems {
		fmt.Printf("Question #%d: %s = ?\n", index+1, problem.question)
		var answerChannel = make(chan string)
		go func() {
			var answer string
			a, _, _ := bufio.NewReader(os.Stdin).ReadLine()
			answer = strings.TrimSpace(string(a))
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			return correctAnswers, true
		case answer := <-answerChannel:
			if answer == problem.answer {
				correctAnswers++
			}
			continue
		}
	}
	return correctAnswers, false
}
