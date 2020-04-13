package main

import "fmt"

func main() {
	// Go의 유일한 반복문 for
	//예제
	var a int16 = 20
	if a >= 15 {
		fmt.Println("15이상")
	}

	b := 30
	if b >= 25 {
		fmt.Println("25이상")
	}

	if c := 100; c < 25 {
		fmt.Println("25미만")
	} else if c > 50 && c < 70 {
		fmt.Println("50초과 70미만 ")
	} else {
		fmt.Println("70이상")
	}

	//늘리거나 {}블럭으로 안감싸면 지멋대로 ; 넣어버려서 안됨

}
