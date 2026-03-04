// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleVersionsValueComponentsValue interface {
	dara.Model
	String() string
	GoString() string
	SetComponentCode(v string) *ModuleVersionsValueComponentsValue
	GetComponentCode() *string
	SetComponentName(v string) *ModuleVersionsValueComponentsValue
	GetComponentName() *string
	SetInstanceProperty(v []*ModuleVersionsValueComponentsValueInstanceProperty) *ModuleVersionsValueComponentsValue
	GetInstanceProperty() []*ModuleVersionsValueComponentsValueInstanceProperty
	SetProperties(v map[string]*ModuleVersionsValueComponentsValuePropertiesValue) *ModuleVersionsValueComponentsValue
	GetProperties() map[string]*ModuleVersionsValueComponentsValuePropertiesValue
	SetModuleAttrStatus(v int32) *ModuleVersionsValueComponentsValue
	GetModuleAttrStatus() *int32
}

type ModuleVersionsValueComponentsValue struct {
	// example:
	//
	// ComponentCode
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// example:
	//
	// ComponentName
	ComponentName    *string                                                       `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	InstanceProperty []*ModuleVersionsValueComponentsValueInstanceProperty         `json:"InstanceProperty,omitempty" xml:"InstanceProperty,omitempty" type:"Repeated"`
	Properties       map[string]*ModuleVersionsValueComponentsValuePropertiesValue `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// attr
	ModuleAttrStatus *int32 `json:"ModuleAttrStatus,omitempty" xml:"ModuleAttrStatus,omitempty"`
}

func (s ModuleVersionsValueComponentsValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleVersionsValueComponentsValue) GoString() string {
	return s.String()
}

func (s *ModuleVersionsValueComponentsValue) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *ModuleVersionsValueComponentsValue) GetComponentName() *string {
	return s.ComponentName
}

func (s *ModuleVersionsValueComponentsValue) GetInstanceProperty() []*ModuleVersionsValueComponentsValueInstanceProperty {
	return s.InstanceProperty
}

func (s *ModuleVersionsValueComponentsValue) GetProperties() map[string]*ModuleVersionsValueComponentsValuePropertiesValue {
	return s.Properties
}

func (s *ModuleVersionsValueComponentsValue) GetModuleAttrStatus() *int32 {
	return s.ModuleAttrStatus
}

func (s *ModuleVersionsValueComponentsValue) SetComponentCode(v string) *ModuleVersionsValueComponentsValue {
	s.ComponentCode = &v
	return s
}

func (s *ModuleVersionsValueComponentsValue) SetComponentName(v string) *ModuleVersionsValueComponentsValue {
	s.ComponentName = &v
	return s
}

func (s *ModuleVersionsValueComponentsValue) SetInstanceProperty(v []*ModuleVersionsValueComponentsValueInstanceProperty) *ModuleVersionsValueComponentsValue {
	s.InstanceProperty = v
	return s
}

func (s *ModuleVersionsValueComponentsValue) SetProperties(v map[string]*ModuleVersionsValueComponentsValuePropertiesValue) *ModuleVersionsValueComponentsValue {
	s.Properties = v
	return s
}

func (s *ModuleVersionsValueComponentsValue) SetModuleAttrStatus(v int32) *ModuleVersionsValueComponentsValue {
	s.ModuleAttrStatus = &v
	return s
}

func (s *ModuleVersionsValueComponentsValue) Validate() error {
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

type ModuleVersionsValueComponentsValueInstanceProperty struct {
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
	Value  *string                                                     `json:"Value,omitempty" xml:"Value,omitempty"`
	Values []*ModuleVersionsValueComponentsValueInstancePropertyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ModuleVersionsValueComponentsValueInstanceProperty) String() string {
	return dara.Prettify(s)
}

func (s ModuleVersionsValueComponentsValueInstanceProperty) GoString() string {
	return s.String()
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) GetCode() *string {
	return s.Code
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) GetName() *string {
	return s.Name
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) GetValue() *string {
	return s.Value
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) GetValues() []*ModuleVersionsValueComponentsValueInstancePropertyValues {
	return s.Values
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) SetCode(v string) *ModuleVersionsValueComponentsValueInstanceProperty {
	s.Code = &v
	return s
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) SetName(v string) *ModuleVersionsValueComponentsValueInstanceProperty {
	s.Name = &v
	return s
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) SetValue(v string) *ModuleVersionsValueComponentsValueInstanceProperty {
	s.Value = &v
	return s
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) SetValues(v []*ModuleVersionsValueComponentsValueInstancePropertyValues) *ModuleVersionsValueComponentsValueInstanceProperty {
	s.Values = v
	return s
}

func (s *ModuleVersionsValueComponentsValueInstanceProperty) Validate() error {
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

type ModuleVersionsValueComponentsValueInstancePropertyValues struct {
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

func (s ModuleVersionsValueComponentsValueInstancePropertyValues) String() string {
	return dara.Prettify(s)
}

func (s ModuleVersionsValueComponentsValueInstancePropertyValues) GoString() string {
	return s.String()
}

func (s *ModuleVersionsValueComponentsValueInstancePropertyValues) GetCode() *string {
	return s.Code
}

func (s *ModuleVersionsValueComponentsValueInstancePropertyValues) GetValue() *string {
	return s.Value
}

func (s *ModuleVersionsValueComponentsValueInstancePropertyValues) GetName() *string {
	return s.Name
}

func (s *ModuleVersionsValueComponentsValueInstancePropertyValues) SetCode(v string) *ModuleVersionsValueComponentsValueInstancePropertyValues {
	s.Code = &v
	return s
}

func (s *ModuleVersionsValueComponentsValueInstancePropertyValues) SetValue(v string) *ModuleVersionsValueComponentsValueInstancePropertyValues {
	s.Value = &v
	return s
}

func (s *ModuleVersionsValueComponentsValueInstancePropertyValues) SetName(v string) *ModuleVersionsValueComponentsValueInstancePropertyValues {
	s.Name = &v
	return s
}

func (s *ModuleVersionsValueComponentsValueInstancePropertyValues) Validate() error {
	return dara.Validate(s)
}
