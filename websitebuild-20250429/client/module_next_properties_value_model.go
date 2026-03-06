// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleNextPropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleNextPropertiesValue
	GetCode() *string
	SetValues(v []*ModuleNextPropertiesValueValues) *ModuleNextPropertiesValue
	GetValues() []*ModuleNextPropertiesValueValues
}

type ModuleNextPropertiesValue struct {
	// example:
	//
	// placeholder
	Code   *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Values []*ModuleNextPropertiesValueValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ModuleNextPropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleNextPropertiesValue) GoString() string {
	return s.String()
}

func (s *ModuleNextPropertiesValue) GetCode() *string {
	return s.Code
}

func (s *ModuleNextPropertiesValue) GetValues() []*ModuleNextPropertiesValueValues {
	return s.Values
}

func (s *ModuleNextPropertiesValue) SetCode(v string) *ModuleNextPropertiesValue {
	s.Code = &v
	return s
}

func (s *ModuleNextPropertiesValue) SetValues(v []*ModuleNextPropertiesValueValues) *ModuleNextPropertiesValue {
	s.Values = v
	return s
}

func (s *ModuleNextPropertiesValue) Validate() error {
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

type ModuleNextPropertiesValueValues struct {
	// placeholder
	//
	// example:
	//
	// placeholder
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModuleNextPropertiesValueValues) String() string {
	return dara.Prettify(s)
}

func (s ModuleNextPropertiesValueValues) GoString() string {
	return s.String()
}

func (s *ModuleNextPropertiesValueValues) GetValue() *string {
	return s.Value
}

func (s *ModuleNextPropertiesValueValues) SetValue(v string) *ModuleNextPropertiesValueValues {
	s.Value = &v
	return s
}

func (s *ModuleNextPropertiesValueValues) Validate() error {
	return dara.Validate(s)
}
