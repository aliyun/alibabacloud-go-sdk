// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleDataComponentsValuePropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleDataComponentsValuePropertiesValue
	GetCode() *string
	SetName(v string) *ModuleDataComponentsValuePropertiesValue
	GetName() *string
	SetValue(v string) *ModuleDataComponentsValuePropertiesValue
	GetValue() *string
	SetValues(v []*ModuleDataComponentsValuePropertiesValueValues) *ModuleDataComponentsValuePropertiesValue
	GetValues() []*ModuleDataComponentsValuePropertiesValueValues
}

type ModuleDataComponentsValuePropertiesValue struct {
	// example:
	//
	// placeholder
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// placeholder
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// placeholder
	Value  *string                                           `json:"Value,omitempty" xml:"Value,omitempty"`
	Values []*ModuleDataComponentsValuePropertiesValueValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ModuleDataComponentsValuePropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleDataComponentsValuePropertiesValue) GoString() string {
	return s.String()
}

func (s *ModuleDataComponentsValuePropertiesValue) GetCode() *string {
	return s.Code
}

func (s *ModuleDataComponentsValuePropertiesValue) GetName() *string {
	return s.Name
}

func (s *ModuleDataComponentsValuePropertiesValue) GetValue() *string {
	return s.Value
}

func (s *ModuleDataComponentsValuePropertiesValue) GetValues() []*ModuleDataComponentsValuePropertiesValueValues {
	return s.Values
}

func (s *ModuleDataComponentsValuePropertiesValue) SetCode(v string) *ModuleDataComponentsValuePropertiesValue {
	s.Code = &v
	return s
}

func (s *ModuleDataComponentsValuePropertiesValue) SetName(v string) *ModuleDataComponentsValuePropertiesValue {
	s.Name = &v
	return s
}

func (s *ModuleDataComponentsValuePropertiesValue) SetValue(v string) *ModuleDataComponentsValuePropertiesValue {
	s.Value = &v
	return s
}

func (s *ModuleDataComponentsValuePropertiesValue) SetValues(v []*ModuleDataComponentsValuePropertiesValueValues) *ModuleDataComponentsValuePropertiesValue {
	s.Values = v
	return s
}

func (s *ModuleDataComponentsValuePropertiesValue) Validate() error {
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

type ModuleDataComponentsValuePropertiesValueValues struct {
	// example:
	//
	// placeholder
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// placeholder
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// placeholder
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModuleDataComponentsValuePropertiesValueValues) String() string {
	return dara.Prettify(s)
}

func (s ModuleDataComponentsValuePropertiesValueValues) GoString() string {
	return s.String()
}

func (s *ModuleDataComponentsValuePropertiesValueValues) GetCode() *string {
	return s.Code
}

func (s *ModuleDataComponentsValuePropertiesValueValues) GetValue() *string {
	return s.Value
}

func (s *ModuleDataComponentsValuePropertiesValueValues) GetName() *string {
	return s.Name
}

func (s *ModuleDataComponentsValuePropertiesValueValues) SetCode(v string) *ModuleDataComponentsValuePropertiesValueValues {
	s.Code = &v
	return s
}

func (s *ModuleDataComponentsValuePropertiesValueValues) SetValue(v string) *ModuleDataComponentsValuePropertiesValueValues {
	s.Value = &v
	return s
}

func (s *ModuleDataComponentsValuePropertiesValueValues) SetName(v string) *ModuleDataComponentsValuePropertiesValueValues {
	s.Name = &v
	return s
}

func (s *ModuleDataComponentsValuePropertiesValueValues) Validate() error {
	return dara.Validate(s)
}
