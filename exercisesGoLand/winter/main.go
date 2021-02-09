package winter

import "fmt"

func soluction(T []int) int {

	maxInv := T[0]
	var indexStartSummer int
	maxTemp := T[10]

	for i := range T {
		if maxTemp < T[i] {
			maxTemp = T[i]
		}
		if maxInv >= T[i] {
			indexStartSummer = -1
			maxInv = maxTemp

		} else {
			if indexStartSummer == -1 {
				indexStartSummer = i
			}
		}
	}

	return indexStartSummer
}

func main() {
	t := []int{-10, 5, -12, 3, 8, 9}
	fmt.Print("length of winter", soluction(t))

}
