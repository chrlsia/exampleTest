package example

import "fmt"

func PrintMap() {
	m := map[int]bool{
		1: true,
		3: true,
		2: false,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
