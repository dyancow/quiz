package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func error_handle(e error) {
	if e != nil {
		log.Fatal(e)
	}
	return
}

func main() {
	//read file setup
	f, err := os.Open("problems.csv")

	error_handle(err)

	//parsing file
	r := csv.NewReader(f)

	//get commandline input setup
	usr_reader := bufio.NewReader(os.Stdin)

	//setting up counter variable for number of total items and items gotten right
	total, num_correct := 0, 0

	//iterating through
	for {
		//read per line from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		error_handle(err)
		total += 1

		//ask question and give prompt line
		fmt.Printf("What is %s?\n", record[0])
		fmt.Printf("-->")
		answer_timer := time.NewTimer(3 * time.Second) //new timer

		//grab answer and parse
		answerCh := make(chan string)
		go func() {
			answer, _ := usr_reader.ReadString('\n')
			answer = strings.TrimSuffix(answer, "\n")
			answerCh <- answer
		}()

		select {
		case <-answer_timer.C:
			fmt.Printf("You ran out of time! You got %d answers out of %d.\n", num_correct, total)
			return
		case answer := <-answerCh:
			//keep record format as a string bc you can have words or integers.
			if answer == record[1] {
				num_correct += 1
				fmt.Printf("Good job! You got this one right.\n")
			} else {
				fmt.Printf("Incorrect. The right answer is %s.\n", record[1])
			}
		}

	}
	fmt.Printf("End of game: you got %d answers out of %d.\n", num_correct, total)
}
