package main

import "fmt"

func Age(age int) bool {
	if age <= 18 {
		return false
	}
	return true
}

func main() {
	fmt.Printf("", Age(18))
}
