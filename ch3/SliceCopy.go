package main

import "fmt";


func main (){
	a := []int{30, 20, 50, 10, 40}
	b := make([]int, len(a))
	for i := range a {
		b[i] = a[i]
	}
    fmt.Println(b)
}