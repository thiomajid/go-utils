package itertools

import (
	"reflect"
	"testing"
)

func TestCountForExistingElement(t *testing.T) {
	slice := []int{1, 2, 3, 5, 1, 5}
	result := Count(slice, 1)
	expected := 2

	if result != expected {
		t.Errorf("Expected to found %d occurences but found %d occurences", expected, result)
	}
}

func TestCountForNonExistingElement(t *testing.T) {
	slice := []int{1, 2, 3, 5, 1, 5}
	value := 45
	result := Count(slice, 45)
	expected := 0

	if result != expected {
		t.Errorf("The slice doesn't contain %d but found %d occurences of %d", value, result, value)
	}
}

func TestAllWithSliceOfPositiveIntegers(t *testing.T) {
	slice := []int{1, 2, 3, 5}
	predictate := func(x int) bool {
		return x > 0
	}

	expected := true
	result := All(slice, predictate)

	if result != expected {
		t.Errorf("The slice contains at least a negative element")
	}
}

func TestAnyWithAtLeastOneZeroElement(t *testing.T) {
	slice := []int{1, 0, 2, 3, 0, 5}
	predictate := func(x int) bool {
		return x == 0
	}

	expected := true
	result := Any(slice, predictate)

	if result != expected {
		t.Errorf("The slice contains at least a negative element")
	}
}

func TestTakeWhileWithVaryingSizeStrings(t *testing.T) {
	slice := []string{"foo", "bar", "aba", "z", "45"}
	predicate := func(s string) bool {
		return len(s) == 3
	}

	expectedSliceSize := 3
	outSlice := TakeWhile(slice, predicate)

	if expectedSliceSize != len(outSlice) {
		t.Errorf("The resulting slice should contain %d elements but %d elements were found", expectedSliceSize, len(outSlice))
	}
}

func TestDropWhileWithVaryingSizeStrings(t *testing.T) {
	slice := []string{"foo", "bar", "aba", "z", "45"}
	predicate := func(s string) bool {
		return len(s) == 3
	}

	expectedSliceSize := 2
	outSlice := SkipWhile(slice, predicate)

	if expectedSliceSize != len(outSlice) {
		t.Errorf("The resulting slice should contain %d elements but %d elements were found", expectedSliceSize, len(outSlice))
	}
}

func TestMapBySquaringIntegers(t *testing.T) {
	slice := []int{1, 2, 3}
	squareFn := func(i int) int {
		return i * i
	}

	expectedSlice := []int{1, 4, 9}
	outSlice := Map(slice, squareFn)

	if len(outSlice) != len(expectedSlice) {
		t.Errorf("The resulting slice has %d items while it should contain %d items", len(outSlice), len(expectedSlice))
	}

	for idx := 0; idx < len(slice); idx++ {
		current := outSlice[idx]
		expected := expectedSlice[idx]

		if expected != current {
			t.Errorf("%d is not the square of %d", current, expected)
		}
	}
}

func TestFilterToKeepNonEmptyStrings(t *testing.T) {
	slice := []string{"", "foo", "bar", "", "baz"}
	predicate := func(s string) bool {
		return len(s) > 0
	}

	expectedSlice := []string{"foo", "bar", "baz"}
	expectedSize := 3
	outSlice := Filter(slice, predicate)

	if len(expectedSlice) != len(outSlice) {
		t.Errorf("The filtered slice contains %d items while it should contain %d items", len(outSlice), len(expectedSlice))
	}

	if len(outSlice) != 3 {
		t.Errorf("The filtered slice should have %d elements but only %d were found", expectedSize, len(outSlice))
	}

	for idx := 0; idx < len(outSlice); idx++ {
		if outSlice[idx] != expectedSlice[idx] {
			t.Errorf("%s has been modified to %s", expectedSlice[idx], outSlice[idx])
		}
	}
}

func TestFlattenWithOneEmptyNestedSlice(t *testing.T) {
	slice := [][]int{{1, 2, 3}, {}, {4, 5, 6}}
	expectedSlice := []int{1, 2, 3, 4, 5, 6}
	flattenedSlice := Flatten(slice)

	if len(flattenedSlice) != len(expectedSlice) {
		t.Errorf("The flattened slice should contain %d elements but only %d were found", len(expectedSlice), len(flattenedSlice))
	}

	for idx := 0; idx < len(flattenedSlice); idx++ {
		r := flattenedSlice[idx]
		e := expectedSlice[idx]

		if r != e {
			t.Errorf("Elements order has not been preserved, %d != %d", r, e)
		}
	}
}

func TestChunkToProduceTwoBatches(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	chunkSize := 3

	expectedChunks := [][]int{{1, 2, 3}, {4, 5, 6}}
	expected := ChunkResult[int]{
		Chunks:    expectedChunks,
		ChunkSize: chunkSize,
		Total:     len(slice),
		Remainder: len(slice) % chunkSize,
	}
	result, err := Chunk(slice, chunkSize)

	if err != nil {
		t.Error(err)
	}

	if result.ChunkSize != expected.ChunkSize {
		t.Errorf("The chunking result has a chunk size of %d while it should be %d", result.ChunkSize, expected.ChunkSize)
	}

	if result.Total != expected.Total {
		t.Errorf("The chunking result has a items %d while it should be %d", result.Total, expected.Total)
	}

	if result.Remainder != expected.Remainder {
		t.Errorf("The chunking result has a remainder of %d while it should be %d", result.Remainder, expected.Remainder)
	}

	flattenedExpectation := Flatten(expectedChunks)
	flattenedResult := Flatten(result.Chunks)

	for idx := 0; idx < len(flattenedExpectation); idx++ {
		r := flattenedResult[idx]
		e := flattenedExpectation[idx]

		if r != e {
			t.Errorf("Elements order has not been preserved, %d != %d", r, e)
		}
	}

}

func TestGroupByWithStringOfVaryingLength(t *testing.T) {
	input := []string{"a", "aa", "b", "bbb"}
	expected := map[int][]string{
		1: {"a", "b"},
		2: {"aa"},
		3: {"bbb"},
	}

	keyFn := func(s string) int {
		return len(s)
	}

	output := GroupBy(input, keyFn)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("An incorrect map has been produced")
	}
}
