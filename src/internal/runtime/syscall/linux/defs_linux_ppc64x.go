// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (ppc64 || ppc64le)

package linux

const (
	SYS_CLOSE         = 6
	SYS_FCNTL         = 55
	SYS_MPROTECT      = 125
	SYS_PRCTL         = 171
	SYS_EPOLL_CTL     = 237
	SYS_EPOLL_PWAIT   = 303
	SYS_EPOLL_CREATE1 = 315
	SYS_EPOLL_PWAIT2  = 441
	SYS_EVENTFD2      = 314
	SYS_OPENAT        = 286
	SYS_PREAD64       = 179
	SYS_READ          = 3

	EFD_NONBLOCK = 0x800

	O_LARGEFILE = 0x0
)

type EpollEvent struct {
	Events    uint32
	pad_cgo_0 [4]byte
	Data      [8]byte // unaligned uintptr
}
