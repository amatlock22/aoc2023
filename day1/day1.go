package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		firstNum, lastNum := "", ""
		foundFirst := false
		line := scanner.Text()

		for _, char := range line {
			if isNum(char) {
				strNum := string(char)

				if foundFirst {
					lastNum = strNum
				} else {
					firstNum, lastNum = strNum, strNum
				}

				foundFirst = true
			}
		}

		totalNumLine, _ := strconv.Atoi(firstNum + lastNum)
		total += totalNumLine
	}

	return total
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	strNums := map[string][]string{
		"o": {"one"},
		"t": {"two", "three"},
		"f": {"four", "five"},
		"s": {"six", "seven"},
		"e": {"eight"},
		"n": {"nine"},
	}
	numsMap := map[string]string{
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

	for scanner.Scan() {
		firstNum, lastNum := "", ""
		foundFirst := false
		line := scanner.Text()

		for i, char := range line {
			str := string(char)

			// number character check
			if isNum(char) {
				if foundFirst {
					lastNum = str
				} else {
					firstNum, lastNum = str, str
					foundFirst = true
				}
			} else { // string character check
				if _, ok := strNums[str]; ok {
					// get possible string matches for starting letter
					possibleMatches := strNums[str]

					for _, pm := range possibleMatches {
						// out of index check
						if i+len(pm) > len(line) {
							continue
						}

						// possible match check
						if string(line[i:i+len(pm)]) == pm {
							if foundFirst {
								lastNum = numsMap[pm]
							} else {
								firstNum, lastNum = numsMap[pm], numsMap[pm]
								foundFirst = true
							}
						}
					}
				}
			}
		}

		totalNumLine, _ := strconv.Atoi(firstNum + lastNum)
		total += totalNumLine
	}

	return total
}

func isNum(char rune) bool {
	return char >= '0' && char <= '9'
}
