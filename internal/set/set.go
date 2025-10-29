package set

type Set[T comparable] interface {
    Add(value T)
    Has(value T) bool
    Delete(value T)
    Len() int
	Clear()
	// Возвращает новое множество, содержащее 
	// только общие элементы обоих множеств
	intersection(other Set[T]) Set[T]
	// Возвращает новое множество, 
	// содержащее элементы, которые есть в s1, но нет в s2
	difference(s Set[T], s2 Set[T]) Set[T]
	//  Возвращает новое множество, содержащее 
	// элементы, которые присутствуют только в 
	// одном из множеств
	symmetricDifference(other Set[T]) Set[T]
	// Проверяет, является ли текущее 
	// множество подмножеством другого
	isSubsetOf(other Set[T]) bool

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

func (s *set[T]) Delete(value T) {
	delete(s.values, value)
}

func (s *set[T]) Clear() {
	s.values = make(map[T]bool)
}