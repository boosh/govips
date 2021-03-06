package vips

// #cgo pkg-config: vips
// #include "create.h"
import "C"

// https://libvips.github.io/libvips/API/current/libvips-create.html#vips-xyz
func vipsXYZ(width int, height int) (*C.VipsImage, error) {
	var out *C.VipsImage

	if err := C.xyz(&out, C.int(width), C.int(height)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}
