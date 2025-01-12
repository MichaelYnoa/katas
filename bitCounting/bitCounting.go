/*
Write a function that takes an integer as input,
and returns the number of bits that are equal to one in the
binary representation of that number. You can guarantee that
input is non-negative.

Example: The binary representation
of 1234 is 10011010010, so the function
should return 5 in this case

*/

package main

import "fmt"

func CountBits(n uint) int {

	// n no puede ser negativo

	var result int

	if n <= 0 {
		return 0
	} else {
		//convertir n a binario
		binario := fmt.Sprintf("%b", n)
		//recorrer n e ir sumando los #1 que encuentre
		for _, v := range binario {
			if v != 48 { //48 es el codigo para 0 en unicode
				result++
			}
		}
		// retornar el numero de unos que encontre
		return result
	}

}

func main() {
	fmt.Println(CountBits(1234))
}
