package examples

import "fmt"

// a pointer is a var that holds
// a location in memory where that value
// is stored

func Pointer() {
	var x int32 = 100
	var y bool = false
	// & lets us access that memory location
	pX := &x
	pY := &y
	fmt.Println("memory locations")
	fmt.Println(pX, pY)

}
