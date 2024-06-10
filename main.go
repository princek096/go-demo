package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	// fmt.Println(mascot.Best())
	// fmt.Println(quote.Go())
	// fmt.Println("ABC")
	// var a int
	// var s string = "abc"
	// fmt.Println(a)
	// fmt.Println(s)

	// fmt.Println("ABC", "CDE")

	// fmt.Print("ABC", "CDE ")
	// var ac = false
	// if ac {
	// 	fmt.Println("True  is")
	// } else {
	// 	fmt.Println("false is ")
	// }
	// age := 35
	// if age >= 62 {
	// 	fmt.Println("You are retiring")
	// }
	// if age >= 30 && age <= 40 {
	// 	fmt.Println("Promotion is due")
	// }
	rand.Seed(time.Now().UnixNano())
	_, file := path.Split("css/main.css")
	fmt.Println(file)
	fmt.Printf("%#v\n", os.Args[1])
	fmt.Println(len(os.Args))
	rp := len(os.Args[1])
	ms := strings.ToUpper(os.Args[1]) + strings.Repeat("!", rp)
	fmt.Println(ms)
}
