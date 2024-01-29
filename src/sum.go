// sum.go file
package main

import "C"
import "time"

//export sum
func sum(a C.int, b C.int) C.int {
	return a + b
}

//export sum_long_running
func sum_long_running(a C.int, b C.int) C.int {
	time.Sleep(5 * time.Second)
	return a + b
}

//export enforce_binding
func enforce_binding() {}

func main() {}
