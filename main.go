package main

import "fmt"

func main()  {
	var s []int
	var n int
	var k int
	var suma int

	//fmt.Println("Tamaño del Slice: ")
	fmt.Scanln(&n)

	for i:= 0; i < n; i++ {
		//fmt.Println("Ingresa un numero: ")
		fmt.Scanln(&k)
		s = append(s,k)
	}

	//fmt.Println("Suma de Slide: ")
	
	for j:= 0; j < n; j++{
		suma += s[j]
	}
	
	fmt.Println(suma)

}