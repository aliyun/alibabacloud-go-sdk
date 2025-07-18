// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpCustomField interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *IdpCustomField
	GetDescription() *string
	SetKey(v string) *IdpCustomField
	GetKey() *string
	SetName(v string) *IdpCustomField
	GetName() *string
	SetType(v string) *IdpCustomField
	GetType() *string
	SetValue(v string) *IdpCustomField
	GetValue() *string
}

type IdpCustomField struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s IdpCustomField) String() string {
	return dara.Prettify(s)
}

func (s IdpCustomField) GoString() string {
	return s.String()
}

func (s *IdpCustomField) GetDescription() *string {
	return s.Description
}

func (s *IdpCustomField) GetKey() *string {
	return s.Key
}

func (s *IdpCustomField) GetName() *string {
	return s.Name
}

func (s *IdpCustomField) GetType() *string {
	return s.Type
}

func (s *IdpCustomField) GetValue() *string {
	return s.Value
}

func (s *IdpCustomField) SetDescription(v string) *IdpCustomField {
	s.Description = &v
	return s
}

func (s *IdpCustomField) SetKey(v string) *IdpCustomField {
	s.Key = &v
	return s
}

func (s *IdpCustomField) SetName(v string) *IdpCustomField {
	s.Name = &v
	return s
}

func (s *IdpCustomField) SetType(v string) *IdpCustomField {
	s.Type = &v
	return s
}

func (s *IdpCustomField) SetValue(v string) *IdpCustomField {
	s.Value = &v
	return s
}

func (s *IdpCustomField) Validate() error {
	return dara.Validate(s)
}
