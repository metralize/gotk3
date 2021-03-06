// Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
//
// This file originated from: http://opensource.conformal.com/
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// This file includes wrapers for symbols included since GTK 3.8, and
// and should not be included in a build intended to target any older GTK
// versions.  To target an older build, such as 3.8, use
// 'go build -tags gtk_3_8'.  Otherwise, if no build tags are used, GTK 3.18
// is assumed and this file is built.
// +build !gtk_3_6

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"sync"

	"github.com/gotk3/gotk3/gdk"
)

/*
 * Constants
 */

const (
	STATE_FLAG_DIR_LTR StateFlags = C.GTK_STATE_FLAG_DIR_LTR
	STATE_FLAG_DIR_RTL StateFlags = C.GTK_STATE_FLAG_DIR_RTL
)

/*
 * GtkTickCallback
 */

// TickCallback is a representation of GtkTickCallback
type TickCallback func(widget *Widget, frameClock *gdk.FrameClock, userData ...interface{}) bool

type tickCallbackData struct {
	fn       TickCallback
	userData []interface{}
}

var (
	tickCallbackRegistry = struct {
		sync.RWMutex
		next int
		m    map[int]tickCallbackData
	}{
		next: 1,
		m:    make(map[int]tickCallbackData),
	}
)
