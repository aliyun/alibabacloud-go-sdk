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
	// The build configuration of the template. The build process is specified in blocks by action. If this parameter is not specified, the service automatically infers the configuration based on the source image: public runtime images are used directly without building, and other images automatically execute copy and envdInject.
	BuildConfig *CreateTemplateBuildConfig `json:"buildConfig,omitempty" xml:"buildConfig,omitempty"`
	// The template name, which must be unique within the same team.
	//
	// example:
	//
	// python-data-analysis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The runtime configuration of the template. This configuration determines the specifications and environment of sandbox instances created from this template.
	RuntimeConfig *CreateTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
	// The ID of the team to which the template belongs, in UUID format. If this parameter is not specified, the default team of the current account is used.
	//
	// example:
	//
	// 5f4a2c18-****
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
