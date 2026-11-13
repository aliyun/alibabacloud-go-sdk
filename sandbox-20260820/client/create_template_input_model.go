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
	SetResourceTags(v []*CreateTemplateInputResourceTags) *CreateTemplateInput
	GetResourceTags() []*CreateTemplateInputResourceTags
	SetRuntimeConfig(v *CreateTemplateRuntimeConfig) *CreateTemplateInput
	GetRuntimeConfig() *CreateTemplateRuntimeConfig
	SetTeamID(v string) *CreateTemplateInput
	GetTeamID() *string
	SetTemplateType(v string) *CreateTemplateInput
	GetTemplateType() *string
}

type CreateTemplateInput struct {
	// The build configuration of the template. This configuration specifies the build process in action-based chunks. If this parameter is not specified, the service automatically derives the configuration based on the source image. Public runtime images are used directly without building. For other images, the copy and envdInject actions are automatically executed.
	BuildConfig *CreateTemplateBuildConfig `json:"buildConfig,omitempty" xml:"buildConfig,omitempty"`
	// The name of the template. The name must be unique within the same team.
	//
	// example:
	//
	// python-data-analysis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource tags.
	ResourceTags []*CreateTemplateInputResourceTags `json:"resourceTags" xml:"resourceTags" type:"Repeated"`
	// The runtime configuration of the template. This configuration determines the specifications and environment of the sandbox instances created from the template.
	RuntimeConfig *CreateTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
	// The ID of the team to which the template belongs. The ID must be in UUID format. If this parameter is not specified, the default team of the current account is used.
	//
	// example:
	//
	// 5f4a2c18-****
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The templatetype. Set this parameter to layer to build a layer template.
	//
	// example:
	//
	// layer
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
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

func (s *CreateTemplateInput) GetResourceTags() []*CreateTemplateInputResourceTags {
	return s.ResourceTags
}

func (s *CreateTemplateInput) GetRuntimeConfig() *CreateTemplateRuntimeConfig {
	return s.RuntimeConfig
}

func (s *CreateTemplateInput) GetTeamID() *string {
	return s.TeamID
}

func (s *CreateTemplateInput) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateTemplateInput) SetBuildConfig(v *CreateTemplateBuildConfig) *CreateTemplateInput {
	s.BuildConfig = v
	return s
}

func (s *CreateTemplateInput) SetName(v string) *CreateTemplateInput {
	s.Name = &v
	return s
}

func (s *CreateTemplateInput) SetResourceTags(v []*CreateTemplateInputResourceTags) *CreateTemplateInput {
	s.ResourceTags = v
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

func (s *CreateTemplateInput) SetTemplateType(v string) *CreateTemplateInput {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateInput) Validate() error {
	if s.BuildConfig != nil {
		if err := s.BuildConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceTags != nil {
		for _, item := range s.ResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuntimeConfig != nil {
		if err := s.RuntimeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTemplateInputResourceTags struct {
	// The key of the tag.
	//
	// example:
	//
	// project
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// demo
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateTemplateInputResourceTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateInputResourceTags) GoString() string {
	return s.String()
}

func (s *CreateTemplateInputResourceTags) GetKey() *string {
	return s.Key
}

func (s *CreateTemplateInputResourceTags) GetValue() *string {
	return s.Value
}

func (s *CreateTemplateInputResourceTags) SetKey(v string) *CreateTemplateInputResourceTags {
	s.Key = &v
	return s
}

func (s *CreateTemplateInputResourceTags) SetValue(v string) *CreateTemplateInputResourceTags {
	s.Value = &v
	return s
}

func (s *CreateTemplateInputResourceTags) Validate() error {
	return dara.Validate(s)
}
