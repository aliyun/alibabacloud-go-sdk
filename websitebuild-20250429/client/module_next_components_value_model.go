// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleNextComponentsValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleNextComponentsValue
	GetCode() *string
	SetName(v string) *ModuleNextComponentsValue
	GetName() *string
	SetInstanceProperty(v []*ModuleNextComponentsValueInstanceProperty) *ModuleNextComponentsValue
	GetInstanceProperty() []*ModuleNextComponentsValueInstanceProperty
	SetProperties(v map[string]*ModuleNextComponentsValuePropertiesValue) *ModuleNextComponentsValue
	GetProperties() map[string]*ModuleNextComponentsValuePropertiesValue
	SetModuleAttrStatus(v int32) *ModuleNextComponentsValue
	GetModuleAttrStatus() *int32
}

type ModuleNextComponentsValue struct {
	// example:
	//
	// placeholder
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// placeholder
	Name             *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	InstanceProperty []*ModuleNextComponentsValueInstanceProperty         `json:"InstanceProperty,omitempty" xml:"InstanceProperty,omitempty" type:"Repeated"`
	Properties       map[string]*ModuleNextComponentsValuePropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// placeholder
	ModuleAttrStatus *int32 `json:"ModuleAttrStatus,omitempty" xml:"ModuleAttrStatus,omitempty"`
}

func (s ModuleNextComponentsValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleNextComponentsValue) GoString() string {
	return s.String()
}

func (s *ModuleNextComponentsValue) GetCode() *string {
	return s.Code
}

func (s *ModuleNextComponentsValue) GetName() *string {
	return s.Name
}

func (s *ModuleNextComponentsValue) GetInstanceProperty() []*ModuleNextComponentsValueInstanceProperty {
	return s.InstanceProperty
}

func (s *ModuleNextComponentsValue) GetProperties() map[string]*ModuleNextComponentsValuePropertiesValue {
	return s.Properties
}

func (s *ModuleNextComponentsValue) GetModuleAttrStatus() *int32 {
	return s.ModuleAttrStatus
}

func (s *ModuleNextComponentsValue) SetCode(v string) *ModuleNextComponentsValue {
	s.Code = &v
	return s
}

func (s *ModuleNextComponentsValue) SetName(v string) *ModuleNextComponentsValue {
	s.Name = &v
	return s
}

func (s *ModuleNextComponentsValue) SetInstanceProperty(v []*ModuleNextComponentsValueInstanceProperty) *ModuleNextComponentsValue {
	s.InstanceProperty = v
	return s
}

func (s *ModuleNextComponentsValue) SetProperties(v map[string]*ModuleNextComponentsValuePropertiesValue) *ModuleNextComponentsValue {
	s.Properties = v
	return s
}

func (s *ModuleNextComponentsValue) SetModuleAttrStatus(v int32) *ModuleNextComponentsValue {
	s.ModuleAttrStatus = &v
	return s
}

func (s *ModuleNextComponentsValue) Validate() error {
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

type ModuleNextComponentsValueInstanceProperty struct {
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
	Values []*ModuleNextComponentsValueInstancePropertyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ModuleNextComponentsValueInstanceProperty) String() string {
	return dara.Prettify(s)
}

func (s ModuleNextComponentsValueInstanceProperty) GoString() string {
	return s.String()
}

func (s *ModuleNextComponentsValueInstanceProperty) GetCode() *string {
	return s.Code
}

func (s *ModuleNextComponentsValueInstanceProperty) GetName() *string {
	return s.Name
}

func (s *ModuleNextComponentsValueInstanceProperty) GetValue() *string {
	return s.Value
}

func (s *ModuleNextComponentsValueInstanceProperty) GetValues() []*ModuleNextComponentsValueInstancePropertyValues {
	return s.Values
}

func (s *ModuleNextComponentsValueInstanceProperty) SetCode(v string) *ModuleNextComponentsValueInstanceProperty {
	s.Code = &v
	return s
}

func (s *ModuleNextComponentsValueInstanceProperty) SetName(v string) *ModuleNextComponentsValueInstanceProperty {
	s.Name = &v
	return s
}

func (s *ModuleNextComponentsValueInstanceProperty) SetValue(v string) *ModuleNextComponentsValueInstanceProperty {
	s.Value = &v
	return s
}

func (s *ModuleNextComponentsValueInstanceProperty) SetValues(v []*ModuleNextComponentsValueInstancePropertyValues) *ModuleNextComponentsValueInstanceProperty {
	s.Values = v
	return s
}

func (s *ModuleNextComponentsValueInstanceProperty) Validate() error {
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

type ModuleNextComponentsValueInstancePropertyValues struct {
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

func (s ModuleNextComponentsValueInstancePropertyValues) String() string {
	return dara.Prettify(s)
}

func (s ModuleNextComponentsValueInstancePropertyValues) GoString() string {
	return s.String()
}

func (s *ModuleNextComponentsValueInstancePropertyValues) GetCode() *string {
	return s.Code
}

func (s *ModuleNextComponentsValueInstancePropertyValues) GetValue() *string {
	return s.Value
}

func (s *ModuleNextComponentsValueInstancePropertyValues) GetName() *string {
	return s.Name
}

func (s *ModuleNextComponentsValueInstancePropertyValues) SetCode(v string) *ModuleNextComponentsValueInstancePropertyValues {
	s.Code = &v
	return s
}

func (s *ModuleNextComponentsValueInstancePropertyValues) SetValue(v string) *ModuleNextComponentsValueInstancePropertyValues {
	s.Value = &v
	return s
}

func (s *ModuleNextComponentsValueInstancePropertyValues) SetName(v string) *ModuleNextComponentsValueInstancePropertyValues {
	s.Name = &v
	return s
}

func (s *ModuleNextComponentsValueInstancePropertyValues) Validate() error {
	return dara.Validate(s)
}
