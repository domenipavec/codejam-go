package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeapTuple(t *testing.T) {
	mht := NewMinHeapTuple(NewSliceTuple(3, 2, 5))
	mht1 := mht.Copy()

	assert.Equal(t, Tuple(2), mht.Min())
	assert.Equal(t, Tuple(2), mht1.Min())

	mht.Push(Tuple(1))

	assert.Equal(t, Tuple(1), mht.Min())
	assert.Equal(t, Tuple(2), mht1.Min())

	assert.Equal(t, Tuple(1), mht.Pop())
	assert.Equal(t, Tuple(2), mht1.Min())

	assert.Equal(t, Tuple(2), mht.Min())
	assert.Equal(t, Tuple(2), mht1.Min())

	mht.Min()[0] = 6
	mht.FixMin()

	assert.Equal(t, Tuple(3), mht.Pop())
	assert.Equal(t, Tuple(2), mht1.Min())
	assert.Equal(t, Tuple(5), mht.Pop())
	assert.Equal(t, Tuple(2), mht1.Min())
	assert.Equal(t, Tuple(6), mht.Pop())
	assert.Equal(t, Tuple(2), mht1.Min())
}