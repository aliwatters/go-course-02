package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[') // use WriteRune for non-ascii
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']') // use WriteRune for non-ascii
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3, 56})) //  "[1, 2, 3]"
}
