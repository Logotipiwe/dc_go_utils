package dc_go_utils

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func ToMap[K string, V any](arr []V, groupingFunc func(val V) K) map[K]V {
	result := make(map[K]V)
	for _, v := range arr {
		key := groupingFunc(v)
		result[key] = v
	}
	return result
}
