package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: \n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}

func mayPanic() {
	panic("a problem")
}
