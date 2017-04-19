package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"strings"

	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/net/idna"

	"sort"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in dots Per inch")
	fontfile = flag.String("font-file", "", "filename of the ttf font")
	size     = flag.Float64("size", 12, "font size in points")
)

var font2 *truetype.Font

type program struct {
}

func drawRune(r rune) ([]uint8, error) {
	fg, bg := image.Black, image.White

	rgba := image.NewRGBA(image.Rect(0, 0, 20, 20))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(font2)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	pt := freetype.Pt(0, 10)

	_, err := c.DrawString(string(r), pt)
	if err != nil {
		return []byte{}, err
	}

	return rgba.Pix, nil
}

func x(homographs map[rune][]rune, c chan string, root string, j int) {
	if j > len([]rune(root)) {
		c <- root
		return
	}

	for i, r := range []rune(root) {
		// find all possible options for each letter
		if alts, ok := homographs[r]; !ok {
			x(homographs, c, root, j+1)
			continue
		} else if len(alts) == 0 {
			x(homographs, c, root, j+1)
			continue
		}

		// find alternatives
		for _, ar := range homographs[r] {
			bla := []rune(root)
			bla[i] = ar
			x(homographs, c, string(bla), j+1)
		}
	}
}

// RuneSlice attaches the methods of Interface to []int, sorting in increasing order.
type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func usage() {
	fmt.Println("Usage:")
	fmt.Println("homograph --font-file ./Arial.ttf apple.com")
}

func main() {
	flag.Parse()

	if *fontfile == "" {
		usage()
		return
	}

	if len(flag.Args()) == 0 {
		usage()
		return
	}

	domain := flag.Args()[0]

	root := ""

	parts := strings.Split(domain, ".")
	if len(parts) == 1 {
		fmt.Printf("Invalid domain name: %s\n", domain)
		return
	}

	root = parts[len(parts)-2]

	fmt.Println("Homographs: brutefind homographs within a font")
	fmt.Println("by Remco Verhoef (@remco_verhoef) of DutchSec")
	fmt.Println("---------------------------------------------------")

	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		fmt.Printf("Could not read font: %s\n", err.Error())
		return
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("Could not parse font: %s", err.Error())
		return
	}

	font2 = f

	homographs := map[rune][]rune{}

	for _, i := range []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") {
		homographs[i] = []rune{}

		b1, err := drawRune(i)
		if err != nil {
			fmt.Println("Could not draw rune 0x%x: %s", i, err.Error())
			continue
		}

		for j := rune(0); j < 5000; j++ {
			if i == j {
				continue
			}

			b2, err := drawRune(rune(j))
			if err != nil {
				fmt.Println("Could not draw rune 0x%x: %s", j, err.Error())
				continue
			}

			// drawed runes are not equal
			if bytes.Compare(b1, b2) != 0 {
				continue
			}

			// found one
			homographs[i] = append(homographs[i], j)
		}
	}

	// sort
	sortedKeys := make([]rune, 0, len(homographs))
	for key := range homographs {
		sortedKeys = append(sortedKeys, key)
	}

	sort.Sort(RuneSlice(sortedKeys))

	fmt.Println("[+] Homograph table")

	// print
	for _, k := range sortedKeys {
		runes := homographs[k]

		if len(runes) == 0 {
			continue
		}

		fmt.Printf("%c(0x%x) | ", k, k)

		for i, r := range runes {
			if i > 0 {
				fmt.Printf(", ")
			}

			fmt.Printf("%c(0x%x)", r, r)
		}

		fmt.Println()
	}

	vals := []string{}

	c := make(chan string)
	go func() {
		for s := range c {
			found := false

			// inefficient way of deduplication
			for _, s2 := range vals {
				found = found || (s2 == s)
			}

			if found {
				continue
			}

			vals = append(vals, s)
		}
	}()

	x(homographs, c, root, 0)

	fmt.Println("")

	fmt.Println("[+] Homograph domains")

	for _, s := range vals {
		u, err := idna.ToASCII(s)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("%s %s\n", s+"."+parts[len(parts)-1:][0], u+"."+parts[len(parts)-1:][0])
	}

	fmt.Println("")
}
