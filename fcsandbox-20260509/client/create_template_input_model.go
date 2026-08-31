// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateInput interface {
	dara.Model
	String() string
	GoString() string
	SetBuildConfig(v *CreateTemplateBuildConfig) *CreateTemplateInput
	GetBuildConfig() *CreateTemplateBuildConfig
	SetName(v string) *CreateTemplateInput
	GetName() *string
	SetRuntimeConfig(v *CreateTemplateRuntimeConfig) *CreateTemplateInput
	GetRuntimeConfig() *CreateTemplateRuntimeConfig
	SetTeamID(v string) *CreateTemplateInput
	GetTeamID() *string
}

type CreateTemplateInput struct {
	// The template build configuration.
	BuildConfig *CreateTemplateBuildConfig `json:"buildConfig,omitempty" xml:"buildConfig,omitempty"`
	// The template name.
	//
	// example:
	//
	// my-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The template runtime configuration.
	RuntimeConfig *CreateTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
	// The unique identifier of the team.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s CreateTemplateInput) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateInput) GoString() string {
	return s.String()
}

func (s *CreateTemplateInput) GetBuildConfig() *CreateTemplateBuildConfig {
	return s.BuildConfig
}

func (s *CreateTemplateInput) GetName() *string {
	return s.Name
}

func (s *CreateTemplateInput) GetRuntimeConfig() *CreateTemplateRuntimeConfig {
	return s.RuntimeConfig
}

func (s *CreateTemplateInput) GetTeamID() *string {
	return s.TeamID
}

func (s *CreateTemplateInput) SetBuildConfig(v *CreateTemplateBuildConfig) *CreateTemplateInput {
	s.BuildConfig = v
	return s
}

func (s *CreateTemplateInput) SetName(v string) *CreateTemplateInput {
	s.Name = &v
	return s
}

func (s *CreateTemplateInput) SetRuntimeConfig(v *CreateTemplateRuntimeConfig) *CreateTemplateInput {
	s.RuntimeConfig = v
	return s
}

func (s *CreateTemplateInput) SetTeamID(v string) *CreateTemplateInput {
	s.TeamID = &v
	return s
}

func (s *CreateTemplateInput) Validate() error {
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
