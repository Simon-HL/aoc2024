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

	calculateSum(listA, listB)

}

func calculateSum(listA []int, listB []int) {
	var sum int
	for i := 0; i < len(listA); i++ {
		count := 0
		for j := 0; j < len(listB); j++ {
			if listA[i] == listB[j] {
				count++
			}
		}
		listA[i] = listA[i] * count

		sum += listA[i]
	}

	println(sum)
}

func getDiff(listA []int, listB []int) {
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
