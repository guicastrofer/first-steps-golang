package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"

	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "var3"
		variavel4 string = "var4"
	)
	fmt.Println(variavel3, variavel4)

	var5, var6 := "var5", "var6"
	fmt.Println(var5, var6)

	const constante1 string = "Constante1"
	fmt.Println(constante1)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}
