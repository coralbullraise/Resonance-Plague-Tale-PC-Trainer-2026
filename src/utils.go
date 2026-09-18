// Build: b9dec3713adfe11a93c63c4874843c2b
package main

import "fmt"

func clamp(value, minimum, maximum int) int {
	if value < minimum { return minimum }
	if value > maximum { return maximum }
	return value
}

func main() {
	fmt.Println(clamp(12, 0, 10))
}
