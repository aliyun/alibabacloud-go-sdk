// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProperty interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v string) *Property
	GetDefaultValue() *string
	SetDefinesFormat(v bool) *Property
	GetDefinesFormat() *bool
	SetDescription(v string) *Property
	GetDescription() *string
	SetKey(v string) *Property
	GetKey() *string
	SetRequired(v bool) *Property
	GetRequired() *bool
	SetSensitive(v bool) *Property
	GetSensitive() *bool
}

type Property struct {
	DefaultValue  *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	DefinesFormat *bool   `json:"definesFormat,omitempty" xml:"definesFormat,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	Key           *string `json:"key,omitempty" xml:"key,omitempty"`
	Required      *bool   `json:"required,omitempty" xml:"required,omitempty"`
	Sensitive     *bool   `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
}

func (s Property) String() string {
	return dara.Prettify(s)
}

func (s Property) GoString() string {
	return s.String()
}

func (s *Property) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *Property) GetDefinesFormat() *bool {
	return s.DefinesFormat
}

func (s *Property) GetDescription() *string {
	return s.Description
}

func (s *Property) GetKey() *string {
	return s.Key
}

func (s *Property) GetRequired() *bool {
	return s.Required
}

func (s *Property) GetSensitive() *bool {
	return s.Sensitive
}

func (s *Property) SetDefaultValue(v string) *Property {
	s.DefaultValue = &v
	return s
}

func (s *Property) SetDefinesFormat(v bool) *Property {
	s.DefinesFormat = &v
	return s
}

func (s *Property) SetDescription(v string) *Property {
	s.Description = &v
	return s
}

func (s *Property) SetKey(v string) *Property {
	s.Key = &v
	return s
}

func (s *Property) SetRequired(v bool) *Property {
	s.Required = &v
	return s
}

func (s *Property) SetSensitive(v bool) *Property {
	s.Sensitive = &v
	return s
}

func (s *Property) Validate() error {
	return dara.Validate(s)
}
