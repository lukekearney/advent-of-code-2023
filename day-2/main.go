package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type parseFn func(line string) int

type game struct {
	red   int
	green int
	blue  int
}

func Parse(filename string, fn parseFn) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += fn(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func partA(line string) int {
	allowedBalls := game{
		red:   12,
		green: 13,
		blue:  14,
	}
	// extract the id
	parts := strings.Split(line, ": ")
	id, err := strconv.Atoi(parts[0][5:])
	if err != nil {
		log.Fatal(err)
		return 0
	}

	// now handle each game
	for _, round := range strings.Split(parts[1], "; ") {
		ballParts := strings.Split(round, ", ")
		for _, color := range ballParts {
			numBalls := strings.Split(color, " ")
			if numBalls[1] == "red" {
				red, err := strconv.Atoi(numBalls[0])
				if err != nil {
					log.Fatal(err)
					return 0
				}
				if red > allowedBalls.red {
					return 0
				}
			} else if numBalls[1] == "green" {
				green, err := strconv.Atoi(numBalls[0])
				if err != nil {
					log.Fatal(err)
					return 0
				}
				if green > allowedBalls.green {
					return 0
				}
			} else {
				// must be blue
				blue, err := strconv.Atoi(numBalls[0])
				if err != nil {
					log.Fatal(err)
					return 0
				}
				if blue > allowedBalls.blue {
					return 0
				}
			}
		}
	}
	return id
}

func partB(line string) int {
	minimumRequired := game{
		red:   0,
		green: 0,
		blue:  0,
	}
	// extract the id
	game := line[strings.Index(line, ": ")+2:]
	// now handle each game
	for _, round := range strings.Split(game, "; ") {
		ballParts := strings.Split(round, ", ")
		for _, color := range ballParts {
			numBalls := strings.Split(color, " ")
			if numBalls[1] == "red" {
				red, err := strconv.Atoi(numBalls[0])
				if err != nil {
					log.Fatal(err)
					return 0
				}
				if red > minimumRequired.red {
					minimumRequired.red = red
				}
			} else if numBalls[1] == "green" {
				green, err := strconv.Atoi(numBalls[0])
				if err != nil {
					log.Fatal(err)
					return 0
				}
				if green > minimumRequired.green {
					minimumRequired.green = green
				}
			} else {
				// must be blue
				blue, err := strconv.Atoi(numBalls[0])
				if err != nil {
					log.Fatal(err)
					return 0
				}
				if blue > minimumRequired.blue {
					minimumRequired.blue = blue
				}
			}
		}
	}
	return minimumRequired.red * minimumRequired.blue * minimumRequired.green
}

func main() {
	fmt.Println("Day 2 Part A")
	fmt.Println(Parse("day-2/in.txt", partA))
	fmt.Println("Day 2 Part B")
	fmt.Println(Parse("day-2/in.txt", partB))
}
