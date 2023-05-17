package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const separator = ","

func exitWithMessage(msg string) {
	fmt.Println(msg)
	os.Exit(-1)
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		exitWithMessage("no argument is provided, expecting markdown file")
	}
	readFile, err := os.Open(args[0])
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	csvHeader := fileScanner.Text()
	var sb strings.Builder
	columns := strings.Split(csvHeader, separator)

	writeCSVLine(columns, &sb)
	writeHeaderSeparator(len(columns), &sb)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		columns = strings.Split(line, separator)
		writeCSVLine(columns, &sb)
	}

	err = os.WriteFile("output.md", []byte(sb.String()), os.ModePerm)
	if err != nil {
		fmt.Println("error occurred while writing to output file ", err)
	}
}

func writeCSVLine(columns []string, sb *strings.Builder) {
	sb.WriteString("| ")
	for _, val := range columns {
		sb.WriteString(val + " |")
	}
	sb.WriteString("\n")
}

func writeHeaderSeparator(columnCount int, sb *strings.Builder) {
	sb.WriteString("| ")
	for i := 0; i < columnCount; i++ {
		sb.WriteString("---- |")
	}
	sb.WriteString("\n")
}
