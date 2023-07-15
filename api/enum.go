package pb

func inverseMap[K comparable, V comparable](m map[K]V) map[V]K {
	im := make(map[V]K)
	for k, v := range m {
		im[v] = k
	}

	return im
}
