// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateSandboxConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *PublicUpdateTemplateSandboxConfig
	GetAcrInstanceId() *string
	SetGeneration(v int32) *PublicUpdateTemplateSandboxConfig
	GetGeneration() *int32
	SetImage(v string) *PublicUpdateTemplateSandboxConfig
	GetImage() *string
	SetOsType(v string) *PublicUpdateTemplateSandboxConfig
	GetOsType() *string
	SetReadyCommand(v string) *PublicUpdateTemplateSandboxConfig
	GetReadyCommand() *string
	SetRegistryConfig(v *PublicUpdateTemplateRegistryConfig) *PublicUpdateTemplateSandboxConfig
	GetRegistryConfig() *PublicUpdateTemplateRegistryConfig
	SetRegistryType(v string) *PublicUpdateTemplateSandboxConfig
	GetRegistryType() *string
	SetStartCommand(v string) *PublicUpdateTemplateSandboxConfig
	GetStartCommand() *string
}

type PublicUpdateTemplateSandboxConfig struct {
	AcrInstanceId  *string                             `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	Generation     *int32                              `json:"generation,omitempty" xml:"generation,omitempty"`
	Image          *string                             `json:"image,omitempty" xml:"image,omitempty"`
	OsType         *string                             `json:"osType,omitempty" xml:"osType,omitempty"`
	ReadyCommand   *string                             `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	RegistryConfig *PublicUpdateTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	RegistryType   *string                             `json:"registryType,omitempty" xml:"registryType,omitempty"`
	StartCommand   *string                             `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
}

func (s PublicUpdateTemplateSandboxConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateSandboxConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateSandboxConfig) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *PublicUpdateTemplateSandboxConfig) GetGeneration() *int32 {
	return s.Generation
}

func (s *PublicUpdateTemplateSandboxConfig) GetImage() *string {
	return s.Image
}

func (s *PublicUpdateTemplateSandboxConfig) GetOsType() *string {
	return s.OsType
}

func (s *PublicUpdateTemplateSandboxConfig) GetReadyCommand() *string {
	return s.ReadyCommand
}

func (s *PublicUpdateTemplateSandboxConfig) GetRegistryConfig() *PublicUpdateTemplateRegistryConfig {
	return s.RegistryConfig
}

func (s *PublicUpdateTemplateSandboxConfig) GetRegistryType() *string {
	return s.RegistryType
}

func (s *PublicUpdateTemplateSandboxConfig) GetStartCommand() *string {
	return s.StartCommand
}

func (s *PublicUpdateTemplateSandboxConfig) SetAcrInstanceId(v string) *PublicUpdateTemplateSandboxConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) SetGeneration(v int32) *PublicUpdateTemplateSandboxConfig {
	s.Generation = &v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) SetImage(v string) *PublicUpdateTemplateSandboxConfig {
	s.Image = &v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) SetOsType(v string) *PublicUpdateTemplateSandboxConfig {
	s.OsType = &v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) SetReadyCommand(v string) *PublicUpdateTemplateSandboxConfig {
	s.ReadyCommand = &v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) SetRegistryConfig(v *PublicUpdateTemplateRegistryConfig) *PublicUpdateTemplateSandboxConfig {
	s.RegistryConfig = v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) SetRegistryType(v string) *PublicUpdateTemplateSandboxConfig {
	s.RegistryType = &v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) SetStartCommand(v string) *PublicUpdateTemplateSandboxConfig {
	s.StartCommand = &v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) Validate() error {
	if s.RegistryConfig != nil {
		if err := s.RegistryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
