// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateBuildConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCopy(v *CreateTemplateCopyAction) *CreateTemplateBuildConfig
	GetCopy() *CreateTemplateCopyAction
	SetEnvdInject(v *CreateTemplateEnvdInjectAction) *CreateTemplateBuildConfig
	GetEnvdInject() *CreateTemplateEnvdInjectAction
}

type CreateTemplateBuildConfig struct {
	// The image copy build action.
	Copy *CreateTemplateCopyAction `json:"copy,omitempty" xml:"copy,omitempty"`
	// The envd injection build action.
	EnvdInject *CreateTemplateEnvdInjectAction `json:"envdInject,omitempty" xml:"envdInject,omitempty"`
}

func (s CreateTemplateBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateBuildConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateBuildConfig) GetCopy() *CreateTemplateCopyAction {
	return s.Copy
}

func (s *CreateTemplateBuildConfig) GetEnvdInject() *CreateTemplateEnvdInjectAction {
	return s.EnvdInject
}

func (s *CreateTemplateBuildConfig) SetCopy(v *CreateTemplateCopyAction) *CreateTemplateBuildConfig {
	s.Copy = v
	return s
}

func (s *CreateTemplateBuildConfig) SetEnvdInject(v *CreateTemplateEnvdInjectAction) *CreateTemplateBuildConfig {
	s.EnvdInject = v
	return s
}

func (s *CreateTemplateBuildConfig) Validate() error {
	if s.Copy != nil {
		if err := s.Copy.Validate(); err != nil {
			return err
		}
	}
	if s.EnvdInject != nil {
		if err := s.EnvdInject.Validate(); err != nil {
			return err
		}
	}
	return nil
}
