package yukpiz

func UniqInts(values []int) []int {
	m := make(map[int]bool)
	uniq := []int{}
	for _, value := range values {
		if !m[value] {
			m[value] = true
			uniq = append(uniq, value)
		}
	}
	return uniq
}
