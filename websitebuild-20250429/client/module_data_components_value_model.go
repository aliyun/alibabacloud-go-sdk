// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleDataComponentsValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleDataComponentsValue
	GetCode() *string
	SetName(v string) *ModuleDataComponentsValue
	GetName() *string
	SetInstanceProperty(v []*ModuleDataComponentsValueInstanceProperty) *ModuleDataComponentsValue
	GetInstanceProperty() []*ModuleDataComponentsValueInstanceProperty
	SetProperties(v map[string]*ModuleDataComponentsValuePropertiesValue) *ModuleDataComponentsValue
	GetProperties() map[string]*ModuleDataComponentsValuePropertiesValue
	SetModuleAttrStatus(v int32) *ModuleDataComponentsValue
	GetModuleAttrStatus() *int32
}

type ModuleDataComponentsValue struct {
	// example:
	//
	// placeholder
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// placeholder
	Name             *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	InstanceProperty []*ModuleDataComponentsValueInstanceProperty         `json:"InstanceProperty,omitempty" xml:"InstanceProperty,omitempty" type:"Repeated"`
	Properties       map[string]*ModuleDataComponentsValuePropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// placeholder
	ModuleAttrStatus *int32 `json:"ModuleAttrStatus,omitempty" xml:"ModuleAttrStatus,omitempty"`
}

func (s ModuleDataComponentsValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleDataComponentsValue) GoString() string {
	return s.String()
}

func (s *ModuleDataComponentsValue) GetCode() *string {
	return s.Code
}

func (s *ModuleDataComponentsValue) GetName() *string {
	return s.Name
}

func (s *ModuleDataComponentsValue) GetInstanceProperty() []*ModuleDataComponentsValueInstanceProperty {
	return s.InstanceProperty
}

func (s *ModuleDataComponentsValue) GetProperties() map[string]*ModuleDataComponentsValuePropertiesValue {
	return s.Properties
}

func (s *ModuleDataComponentsValue) GetModuleAttrStatus() *int32 {
	return s.ModuleAttrStatus
}

func (s *ModuleDataComponentsValue) SetCode(v string) *ModuleDataComponentsValue {
	s.Code = &v
	return s
}

func (s *ModuleDataComponentsValue) SetName(v string) *ModuleDataComponentsValue {
	s.Name = &v
	return s
}

func (s *ModuleDataComponentsValue) SetInstanceProperty(v []*ModuleDataComponentsValueInstanceProperty) *ModuleDataComponentsValue {
	s.InstanceProperty = v
	return s
}

func (s *ModuleDataComponentsValue) SetProperties(v map[string]*ModuleDataComponentsValuePropertiesValue) *ModuleDataComponentsValue {
	s.Properties = v
	return s
}

func (s *ModuleDataComponentsValue) SetModuleAttrStatus(v int32) *ModuleDataComponentsValue {
	s.ModuleAttrStatus = &v
	return s
}

func (s *ModuleDataComponentsValue) Validate() error {
	if s.InstanceProperty != nil {
		for _, item := range s.InstanceProperty {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModuleDataComponentsValueInstanceProperty struct {
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
	Value  *string                                            `json:"Value,omitempty" xml:"Value,omitempty"`
	Values []*ModuleDataComponentsValueInstancePropertyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ModuleDataComponentsValueInstanceProperty) String() string {
	return dara.Prettify(s)
}

func (s ModuleDataComponentsValueInstanceProperty) GoString() string {
	return s.String()
}

func (s *ModuleDataComponentsValueInstanceProperty) GetCode() *string {
	return s.Code
}

func (s *ModuleDataComponentsValueInstanceProperty) GetName() *string {
	return s.Name
}

func (s *ModuleDataComponentsValueInstanceProperty) GetValue() *string {
	return s.Value
}

func (s *ModuleDataComponentsValueInstanceProperty) GetValues() []*ModuleDataComponentsValueInstancePropertyValues {
	return s.Values
}

func (s *ModuleDataComponentsValueInstanceProperty) SetCode(v string) *ModuleDataComponentsValueInstanceProperty {
	s.Code = &v
	return s
}

func (s *ModuleDataComponentsValueInstanceProperty) SetName(v string) *ModuleDataComponentsValueInstanceProperty {
	s.Name = &v
	return s
}

func (s *ModuleDataComponentsValueInstanceProperty) SetValue(v string) *ModuleDataComponentsValueInstanceProperty {
	s.Value = &v
	return s
}

func (s *ModuleDataComponentsValueInstanceProperty) SetValues(v []*ModuleDataComponentsValueInstancePropertyValues) *ModuleDataComponentsValueInstanceProperty {
	s.Values = v
	return s
}

func (s *ModuleDataComponentsValueInstanceProperty) Validate() error {
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

type ModuleDataComponentsValueInstancePropertyValues struct {
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

func (s ModuleDataComponentsValueInstancePropertyValues) String() string {
	return dara.Prettify(s)
}

func (s ModuleDataComponentsValueInstancePropertyValues) GoString() string {
	return s.String()
}

func (s *ModuleDataComponentsValueInstancePropertyValues) GetCode() *string {
	return s.Code
}

func (s *ModuleDataComponentsValueInstancePropertyValues) GetValue() *string {
	return s.Value
}

func (s *ModuleDataComponentsValueInstancePropertyValues) GetName() *string {
	return s.Name
}

func (s *ModuleDataComponentsValueInstancePropertyValues) SetCode(v string) *ModuleDataComponentsValueInstancePropertyValues {
	s.Code = &v
	return s
}

func (s *ModuleDataComponentsValueInstancePropertyValues) SetValue(v string) *ModuleDataComponentsValueInstancePropertyValues {
	s.Value = &v
	return s
}

func (s *ModuleDataComponentsValueInstancePropertyValues) SetName(v string) *ModuleDataComponentsValueInstancePropertyValues {
	s.Name = &v
	return s
}

func (s *ModuleDataComponentsValueInstancePropertyValues) Validate() error {
	return dara.Validate(s)
}
