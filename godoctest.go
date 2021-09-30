// Package godoctest is nonsense
//
// bla
// bla
package godoctest

import "fmt"

// NewAPI1 is a preview function (PREVIEW)
//
// This is the
// the longer
// description
func NewAPI1(x string) {
	// TODO(ansiwen): this is a todo note
	fmt.Println(x)
}

// OldAPI1 is a deprecated function (DEPRECATED)
//
// Deprecated: since v0.10.0
//
// This is the
// the longer
// description
func OldAPI1(x string) {
	fmt.Println(x)
}

// NewAPI2 is a preview function
//  PREVIEW
// This is the
// the longer
// description
func NewAPI2(x string) {
	// BUG(ansiwen): this is a bug note
	fmt.Println(x)
}

// OldAPI2 is a deprecated function
//
// Deprecated: since v0.10.0
//
// This is the
// the longer
// description
func OldAPI2(x string) {
	fmt.Println(x)
}

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
	// BUG(uid)
	// PREVIEW(uid): this is a preview note
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
// PREVIEW(uid): this is a preview note
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

// Test7 is a test for godoc (PREVIEW)
func Test7() {
	fmt.Println("test")
}
