package main

import "fmt"

func main() {
	intArr := []int32{1, 2, 3, 4, 0, 0, 0, 0}
	var intArr2 []int32 = []int32{5, 6, 7, 8, 9}
	fmt.Println(intArr2)
	fmt.Println(intArr)
	// intArr = append(intArr,6)
	fmt.Println(intArr)
	fmt.Printf("Array length: %d\n with capacity: %d", len(intArr), cap(intArr))
	newArr := append(intArr, intArr2...)
	fmt.Println(newArr)

	var SizedArr []int32 = make([]int32, 100, 100000)

	fmt.Printf("Array length: %d\n with capacity: %d", len(SizedArr), cap(SizedArr))
}
