// See https://courses.calhoun.io/lessons/les_goph_01
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
	// TODO Can I make this relative
	defaultFilePath := "/Users/kearls/sources/kevinearls/gophercises/quiz/problem.csv"
	filePath := flag.String("fileName", defaultFilePath, "Quiz file to use")
	flag.Parse()
	fmt.Printf("Using file %s\n", *filePath)

	records := getQuestions(*filePath)

	correctAnswers := 0
	questionsAsked := 0
	reader := bufio.NewReader(os.Stdin)
	for _, row := range records {
		fmt.Printf("%s: ", row[0])
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		response := string(text)
		if response == row[1] {
			fmt.Printf("Correct!!!!\n")
			correctAnswers++
		} else {
			fmt.Println("WRONG!!!\n")
			fmt.Printf("Expected [%s] but got [%s]\n", row[1], text)
		}
		questionsAsked++
	}
	fmt.Printf("Asked %d, Correct %d\n", questionsAsked, correctAnswers)
}

func getQuestions(filePath string) [][] string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()

	fmt.Printf("Opened %s\n", f.Name())

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}

	return records
}


