package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	uni = flag.Bool("u", false, "\\uXXXX style")
)

func getCodePoint(s string) (ret string) {
	r := []rune(s)
	size := len(r)
	for i := 0; i < size; i++ {
		ret += fmt.Sprintf("%U", r[i])
	}
	return ret
}

func getUnicodeScalar(s string) string {
	return strings.ToLower(strings.Replace(s, "U+", "\\u", -1))
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "too few arguments\n")
		os.Exit(1)
	}

	s := flag.Arg(0)
	if strings.HasPrefix(s, "U+") || strings.HasPrefix(s, "\\u") {
		// TODO
	} else {
		if *uni {
			// \uXXXX
			fmt.Printf("%s\n", getUnicodeScalar(getCodePoint(s)))
		} else {
			// U+XXXX
			fmt.Printf("%s\n", getCodePoint(s))
		}
	}
}
