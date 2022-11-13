// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2014 Eric Woroshow

package gl

import (
	"github.com/ebitengine/purego"
)

var opengl = purego.Dlopen("/System/Library/Frameworks/OpenGL.framework/Versions/Current/OpenGL", purego.RTLD_GLOBAL)

func getProcAddress(name string) uintptr {
	return purego.Dlsym(opengl, name)
}
