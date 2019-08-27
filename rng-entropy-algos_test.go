package main

import (
	"fmt"
	"testing"
	"gonum.org/v1/gonum/stat"
)

// Test the degree to which the TimeNowDelta algorithm produces random values.
//
// Evaluate the Chi Square distance between expected and actual frequency of the
// range of characters produced, for various values of N calls to the function.
//
// The output of the TimeNowDelta algorithm for N values is compared against an
// expected uniform distribution of the same number of characters (length of map
// returned from TimeNowDeltaForNValues()) with equal frequency, calculated by
// map length divided by N values (CreateUniformSlice()).
//
// The Chi Square statistic is computed by stat.ChiSquare (from the gonum
// package), which takes slices of equal length of actual and expected
// frequency (where the actual frequency slice is produced by GetValuesSlice()).
//
// TODO: Provide better measure and understanding of the results of the Chi
//   Square statistics produced for each value of N, as pertains to its
//   significance in assessing randomness of the TimeNowDelta() algorithm.

func TimeNowDeltaForNValues(n int) map[string]int {
	valuesmap := make(map[string]int)

	for i:=0; i < n; i++ {
		char := TimeNowDelta()

		_, ok := valuesmap[char]
		if ok {
			valuesmap[char] += 1
		} else {
			valuesmap[char] = 1
		}
	}

	return valuesmap
}

func GetValuesSlice(m map[string]int) []float64 {
	values := make([]float64, 0, len(m))

	for _, v := range m {
		values = append(values, float64(v))
	}

	return values
}

func CreateUniformSlice(n, actuallength int) []float64 {
	uniform := make([]float64, actuallength)

	for i := range uniform {
		uniform[i] = float64(actuallength) / float64(n)
	}

	return uniform
}

func TestTimeNowDeltaOneThousandValues(t *testing.T) {
	n := 1000
	actualmap := TimeNowDeltaForNValues(n)

	actual := GetValuesSlice(actualmap)
	expected := CreateUniformSlice(n, len(actual))

	chisquare := stat.ChiSquare(actual, expected)
	fmt.Printf("chisquare: %v,\tn: %v,\tcharacters: %v\n", chisquare, n, len(actual))
}

func TestTimeNowDeltaTenThousandValues(t *testing.T) {
	n := 10000
	actualmap := TimeNowDeltaForNValues(n)

	actual := GetValuesSlice(actualmap)
	expected := CreateUniformSlice(n, len(actual))

	chisquare := stat.ChiSquare(actual, expected)
	fmt.Printf("chisquare: %v,\tn: %v,\tcharacters: %v\n", chisquare, n, len(actual))
}

func TestTimeNowDeltaOneHundredThousandValues(t *testing.T) {
	n := 100000
	actualmap := TimeNowDeltaForNValues(n)

	actual := GetValuesSlice(actualmap)
	expected := CreateUniformSlice(n, len(actual))

	chisquare := stat.ChiSquare(actual, expected)
	fmt.Printf("chisquare: %v,\tn: %v,\tcharacters: %v\n", chisquare, n, len(actual))
}

func TestTimeNowDeltaOneMillionValues(t *testing.T) {
	n := 1000000
	actualmap := TimeNowDeltaForNValues(n)

	actual := GetValuesSlice(actualmap)
	expected := CreateUniformSlice(n, len(actual))

	chisquare := stat.ChiSquare(actual, expected)
	fmt.Printf("chisquare: %v,\tn: %v,\tcharacters: %v\n", chisquare, n, len(actual))
}

func TestTimeNowDeltaTenMillionValues(t *testing.T) {
	n := 10000000
	actualmap := TimeNowDeltaForNValues(n)

	actual := GetValuesSlice(actualmap)
	expected := CreateUniformSlice(n, len(actual))

	chisquare := stat.ChiSquare(actual, expected)
	fmt.Printf("chisquare: %v,\tn: %v,\tcharacters: %v\n", chisquare, n, len(actual))
}

func TestTimeNowDeltaOneHundredMillionValues(t *testing.T) {
	n := 100000000
	actualmap := TimeNowDeltaForNValues(n)

	actual := GetValuesSlice(actualmap)
	expected := CreateUniformSlice(n, len(actual))

	chisquare := stat.ChiSquare(actual, expected)
	fmt.Printf("chisquare: %v,\tn: %v,\tcharacters: %v\n", chisquare, n, len(actual))
}

func TestTimeNowDeltaOneBillionValues(t *testing.T) {
	n := 1000000000
	actualmap := TimeNowDeltaForNValues(n)

	actual := GetValuesSlice(actualmap)
	expected := CreateUniformSlice(n, len(actual))

	chisquare := stat.ChiSquare(actual, expected)
	fmt.Printf("chisquare: %v,\tn: %v,\tcharacters: %v\n", chisquare, n, len(actual))
}
