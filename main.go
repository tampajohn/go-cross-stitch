package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lindsaylandry/go-cross-stitch/src/convert"
)

//go:generate go run scripts/setupColors.go

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input image provided")
		os.Exit(0)
	}

	num := flag.Int("n", 6, "number of colors to attempt to match (2^n)")
	rgb := flag.Bool("rgb", false, "use rgb color space")
	all := flag.Bool("all", false, "use all thread colors available (currently broken)")
	pal := flag.String("color", "dmc", "color palette to use (OPTIONS: dmc, lego)")
	flag.Parse()

	c, err := convert.NewConverter(flag.Args()[0], *num, *rgb, *all, *pal)
	if err != nil {
		panic(err)
	}

	if err := c.DMC(); err != nil {
		panic(err)
	}
}
