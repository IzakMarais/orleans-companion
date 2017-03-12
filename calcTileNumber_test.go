package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type testData struct {
	required int
	limits   []int
}

func TestCalcTileNumbers(t *testing.T) {
	for index := 0; index < 100; index++ {
		test := genTestData()
		output := calcTileNumbers(test.required, test.limits)

		if len(output) != len(test.limits) {
			t.Errorf("Expected lengths to equal: Exp %v, got: %v",
				len(test.limits), len(output))
		}

		total := 0
		for i := 0; i < len(test.limits); i++ {
			if output[i] < 0 {
				t.Errorf("Output %v must be > 0, got: %v", i, output[i])
			}

			if output[i] > test.limits[i] {
				t.Errorf("Output %v must be < limit %v, got: %v",
					i, test.limits[i], output[i])
			}
			total += output[i]
		}
		if total != test.required {
			t.Errorf("Total should equal %v, got %v", test.required, total)
		}
		fmt.Printf("req:%v, limits:%v, result:%v\n", test.required, test.limits, output)
	}
}

func genTestData() testData {
	limits := make([]int, 4)
	totLimit := 0
	for i := 0; i < len(limits); i++ {
		limits[i] = int(rand.Intn(30))
		totLimit += limits[i]
	}
	return testData{
		required: int(rand.Intn(totLimit)),
		limits:   limits,
	}
}
