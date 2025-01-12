package main

import "fmt"

/*
Complete the function that accepts a string parameter, and reverses
each word in the string. All spaces in the string should be retained.

Examples

"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"

*/

/*
mi codigo del kata

func ReverseWords(str string) string {

	var str1 string
	var str2 string
	tamano := len(str)

	for i := 0; i < tamano; i++ {
		// vamos a ir guardando en str1 hasta que aparezca un espacio
		if str[i] != ' ' {
			str1 += string(str[i])
		}
		//cuando aparezca un espacio o se termine la palabra
		if str[i] == ' ' || i == tamano-1 {

			//manejo para cuando termina con espacio

			if i == tamano-1 && str[i] == ' ' {
				str1 += string(str[i])
			}

			// invertimos la palabra numero N
			for j := len(str1) - 1; j >= 0; j-- {
				str2 += string(str1[j])
			}

			//agregamos los espacios a str2 para que no siga en bucle
			if str[i] == ' ' {
				str2 += " "
			}

			// vaceamos str1 para la proxima palabra
			str1 = ""

		}

	}

	//fmt.Println(str2)

	str = str2

	return str // reverse those words
}

*/

// mejor codigo del kata
func ReverseWords(str string) string {
	var rev string
	var word string

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
}
func main() {

	fmt.Println(ReverseWords("uno dos tres cuatro"))

}
