package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main1() {
	file, err := os.Open("inputDay1-1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		digits := []string{}
		for _, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, string(c))
			}
		}
		first := digits[0]
		last := digits[len(digits)-1]
		digit, _ := strconv.Atoi(first + last)
		sum += digit
	}

	fmt.Println(sum)
}
