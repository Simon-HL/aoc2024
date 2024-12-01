package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	listA, listB := parseInput()

	sort.Ints(listA)
	sort.Ints(listB)

	var totalDiff int

	for i := 0; i < len(listA); i++ {
		diff := listA[i] - listB[i]

		if diff < 0 {
			diff = -diff
		}

		totalDiff += diff
	}

	println(totalDiff)
}

func parseInput() ([]int, []int) {
	var listA []int
	var listB []int
	file, _ := os.Open("input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, err := r.ReadString('\n')
		if len(line) > 0 {
			values := strings.Fields(line)
			valueA, _ := strconv.ParseInt(values[0], 10, 0)
			valueB, _ := strconv.ParseInt(values[1], 10, 0)

			listA = append(listA, int(valueA))
			listB = append(listB, int(valueB))
		}
		if err != nil {
			break
		}
	}

	return listA, listB

}
