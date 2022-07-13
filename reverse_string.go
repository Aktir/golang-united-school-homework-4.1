package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	bytes := []rune(input)
	for i, v := 0, len(bytes)-1; i < v; i, v = i+1, v-1 {
		bytes[i], bytes[v] = bytes[v], bytes[i]
	}
	return string(bytes)

}
