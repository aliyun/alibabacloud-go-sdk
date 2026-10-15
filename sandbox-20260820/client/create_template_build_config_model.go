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
	SetInstallStorageClients(v bool) *CreateTemplateBuildConfig
	GetInstallStorageClients() *bool
}

type CreateTemplateBuildConfig struct {
	// The action configuration for pulling the source image and pushing it to the destination repository.
	Copy *CreateTemplateCopyAction `json:"copy,omitempty" xml:"copy,omitempty"`
	// The action configuration for packaging the envd daemon process into the image. This is an optional intermediate step in the copy process and is supported only by first-generation sandboxes.
	EnvdInject *CreateTemplateEnvdInjectAction `json:"envdInject,omitempty" xml:"envdInject,omitempty"`
	// The installation of clients required for storage mounting.
	InstallStorageClients *bool `json:"installStorageClients,omitempty" xml:"installStorageClients,omitempty"`
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

func (s *CreateTemplateBuildConfig) GetInstallStorageClients() *bool {
	return s.InstallStorageClients
}

func (s *CreateTemplateBuildConfig) SetCopy(v *CreateTemplateCopyAction) *CreateTemplateBuildConfig {
	s.Copy = v
	return s
}

func (s *CreateTemplateBuildConfig) SetEnvdInject(v *CreateTemplateEnvdInjectAction) *CreateTemplateBuildConfig {
	s.EnvdInject = v
	return s
}

func (s *CreateTemplateBuildConfig) SetInstallStorageClients(v bool) *CreateTemplateBuildConfig {
	s.InstallStorageClients = &v
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
