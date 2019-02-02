package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L.build/debug/ -lgo_swift
#include <stdlib.h>
#include "include/go_swift.h"
*/
import "C"
import "unsafe"

func main() {
	// Convert Go string to C string
	cstr := C.CString("Bob")

	// Call Swift function `sayHello``
	C.sayHello(cstr)

	// The C string is allocated in the C heap using malloc.
	// Therefore memory must be freed when we're done
	C.free(unsafe.Pointer(cstr))
}
