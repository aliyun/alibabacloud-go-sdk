// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v []*string) *ContainerConfiguration
	GetCommand() []*string
	SetImage(v string) *ContainerConfiguration
	GetImage() *string
}

type ContainerConfiguration struct {
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
}

func (s ContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ContainerConfiguration) GoString() string {
	return s.String()
}

func (s *ContainerConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *ContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *ContainerConfiguration) SetCommand(v []*string) *ContainerConfiguration {
	s.Command = v
	return s
}

func (s *ContainerConfiguration) SetImage(v string) *ContainerConfiguration {
	s.Image = &v
	return s
}

func (s *ContainerConfiguration) Validate() error {
	return dara.Validate(s)
}
