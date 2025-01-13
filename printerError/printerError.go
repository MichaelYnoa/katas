/*
In a factory a printer prints labels for boxes. For one kind of boxes the printer has
to use colors which, for the sake of simplicity, are named with letters from a to m.

The colors used by the printer are recorded in a control string. For example a "good"
control string would be aaabbbbhaijjjm meaning that the printer used three times color a,
four times color b, one time color h then one time color a...

Sometimes there are problems: lack of colors, technical malfunction and a "bad" control string is
produced e.g. aaaxbbbbyyhwawiwjjjwwm with letters not from a to m.

You have to write a function printer_error which given a string will return the error rate of the
printer as a string representing a rational whose numerator is the number of errors and the
denominator the length of the control string. Don't reduce this fraction to a simpler expression.

The string has a length greater or equal to one and contains only letters from a to z.

Examples:

s="aaabbbbhaijjjm"
printer_error(s) => "0/14"

s="aaaxbbbbyyhwawiwjjjwwm"
printer_error(s) => "8/22"
*/
package main

import (
	"fmt"
)

func PrinterError(s string) string {
	numerador := 0
	denominador := 0

	// leer el string de caracteres
	for i := range s {
		//identificar los casos malos, los que NO son de a to m (a, b, c, d, e, f, g, h, i, j, k, l, m)
		item := string(s[i])

		switch item {
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m":
			continue
		default:
			numerador++
		}

	}

	// tener el numero total del string para el denominador
	denominador = len(s)

	// necesito manejar mayusculas ? no estoy seguro

	// retornar la fraccion como string

	return string(fmt.Sprintf("%d/%d", numerador, denominador))
}

func main() {
	fmt.Println(PrinterError("aaaxbbbbyyhwawiwjjjwwm"))
}
