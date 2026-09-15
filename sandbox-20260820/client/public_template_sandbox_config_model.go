// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateSandboxConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *PublicTemplateSandboxConfig
	GetAcrInstanceId() *string
	SetGeneration(v int32) *PublicTemplateSandboxConfig
	GetGeneration() *int32
	SetImage(v string) *PublicTemplateSandboxConfig
	GetImage() *string
	SetOsType(v string) *PublicTemplateSandboxConfig
	GetOsType() *string
	SetReadyCommand(v string) *PublicTemplateSandboxConfig
	GetReadyCommand() *string
	SetRegistryConfig(v *PublicTemplateRegistryConfig) *PublicTemplateSandboxConfig
	GetRegistryConfig() *PublicTemplateRegistryConfig
	SetStartCommand(v string) *PublicTemplateSandboxConfig
	GetStartCommand() *string
}

type PublicTemplateSandboxConfig struct {
	AcrInstanceId  *string                       `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	Generation     *int32                        `json:"generation,omitempty" xml:"generation,omitempty"`
	Image          *string                       `json:"image,omitempty" xml:"image,omitempty"`
	OsType         *string                       `json:"osType,omitempty" xml:"osType,omitempty"`
	ReadyCommand   *string                       `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	RegistryConfig *PublicTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	StartCommand   *string                       `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
}

func (s PublicTemplateSandboxConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateSandboxConfig) GoString() string {
	return s.String()
}

func (s *PublicTemplateSandboxConfig) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *PublicTemplateSandboxConfig) GetGeneration() *int32 {
	return s.Generation
}

func (s *PublicTemplateSandboxConfig) GetImage() *string {
	return s.Image
}

func (s *PublicTemplateSandboxConfig) GetOsType() *string {
	return s.OsType
}

func (s *PublicTemplateSandboxConfig) GetReadyCommand() *string {
	return s.ReadyCommand
}

func (s *PublicTemplateSandboxConfig) GetRegistryConfig() *PublicTemplateRegistryConfig {
	return s.RegistryConfig
}

func (s *PublicTemplateSandboxConfig) GetStartCommand() *string {
	return s.StartCommand
}

func (s *PublicTemplateSandboxConfig) SetAcrInstanceId(v string) *PublicTemplateSandboxConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *PublicTemplateSandboxConfig) SetGeneration(v int32) *PublicTemplateSandboxConfig {
	s.Generation = &v
	return s
}

func (s *PublicTemplateSandboxConfig) SetImage(v string) *PublicTemplateSandboxConfig {
	s.Image = &v
	return s
}

func (s *PublicTemplateSandboxConfig) SetOsType(v string) *PublicTemplateSandboxConfig {
	s.OsType = &v
	return s
}

func (s *PublicTemplateSandboxConfig) SetReadyCommand(v string) *PublicTemplateSandboxConfig {
	s.ReadyCommand = &v
	return s
}

func (s *PublicTemplateSandboxConfig) SetRegistryConfig(v *PublicTemplateRegistryConfig) *PublicTemplateSandboxConfig {
	s.RegistryConfig = v
	return s
}

func (s *PublicTemplateSandboxConfig) SetStartCommand(v string) *PublicTemplateSandboxConfig {
	s.StartCommand = &v
	return s
}

func (s *PublicTemplateSandboxConfig) Validate() error {
	if s.RegistryConfig != nil {
		if err := s.RegistryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
