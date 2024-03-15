// Defines a set of functions that can be applied on slices

package itertools

// Counts how many times elt appears in the slice
func Count[T comparable](iterable []T, elt T) int {
	occurence := 0
	for _, item := range iterable {
		if elt == item {
			occurence++
		}
	}

	return occurence
}

// Checks if all elements withing the slice satifies the provided predicate
func All[T any](iterable []T, predicate func(T) bool) bool {
	for _, elt := range iterable {
		if !predicate(elt) {
			return false
		}
	}

	return true
}

// Checks if at least one element within the slice satisfies the predicate
func Any[T any](iterable []T, predicate func(T) bool) bool {
	for _, elt := range iterable {
		if predicate(elt) {
			return true
		}
	}

	return false
}

/*
 ELEMENT-WISE OPERATIONS
*/

// Selects all elements held in the slice satisfying the predicate until at least one
// of them doesn't satisfy the predicate.
func TakeWhile[T any](iterable []T, predicate func(T) bool) []T {
	result := make([]T, 1)
	for _, element := range iterable {
		if predicate(element) {
			result = append(result, element)
		} else {
			break
		}
	}

	return result
}

// Skips all elements held in the slice satisfying the predicate reaching at least one of them
// which satisfies it. The function will return a new slice going from this found item till the
// end of the original slice.
func DropWhile[T any](iterable []T, predicate func(T) bool) []T {
	result := make([]T, 1)
	for i := 0; i < len(iterable); i++ {
		if predicate(iterable[i]) {
			result = append(result, iterable[i+1:]...)
			break
		}
	}

	return result
}

// Calls the given function for each element within the slice
func ForEach[T any](iterable []T, fn func(T)) {
	for _, element := range iterable {
		fn(element)
	}
}

// Transforms each element within the provide iterable into TOut elements by applying the provided
// transformation function.
func Map[TIn any, TOut any](iterable []TIn, transformFn func(TIn) TOut) []TOut {
	result := make([]TOut, len(iterable))
	for _, element := range iterable {
		result = append(result, transformFn(element))
	}

	return result
}

// Removes from the slice elements that don't satisfy the predicate
func Filter[T any](iterable []T, predicate func(T) bool) []T {
	result := make([]T, len(iterable))
	for _, element := range iterable {
		ok := predicate(element)
		if ok {
			result = append(result, element)
		}
	}

	return result
}

// Takes a 2 dimensional slice and returns a one dimensional slice containing elements
// from each of the nested original slices.
func Flatten[T any](iterable [][]T) []T {
	nElements := 0
	for _, nestedIterable := range iterable {
		nElements += len(nestedIterable)
	}

	flattened := make([]T, nElements)
	for _, nestedIterable := range iterable {
		flattened = append(flattened, nestedIterable...)
	}

	return flattened
}

// Divides a slice into groups of at most `size` elements and returns a slice of slices.
func Chunk[T any](iterable []T, size int) *ChunkResult[T] {
	nBatches := (len(iterable) + size - 1) / size
	remainder := len(iterable) % size

	result := make([][]T, nBatches)
	for i := 0; i < nBatches; i++ {
		start := i * size
		end := start + size
		if i == nBatches-1 && remainder > 0 {
			end = start + remainder
		}
		result[i] = iterable[start:end]
	}

	return &ChunkResult[T]{
		Chunks:    result,
		ChunkSize: size,
		Total:     len(iterable),
		Remainder: remainder,
	}
}

// func GroupBy[TKey comparable, TValue any](iterable []TValue, keyFn func(TValue) TKey) map[TKey][]TValue {
// 	result := make(map[TKey][]TValue)
// 	for _, element := range iterable {
// 		key := keyFn(element)
// 		group, ok := result[key]
// 		if !ok {
// 			group = make([]TValue, 0)
// 			result[key] = group
// 		}
// 		group = append(group, element)
// 	}
// 	return result
// }
