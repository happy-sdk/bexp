// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Happy Authors

package app

func osmain(ch chan struct{}) {
	if ch != nil {
		<-ch
	} else {
		select {}
	}
}
