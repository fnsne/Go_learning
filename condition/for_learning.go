package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	i = 0
	for i < 10 {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
