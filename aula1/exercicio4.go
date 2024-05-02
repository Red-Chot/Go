package main

import (
	"errors"
	"fmt"
)

func divide(p, q int) (int, error) {
	if q == 0 {
		return 0, errors.New("division by 0 is not allowed")
	}
	return p / q, nil
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i := 0; i < len(a); i++ {
		v, err := divide(a[i], b[i])
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println(a[i], "/", b[i], "=", v)
	}
}
