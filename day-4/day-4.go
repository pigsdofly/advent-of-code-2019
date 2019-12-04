package main

import (
	"fmt"
	"strconv"
)

func is_key(str_int string) bool {

	eq := make(map[int]int)
	for j := 0; j < len(str_int)-1; j++ {
		current, _ := strconv.Atoi(string(str_int[j]))
		next, _ := strconv.Atoi(string(str_int[j+1]))

		if current > next {
			return false
		}

		if eq[current] == 0 {
			eq[current] = 1
		}
		if next == current {
			eq[current]++
		}

	}

	for k1 := range eq {
		if eq[k1] == 2 {
			return true
		}
	}
	return false
}

func main() {
	low := 240920
	high := 789857
	total := 0
	for i := low; i <= high; i++ {
		str_int := fmt.Sprintf("%d", i)

		if is_key(str_int) {
			total++
		}
	}
	fmt.Println(total)
}
