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
	p := "../problems.csv"
	h := flag.Bool("h", false, "help text")
	flag.Parse()

	if *h {
		fmt.Print("-csv string\n\ta csv file in the format of 'question,answer' (default \"problems.csv\")) " +
			"\n-limit int\n\tthe time limit for the quiz in seconds (default 30)")
	}

	f, err := os.Open(p)
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
