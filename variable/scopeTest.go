package main

import "fmt"

func func1() {
	x := 1
	fmt.Println("in scopeTest.go")
	{
		fmt.Println(x)
	}
}
