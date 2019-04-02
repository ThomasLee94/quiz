package main

import(
	"flag"
	"os"
	"fmt"
)

func main() {
	openFile("problems.csv")
	r := csv.NewReader(file)
}

// OPENING ANY CSV FILE
func openFile(fileName string) {
	csvFile := flag.String("csv", fileName, "a csv file in the format of 'question, answer' ")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("failed to load csv file: %s\n", *csvFile))
	}

	_ = file
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}