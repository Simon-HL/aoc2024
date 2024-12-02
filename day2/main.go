package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := parseInput()
	part1(reports)
}

func part1(reports [][]int) {
	var safeCount int
	for _, report := range reports {
		if isReportSafe(report) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isReportSafe(levels []int) bool {
	prevValue := 0
	isDecreasing := false
	isIncreasing := false
	isSafe := false

	for i, level := range levels {
		isSafe = true

		if i == 0 {
			prevValue = level
			continue
		}

		if prevValue == level {
			isSafe = false
			break
		}

		if prevValue < level {
			isIncreasing = true
		} else {
			isDecreasing = true
		}

		if !isWithinRange(prevValue, level) || (isDecreasing && isIncreasing) {
			isSafe = false
			break
		}

		prevValue = level
	}
	return isSafe
}

func isWithinRange(v1 int, v2 int) bool {
	diff := v1 - v2
	if diff < 0 {
		diff = -diff
	}

	return diff >= 1 && diff <= 3
}

func parseInput() [][]int {
	var reports [][]int
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		stringValues := strings.Fields(line)
		var levels []int

		for _, stringValue := range stringValues {
			value, err := strconv.Atoi(stringValue)
			if err != nil {
				fmt.Printf("Error parsing '%s' as int: %v\n", stringValue, err)
				break
			}
			levels = append(levels, value)
		}
		reports = append(reports, levels)
	}

	return reports
}
