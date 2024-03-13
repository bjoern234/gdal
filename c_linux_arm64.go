//go:build linux && arm64
// +build linux,arm64

package gdal

/*
#cgo CFLAGS: -I/usr/include/gdal
#cgo LDFLAGS: -lgdal
*/

import "C"
