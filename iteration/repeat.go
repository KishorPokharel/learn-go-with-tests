package iteration

// Repeat takes in a string and repeats it by the count and returns the repeated string
func Repeat(character string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
