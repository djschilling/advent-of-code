package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main1() {
	file, err := os.Open("input-1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	startSequenceLength := len("Game 1: ")

	maxColors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	validGames := 0
	gameIndex := 0

	for fileScanner.Scan() {
		gameIndex++
		//fmt.Print(gameIndex)
		line := fileScanner.Text()
		lineWithoutStartSequence := line[startSequenceLength:]
		colorCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, gameString := range strings.Split(lineWithoutStartSequence, ";") {
			regex := regexp.MustCompile(`(?P<count>\d+)\s(?P<color>red|green|blue)`)
			matches := regex.FindAllStringSubmatch(gameString, -1)
			for _, match := range matches {
				count, _ := strconv.Atoi(match[1])
				color := match[2]
				if colorCount[color] < count {
					colorCount[color] = count
					//fmt.Print(colorCount)
				}
			}
		}
		if colorCount["red"] > maxColors["red"] || colorCount["green"] > maxColors["green"] || colorCount["blue"] > maxColors["blue"] {
			//fmt.Println("Invalid")
		} else {
			validGames += gameIndex
		}
		//fmt.Println()
	}
	fmt.Println(validGames)

}

func main() {
	file, err := os.Open("input-1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	startSequenceLength := len("Game 1: ")

	gameIndex := 0
	sum := 0

	for fileScanner.Scan() {
		gameIndex++
		//fmt.Print(gameIndex)
		line := fileScanner.Text()
		lineWithoutStartSequence := line[startSequenceLength:]
		colorCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, gameString := range strings.Split(lineWithoutStartSequence, ";") {
			regex := regexp.MustCompile(`(?P<count>\d+)\s(?P<color>red|green|blue)`)
			matches := regex.FindAllStringSubmatch(gameString, -1)
			for _, match := range matches {
				count, _ := strconv.Atoi(match[1])
				color := match[2]
				if colorCount[color] < count {
					colorCount[color] = count
				}
			}
		}
		sum += (colorCount["red"] * colorCount["green"] * colorCount["blue"])
	}
	fmt.Println(sum)

}
