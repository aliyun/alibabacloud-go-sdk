// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateSandboxConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *CreateTemplateSandboxConfig
	GetAcrInstanceId() *string
	SetGeneration(v int32) *CreateTemplateSandboxConfig
	GetGeneration() *int32
	SetImage(v string) *CreateTemplateSandboxConfig
	GetImage() *string
	SetOsType(v string) *CreateTemplateSandboxConfig
	GetOsType() *string
	SetReadyCommand(v string) *CreateTemplateSandboxConfig
	GetReadyCommand() *string
	SetRegistryConfig(v *CreateTemplateRegistryConfig) *CreateTemplateSandboxConfig
	GetRegistryConfig() *CreateTemplateRegistryConfig
	SetRegistryType(v string) *CreateTemplateSandboxConfig
	GetRegistryType() *string
	SetStartCommand(v string) *CreateTemplateSandboxConfig
	GetStartCommand() *string
	SetSteps(v []*CreateTemplateStep) *CreateTemplateSandboxConfig
	GetSteps() []*CreateTemplateStep
}

type CreateTemplateSandboxConfig struct {
	// The Container Registry Enterprise instance ID.
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
	RegistryConfig *CreateTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// The image repository type.
	//
	// example:
	//
	// acr
	RegistryType *string `json:"registryType,omitempty" xml:"registryType,omitempty"`
	// The sandbox startup command.
	//
	// example:
	//
	// sleep infinity
	StartCommand *string `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
	// The list of custom build steps.
	Steps []*CreateTemplateStep `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s CreateTemplateSandboxConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateSandboxConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateSandboxConfig) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateTemplateSandboxConfig) GetGeneration() *int32 {
	return s.Generation
}

func (s *CreateTemplateSandboxConfig) GetImage() *string {
	return s.Image
}

func (s *CreateTemplateSandboxConfig) GetOsType() *string {
	return s.OsType
}

func (s *CreateTemplateSandboxConfig) GetReadyCommand() *string {
	return s.ReadyCommand
}

func (s *CreateTemplateSandboxConfig) GetRegistryConfig() *CreateTemplateRegistryConfig {
	return s.RegistryConfig
}

func (s *CreateTemplateSandboxConfig) GetRegistryType() *string {
	return s.RegistryType
}

func (s *CreateTemplateSandboxConfig) GetStartCommand() *string {
	return s.StartCommand
}

func (s *CreateTemplateSandboxConfig) GetSteps() []*CreateTemplateStep {
	return s.Steps
}

func (s *CreateTemplateSandboxConfig) SetAcrInstanceId(v string) *CreateTemplateSandboxConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateTemplateSandboxConfig) SetGeneration(v int32) *CreateTemplateSandboxConfig {
	s.Generation = &v
	return s
}

func (s *CreateTemplateSandboxConfig) SetImage(v string) *CreateTemplateSandboxConfig {
	s.Image = &v
	return s
}

func (s *CreateTemplateSandboxConfig) SetOsType(v string) *CreateTemplateSandboxConfig {
	s.OsType = &v
	return s
}

func (s *CreateTemplateSandboxConfig) SetReadyCommand(v string) *CreateTemplateSandboxConfig {
	s.ReadyCommand = &v
	return s
}

func (s *CreateTemplateSandboxConfig) SetRegistryConfig(v *CreateTemplateRegistryConfig) *CreateTemplateSandboxConfig {
	s.RegistryConfig = v
	return s
}

func (s *CreateTemplateSandboxConfig) SetRegistryType(v string) *CreateTemplateSandboxConfig {
	s.RegistryType = &v
	return s
}

func (s *CreateTemplateSandboxConfig) SetStartCommand(v string) *CreateTemplateSandboxConfig {
	s.StartCommand = &v
	return s
}

func (s *CreateTemplateSandboxConfig) SetSteps(v []*CreateTemplateStep) *CreateTemplateSandboxConfig {
	s.Steps = v
	return s
}

func (s *CreateTemplateSandboxConfig) Validate() error {
	if s.RegistryConfig != nil {
		if err := s.RegistryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Steps != nil {
		for _, item := range s.Steps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
