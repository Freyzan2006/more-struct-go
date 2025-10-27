package set

type Set[T comparable] interface {
    Add(value T)
    Has(value T) bool
    Delete(value T)
    Len() int
}


type set[T comparable] struct {
	values map[T]bool
}

func NewSet[T comparable](values []T) *set[T] {

	if len(values) == 0 {
		return &set[T]{values: make(map[T]bool)}
	}

	s := &set[T]{values: make(map[T]bool, len(values))}
	s.init(values)

    return s
}

func (s *set[T]) init(values []T) {
	for _, value := range values {
		s.values[value] = true
	}
}


func (s *set[T]) Add(value T) {
	if s.Has(value) { return }
	s.values[value] = true;
}

func (s *set[T]) Has(value T) bool {
	return s.values[value]
}

func (s *set[T]) Len() int {
	return len(s.values)
}