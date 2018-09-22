/*This program reads in a wordlist and selects a random guess word*/
package main

import (
	"fmt" // For Printing
	"os" // For error handling
	"bufio"// To read the wordlist
	"time" // To use as a rand.Seed Arg
	"math/rand" // For picking a word
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var input string

	fmt.Print("Enter password: ")
	_, err :=fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error:",err.Error())
		os.Exit(0)
	}

	generated := random_gen()
	if !(input == generated) {
		fmt.Println("Wrong Password!")
		fmt.Println("Password:",generated)
		os.Exit(0)
	} else { fmt.Println("Access granted!") }

}

func random_gen() string {
	f, err := os.Open("namelist.txt")
	if err != nil {
		fmt.Println("Error:",err.Error())
		os.Exit(0)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var pass []string
	for scanner.Scan() {
		pass = append(pass, scanner.Text())
	}

	var p *string
	for {
		p = &pass[rand.Intn(len(pass))]
		if !(len(*p) > 4) {
			p = &pass[rand.Intn(len(pass))]
		} else { break }
	}
	return *p
}
