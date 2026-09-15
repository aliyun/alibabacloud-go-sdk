// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateBuildConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCopy(v *PublicUpdateTemplateCopyAction) *PublicUpdateTemplateBuildConfig
	GetCopy() *PublicUpdateTemplateCopyAction
	SetEnvdInject(v *PublicUpdateTemplateEnvdInjectAction) *PublicUpdateTemplateBuildConfig
	GetEnvdInject() *PublicUpdateTemplateEnvdInjectAction
}

type PublicUpdateTemplateBuildConfig struct {
	Copy       *PublicUpdateTemplateCopyAction       `json:"copy,omitempty" xml:"copy,omitempty"`
	EnvdInject *PublicUpdateTemplateEnvdInjectAction `json:"envdInject,omitempty" xml:"envdInject,omitempty"`
}

func (s PublicUpdateTemplateBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateBuildConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateBuildConfig) GetCopy() *PublicUpdateTemplateCopyAction {
	return s.Copy
}

func (s *PublicUpdateTemplateBuildConfig) GetEnvdInject() *PublicUpdateTemplateEnvdInjectAction {
	return s.EnvdInject
}

func (s *PublicUpdateTemplateBuildConfig) SetCopy(v *PublicUpdateTemplateCopyAction) *PublicUpdateTemplateBuildConfig {
	s.Copy = v
	return s
}

func (s *PublicUpdateTemplateBuildConfig) SetEnvdInject(v *PublicUpdateTemplateEnvdInjectAction) *PublicUpdateTemplateBuildConfig {
	s.EnvdInject = v
	return s
}

func (s *PublicUpdateTemplateBuildConfig) Validate() error {
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
