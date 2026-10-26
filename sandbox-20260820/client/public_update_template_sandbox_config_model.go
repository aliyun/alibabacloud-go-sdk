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
	SetSteps(v []*PublicUpdateTemplateStep) *PublicUpdateTemplateSandboxConfig
	GetSteps() []*PublicUpdateTemplateStep
}

type PublicUpdateTemplateSandboxConfig struct {
	// The ID of the ACR Enterprise instance.
	//
	// example:
	//
	// cri-abc123
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The sandbox generation. Valid values: 1: rund. 2: micro.
	//
	// example:
	//
	// 2
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The image address.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/my/ns:v1
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// linux
	OsType *string `json:"osType,omitempty" xml:"osType,omitempty"`
	// The readiness probe command for the micro sandbox. This command is supported only by the second generation.
	//
	// example:
	//
	// /ready
	ReadyCommand *string `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	// The image repository configuration.
	RegistryConfig *PublicUpdateTemplateRegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// The image repository type.
	//
	// example:
	//
	// acr
	RegistryType *string `json:"registryType,omitempty" xml:"registryType,omitempty"`
	// The startup command for the micro sandbox. This command is supported only by the second generation.
	//
	// example:
	//
	// /start-coroutines.sh
	StartCommand *string `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
	// The list of custom build steps. If you do not specify this parameter, the original steps are retained. If you pass an empty array, the steps are cleared.
	Steps []*PublicUpdateTemplateStep `json:"steps" xml:"steps" type:"Repeated"`
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

func (s *PublicUpdateTemplateSandboxConfig) GetSteps() []*PublicUpdateTemplateStep {
	return s.Steps
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

func (s *PublicUpdateTemplateSandboxConfig) SetSteps(v []*PublicUpdateTemplateStep) *PublicUpdateTemplateSandboxConfig {
	s.Steps = v
	return s
}

func (s *PublicUpdateTemplateSandboxConfig) Validate() error {
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
