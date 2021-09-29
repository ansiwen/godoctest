// Package godoctest is nonsense
//
// bla
// bla
package godoctest

import "fmt"

// Test is a test for godoc
//
// This is another span that
// should be rendered differently than
// the first line.
//
// Can we do *bold* or _italic_ test? For sure we can do
//  preformatted text
//  -----------------
// because that's documented.
//
// Preview: would be awesome if that would work
//
func Test() {
	fmt.Println("test")
}

// Test2 is a test for godoc
//
// DEPRECATED: does it work with all caps?
//
func Test2() {
	fmt.Println("test")
}

// Test3 bla
//
// This is a Header
//
// This
// is
// ~one~
// paragraph.
//
// <b>PREVIEW</b>
//
func Test3() {
	fmt.Println("test")
}

// Test4 is a test for godoc
//
// Deprecated: this should work.
//
func Test4() {
	fmt.Println("test")
}

// Test5 is a test for godoc
//
//  PREVIEW
//
func Test5() {
	fmt.Println("test")
}

// Test6 is a test for godoc
//
// PREVIEW
//
// header works only with following paragraph
func Test6() {
	fmt.Println("test")
}
