// Stack
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func Eval(expr string) int {
	var ops []string
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				return
			}
		
		    ops = ops[:len(ops)-1]
		    if op == "(" {
		    	return
		    }
		    b, a := pop(), pop()
		    switch op {
			    case "+":
			        nums = append(nums,a+b)
			    case "-":
			        nums = append(nums,a-b)
			    case "*":
			        nums = append(nums,a*b)
			    case "/":
			        nums = append(nums,a/b)
		    }
	    }
	}
	for _, token := range strings.Split(expr, " "){
		switch token {
			case "(":
			    ops = append(ops, token)
		        case "+", "-":	
			    reduce("+-&/")
				ops = append(ops, token)
                        case "*", "/":	
			    reduce("*/")
				ops = append(ops, token)
                        case ")":
			    reduce("+-*/(")
		        default:
			    num, _ := strconv.Atoi(token)
				nums = append(nums, num)
		}
	}
    reduce("+-*/")
    return nums[0]
}

func main (){
	fmt.Println(Eval("5"))
        fmt.Println(Eval("1+2"))
	fmt.Println(Eval("1-2+3"))
	fmt.Println(Eval("3*(3+1*3)/2"))
	fmt.Println(Eval("3*((3+1)*3)/2"))

}
