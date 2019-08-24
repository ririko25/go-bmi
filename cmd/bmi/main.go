package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ririko25/go-bmi"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("体重(kg)、身長(cm)の順で入力してください。")
		os.Exit(1)
	}

	weight, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("体重(kg)は数字で入力してください。")
		os.Exit(1)
	}
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("身長(cm)は数字で入力してください。")
		os.Exit(1)
	}
	bmi, err := bmi.Calc(weight, height)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("あなたのbmi値は%vです。", bmi)
}
