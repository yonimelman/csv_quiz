package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type question struct {
	Question string
	Answer   int
}

func (q question) Ask() int {
	var r int
	score := 0
	fmt.Println(q.Question, "?")
	_, err := fmt.Scan(&r)
	if err != nil {
		fmt.Println(err)
	}
	if q.Answer == r {
		score++
	}
	return score

}

func bulidQuestion(r []string) question {
	ans, err := strconv.Atoi(r[1])
	if err != nil {
		fmt.Println(err)
	}
	return question{Question: r[0], Answer: ans}
}

func main() {
	f, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Println("something went wrong")
	}
	r := csv.NewReader(f)
	score := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		q := bulidQuestion(record)
		score += q.Ask()
	}
	println("total score: ", score)
}
