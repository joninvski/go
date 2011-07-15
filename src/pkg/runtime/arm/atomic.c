// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "runtime.h"

// Atomic add and return new value.
#pragma textflag 7
uint32
runtime·xadd(uint32 volatile *val, int32 delta)
{
	uint32 oval, nval;

	for(;;){
		oval = *val;
		nval = oval + delta;
		if(runtime·cas(val, oval, nval))
			return nval;
	}
}

#pragma textflag 7
uint32
runtime·atomicload(uint32 volatile* addr)
{
	return runtime·xadd(addr, 0);
}

#pragma textflag 7
void*
runtime·atomicloadp(void* volatile* addr)
{
	return (void*)runtime·xadd((uint32 volatile*)addr, 0);
}

#pragma textflag 7
void
runtime·atomicstorep(void* volatile* addr, void* v)
{
	void *old;

	for(;;) {
		old = *addr;
		if(runtime·casp(addr, old, v))
			return;
	}
}
