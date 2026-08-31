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
	// The ACR Enterprise instance ID.
	//
	// example:
	//
	// cri-abcd1234efgh
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The sandbox generation. A value of 1 indicates the first-generation sandbox. A value of 2 indicates the second-generation sandbox.
	//
	// example:
	//
	// 2
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The image address.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/demo/nginx:latest
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// linux
	OsType *string `json:"osType,omitempty" xml:"osType,omitempty"`
	// The sandbox readiness probe command.
	//
	// example:
	//
	// curl -sf http://127.0.0.1:8080/healthz
	ReadyCommand *string `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	// The image repository configuration.
	RegistryConfig *PublicTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// The sandbox startup command.
	//
	// example:
	//
	// sleep infinity
	StartCommand *string `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
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
