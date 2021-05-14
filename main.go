// Unit-convert is a general-purpose unit-conversion program that reads number
// from command-line arguments or from the standard input if there are not
// arguments, and converts each number into units.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unit-convert/length"
	"unit-convert/temp"
	"unit-convert/weight"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			num, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "unit-conver: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(convert(num))
		}
	}
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unit-convert: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(convert(num))
	}
}

func convert(x float64) string {
	c := temp.Celsius(x)
	f := temp.Fahrenheit(x)
	k := weight.Kilogram(x)
	p := weight.Pound(x)
	m := length.Meter(x)
	ft := length.Foot(x)
	return fmt.Sprintf("Temperature: %s = %s, %s = %s\nWeight: %s = %s, %s = %s\nLength: %s = %s, %s = %s", c, temp.CToF(c), f, temp.FToC(f), k, weight.KToP(k), p, weight.PToK(p), m, length.MToF(m), ft, length.FToM(ft))

}
