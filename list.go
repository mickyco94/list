package list

import (
	"errors"
)

var (
	ErrElementNotFound = errors.New("Element not found")
)

// List is an implementation of .NET List<T>
// written in GO. This is not a fully-featured implementation
// and supports only the core methods
type List[T comparable] struct {
	inner  []T
	length int
}

// defaultCapacity is the size of the internal data structure
// that is used when `New()` is used to construct a new list
var defaultCapacity = 4

// New constructs an empty List of elements T
func New[T comparable]() List[T] {
	return List[T]{
		inner: make([]T, defaultCapacity),
	}
}

// WithCapacity constructs a list with the specified capacity.
func WithCapacity[T comparable](capacity int) List[T] {
	return List[T]{
		inner: make([]T, capacity),
	}
}

// FromArray constructs a List from an existing array
func FromArray[T comparable](arr []T) List[T] {
	return List[T]{
		inner:  arr,
		length: len(arr),
	}
}

// capacity returns the number of elements that can be added to this List
// without the need for resizing. Upon resizing capacity will grow.
func (list *List[T]) capacity() int {
	return len(list.inner)
}

func (list *List[T]) ToArray() []T {
	return list.inner[:list.length]
}

// Add adds a new element to the list.
// This increases the length of the list by 1, if needed
// the underlying data structure is resized to make room for this new element
func (list *List[T]) Add(element T) {

	if list.length == list.capacity() {
		newInner := make([]T, len(list.inner)*2)
		copy(newInner, list.inner)
		list.inner = newInner
	}

	list.inner[list.length] = element
	list.length++
}

// Remove removes the specified element from the list if it is present.
// If element is not present in the list an ErrElementNotFound is returned
func (list *List[T]) Remove(element T) error {
	index, err := list.IndexOf(element)
	if err != nil {
		return err
	}

	copy(list.inner[index:], list.inner[index+1:])
	list.length--

	return nil
}

// IndexOf returns the index of the first element specified found in the List.
func (list *List[T]) IndexOf(element T) (int, error) {
	for i, v := range list.inner {
		if v == element {
			return i, nil
		}
	}
	return -1, ErrElementNotFound
}

// Contains indicates if the specified element is present in the List.
func (list *List[T]) Contains(element T) bool {
	_, err := list.IndexOf(element)
	return err == nil
}

// Reverse reverses the list in place
func (list *List[T]) Reverse() {
	split := list.length / 2

	for i := 0; i < split; i++ {
		temp := list.inner[i]
		list.inner[i] = list.inner[list.length-1-i]
		list.inner[list.length-1-i] = temp
	}
}

// Len returns the length of the list, i.e. the number of elemnts currently
// in the list.
func (list *List[T]) Len() int {
	return list.length
}
