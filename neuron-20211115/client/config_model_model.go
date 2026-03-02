// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigModel interface {
	dara.Model
	String() string
	GoString() string
	SetAttribute(v string) *ConfigModel
	GetAttribute() *string
	SetType(v string) *ConfigModel
	GetType() *string
	SetValue(v string) *ConfigModel
	GetValue() *string
}

type ConfigModel struct {
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ConfigModel) String() string {
	return dara.Prettify(s)
}

func (s ConfigModel) GoString() string {
	return s.String()
}

func (s *ConfigModel) GetAttribute() *string {
	return s.Attribute
}

func (s *ConfigModel) GetType() *string {
	return s.Type
}

func (s *ConfigModel) GetValue() *string {
	return s.Value
}

func (s *ConfigModel) SetAttribute(v string) *ConfigModel {
	s.Attribute = &v
	return s
}

func (s *ConfigModel) SetType(v string) *ConfigModel {
	s.Type = &v
	return s
}

func (s *ConfigModel) SetValue(v string) *ConfigModel {
	s.Value = &v
	return s
}

func (s *ConfigModel) Validate() error {
	return dara.Validate(s)
}
