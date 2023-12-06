package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type parseFn func(string) int

func Parse(filename string, fn parseFn) string {
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
	return strconv.Itoa(sum)
}

func partA(line string) int {
	i := 0
	j := len(line) - 1
	first := -1
	last := -1
	for first == -1 || last == -1 {
		if first == -1 {
			if unicode.IsDigit(rune(line[i])) {
				first = int(line[i]) - '0'
				continue
			}
			i++
		}

		if last == -1 {
			if unicode.IsDigit(rune(line[j])) {
				last = int(line[j]) - '0'
				continue
			}
			j--
		}
	}
	return first*10 + last
}

func getCorpus() []string {
	return []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
}

/*
* returns 0 if not a valid number. Otherwise checks if the character at i is the start of a number string
* (or end if the reverse arg is given)
 */
func validNumber(input string, startIndex int, checkReverse bool) (value int, jump int) {
	c := rune(input[startIndex])
	// easy if it's just a number
	if unicode.IsDigit(c) {
		return int(c) - '0', 1
	}
	//"two1nine"

	corpus := getCorpus()
	for i, num := range corpus {
		if !checkReverse {
			// make sure we're not going out of bounds
			if startIndex+len(num) > len(input) {
				continue
			}
			if rune(num[0]) == c && input[startIndex:startIndex+len(num)] == num {
				return 1 + i, len(num)
			}
		} else {
			// make sure we're not going out of bounds
			if startIndex-len(num) < 0 {
				continue
			}
			// if last letter of our number string is equal to the given character
			if rune(num[len(num)-1]) == c && input[startIndex-len(num)+1:startIndex+1] == num {
				return 1 + i, len(num)
			}
		}
	}
	return 0, 1
}

func partB(line string) int {
	i := 0
	j := len(line) - 1
	first := -1
	last := -1
	for first == -1 || last == -1 {
		//"two1nine"
		if first == -1 {
			n, inc := validNumber(line, i, false)
			if n != 0 {
				fmt.Printf("n=%d\n", n)
				first = n
				continue
			}
			i += inc
		}

		if last == -1 {
			n, dec := validNumber(line, j, true)
			if n != 0 {
				fmt.Printf("n=%d\n", n)
				last = n
				continue
			}
			j -= dec
		}
	}
	fmt.Printf("%s => %d\n", line, first*10+last)
	return first*10 + last
}

func main() {
	fmt.Println("Day 1 Part A")
	fmt.Println(Parse("day-1/in.txt", partA))
	// fmt.Println(ValidNumber("1nine", 4, true))
	// fmt.Println(PartB("1nine"))
	fmt.Println("Day 1 Part B")
	fmt.Println(Parse("day-1/in.txt", partB))
}
