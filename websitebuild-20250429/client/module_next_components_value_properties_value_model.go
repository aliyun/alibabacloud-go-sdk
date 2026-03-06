// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleNextComponentsValuePropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleNextComponentsValuePropertiesValue
	GetCode() *string
	SetName(v string) *ModuleNextComponentsValuePropertiesValue
	GetName() *string
	SetValue(v string) *ModuleNextComponentsValuePropertiesValue
	GetValue() *string
	SetValues(v []*ModuleNextComponentsValuePropertiesValueValues) *ModuleNextComponentsValuePropertiesValue
	GetValues() []*ModuleNextComponentsValuePropertiesValueValues
}

type ModuleNextComponentsValuePropertiesValue struct {
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
	Values []*ModuleNextComponentsValuePropertiesValueValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ModuleNextComponentsValuePropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleNextComponentsValuePropertiesValue) GoString() string {
	return s.String()
}

func (s *ModuleNextComponentsValuePropertiesValue) GetCode() *string {
	return s.Code
}

func (s *ModuleNextComponentsValuePropertiesValue) GetName() *string {
	return s.Name
}

func (s *ModuleNextComponentsValuePropertiesValue) GetValue() *string {
	return s.Value
}

func (s *ModuleNextComponentsValuePropertiesValue) GetValues() []*ModuleNextComponentsValuePropertiesValueValues {
	return s.Values
}

func (s *ModuleNextComponentsValuePropertiesValue) SetCode(v string) *ModuleNextComponentsValuePropertiesValue {
	s.Code = &v
	return s
}

func (s *ModuleNextComponentsValuePropertiesValue) SetName(v string) *ModuleNextComponentsValuePropertiesValue {
	s.Name = &v
	return s
}

func (s *ModuleNextComponentsValuePropertiesValue) SetValue(v string) *ModuleNextComponentsValuePropertiesValue {
	s.Value = &v
	return s
}

func (s *ModuleNextComponentsValuePropertiesValue) SetValues(v []*ModuleNextComponentsValuePropertiesValueValues) *ModuleNextComponentsValuePropertiesValue {
	s.Values = v
	return s
}

func (s *ModuleNextComponentsValuePropertiesValue) Validate() error {
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

type ModuleNextComponentsValuePropertiesValueValues struct {
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

func (s ModuleNextComponentsValuePropertiesValueValues) String() string {
	return dara.Prettify(s)
}

func (s ModuleNextComponentsValuePropertiesValueValues) GoString() string {
	return s.String()
}

func (s *ModuleNextComponentsValuePropertiesValueValues) GetCode() *string {
	return s.Code
}

func (s *ModuleNextComponentsValuePropertiesValueValues) GetValue() *string {
	return s.Value
}

func (s *ModuleNextComponentsValuePropertiesValueValues) GetName() *string {
	return s.Name
}

func (s *ModuleNextComponentsValuePropertiesValueValues) SetCode(v string) *ModuleNextComponentsValuePropertiesValueValues {
	s.Code = &v
	return s
}

func (s *ModuleNextComponentsValuePropertiesValueValues) SetValue(v string) *ModuleNextComponentsValuePropertiesValueValues {
	s.Value = &v
	return s
}

func (s *ModuleNextComponentsValuePropertiesValueValues) SetName(v string) *ModuleNextComponentsValuePropertiesValueValues {
	s.Name = &v
	return s
}

func (s *ModuleNextComponentsValuePropertiesValueValues) Validate() error {
	return dara.Validate(s)
}
