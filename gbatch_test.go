package gbatch

import (
    "github.com/stretchr/testify/require"
    "testing"
)

func TestBatch(t *testing.T) {
    const size = 3
    batch := New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, size)

    require.Equal(t, []int{1, 2, 3}, batch.Batch())
    require.Equal(t, []int{4, 5, 6}, batch.Batch())
    require.Equal(t, []int{7, 8, 9}, batch.Batch())
    require.Equal(t, []int{10, 11}, batch.Batch())
    require.Empty(t, batch.Batch())
}

func TestEmptySlice(t *testing.T) {
    const size = 5
    batch := New([]string{}, size)

    require.Empty(t, batch.Batch())
}

func TestSizeGreaterThanSliceLen(t *testing.T) {
    const size = 100000000
    batch := New([]float64{1, 2, 3, 4, 5, 6, 7}, size)

    require.Equal(t, []float64{1, 2, 3, 4, 5, 6, 7}, batch.Batch())
    require.Empty(t, batch.Batch())
}

func TestNext(t *testing.T) {
    const size = 3
    batch := New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, size)

    require.Equal(t, []int{1, 2, 3}, batch.Curr())
    require.Equal(t, []int{1, 2, 3}, batch.Curr())
    require.True(t, batch.Next())
    require.Equal(t, []int{4, 5, 6}, batch.Curr())
    require.True(t, batch.Next())
    require.Equal(t, []int{7, 8, 9}, batch.Curr())
    require.True(t, batch.Next())
    require.Equal(t, []int{10, 11, 12}, batch.Curr())
    require.False(t, batch.Next())
}
