// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFieldInputConfig interface {
	dara.Model
	String() string
	GoString() string
	SetArrayed(v bool) *FieldInputConfig
	GetArrayed() *bool
	SetDefaultValue(v string) *FieldInputConfig
	GetDefaultValue() *string
	SetFieldCheckRegex(v string) *FieldInputConfig
	GetFieldCheckRegex() *string
	SetFieldClass(v string) *FieldInputConfig
	GetFieldClass() *string
	SetFieldConfigs(v []*FieldInputConfig) *FieldInputConfig
	GetFieldConfigs() []*FieldInputConfig
	SetFieldDescription(v string) *FieldInputConfig
	GetFieldDescription() *string
	SetFieldExample(v string) *FieldInputConfig
	GetFieldExample() *string
	SetFieldName(v string) *FieldInputConfig
	GetFieldName() *string
	SetFieldPath(v string) *FieldInputConfig
	GetFieldPath() *string
	SetFieldType(v string) *FieldInputConfig
	GetFieldType() *string
	SetRequired(v bool) *FieldInputConfig
	GetRequired() *bool
}

type FieldInputConfig struct {
	Arrayed          *bool               `json:"Arrayed,omitempty" xml:"Arrayed,omitempty"`
	DefaultValue     *string             `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	FieldCheckRegex  *string             `json:"FieldCheckRegex,omitempty" xml:"FieldCheckRegex,omitempty"`
	FieldClass       *string             `json:"FieldClass,omitempty" xml:"FieldClass,omitempty"`
	FieldConfigs     []*FieldInputConfig `json:"FieldConfigs,omitempty" xml:"FieldConfigs,omitempty" type:"Repeated"`
	FieldDescription *string             `json:"FieldDescription,omitempty" xml:"FieldDescription,omitempty"`
	FieldExample     *string             `json:"FieldExample,omitempty" xml:"FieldExample,omitempty"`
	FieldName        *string             `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FieldPath        *string             `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
	FieldType        *string             `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	Required         *bool               `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s FieldInputConfig) String() string {
	return dara.Prettify(s)
}

func (s FieldInputConfig) GoString() string {
	return s.String()
}

func (s *FieldInputConfig) GetArrayed() *bool {
	return s.Arrayed
}

func (s *FieldInputConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *FieldInputConfig) GetFieldCheckRegex() *string {
	return s.FieldCheckRegex
}

func (s *FieldInputConfig) GetFieldClass() *string {
	return s.FieldClass
}

func (s *FieldInputConfig) GetFieldConfigs() []*FieldInputConfig {
	return s.FieldConfigs
}

func (s *FieldInputConfig) GetFieldDescription() *string {
	return s.FieldDescription
}

func (s *FieldInputConfig) GetFieldExample() *string {
	return s.FieldExample
}

func (s *FieldInputConfig) GetFieldName() *string {
	return s.FieldName
}

func (s *FieldInputConfig) GetFieldPath() *string {
	return s.FieldPath
}

func (s *FieldInputConfig) GetFieldType() *string {
	return s.FieldType
}

func (s *FieldInputConfig) GetRequired() *bool {
	return s.Required
}

func (s *FieldInputConfig) SetArrayed(v bool) *FieldInputConfig {
	s.Arrayed = &v
	return s
}

func (s *FieldInputConfig) SetDefaultValue(v string) *FieldInputConfig {
	s.DefaultValue = &v
	return s
}

func (s *FieldInputConfig) SetFieldCheckRegex(v string) *FieldInputConfig {
	s.FieldCheckRegex = &v
	return s
}

func (s *FieldInputConfig) SetFieldClass(v string) *FieldInputConfig {
	s.FieldClass = &v
	return s
}

func (s *FieldInputConfig) SetFieldConfigs(v []*FieldInputConfig) *FieldInputConfig {
	s.FieldConfigs = v
	return s
}

func (s *FieldInputConfig) SetFieldDescription(v string) *FieldInputConfig {
	s.FieldDescription = &v
	return s
}

func (s *FieldInputConfig) SetFieldExample(v string) *FieldInputConfig {
	s.FieldExample = &v
	return s
}

func (s *FieldInputConfig) SetFieldName(v string) *FieldInputConfig {
	s.FieldName = &v
	return s
}

func (s *FieldInputConfig) SetFieldPath(v string) *FieldInputConfig {
	s.FieldPath = &v
	return s
}

func (s *FieldInputConfig) SetFieldType(v string) *FieldInputConfig {
	s.FieldType = &v
	return s
}

func (s *FieldInputConfig) SetRequired(v bool) *FieldInputConfig {
	s.Required = &v
	return s
}

func (s *FieldInputConfig) Validate() error {
	if s.FieldConfigs != nil {
		for _, item := range s.FieldConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
