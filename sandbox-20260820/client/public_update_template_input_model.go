// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateInput interface {
	dara.Model
	String() string
	GoString() string
	SetBuildConfig(v *PublicUpdateTemplateBuildConfig) *PublicUpdateTemplateInput
	GetBuildConfig() *PublicUpdateTemplateBuildConfig
	SetRuntimeConfig(v *PublicUpdateTemplateRuntimeConfig) *PublicUpdateTemplateInput
	GetRuntimeConfig() *PublicUpdateTemplateRuntimeConfig
}

type PublicUpdateTemplateInput struct {
	BuildConfig   *PublicUpdateTemplateBuildConfig   `json:"buildConfig,omitempty" xml:"buildConfig,omitempty"`
	RuntimeConfig *PublicUpdateTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
}

func (s PublicUpdateTemplateInput) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateInput) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateInput) GetBuildConfig() *PublicUpdateTemplateBuildConfig {
	return s.BuildConfig
}

func (s *PublicUpdateTemplateInput) GetRuntimeConfig() *PublicUpdateTemplateRuntimeConfig {
	return s.RuntimeConfig
}

func (s *PublicUpdateTemplateInput) SetBuildConfig(v *PublicUpdateTemplateBuildConfig) *PublicUpdateTemplateInput {
	s.BuildConfig = v
	return s
}

func (s *PublicUpdateTemplateInput) SetRuntimeConfig(v *PublicUpdateTemplateRuntimeConfig) *PublicUpdateTemplateInput {
	s.RuntimeConfig = v
	return s
}

func (s *PublicUpdateTemplateInput) Validate() error {
	if s.BuildConfig != nil {
		if err := s.BuildConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeConfig != nil {
		if err := s.RuntimeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
