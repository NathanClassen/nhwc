package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type countdata struct {
	filename string
	data     []byte
	linec    int
	bytec    int
	wordc    int
	charc    int
}

var c *bool = flag.Bool("c", false, "the number of bytes in each input file is written to the standard output.  This will cancel out any prior usage of the -m option.")
var l *bool = flag.Bool("l", false, "The number of lines in each input file is written to the standard output.")
var m *bool = flag.Bool("m", false, "The number of characters in each input file is written to the standard output.  If the current locale does not support multibyte characters, this is equivalent to the -c option.  This will cancel out any prior usage of the -c option.")
var w *bool = flag.Bool("w", false, "The number of words in each input file is written to the standard output.")

var b strings.Builder

var allF bool

func main() {
	flag.Parse()

	if *c && *m {
		*c = false
	}
	if !(*l || *w || *m || *c) {
		allF = true
	}
	cntobjects, err := getCountData()
	if err != nil {
		log.Fatalln(err)
	}

	processCounts(cntobjects)
	writeres(cntobjects)

	os.Exit(0)
}

func writeres(cnts []*countdata) {
	ti := 0
	tl := 0
	tw := 0
	tb := 0
	tc := 0

	for _, cnt := range cnts {
		ti++
		if *l || allF {
			fmt.Fprintf(&b, " %7d", cnt.linec)
			tl += cnt.linec
		}
		if *w || allF {
			fmt.Fprintf(&b, " %7d", cnt.wordc)
			tw += cnt.wordc
		}
		if *c || allF {
			fmt.Fprintf(&b, " %7d", cnt.bytec)
			tb += cnt.bytec
		}
		if *m {
			fmt.Fprintf(&b, " %7d", cnt.charc)
			tc += cnt.charc
		}
		fmt.Fprintf(&b, " %s\n", cnt.filename)
	}

	if ti > 1 {
		if *l || allF {
			fmt.Fprintf(&b, " %7d", tl)
		}
		if *w || allF {
			fmt.Fprintf(&b, " %7d", tw)
		}
		if *c || allF {
			fmt.Fprintf(&b, " %7d", tb)
		}
		if *m {
			fmt.Fprintf(&b, " %7d", tc)
		}
		fmt.Fprintf(&b, " total\n")
	}

	fmt.Print(b.String())
}

func processCounts(cntobjects []*countdata) {
	for _, cd := range cntobjects {
		if allF {
			cd.lres()
			cd.wres()
			cd.cres()
			continue
		}
		if *l {
			cd.lres()
		}
		if *w {
			cd.wres()
		}
		if *c {
			cd.cres()
		}
		if *m {
			cd.mres()
		}
	}
}

func (cd *countdata) lres() {
	cd.linec = bytes.Count(cd.data, []byte{'\n'})
}

func (cd *countdata) wres() {
	wc := 0
	fc := cd.data[0]
	inw := !unicode.IsSpace(rune(fc))
	for _, b := range cd.data {
		isws := unicode.IsSpace(rune(b))
		if inw {
			if isws {
				wc++
				inw = false
			}
		}
		if !inw {
			if !isws {
				inw = true
			}
		}
	}
	cd.wordc = wc
}

func (cd *countdata) mres() {
	cd.charc = utf8.RuneCount(cd.data)
}

func (cd *countdata) cres() {
	cd.bytec = len(cd.data)
}

func getCountData() ([]*countdata, error) {
	args := flag.Args()

	var data []*countdata

	if len(args) == 0 {
		return getStdinAsInput()
	}

	for _, filen := range args {
		d, e := os.ReadFile(filen)

		if e != nil {
			return nil, e
		}
		data = append(data, &countdata{
			filename: filen,
			data:     d,
		})
	}

	return data, nil
}

func getStdinAsInput() ([]*countdata, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	return []*countdata{{filename: "", data: data}}, nil
}
