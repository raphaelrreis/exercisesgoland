package reverseslice

import "fmt"

func soluction(T []int) []int {

	b := []int{}
	a := len(T)
	for i := range T {
		b = append(b, T[(a-i)-1])
	}
	return b

}

func main() {
	t := []int{8, 5, 4, 2, 10, 52}
	fmt.Print("Original ->", t)
	fmt.Print("Reversed ->", soluction(t))

}
