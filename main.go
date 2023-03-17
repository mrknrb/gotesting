package main

import (
	"fmt"
	"github.com/pkg/browser"
	"gotesting/helpers"
)

func main() {
	helpers.Tes("gdfg", "fghd")
	helpers.Tes2("gdfg", "fghd")
	helpers.Tes27("HH", "gdfg")
	fmt.Println("a, b")
	browser.OpenURL()

}
