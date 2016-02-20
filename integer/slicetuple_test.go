package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceTuple(t *testing.T) {
	st := NewSliceTuple(2, 3)

	assert.Equal(t, 2, len(st))
	assert.Equal(t, Tuple(2), st[0])
	assert.Equal(t, Tuple(3), st[1])

	st.Prepend(Tuple(4))

	assert.Equal(t, 3, len(st))
	assert.Equal(t, Tuple(4), st[0])
	assert.Equal(t, Tuple(2), st[1])
	assert.Equal(t, Tuple(3), st[2])

	st.Append(Tuple(5))

	assert.Equal(t, 4, len(st))
	assert.Equal(t, Tuple(4), st[0])
	assert.Equal(t, Tuple(2), st[1])
	assert.Equal(t, Tuple(3), st[2])
	assert.Equal(t, Tuple(5), st[3])

	st.Insert(2, Tuple(1))

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(4), st[0])
	assert.Equal(t, Tuple(2), st[1])
	assert.Equal(t, Tuple(1), st[2])
	assert.Equal(t, Tuple(3), st[3])
	assert.Equal(t, Tuple(5), st[4])

	st.PrefixConst(-1)

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(-1, 4), st[0])
	assert.Equal(t, Tuple(-1, 2), st[1])
	assert.Equal(t, Tuple(-1, 1), st[2])
	assert.Equal(t, Tuple(-1, 3), st[3])
	assert.Equal(t, Tuple(-1, 5), st[4])

	st.PrefixIndex()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4), st[0])
	assert.Equal(t, Tuple(1, -1, 2), st[1])
	assert.Equal(t, Tuple(2, -1, 1), st[2])
	assert.Equal(t, Tuple(3, -1, 3), st[3])
	assert.Equal(t, Tuple(4, -1, 5), st[4])

	st.PostfixConst(1)

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4, 1), st[0])
	assert.Equal(t, Tuple(1, -1, 2, 1), st[1])
	assert.Equal(t, Tuple(2, -1, 1, 1), st[2])
	assert.Equal(t, Tuple(3, -1, 3, 1), st[3])
	assert.Equal(t, Tuple(4, -1, 5, 1), st[4])

	st.PostfixIndex()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st[0])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[1])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st[2])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[3])
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st[4])

	st.Swap(2, 3)

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st[0])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[1])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[2])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st[3])
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st[4])

	st.Reverse()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st[0])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st[1])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[2])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[3])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st[4])

	st.SortAsc()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st[0])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[1])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st[2])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[3])
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st[4])

	st.SortDesc()

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st[0])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[1])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st[2])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[3])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st[4])

	st1 := st.Copy()
	st[0][0] = 3

	assert.Equal(t, 5, len(st1))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st1[0])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st1[1])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st1[2])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st1[3])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st1[4])

	assert.Equal(t, 5, len(st))
	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st[0])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[1])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st[2])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[3])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st[4])

	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st.Delete(2))

	assert.Equal(t, 5, len(st1))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st1[0])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st1[1])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st1[2])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st1[3])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st1[4])

	assert.Equal(t, 4, len(st))
	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st[0])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[1])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[2])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st[3])

	assert.Equal(t, Tuple(3, -1, 5, 1, 4), st.DeleteFirst())

	assert.Equal(t, 5, len(st1))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st1[0])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st1[1])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st1[2])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st1[3])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st1[4])

	assert.Equal(t, 3, len(st))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[0])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[1])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st[2])

	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st.DeleteLast())

	assert.Equal(t, 5, len(st1))
	assert.Equal(t, Tuple(4, -1, 5, 1, 4), st1[0])
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st1[1])
	assert.Equal(t, Tuple(2, -1, 1, 1, 2), st1[2])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st1[3])
	assert.Equal(t, Tuple(0, -1, 4, 1, 0), st1[4])

	assert.Equal(t, 2, len(st))
	assert.Equal(t, Tuple(3, -1, 3, 1, 3), st[0])
	assert.Equal(t, Tuple(1, -1, 2, 1, 1), st[1])
}