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

	var a = [3]int{1, 2, 3}
	var b = [...]int{1, 2, 3}
	fmt.Println(a == b)

	var c []int
	// [:] lets you convert an array to a slice
	aSlice := a[:]
	// now we can append to the slice 'c'
	c = append(c, aSlice...)
}
