//
// File: signal.go
// Created by Dizzrt on 2023/02/24.
//
// Copyright (C) 2023 The oset Authors.
// This source code is licensed under the MIT license found in
// the LICENSE file in the root directory of this source tree.
//

package sig

import "sync"

type foundation struct {
	quit chan struct{}
}

var (
	f    *foundation
	once sync.Once
)

func init() {
	f = &foundation{}
}

func Quit() <-chan struct{} {
	once.Do(func() {
		f.quit = make(chan struct{})
	})

	return f.quit
}

func DoQuit() {
	close(f.quit)
}
