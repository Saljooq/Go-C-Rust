package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -L. -loutput
#include <stdlib.h>
extern char* process_data(char* input);
*/
import (
	"C"
)

func main() {

	h := "Hello World!"
	fmt.Println("Processing data: ", h)

	fmt.Println()
	fmt.Print("Ouput: ", ProcessData(h))
}

func ProcessData(input string) string {

	c_input := C.CString(input)
	defer C.free(unsafe.Pointer(c_input))

	c_output := C.process_data(c_input)
	defer C.free(unsafe.Pointer(c_output))

	return C.GoString(c_output)
}
