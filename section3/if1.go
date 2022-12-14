package main

import "fmt"

func main() {
	// 제어문(조건문) : if 문을 사용할 수 있다.
	// 조건에는 반드시 Boolean 타입을 사용해야 한다. (1, 0 등의 값은 자동 형변환 되지 않으며 사용할 수 없다.)
	// 소괄호를 사용하지 않는다.

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("a 는 15 이상이다.")
	}

	if b >= 25 {
		fmt.Println("b 는 25 이상이다.")
	}

	/*
		[ 에러가 발생하는 경우 ]
		1. Go 언어는 소스코드를 바이트코드로 변환할 때 문장의 마지막에 세미콜론(;)을 붙이기 때문에
		아래와 같이 if 문 중괄호 이전에 줄바꿈이 들어가면 에러가 발생한다.
			if a >= 15
			{
				fmt.Println("a 는 15 이상이다.")
			}

		2. 중괄호를 생략해도 에러가 발생한다.
			if a >= 25
				fmt.Println("a 는 15 이상이다.")

		3. Boolean 값 대신 1, 0 의 값을 사용할 수 없다.
			if c:= 1; c {
				fmt.Println("True")
			}
	*/

	if c := 40; c >= 35 {
		fmt.Println("c 는 35 이상이다.")
	}
	// c = 10 // 에러발생 : c는 if 문 안에서 사용된 이후 소멸되기 때문에 재사용할 수 없다.

}
