//   #################################################################   //
//                Exercise #1: Loops and Fuctions                        //
//   #################################################################   //

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	error := 1/math.Pow(10,10)
	for pre:=1.1; math.Abs(pre-z)>=error;{
		pre = z
		z = z - (z*z - x)/(2*z)
	}
	return z
	
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

//   #################################################################   //
//                Exercise #2: Slices                                    //
//   #################################################################   //

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	var img = make([][]uint8, dy)
	for i := range img {
		img[i] = make([]uint8, dx)
		for j := range img[i] {
			//img[i][j] = uint8((i+j)/2)
			//img[i][j] = uint8(i*j)	
			img[i][j] = uint8(i^j)	
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
}

//   #################################################################   //
//                Exercise #3: Maps                                      //
//   #################################################################   //

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words = strings.Fields(s)
	var m = make(map[string]int)
	for _,w := range words {
		m[w] += 1 
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}

//   #################################################################   //
//                Exercise #4: Fibonacci closure                         //
//   #################################################################   //

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i0 := 0
	i1 := 1
	return func() int {
		aux := i0
		i0 = i1
		i1 += aux
		return i1-i0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


//   #################################################################   //
//                Exercise #5: Stringers                                 //
//   #################################################################   //

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipv4 IPAddr) String() string {
	//return fmt.Sprintf("%v.%v.%v.%v",ipv4[0],ipv4[1],ipv4[2],ipv4[3])
	out := fmt.Sprintf("%v",ipv4[0])
	for i:= range ipv4[1:] {
		out = out + fmt.Sprintf(".%v",ipv4[i])
	}
	return out
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}



//   #################################################################   //
//                Exercise #6: Errors                                    //
//   #################################################################   //

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x<0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	err := 1/math.Pow(10,10)
	for pre:=1.1; math.Abs(pre-z)>=err;{
		pre = z
		z = z - (z*z - x)/(2*z)
	}
	return z , nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}


//   #################################################################   //
//                Exercise #7: Readers                                   //
//   #################################################################   //

package main

import "golang.org/x/tour/reader"

type MyReader struct{
	letter byte
}

func NewReader(s byte) *MyReader {return &MyReader{s}}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	b[0] = NewReader('A').letter
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}


//   #################################################################   //
//                Exercise #8: rot13Reader                               //
//   #################################################################   //

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	var base byte;
	for i := 0 ; i<len(p) ; i++ {
		if p[i]>='A' && p[i]<='Z' {
			base = 'A'
		} else if p[i]>='a' && p[i]<='z' {
			base = 'a'
		} else {
			base = 0
		}
		if base != 0 {
			p[i] = ((p[i] - base) + 13) % 26 + base
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}


//   #################################################################   //
//                Exercise #9: Images                                    //
//   #################################################################   //

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)
	
type Image struct{
	rect image.Rectangle
	color uint8
}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (i Image) Bounds() image.Rectangle {
	return i.rect
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color + uint8(x+y), i.color + uint8(x+y), 255, 255}
}

func main() {
	m := Image{image.Rect(0, 0, 128, 128), 255}
	pic.ShowImage(m)
}


//   #################################################################   //
//                Exercise #10: Equivalent Binary Trees                  //
//   #################################################################   //

package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	recWalk(t, ch)
	close(ch)
}

func recWalk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		recWalk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		recWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if equals:= (<-ch1 == <-ch2); !equals {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	fmt.Println("Walk Test:")
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	
	test := make([]bool, 2)
	test[0] = Same(tree.New(1), tree.New(1))
	test[1] = Same(tree.New(1), tree.New(2))
	for i,v := range test {
		fmt.Printf("Same Test %v:", i)
		if v {
			fmt.Println("They have the same elements")
		} else {
			fmt.Println("They have different elements")
		}
	}
}


//   #################################################################   //
//                Exercise #11: Web Crawler                              //
//   #################################################################   //

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type HashTable struct {
	Map map[string]bool
	mux sync.Mutex
}

type Result struct {
	url string
	body string
	urls []string
}
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan Result,
		  visited HashTable) {
	// This implementation doesn't do either:
	defer close(ch)
	
	visited.mux.Lock()
	if depth <= 0 || visited.Map[url] {
		return
	}
	visited.Map[url] = true
	visited.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- Result{url, body, urls}
	subch := make([]chan Result, len(urls))
	for i, u := range urls {
		subch[i] = make(chan Result)
		go Crawl(u, depth-1, fetcher, subch[i], visited)
	}
	for i:= range subch {
		for resp := range subch[i] {
			ch <- resp
		}
	}
	
	return
}

func main() {
	ch := make(chan Result)
	visited := HashTable{Map: make(map[string] bool)}
	go Crawl("http://golang.org/", 4, fetcher, ch, visited)
	for res := range ch {
		fmt.Printf("found: %v %q\n", res.url, res.body)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
