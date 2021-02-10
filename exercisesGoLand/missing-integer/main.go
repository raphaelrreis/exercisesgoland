/**
Return the smallest positive integer greater than 0 that does not occur in slice
- the slice in generated randomly
*/

package missing_integer

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Create slice talking the number of elements and the minRange, maxRange) for the randNum Function,
returns the positive integer greater than 0 that does not occur in slice
*/
func createSlice(elements, minRange, maxRange int) int {
	var s []int

	for i := 0; i <= elements; i++ {
		num := randNum(minRange, maxRange)
		s = append(s, num)
	}
	fmt.Print(s)

	a := createMap(s)

	return a
}

/**
Create a random number
*/
func randNum(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

/**
Create a map from the slice and returns the positive integer greater than 0
*/
func createMap(s []int) int {
	m := make(map[int]int)

	b := false

	positive := 1

	for _, val := range s {
		if val > 0 {
			m[val] = val
		}

		for !b {
			if _, ok := m[positive]; !ok {
				b = true
			} else {
				positive++
			}
		}

	}

	return positive
}

func main() {

	pos := createSlice(10, -50, 20)
	fmt.Print(pos)

}
