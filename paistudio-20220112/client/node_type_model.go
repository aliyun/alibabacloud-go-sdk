// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeType interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *NodeType
	GetAcceleratorType() *string
	SetCPU(v string) *NodeType
	GetCPU() *string
	SetGPU(v string) *NodeType
	GetGPU() *string
	SetGPUMemory(v string) *NodeType
	GetGPUMemory() *string
	SetGPUType(v string) *NodeType
	GetGPUType() *string
	SetMemory(v string) *NodeType
	GetMemory() *string
	SetNodeType(v string) *NodeType
	GetNodeType() *string
}

type NodeType struct {
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// example:
	//
	// 16
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 0
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// 80G
	GPUMemory *string `json:"GPUMemory,omitempty" xml:"GPUMemory,omitempty"`
	GPUType   *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 64Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// ecs.g6.4xlarge
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s NodeType) String() string {
	return dara.Prettify(s)
}

func (s NodeType) GoString() string {
	return s.String()
}

func (s *NodeType) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *NodeType) GetCPU() *string {
	return s.CPU
}

func (s *NodeType) GetGPU() *string {
	return s.GPU
}

func (s *NodeType) GetGPUMemory() *string {
	return s.GPUMemory
}

func (s *NodeType) GetGPUType() *string {
	return s.GPUType
}

func (s *NodeType) GetMemory() *string {
	return s.Memory
}

func (s *NodeType) GetNodeType() *string {
	return s.NodeType
}

func (s *NodeType) SetAcceleratorType(v string) *NodeType {
	s.AcceleratorType = &v
	return s
}

func (s *NodeType) SetCPU(v string) *NodeType {
	s.CPU = &v
	return s
}

func (s *NodeType) SetGPU(v string) *NodeType {
	s.GPU = &v
	return s
}

func (s *NodeType) SetGPUMemory(v string) *NodeType {
	s.GPUMemory = &v
	return s
}

func (s *NodeType) SetGPUType(v string) *NodeType {
	s.GPUType = &v
	return s
}

func (s *NodeType) SetMemory(v string) *NodeType {
	s.Memory = &v
	return s
}

func (s *NodeType) SetNodeType(v string) *NodeType {
	s.NodeType = &v
	return s
}

func (s *NodeType) Validate() error {
	return dara.Validate(s)
}
