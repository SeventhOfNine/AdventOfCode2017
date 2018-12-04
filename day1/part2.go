package main

import (
	"fmt"
	"strconv"
)

func main2(line []byte) {
	counter := 0

	for i := 0; i < len(line); i++ {
		a, _ := strconv.Atoi(string(line[i%len(line)]))
		b, _ := strconv.Atoi(string(line[(i+len(line)/2)%len(line)]))

		if a == b {
			counter += a
		}
	}

	fmt.Println("2nd counter is: ", counter)
}
