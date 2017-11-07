// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package tag

import (
	"strconv"
	"reflect"
)

func NewParamString(v string) Param{
	return ParamString(v[1 : len(v) - 1])
}

func NewParamEscapeString(v string) Param{
	v,err := strconv.Unquote(v)
	if nil != err {
		panic(err)
	}
	return ParamString(v)
}

func NewParamNumber(v string) Param{
	return ParamNumber(ParamString(v).Number())
}

func NewParamFloat(v string) Param{
	return ParamFloat(ParamString(v).Float())
}

func NewParamNull(v bool) Param{
	return ParamNull(v)
}

func(ps ParamString)String() string {
	return string(ps)
}

func(ps ParamString)Number() int64 {
	v, err := strconv.ParseInt(ps.String(), 10, 64)
	if nil != err {
		panic(err)
	}
	return v
}

func(ps ParamString)Float() float64{
	v, err := strconv.ParseFloat(ps.String(), 64)
	if nil != err {
		panic(err)
	}
	return v
}

func(ps ParamString)Value() reflect.Value{
	return reflect.ValueOf(ps.String())
}

func(ps ParamString)Boolean() bool {
	return "" == ps
}

func(ps ParamString)Kind() Kind {
	return String
}

func(pn ParamNumber)String() string {
	return strconv.FormatInt(pn.Number(), 10)
}

func(pn ParamNumber)Number() int64 {
	return int64(pn)
}

func(pn ParamNumber)Float() float64{
	return float64(pn)
}

func(pn ParamNumber)Boolean() bool {
	return 0 != pn
}

func(pn ParamNumber)Kind() Kind {
	return Int
}

func(pn ParamNumber)Value() reflect.Value{
	return reflect.ValueOf(pn.Number())
}

func(pf ParamFloat)String() string {
	return strconv.FormatFloat(pf.Float(),'f',-1,64)
}

func(pf ParamFloat)Number() int64 {
	return int64(pf)
}

func(pf ParamFloat)Float() float64{
	return float64(pf)
}

func(pf ParamFloat)Boolean() bool {
	return 0 != pf
}

func(pf ParamFloat)Kind() Kind {
	return Float
}

func(pf ParamFloat)Value() reflect.Value{
	return reflect.ValueOf(pf.Float())
}

func(pe ParamNull)String() string {
	return strconv.FormatFloat(pe.Float(),'f',-1,64)
}

func(pe ParamNull)Number() int64 {
	return 0
}

func(pe ParamNull)Float() float64{
	return 0
}

func(pe ParamNull)Boolean() bool {
	return bool(pe)
}

func(pe ParamNull)Kind() Kind {
	return Null
}

func(pe ParamNull)Value() reflect.Value{
	return reflect.Value{}
}