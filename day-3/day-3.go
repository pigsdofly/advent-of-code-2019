package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x float64
	y float64
}

func get_input() []string {
	filename := "input"
	dat, _ := ioutil.ReadFile(filename)
	//	if err != nil {
	//		panic(err)
	//	}

	s := string(dat)
	s_a := strings.Split(s, "\n")
	return s_a
}

func get_distance(ins []string) map[string]int {
	var x int
	var y int
	x = 0
	y = 0
	length := 0
	coordinates := make(map[string]int)
	dx := map[string]int{
		"D": 0,
		"U": 0,
		"R": 1,
		"L": -1,
	}
	dy := map[string]int{
		"D": -1,
		"U": 1,
		"R": 0,
		"L": 0,
	}

	for i := 0; i < len(ins); i++ {
		command := string(ins[i][0])

		fmt.Println(y)
		num, _ := strconv.Atoi(string(ins[i][1:]))
		for j := 0; j < num; j++ {
			x += dx[command]
			y += dy[command]
			length += 1
			coord := fmt.Sprintf("%d,%d", x, y)
			//fmt.Println(coord)
			if coordinates[coord] == 0 {
				coordinates[coord] = length

			}
		}
	}
	return coordinates
}

func main() {
	instructions := get_input()
	coords_a := get_distance(strings.Split(instructions[0], ","))
	coords_b := get_distance(strings.Split(instructions[1], ","))
	matches := []int{}

	for k1 := range coords_a {

		if coords_b[k1] != 0 {
			//x, _ := strconv.Atoi(strings.Split(k1, ",")[0])
			//y, _ := strconv.Atoi(strings.Split(k1, ",")[1])
			dist := coords_b[k1] + coords_a[k1]
			matches = append(matches, dist)
		}
	}
	fmt.Println(len(matches))
	sort.Ints(matches)
	fmt.Println(matches[0], matches[len(matches)-1])
}
