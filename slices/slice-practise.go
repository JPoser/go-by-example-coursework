package main

import "fmt"

func main() {
	is := make([]int, 3)

	is[0] = 1
	is[1] = 2
	is[2] = 3

	fmt.Println(is)

	v := 0

	for _, val := range is {
		v += val
	}
	fmt.Println(v)

	strings := []string{"a", "b", "c"}
	fmt.Println(strings)

	for _, letter := range strings {
		fmt.Println(letter)
	}
}
