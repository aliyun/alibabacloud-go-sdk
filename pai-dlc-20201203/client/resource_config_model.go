// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v string) *ResourceConfig
	GetCPU() *string
	SetGPU(v string) *ResourceConfig
	GetGPU() *string
	SetGPUType(v string) *ResourceConfig
	GetGPUType() *string
	SetMemory(v string) *ResourceConfig
	GetMemory() *string
	SetSharedMemory(v string) *ResourceConfig
	GetSharedMemory() *string
}

type ResourceConfig struct {
	// example:
	//
	// 10
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 3
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// Tesla-V100-16G
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 10Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 5Gi
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s ResourceConfig) String() string {
	return dara.Prettify(s)
}

func (s ResourceConfig) GoString() string {
	return s.String()
}

func (s *ResourceConfig) GetCPU() *string {
	return s.CPU
}

func (s *ResourceConfig) GetGPU() *string {
	return s.GPU
}

func (s *ResourceConfig) GetGPUType() *string {
	return s.GPUType
}

func (s *ResourceConfig) GetMemory() *string {
	return s.Memory
}

func (s *ResourceConfig) GetSharedMemory() *string {
	return s.SharedMemory
}

func (s *ResourceConfig) SetCPU(v string) *ResourceConfig {
	s.CPU = &v
	return s
}

func (s *ResourceConfig) SetGPU(v string) *ResourceConfig {
	s.GPU = &v
	return s
}

func (s *ResourceConfig) SetGPUType(v string) *ResourceConfig {
	s.GPUType = &v
	return s
}

func (s *ResourceConfig) SetMemory(v string) *ResourceConfig {
	s.Memory = &v
	return s
}

func (s *ResourceConfig) SetSharedMemory(v string) *ResourceConfig {
	s.SharedMemory = &v
	return s
}

func (s *ResourceConfig) Validate() error {
	return dara.Validate(s)
}
