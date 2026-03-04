// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleVersionsValueComponentsValuePropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleVersionsValueComponentsValuePropertiesValue
	GetCode() *string
	SetName(v string) *ModuleVersionsValueComponentsValuePropertiesValue
	GetName() *string
	SetValue(v string) *ModuleVersionsValueComponentsValuePropertiesValue
	GetValue() *string
	SetValues(v []*ModuleVersionsValueComponentsValuePropertiesValueValues) *ModuleVersionsValueComponentsValuePropertiesValue
	GetValues() []*ModuleVersionsValueComponentsValuePropertiesValueValues
}

type ModuleVersionsValueComponentsValuePropertiesValue struct {
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// value
	Value  *string                                                    `json:"Value,omitempty" xml:"Value,omitempty"`
	Values []*ModuleVersionsValueComponentsValuePropertiesValueValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ModuleVersionsValueComponentsValuePropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleVersionsValueComponentsValuePropertiesValue) GoString() string {
	return s.String()
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) GetCode() *string {
	return s.Code
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) GetName() *string {
	return s.Name
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) GetValue() *string {
	return s.Value
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) GetValues() []*ModuleVersionsValueComponentsValuePropertiesValueValues {
	return s.Values
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) SetCode(v string) *ModuleVersionsValueComponentsValuePropertiesValue {
	s.Code = &v
	return s
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) SetName(v string) *ModuleVersionsValueComponentsValuePropertiesValue {
	s.Name = &v
	return s
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) SetValue(v string) *ModuleVersionsValueComponentsValuePropertiesValue {
	s.Value = &v
	return s
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) SetValues(v []*ModuleVersionsValueComponentsValuePropertiesValueValues) *ModuleVersionsValueComponentsValuePropertiesValue {
	s.Values = v
	return s
}

func (s *ModuleVersionsValueComponentsValuePropertiesValue) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModuleVersionsValueComponentsValuePropertiesValueValues struct {
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModuleVersionsValueComponentsValuePropertiesValueValues) String() string {
	return dara.Prettify(s)
}

func (s ModuleVersionsValueComponentsValuePropertiesValueValues) GoString() string {
	return s.String()
}

func (s *ModuleVersionsValueComponentsValuePropertiesValueValues) GetCode() *string {
	return s.Code
}

func (s *ModuleVersionsValueComponentsValuePropertiesValueValues) GetValue() *string {
	return s.Value
}

func (s *ModuleVersionsValueComponentsValuePropertiesValueValues) GetName() *string {
	return s.Name
}

func (s *ModuleVersionsValueComponentsValuePropertiesValueValues) SetCode(v string) *ModuleVersionsValueComponentsValuePropertiesValueValues {
	s.Code = &v
	return s
}

func (s *ModuleVersionsValueComponentsValuePropertiesValueValues) SetValue(v string) *ModuleVersionsValueComponentsValuePropertiesValueValues {
	s.Value = &v
	return s
}

func (s *ModuleVersionsValueComponentsValuePropertiesValueValues) SetName(v string) *ModuleVersionsValueComponentsValuePropertiesValueValues {
	s.Name = &v
	return s
}

func (s *ModuleVersionsValueComponentsValuePropertiesValueValues) Validate() error {
	return dara.Validate(s)
}
