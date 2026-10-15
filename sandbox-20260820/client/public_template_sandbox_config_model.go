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
	SetSteps(v []*PublicTemplateStep) *PublicTemplateSandboxConfig
	GetSteps() []*PublicTemplateStep
}

type PublicTemplateSandboxConfig struct {
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
	// The readiness probe command for the sandbox. A successful execution indicates that the sandbox is ready to receive requests. This parameter is supported only by first-generation sandboxes.
	//
	// example:
	//
	// curl -sf http://127.0.0.1:49983/health
	ReadyCommand *string `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	// The configuration used to access the image repository. For security reasons, only the network configuration is returned in query responses, and logon credentials are not returned.
	RegistryConfig *PublicTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// The startup command for the sandbox container. This parameter is supported only by first-generation sandboxes.
	//
	// example:
	//
	// /usr/local/bin/start.sh
	StartCommand *string `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
	// The list of custom build steps.
	Steps []*PublicTemplateStep `json:"steps" xml:"steps" type:"Repeated"`
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

func (s *PublicTemplateSandboxConfig) GetSteps() []*PublicTemplateStep {
	return s.Steps
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

func (s *PublicTemplateSandboxConfig) SetSteps(v []*PublicTemplateStep) *PublicTemplateSandboxConfig {
	s.Steps = v
	return s
}

func (s *PublicTemplateSandboxConfig) Validate() error {
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
