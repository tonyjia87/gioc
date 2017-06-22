// Copyright 2017 Granitic. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be found in the LICENSE file at the root of this project.

package gioc

import (
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/binder"
	"github.com/vlorc/gioc/builder"
	"github.com/vlorc/gioc/container"
	"github.com/vlorc/gioc/depend"
	"github.com/vlorc/gioc/register"
	"github.com/vlorc/gioc/selector"
)

// create a root container
func NewRootContainer() types.Container {
	registerFactory := register.NewRegisterFactory()
	binderFactory := binder.NewBinderFactory()
	dependFactory := depend.NewDependencyFactory()
	builderFactory := builder.NewBuilderFactory()
	selectorFactory := selector.NewSelectorFactory()

	sel,err := selectorFactory.Instance(binderFactory)
	if nil != err {
		panic(err)
	}

	reg, err := registerFactory.Instance(sel)
	if nil != err {
		panic(err)
	}

	reg.RegisterInterface(&registerFactory)
	reg.RegisterInterface(&binderFactory)
	reg.RegisterInterface(&dependFactory)
	reg.RegisterInterface(&builderFactory)
	reg.RegisterInterface(&selectorFactory)

	return container.NewContainer(reg, nil, 30)
}
