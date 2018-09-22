/*This prog reads a file and counts the number of lines in it*/
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func countLines(filename *os.File) {
	filename.Seek(0,0)  //seek to the start of the file
	scanner := bufio.NewScanner(filename)
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}

	fmt.Println("lines count:",count)
}

func checkerr(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(0)
	}
}

func countComments(filename *os.File) {
	filename.Seek(0, 0)
	scanner := bufio.NewScanner(filename)
	scanner.Split(bufio.ScanLines)

	comments := 0
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "//") {
			comments++
		}
		if strings.Contains(scanner.Text(), "/*") {
			comments++
		} else {
			continue
		}
	}
	fmt.Println("Comments:", comments)
}

func countWords(filename *os.File) {
	filename.Seek(0, 0)
	scanner := bufio.NewScanner(filename)
	scanner.Split(bufio.ScanWords)

	words := 0
	for scanner.Scan() {
		words++
	}
	fmt.Println("Words:", words)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func countTab_Spaces(filename *os.File) {
	filename.Seek(0, 0)
	scanner := bufio.NewScanner(filename)
	scanner.Split(bufio.ScanBytes)

	tabs := 0
	spaces := 0
	for scanner.Scan() {
		if  (scanner.Bytes()[0] == ' ') {
			spaces++
		}
		if (scanner.Bytes()[0] == '\t') {
			tabs++
		}
	}
	fmt.Println("Tabs:",tabs)
	fmt.Println("Spaces:",spaces)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s 'filename.txt'\n",os.Args[0])
		os.Exit(0)
	}
	f, err := os.Open(os.Args[1])
	checkerr(err)
	defer f.Close()

	fmt.Println("Analysis of",os.Args[1])
	fmt.Println("--------------------------")

	countLines(f)
	countWords(f)
	countComments(f)
	countTab_Spaces(f)
}
