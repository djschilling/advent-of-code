package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("inputDay1-1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Print(line)
		digitsAsWordsMap := buildDigitsAsWordMapWithIndexes(line)
		digits := map[int]string{}
		for index, c := range line {
			if unicode.IsDigit(c) {
				digits[index] = string(c)
			}
		}
		var first string
		var last string
		firstIndex := 99999999
		lastIndex := -1
		for digit, indexes := range digitsAsWordsMap {
			for _, index := range indexes {
				if index < firstIndex {
					firstIndex = index
					first = digit
				}
				if index > lastIndex {
					lastIndex = index
					last = digit
				}
			}
		}
		for index, digit := range digits {
			if index < firstIndex {
				firstIndex = index
				first = digit
			}
			if index > lastIndex {
				lastIndex = index
				last = digit
			}
		}
		fmt.Print(" First and last:")
		fmt.Println(first, last)
		digit, _ := strconv.Atoi(first + last)
		sum += digit
	}

	fmt.Println(sum)
}

func getAllIndexesOfWordInString(word string, str string) []int {
	indexes := []int{}
	for i := 0; i < len(str); i++ {
		if len(str) >= i+len(word) && str[i:i+len(word)] == word {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func buildDigitsAsWordMapWithIndexes(line string) map[string][]int {
	digitsAsWordsMap := map[string][]int{}
	wordToDigits := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for word, digit := range wordToDigits {
		indexes := getAllIndexesOfWordInString(word, line)
		digitsAsWordsMap[digit] = indexes
	}
	return digitsAsWordsMap
}
