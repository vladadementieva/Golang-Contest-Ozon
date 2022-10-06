package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	returnV := "YES"
	var testCount int
	fmt.Fscan(in, &testCount)
	for i := 0; i < testCount; i++ {
		var goodCount int
		fmt.Fscan(in, &goodCount)
		// Create a dictionary of values for each element
		//mapping := make(map[int]mStruct)
		mapping := make([]mStruct, goodCount)
		doubles := []mStruct{}

		for j := 0; j < goodCount; j++ {
			var elem int
			fmt.Fscan(in, &elem)
			mapping[j] = mStruct{j + 1, elem, 0}
		}
		for i := 0; i < goodCount; i++ {
			for j := i + 1; j < goodCount; j++ {
				if mapping[i].val == mapping[j].val {
					doubles = append(doubles, mapping[i])
					doubles = append(doubles, mapping[j])
					i = i + 2
					j++
				}
			}
		}
		// Сортировка дублей
		// сохраняя оригинальный порядок или равные элементы.
		sort.SliceStable(doubles, func(i, j int) bool {
			return doubles[i].val < doubles[j].val
		})
		if len(doubles) == 0 {
			returnV = "YES"
		} else {
			flag := 0
			returnV = "YES"
			for i, _ := range doubles {
				if int(i+1) < len(doubles) && (flag == 0) {
					if (doubles[i].val == doubles[i+1].val) && ((doubles[i+1].key - doubles[i].key) > 1) {
						returnV = "NO"
						flag = 1
					}
				}
			}
		}
		fmt.Println(returnV)
	}
}
