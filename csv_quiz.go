package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type question struct {
	Question string
	Answer   int
}

func (q question) Ask(c chan int) {
	var r int
	fmt.Printf("%v? ", q.Question)
	_, err := fmt.Scan(&r)
	if err != nil {
		fmt.Println(err)
	}
	if q.Answer == r {
		c <- 1
	} else {
		c <- 0
	}
}

func newQuestion(r []string) question {
	ans, err := strconv.Atoi(r[1])
	if err != nil {
		fmt.Println(err)
	}
	return question{Question: r[0], Answer: ans}
}

func runGame(r *csv.Reader, d int) int {
	c := make(chan int)
	timeout := time.After(time.Duration(d) * time.Second)
	totalScore := 0

	for {
		rec, err := r.Read()
		if err == io.EOF { // no more questions
			return totalScore
		}
		q := newQuestion(rec)
		go q.Ask(c)

		select {
		case score := <-c:
			{
				totalScore += score
			}
		case <-timeout:
			{
				fmt.Println("Time out!")
				return totalScore
			}
		}
	}
}

func main() {
	f, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	duration := 10 // seconds
	score := runGame(r, duration)
	println("total score: ", score)
}
