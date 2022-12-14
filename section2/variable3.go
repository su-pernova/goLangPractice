package main

import "fmt"

func main() {
	// 짧은 선언 : var 키워드를 사용하지 않고도 변수를 선언할 수 있다.
	// 짧은 선언은 선언했던 함수 안에서만 사용할 수 있다. (지역변수 개념)(전역변수로 사용 불가)
	// 선언 후 값을 다시 초기화하면 예외가 발생한다. (Like 상수)
	// 주로 제한된 범위의 함수 내에서 사용 할 경우 코드 가독성을 높일 수 있다.

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false
	// shortVar1 := 10 // 선언 후 값 재할당시 예외 발생
	shortVar1 = 10 // 이건 된다 : 초기화가 아닌 대입이기 때문 -> 초기화와 대입의 차이를 찾아보면 좋을 것 같다.

	fmt.Println("shorVar1 : ", shortVar1, " / shorVar2 : ", shortVar2, " / shorVar3 : ", shortVar3)

	// 활용 예시 : if 문에서의 짧은 선언
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success!")
	}
}
