package lib

import "fmt"

func init() {
	fmt.Println("lib package init start!")
}

func CheckNum(c int32) bool {
	return c > 10
}
