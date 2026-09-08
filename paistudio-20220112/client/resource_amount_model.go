// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceAmount interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v string) *ResourceAmount
	GetCPU() *string
	SetGPU(v string) *ResourceAmount
	GetGPU() *string
	SetGPUMemory(v string) *ResourceAmount
	GetGPUMemory() *string
	SetGPUMemoryBytes(v int64) *ResourceAmount
	GetGPUMemoryBytes() *int64
	SetGPUType(v string) *ResourceAmount
	GetGPUType() *string
	SetMemory(v string) *ResourceAmount
	GetMemory() *string
}

type ResourceAmount struct {
	// Total CPU
	//
	// example:
	//
	// 100
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// Total GPU cards
	//
	// example:
	//
	// 16
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// 80G
	GPUMemory *string `json:"GPUMemory,omitempty" xml:"GPUMemory,omitempty"`
	// example:
	//
	// 85899345920
	GPUMemoryBytes *int64 `json:"GPUMemoryBytes,omitempty" xml:"GPUMemoryBytes,omitempty"`
	// GPU card type
	//
	// example:
	//
	// GPU
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// Total memory
	//
	// example:
	//
	// 100Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ResourceAmount) String() string {
	return dara.Prettify(s)
}

func (s ResourceAmount) GoString() string {
	return s.String()
}

func (s *ResourceAmount) GetCPU() *string {
	return s.CPU
}

func (s *ResourceAmount) GetGPU() *string {
	return s.GPU
}

func (s *ResourceAmount) GetGPUMemory() *string {
	return s.GPUMemory
}

func (s *ResourceAmount) GetGPUMemoryBytes() *int64 {
	return s.GPUMemoryBytes
}

func (s *ResourceAmount) GetGPUType() *string {
	return s.GPUType
}

func (s *ResourceAmount) GetMemory() *string {
	return s.Memory
}

func (s *ResourceAmount) SetCPU(v string) *ResourceAmount {
	s.CPU = &v
	return s
}

func (s *ResourceAmount) SetGPU(v string) *ResourceAmount {
	s.GPU = &v
	return s
}

func (s *ResourceAmount) SetGPUMemory(v string) *ResourceAmount {
	s.GPUMemory = &v
	return s
}

func (s *ResourceAmount) SetGPUMemoryBytes(v int64) *ResourceAmount {
	s.GPUMemoryBytes = &v
	return s
}

func (s *ResourceAmount) SetGPUType(v string) *ResourceAmount {
	s.GPUType = &v
	return s
}

func (s *ResourceAmount) SetMemory(v string) *ResourceAmount {
	s.Memory = &v
	return s
}

func (s *ResourceAmount) Validate() error {
	return dara.Validate(s)
}
