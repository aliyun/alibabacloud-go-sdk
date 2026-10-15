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
	// The ID of the Container Registry Enterprise instance. This parameter is required when registryType is set to acree.
	//
	// example:
	//
	// cri-****
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The sandbox generation. Valid values: 1 (first generation) and 2 (second generation). Default value: 1.
	//
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The address of the sandbox base image.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/sandbox-demo/python:3.12
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The operating system type of the second-generation sandbox. Valid values: linux-amd64, android-36-osworld, and windows-osworld.
	//
	// example:
	//
	// linux-amd64
	OsType *string `json:"osType,omitempty" xml:"osType,omitempty"`
	// The sandbox readiness probe command. A successful execution indicates that the sandbox is ready to accept requests. This parameter is supported only for first-generation sandboxes.
	//
	// example:
	//
	// curl -sf http://127.0.0.1:49983/health
	ReadyCommand *string `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	// The credential, TLS, and network configuration required to access the source image repository.
	RegistryConfig *CreateTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// The image repository type. Valid values: oci (standard OCI repository), acr (Container Registry Personal Edition), and acree (Container Registry Enterprise Edition). If not specified, the type is inferred from the image address.
	//
	// example:
	//
	// acree
	RegistryType *string `json:"registryType,omitempty" xml:"registryType,omitempty"`
	// The startup command for the sandbox container. This parameter is supported only for first-generation sandboxes.
	//
	// example:
	//
	// /usr/local/bin/start.sh
	StartCommand *string `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
	// The list of custom build instructions. This is not supported in the current version. Passing a non-empty list returns an error.
	Steps []*CreateTemplateStep `json:"steps" xml:"steps" type:"Repeated"`
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
