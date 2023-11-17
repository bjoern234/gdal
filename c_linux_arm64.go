//go:build linux && arm64
// +build linux,arm64

package gdal

/*
#cgo LDFLAGS: -L/usr/lib -lgdal_i
#cgo CFLAGS: -I/usr/include
*/

import "C"
