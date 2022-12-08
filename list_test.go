package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithCapacity(t *testing.T) {
	list := WithCapacity[string](5)

	assert.Equal(t, len(list.inner), 5)
}

func TestAddWithinCapacity(t *testing.T) {
	list := New[int]()

	list.Add(1)

	assert.Equal(t, list.inner[0], 1)
	assert.Equal(t, list.length, 1)
}

func TestAddOutsideOfCapacity(t *testing.T) {
	list := New[int]()

	for i := 0; i < defaultCapacity; i++ {
		list.Add(i)
	}

	list.Add(-1)

	assert.Equal(t, list.inner[defaultCapacity], -1)
}

func TestToArray(t *testing.T) {
	arr := make([]int, 0)
	list := New[int]()

	arr = append(arr, 1)
	list.Add(1)

	arr = append(arr, 2)
	list.Add(2)

	assert.Equal(t, list.ToArray(), arr)
}

func TestRemove(t *testing.T) {
	list := New[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	list.Remove(2)

	assert.Equal(t, []int{1, 3, 4}, list.ToArray())
}

func TestRemoveDoesntExist(t *testing.T) {
	list := New[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	err := list.Remove(-1)

	assert.Equal(t, err, ErrElementNotFound)
}

func TestIndexOf(t *testing.T) {
	list := New[string]()

	list.Add("foo")
	list.Add("bar")
	list.Add("")
	list.Add("bar")

	index, err := list.IndexOf("bar")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, index)
}

func TestContains(t *testing.T) {
	type testCase struct {
		list     List[string]
		input    string
		expected bool
	}

	testCases := []testCase{
		{
			list:     FromArray([]string{"foo"}),
			input:    "foo",
			expected: true,
		},
		{
			list:     FromArray([]string{"bar"}),
			input:    "foo",
			expected: false,
		},
		{
			list:     FromArray([]string{"foo", "foo"}),
			input:    "foo",
			expected: true,
		},
		{
			list:     New[string](),
			input:    "foo",
			expected: false,
		},
	}

	for _, testCase := range testCases {
		actual := testCase.list.Contains(testCase.input)
		assert.Equal(t, testCase.expected, actual)
	}
}

func TestLen(t *testing.T) {
	list := WithCapacity[string](5)

	list.Add("smaller")
	list.Add("than")
	list.Add("capacity")

	assert.Equal(t, list.Len(), 3)
}

func TestReverse(t *testing.T) {
	list := New[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Reverse()

	assert.Equal(t, []int{4, 3, 2, 1}, list.ToArray())
}

//---------
// BENCHES
//---------

var existingList List[string] = New[string]()
var existingArr []string = make([]string, 0)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		existingList.Add("foo")
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		existingArr = append(existingArr, "foo")
	}
}

func BenchmarkReverse(b *testing.B) {
	for _, size := range []int{100, 1000, 10000, 100000} {
		list := New[int]()
		for i := 0; i < size; i++ {
			list.Add(i)
		}

		b.Run(fmt.Sprintf("list_size%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.Reverse()
			}
		})
	}
}
