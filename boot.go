package awesomeProject

import "fmt"

func Add(ints ...int) {
	var result int
	for _, v := range ints {
		result += v
	}
	fmt.Println(result)
}
