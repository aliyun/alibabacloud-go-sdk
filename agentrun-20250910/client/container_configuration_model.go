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
}

type ContainerConfiguration struct {
	// 阿里云容器镜像服务（ACR）的实例ID或名称
	//
	// example:
	//
	// cri-xxxxx
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// 在容器中运行的命令（例如：[\"python3\", \"app.py\"]）
	//
	// example:
	//
	// python3,app.py
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/my-namespace/agent-runtime:latest
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// 容器镜像的来源类型，支持ACR（阿里云容器镜像服务）、ACREE（阿里云容器镜像服务企业版）、CUSTOM（自定义镜像仓库）
	//
	// example:
	//
	// ACR
	ImageRegistryType *string `json:"imageRegistryType,omitempty" xml:"imageRegistryType,omitempty"`
	Port              *int32  `json:"port,omitempty" xml:"port,omitempty"`
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

func (s *ContainerConfiguration) Validate() error {
	return dara.Validate(s)
}
