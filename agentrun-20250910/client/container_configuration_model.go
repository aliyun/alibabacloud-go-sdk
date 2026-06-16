// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *ContainerConfiguration
	GetAcrInstanceId() *string
	SetCommand(v []*string) *ContainerConfiguration
	GetCommand() []*string
	SetImage(v string) *ContainerConfiguration
	GetImage() *string
	SetImageRegistryType(v string) *ContainerConfiguration
	GetImageRegistryType() *string
	SetPort(v int32) *ContainerConfiguration
	GetPort() *int32
	SetRegistryConfig(v *RegistryConfig) *ContainerConfiguration
	GetRegistryConfig() *RegistryConfig
}

type ContainerConfiguration struct {
	// The instance ID or name of Alibaba Cloud Container Registry (ACR).
	//
	// example:
	//
	// cri-xxxxx
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The sandbox entrypoint command. For example: [\\"python3\\", \\"app.py\\"].
	//
	// example:
	//
	// python3,app.py
	Command []*string `json:"command" xml:"command" type:"Repeated"`
	// The container image URI.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/my-namespace/agent-runtime:latest
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// The source of the container image. Valid values: ACR (Alibaba Cloud Container Registry), ACREE (Alibaba Cloud Container Registry Enterprise Edition), and CUSTOM (a custom image repository).
	//
	// example:
	//
	// ACR
	ImageRegistryType *string `json:"imageRegistryType,omitempty" xml:"imageRegistryType,omitempty"`
	// The listening port for the sandbox runtime.
	//
	// example:
	//
	// 5000
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// 自定义镜像仓库的配置信息，当imageRegistryType为CUSTOM时使用
	//
	// example:
	//
	// {}
	RegistryConfig *RegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
}

func (s ContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ContainerConfiguration) GoString() string {
	return s.String()
}

func (s *ContainerConfiguration) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *ContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *ContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *ContainerConfiguration) GetImageRegistryType() *string {
	return s.ImageRegistryType
}

func (s *ContainerConfiguration) GetPort() *int32 {
	return s.Port
}

func (s *ContainerConfiguration) GetRegistryConfig() *RegistryConfig {
	return s.RegistryConfig
}

func (s *ContainerConfiguration) SetAcrInstanceId(v string) *ContainerConfiguration {
	s.AcrInstanceId = &v
	return s
}

func (s *ContainerConfiguration) SetCommand(v []*string) *ContainerConfiguration {
	s.Command = v
	return s
}

func (s *ContainerConfiguration) SetImage(v string) *ContainerConfiguration {
	s.Image = &v
	return s
}

func (s *ContainerConfiguration) SetImageRegistryType(v string) *ContainerConfiguration {
	s.ImageRegistryType = &v
	return s
}

func (s *ContainerConfiguration) SetPort(v int32) *ContainerConfiguration {
	s.Port = &v
	return s
}

func (s *ContainerConfiguration) SetRegistryConfig(v *RegistryConfig) *ContainerConfiguration {
	s.RegistryConfig = v
	return s
}

func (s *ContainerConfiguration) Validate() error {
	if s.RegistryConfig != nil {
		if err := s.RegistryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
