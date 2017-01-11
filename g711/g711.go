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
import "unsafe"

// EncodeA function as declared in g711/g711_interface.h:42
func EncodeA(speechIn []int16, len uint, encoded []byte) uint {
	cspeechIn, _ := (*C.int16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&speechIn)).Data)), cgoAllocsUnknown
	clen, _ := (C.size_t)(len), cgoAllocsUnknown
	cencoded, _ := (*C.uint8_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&encoded)).Data)), cgoAllocsUnknown
	__ret := C.WebRtcG711_EncodeA(cspeechIn, clen, cencoded)
	__v := (uint)(__ret)
	return __v
}

// EncodeU function as declared in g711/g711_interface.h:63
func EncodeU(speechIn []int16, len uint, encoded []byte) uint {
	cspeechIn, _ := (*C.int16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&speechIn)).Data)), cgoAllocsUnknown
	clen, _ := (C.size_t)(len), cgoAllocsUnknown
	cencoded, _ := (*C.uint8_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&encoded)).Data)), cgoAllocsUnknown
	__ret := C.WebRtcG711_EncodeU(cspeechIn, clen, cencoded)
	__v := (uint)(__ret)
	return __v
}

// DecodeA function as declared in g711/g711_interface.h:86
func DecodeA(encoded string, len uint, decoded []int16, speechType *int16) uint {
	cencoded, _ := unpackPUint8TString(encoded)
	clen, _ := (C.size_t)(len), cgoAllocsUnknown
	cdecoded, _ := (*C.int16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&decoded)).Data)), cgoAllocsUnknown
	cspeechType, _ := (*C.int16_t)(unsafe.Pointer(speechType)), cgoAllocsUnknown
	__ret := C.WebRtcG711_DecodeA(cencoded, clen, cdecoded, cspeechType)
	__v := (uint)(__ret)
	return __v
}

// DecodeU function as declared in g711/g711_interface.h:110
func DecodeU(encoded string, len uint, decoded []int16, speechType *int16) uint {
	cencoded, _ := unpackPUint8TString(encoded)
	clen, _ := (C.size_t)(len), cgoAllocsUnknown
	cdecoded, _ := (*C.int16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&decoded)).Data)), cgoAllocsUnknown
	cspeechType, _ := (*C.int16_t)(unsafe.Pointer(speechType)), cgoAllocsUnknown
	__ret := C.WebRtcG711_DecodeU(cencoded, clen, cdecoded, cspeechType)
	__v := (uint)(__ret)
	return __v
}

// Version function as declared in g711/g711_interface.h:130
func Version(version []byte, lenBytes int16) int16 {
	cversion, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&version)).Data)), cgoAllocsUnknown
	clenBytes, _ := (C.int16_t)(lenBytes), cgoAllocsUnknown
	__ret := C.WebRtcG711_Version(cversion, clenBytes)
	__v := (int16)(__ret)
	return __v
}
