package maps

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Set(key K, val V) {
	m[key] = val
}

func (m Map[K, V]) Get(key K) V {
	return m[key]
}

func (m Map[K, V]) GetOrDefault(key K, val V) V {
	if m.Has(key) {
		return m.Get(key)
	}

	return val
}

func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func (m Map[K, V]) Values() []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func (m Map[K, V]) Has(key K) bool {
	_, ok := m[key]

	return ok
}

func New[K comparable, V any](vol int) Map[K, V] {
	return make(map[K]V, vol)
}

func FromSlice[K comparable, V any](arr []V, lambda func(V) K) Map[K, V] {
	m := New[K, V](0)

	for _, v := range arr {
		m.Set(lambda(v), v)
	}

	return m
}
