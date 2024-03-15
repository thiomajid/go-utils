package itertools

// Data returned after batching a slice using itertools.Batch
type ChunkResult[T any] struct {
	// A slice containing slices of at most ChunkSize elements
	Chunks [][]T

	// The maximum number of elements that each inner slice with Chunks can hold
	ChunkSize int

	// The number of elements held by the original slice that was divide into smaller chunks
	Total int

	// The size of the smallest chunk. It is equal to 0 if len(iterable) % ChunkSize != 0.
	Remainder int
}
