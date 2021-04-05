package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0
var traceEnabled bool = false

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() {
	traceLevel = traceLevel + 1
}

func decIdent() {
	traceLevel = traceLevel - 1
}

func trace(msg string) string {
	incIdent()
	if traceEnabled {
		tracePrint("BEGIN " + msg)
	}
	return msg
}

func untrace(msg string) {
	if traceEnabled {
		tracePrint("END " + msg)
	}
	decIdent()
}
