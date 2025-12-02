// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFieldOutputConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v string) *FieldOutputConfig
	GetDefaultValue() *string
	SetFieldDescription(v string) *FieldOutputConfig
	GetFieldDescription() *string
	SetFieldExample(v string) *FieldOutputConfig
	GetFieldExample() *string
	SetFieldName(v string) *FieldOutputConfig
	GetFieldName() *string
	SetFieldType(v string) *FieldOutputConfig
	GetFieldType() *string
}

type FieldOutputConfig struct {
	DefaultValue     *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	FieldDescription *string `json:"FieldDescription,omitempty" xml:"FieldDescription,omitempty"`
	FieldExample     *string `json:"FieldExample,omitempty" xml:"FieldExample,omitempty"`
	FieldName        *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FieldType        *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
}

func (s FieldOutputConfig) String() string {
	return dara.Prettify(s)
}

func (s FieldOutputConfig) GoString() string {
	return s.String()
}

func (s *FieldOutputConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *FieldOutputConfig) GetFieldDescription() *string {
	return s.FieldDescription
}

func (s *FieldOutputConfig) GetFieldExample() *string {
	return s.FieldExample
}

func (s *FieldOutputConfig) GetFieldName() *string {
	return s.FieldName
}

func (s *FieldOutputConfig) GetFieldType() *string {
	return s.FieldType
}

func (s *FieldOutputConfig) SetDefaultValue(v string) *FieldOutputConfig {
	s.DefaultValue = &v
	return s
}

func (s *FieldOutputConfig) SetFieldDescription(v string) *FieldOutputConfig {
	s.FieldDescription = &v
	return s
}

func (s *FieldOutputConfig) SetFieldExample(v string) *FieldOutputConfig {
	s.FieldExample = &v
	return s
}

func (s *FieldOutputConfig) SetFieldName(v string) *FieldOutputConfig {
	s.FieldName = &v
	return s
}

func (s *FieldOutputConfig) SetFieldType(v string) *FieldOutputConfig {
	s.FieldType = &v
	return s
}

func (s *FieldOutputConfig) Validate() error {
	return dara.Validate(s)
}
