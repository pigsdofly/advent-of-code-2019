package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func get_input() []string {
	filename := "input"
	dat, _ := ioutil.ReadFile(filename)

	s := string(dat)
	s_a := strings.Split(s, "\n")
	return s_a
}

func make_map(input []string) map[string]string {
	split_store := make(map[string]string)

	for i := 0; i < len(input); i++ {
		t := strings.Split(input[i], ")")
		split_store[t[1]] = t[0]
	}
	return split_store
}

func get_chain(k string, split_store map[string]string) []string {
	var chain []string
	for val, ok := split_store[k]; ok; val, ok = split_store[val] {
		chain = append(chain, val)
	}
	return chain
}

func find_distance(input []string) int {
	split_store := make_map(input)
	from, to := get_chain("YOU", split_store), get_chain("SAN", split_store)
	m := make(map[string]int)
	for i, k := range from {
		m[k] = i
	}
	for i, o := range to {
		if j, ok := m[o]; ok {
			return i + j
		}
	}
	return -1
}

func calculate_orbits(input []string) int {

	split_store := make_map(input)
	sum := 0
	for k := range split_store {
		sum += len(get_chain(k, split_store))
	}

	return sum
}

func main() {
	test_input := strings.Split("COM)B,B)C,C)D,D)E,E)F,B)G,G)H,D)I,E)J,J)K,K)L,K)YOU,I)SAN", ",")

	fmt.Println(find_distance(test_input))
	input := get_input()
	fmt.Println(find_distance(input))
	//fmt.Println(sum)
}
