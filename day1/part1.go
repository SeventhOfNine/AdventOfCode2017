package main

import (
	"AdventOfCode2017/utils"
	"fmt"
	"strconv"
)

func main() {
	line := utils.ReadFile("day1.txt")
	counter := 0

	for i := 0; i < len(line); i++ {
		a, _ := strconv.Atoi(string(line[i]))
		b, _ := strconv.Atoi(string(line[(i+1)%len(line)]))

		if a == b {
			counter += a
		}
	}

	fmt.Println("counter is: ", counter)
	main2(line)
}
