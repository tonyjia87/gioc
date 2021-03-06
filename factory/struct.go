// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package factory

import (
	"github.com/vlorc/gioc/types"
	"reflect"
	"sync"
)

type ValueFactory struct {
	value interface{}
	err   error
}

type MethodFactory struct {
	retIndex int
	errIndex int
	method   reflect.Value
}

type ProxyFactory struct {
	types.Provider
	factory types.BeanFactory
}

type ExportFactory struct {
	provider func() types.Provider
	factory  types.BeanFactory
}

type MutexFactory struct {
	m       sync.Mutex
	factory types.BeanFactory
}

type TypeFactory struct {
	typ reflect.Type
}

type PointerFactory struct {
	value reflect.Value
}

type ConvertFactory struct {
	factory types.BeanFactory
	typ     reflect.Type
}

type ParamFactory int

type FuncFactory func(types.Provider) (interface{}, error)
