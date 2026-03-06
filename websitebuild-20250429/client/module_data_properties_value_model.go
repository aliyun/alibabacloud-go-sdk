// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleDataPropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleDataPropertiesValue
	GetCode() *string
	SetValues(v []*ModuleDataPropertiesValueValues) *ModuleDataPropertiesValue
	GetValues() []*ModuleDataPropertiesValueValues
}

type ModuleDataPropertiesValue struct {
	// example:
	//
	// placeholder
	Code   *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Values []*ModuleDataPropertiesValueValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ModuleDataPropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleDataPropertiesValue) GoString() string {
	return s.String()
}

func (s *ModuleDataPropertiesValue) GetCode() *string {
	return s.Code
}

func (s *ModuleDataPropertiesValue) GetValues() []*ModuleDataPropertiesValueValues {
	return s.Values
}

func (s *ModuleDataPropertiesValue) SetCode(v string) *ModuleDataPropertiesValue {
	s.Code = &v
	return s
}

func (s *ModuleDataPropertiesValue) SetValues(v []*ModuleDataPropertiesValueValues) *ModuleDataPropertiesValue {
	s.Values = v
	return s
}

func (s *ModuleDataPropertiesValue) Validate() error {
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

type ModuleDataPropertiesValueValues struct {
	// placeholder
	//
	// example:
	//
	// placeholder
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModuleDataPropertiesValueValues) String() string {
	return dara.Prettify(s)
}

func (s ModuleDataPropertiesValueValues) GoString() string {
	return s.String()
}

func (s *ModuleDataPropertiesValueValues) GetValue() *string {
	return s.Value
}

func (s *ModuleDataPropertiesValueValues) SetValue(v string) *ModuleDataPropertiesValueValues {
	s.Value = &v
	return s
}

func (s *ModuleDataPropertiesValueValues) Validate() error {
	return dara.Validate(s)
}
