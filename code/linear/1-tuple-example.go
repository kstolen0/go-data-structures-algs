package linear

import (
	"fmt"
)

func powerSeries(a int) (int, int) {

	return a * a, a * a * a
}

func threeElements() (int, int, int) {
	return 1, 2, 3
}

func SimpleTuple() {

	fmt.Println(powerSeries(3))
	fmt.Println(threeElements())
}
