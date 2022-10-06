package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type mStruct struct {
	key  int
	val  int
	flag int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)
	for i := 0; i < testCount; i++ {
		var goodCount int
		fmt.Fscan(in, &goodCount)
		// Create a dictionary of values for each element
		//mapping := make(map[int]mStruct)
		mapping := make([]mStruct, goodCount)
		for j := 0; j < goodCount; j++ {
			var elem int
			fmt.Fscan(in, &elem)
			mapping[j] = mStruct{j + 1, elem, 0}

		}
		for i, _ := range mapping {
			if mapping[i].flag == 0 {
				mapping[i].flag = 1
				goal := mapping[i]
				_, index := find(mapping, goal.val)
				fmt.Println(goal.key, " ", index)
				mapping[index-1].flag = 1
			}

		}
		fmt.Println("  ")
	}
}

func find(numbers1 []mStruct, goal int) ([]mStruct, int) {
	n := []mStruct{}
	numbers := []mStruct{}
	c := 2
	a := 0
	for i, num := range numbers1 {
		if numbers1[i].flag == 0 {
			a = a + 1
			c = num.key
			numbers = append(numbers, numbers1[i])
		}
	}
	if a == 1 {
		return n, c

	} else {
		for i, num := range numbers {
			if int(i+1) < len(numbers) {
				if math.Abs(float64(numbers[i+1].val-goal)) < math.Abs(float64(num.val-goal)) {
					n = append(n, numbers[i+1])
					c = numbers[i+1].key
				} else {
					n = append(n, num)
					c = num.key
				}
			}
		}

		if len(n) > 1 {
			return find(n, goal)

		} else {
			return n, c

		}
	}
}
