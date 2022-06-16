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

func main() {
	csvFilename := flag.String("csv", "../problems.csv", "csv file in a format of 'question,answer'")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	c := 0
	for i, l := range data {
		fmt.Printf("Problem #%d: %v\n", i+1, l[0])
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		if text == l[1] {
			c++
			fmt.Println("Correct!")
		} else {
			fmt.Println("WRONG!")
		}
	}

	fmt.Printf("You got %d correct out of %v", c, len(data))
}
