package yfast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateMockEntries(num int) Entries {
	entries := make(Entries, 0, num)
	for i := uint64(0); i < uint64(num); i++ {
		entries = append(entries, newMockEntry(i, i))
	}

	return entries
}

func TestRTAddSingleDimension(t *testing.T) {
	rt := new(1, uint8(0))
	e1 := newMockEntry(2)
	e2 := newMockEntry(5)

	overwritten := rt.Add(e1, e2)
	assert.Len(t, overwritten, 2)
	assert.Equal(t, Entries{nil, nil}, overwritten)

	assert.Equal(t, Entries{e1}, rt.Get(newMockEntry(2)))
	assert.Equal(t, Entries{e2}, rt.Get(newMockEntry(5)))
	assert.Equal(t, Entries{e1, e2}, rt.Get(newMockEntry(2), newMockEntry(5)))

	assert.Equal(t, Entries{nil, nil}, rt.Get(newMockEntry(18), newMockEntry(19)))
	assert.Equal(t, Entries{e1, nil}, rt.Get(newMockEntry(2), newMockEntry(3)))
}

func TestRTAddSingleDimensionOverwrite(t *testing.T) {
	rt := new(1, uint8(0))
	e1 := newMockEntry(2)
	e2 := newMockEntry(2)

	rt.Add(e1)
	overwritten := rt.Add(e2)

	assert.Equal(t, Entries{e1}, overwritten)
	assert.Equal(t, Entries{e2}, rt.Get(newMockEntry(2)))
}

func TestRTAddMultiDimension(t *testing.T) {
	rt := new(2, uint8(0))

	e1 := newMockEntry(2, 3)
	e2 := newMockEntry(3, 4)

	overwritten := rt.Add(e1, e2)
	assert.Len(t, overwritten, 2)
	assert.Equal(t, Entries{nil, nil}, overwritten)

	assert.Equal(t, Entries{e1}, rt.Get(newMockEntry(2, 3)))
	assert.Equal(t, Entries{e2}, rt.Get(newMockEntry(3, 4)))
	assert.Equal(t, Entries{e1, e2}, rt.Get(newMockEntry(2, 3), newMockEntry(3, 4)))

	assert.Equal(t, Entries{nil}, rt.Get(newMockEntry(2, 4)))
	assert.Equal(t, Entries{e1, nil}, rt.Get(newMockEntry(2, 3), newMockEntry(2, 1)))
	assert.Equal(t, Entries{e2, nil}, rt.Get(newMockEntry(3, 4), newMockEntry(3, 5)))
}

func TestRTAddMultiDimensionOverwrite(t *testing.T) {
	rt := new(2, uint8(0))
	e1 := newMockEntry(2, 3)
	e2 := newMockEntry(2, 3)

	rt.Add(e1)

	overwritten := rt.Add(e2)
	assert.Equal(t, Entries{e1}, overwritten)
	assert.Equal(t, Entries{e2}, rt.Get(newMockEntry(2, 3)))
}

func BenchmarkMultiDimensionalAdd(b *testing.B) {
	rt := new(2, uint64(0))
	entries := generateMockEntries(b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rt.Add(entries[i])
	}
}

func BenchmarkMultiDimensionalAddOverwrite(b *testing.B) {
	rt := new(2, uint64(0))
	entries := generateMockEntries(100000)
	rt.Add(entries...)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rt.Add(entries[i%100000])
	}
}

func BenchmarkMultiDimensionalGet(b *testing.B) {
	rt := new(2, uint64(0))
	entries := generateMockEntries(100000)
	rt.Add(entries...)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rt.Add(entries[i%100000])
	}
}
