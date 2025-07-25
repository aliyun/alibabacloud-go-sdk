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
	SetGPUType(v string) *ResourceAmount
	GetGPUType() *string
	SetMemory(v string) *ResourceAmount
	GetMemory() *string
}

type ResourceAmount struct {
	// example:
	//
	// 100
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 16
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// GPU
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
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
