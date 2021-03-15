package main

import (
	"fmt"

	"github.com/dariolpz/fp-factory/firstpresentent"
)

func main() {
	fmt.Print("AR FIRST PRESENTMENT\n")
	fpFactory := firstpresentent.FPFactory{}
	fp, err := fpFactory.InitFirstPresentment("AR")
	if err != nil {
		fmt.Print(err)
	}

	fp.AddDES()
	fp.AddPDS()
	fmt.Print("\n")

	fmt.Print("BR FIRST PRESENTMENT\n")
	fp, err = fpFactory.InitFirstPresentment("BR")
	if err != nil {
		fmt.Print(err)
	}

	fp.AddDES()
	fp.AddPDS()

}
