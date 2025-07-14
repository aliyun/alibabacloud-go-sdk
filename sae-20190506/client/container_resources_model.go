// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerResources interface {
	dara.Model
	String() string
	GoString() string
	SetCpu(v int32) *ContainerResources
	GetCpu() *int32
	SetMemory(v int32) *ContainerResources
	GetMemory() *int32
}

type ContainerResources struct {
	// This parameter is required.
	//
	// example:
	//
	// 2000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ContainerResources) String() string {
	return dara.Prettify(s)
}

func (s ContainerResources) GoString() string {
	return s.String()
}

func (s *ContainerResources) GetCpu() *int32 {
	return s.Cpu
}

func (s *ContainerResources) GetMemory() *int32 {
	return s.Memory
}

func (s *ContainerResources) SetCpu(v int32) *ContainerResources {
	s.Cpu = &v
	return s
}

func (s *ContainerResources) SetMemory(v int32) *ContainerResources {
	s.Memory = &v
	return s
}

func (s *ContainerResources) Validate() error {
	return dara.Validate(s)
}
