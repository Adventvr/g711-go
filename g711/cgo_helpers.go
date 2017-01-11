// MIT

// WARNING: This file has automatically been generated on Wed, 11 Jan 2017 14:21:04 CST.
// By https://git.io/cgogen. DO NOT EDIT.

package g711

/*
#include "g711/g711.h"
#include "g711/g711_interface.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// unpackPUint8TString represents the data from Go string as *C.uint8_t and avoids copying.
func unpackPUint8TString(str string) (*C.uint8_t, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.uint8_t)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}
