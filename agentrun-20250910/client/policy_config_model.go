// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolicyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAiFallbackConfig(v *AiFallbackConfig) *PolicyConfig
	GetAiFallbackConfig() *AiFallbackConfig
	SetEnable(v bool) *PolicyConfig
	GetEnable() *bool
	SetType(v string) *PolicyConfig
	GetType() *string
}

type PolicyConfig struct {
	AiFallbackConfig *AiFallbackConfig `json:"aiFallbackConfig,omitempty" xml:"aiFallbackConfig,omitempty"`
	Enable           *bool             `json:"enable,omitempty" xml:"enable,omitempty"`
	Type             *string           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PolicyConfig) String() string {
	return dara.Prettify(s)
}

func (s PolicyConfig) GoString() string {
	return s.String()
}

func (s *PolicyConfig) GetAiFallbackConfig() *AiFallbackConfig {
	return s.AiFallbackConfig
}

func (s *PolicyConfig) GetEnable() *bool {
	return s.Enable
}

func (s *PolicyConfig) GetType() *string {
	return s.Type
}

func (s *PolicyConfig) SetAiFallbackConfig(v *AiFallbackConfig) *PolicyConfig {
	s.AiFallbackConfig = v
	return s
}

func (s *PolicyConfig) SetEnable(v bool) *PolicyConfig {
	s.Enable = &v
	return s
}

func (s *PolicyConfig) SetType(v string) *PolicyConfig {
	s.Type = &v
	return s
}

func (s *PolicyConfig) Validate() error {
	return dara.Validate(s)
}
