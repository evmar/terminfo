package main

import (
	"fmt"
	"flag"
	"os"
	"terminfo"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s <terminfo file>\n", os.Args[0])
	}
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	filename := flag.Args()[0]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer f.Close()
	term, err := terminfo.Parse(f)
	if err != nil {
		fmt.Printf("parsing %s: %s\n", filename, err)
		os.Exit(1)
	}
	fmt.Printf("%s: %dx%d\n", term.Names[0],
		term.Numbers[terminfo.Columns], term.Numbers[terminfo.Lines])
	for id := range terminfo.StringAttrNames {
		name := terminfo.StringAttrNames[id]
		text := term.Strings[terminfo.StringAttr(id)]
		if text == "" {
			continue
		}
		fmt.Printf("%s ", name)
		for _, l := range text {
			if l >= 0x20 && l < 0x7F {
				fmt.Printf("%c", l)
			} else if l == 0x1b {
				fmt.Printf("\\E")
			} else {
				fmt.Printf("\\x%02x", l)
			}
		}
		fmt.Printf("\n")
	}
}
