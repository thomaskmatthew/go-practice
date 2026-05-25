package examples

import "fmt"

func Slices() {
	// delcaring an arry
	var x [3]int
	fmt.Println(x)
	x[0] = 100
	x[1] = 100
	x[2] = 300
	fmt.Println(x)

	var y = [12]int{2, 3, 4, 5, 6, 7}
	fmt.Println(y)
}
