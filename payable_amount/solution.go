package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		dict := make(map[int]int)
		for j := 0; j < goodCount; j++ {
			var elem int
			fmt.Fscan(in, &elem)
			// Count uniq elements
			dict[elem] = dict[elem] + 1

		}
		var final_sum int = 0
		for i, num := range dict {
			// Count sum by 3=2
			final_sum = final_sum + i*(num-num/3)
		}
		fmt.Println(final_sum)
	}
}
