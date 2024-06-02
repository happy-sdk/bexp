// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Happy Authors

//go:build (linux && !android) || freebsd || openbsd

package app

import "github.com/happy-sdk/happy/sdk/app/internal/application"

func osmain(ch <-chan application.ShutDown) {
	if ch != nil {
		<-ch
	} else {
		select {}
	}
}
