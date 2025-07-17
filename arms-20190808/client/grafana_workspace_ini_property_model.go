// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceIniProperty interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v string) *GrafanaWorkspaceIniProperty
	GetDefaultValue() *string
	SetDescription(v string) *GrafanaWorkspaceIniProperty
	GetDescription() *string
	SetExample(v string) *GrafanaWorkspaceIniProperty
	GetExample() *string
	SetKey(v string) *GrafanaWorkspaceIniProperty
	GetKey() *string
	SetSecret(v bool) *GrafanaWorkspaceIniProperty
	GetSecret() *bool
	SetValue(v string) *GrafanaWorkspaceIniProperty
	GetValue() *string
}

type GrafanaWorkspaceIniProperty struct {
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	Example      *string `json:"example,omitempty" xml:"example,omitempty"`
	Key          *string `json:"key,omitempty" xml:"key,omitempty"`
	Secret       *bool   `json:"secret,omitempty" xml:"secret,omitempty"`
	Value        *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GrafanaWorkspaceIniProperty) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceIniProperty) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceIniProperty) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GrafanaWorkspaceIniProperty) GetDescription() *string {
	return s.Description
}

func (s *GrafanaWorkspaceIniProperty) GetExample() *string {
	return s.Example
}

func (s *GrafanaWorkspaceIniProperty) GetKey() *string {
	return s.Key
}

func (s *GrafanaWorkspaceIniProperty) GetSecret() *bool {
	return s.Secret
}

func (s *GrafanaWorkspaceIniProperty) GetValue() *string {
	return s.Value
}

func (s *GrafanaWorkspaceIniProperty) SetDefaultValue(v string) *GrafanaWorkspaceIniProperty {
	s.DefaultValue = &v
	return s
}

func (s *GrafanaWorkspaceIniProperty) SetDescription(v string) *GrafanaWorkspaceIniProperty {
	s.Description = &v
	return s
}

func (s *GrafanaWorkspaceIniProperty) SetExample(v string) *GrafanaWorkspaceIniProperty {
	s.Example = &v
	return s
}

func (s *GrafanaWorkspaceIniProperty) SetKey(v string) *GrafanaWorkspaceIniProperty {
	s.Key = &v
	return s
}

func (s *GrafanaWorkspaceIniProperty) SetSecret(v bool) *GrafanaWorkspaceIniProperty {
	s.Secret = &v
	return s
}

func (s *GrafanaWorkspaceIniProperty) SetValue(v string) *GrafanaWorkspaceIniProperty {
	s.Value = &v
	return s
}

func (s *GrafanaWorkspaceIniProperty) Validate() error {
	return dara.Validate(s)
}
