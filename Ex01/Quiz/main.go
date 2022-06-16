package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "../problems.csv", "csv file in a format of 'question,answer'")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	lines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := parseLines(lines)

	reader := bufio.NewReader(os.Stdin)
	c := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %v\n", i+1, p.q)
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		if text == p.a {
			c++
			fmt.Println("Correct!")
		} else {
			fmt.Println("WRONG!")
		}
	}

	fmt.Printf("You got %d correct out of %v", c, len(lines))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
