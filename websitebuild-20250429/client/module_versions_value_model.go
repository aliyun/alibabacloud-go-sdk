// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleVersionsValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleVersionsValue
	GetCode() *string
	SetName(v string) *ModuleVersionsValue
	GetName() *string
	SetComponents(v map[string]*ModuleVersionsValueComponentsValue) *ModuleVersionsValue
	GetComponents() map[string]*ModuleVersionsValueComponentsValue
}

type ModuleVersionsValue struct {
	// code
	//
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// name
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// components
	Components map[string]*ModuleVersionsValueComponentsValue `json:"Components,omitempty" xml:"Components,omitempty"`
}

func (s ModuleVersionsValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleVersionsValue) GoString() string {
	return s.String()
}

func (s *ModuleVersionsValue) GetCode() *string {
	return s.Code
}

func (s *ModuleVersionsValue) GetName() *string {
	return s.Name
}

func (s *ModuleVersionsValue) GetComponents() map[string]*ModuleVersionsValueComponentsValue {
	return s.Components
}

func (s *ModuleVersionsValue) SetCode(v string) *ModuleVersionsValue {
	s.Code = &v
	return s
}

func (s *ModuleVersionsValue) SetName(v string) *ModuleVersionsValue {
	s.Name = &v
	return s
}

func (s *ModuleVersionsValue) SetComponents(v map[string]*ModuleVersionsValueComponentsValue) *ModuleVersionsValue {
	s.Components = v
	return s
}

func (s *ModuleVersionsValue) Validate() error {
	return dara.Validate(s)
}
