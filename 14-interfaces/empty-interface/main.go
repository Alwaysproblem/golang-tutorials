package main

import ("fmt"
		"errors")

type emp interface{}

func maxs(a interface{}, b interface{}) (interface{}, error) {
	v, intTa := a.(int)
	x, intTb := b.(int)
	if intTa && intTb{
		if v > x {
			return v, nil
		}
		return x, nil
	}
	return nil, errors.New("Type error")

}

func main() {
	var a = 1
	var b = 2
	fmt.Println(maxs(a, b))
}