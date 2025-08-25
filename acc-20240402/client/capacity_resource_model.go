// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCapacityResource interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v string) *CapacityResource
	GetCPU() *string
	SetGPU(v string) *CapacityResource
	GetGPU() *string
	SetMemory(v string) *CapacityResource
	GetMemory() *string
}

type CapacityResource struct {
	CPU    *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU    *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CapacityResource) String() string {
	return dara.Prettify(s)
}

func (s CapacityResource) GoString() string {
	return s.String()
}

func (s *CapacityResource) GetCPU() *string {
	return s.CPU
}

func (s *CapacityResource) GetGPU() *string {
	return s.GPU
}

func (s *CapacityResource) GetMemory() *string {
	return s.Memory
}

func (s *CapacityResource) SetCPU(v string) *CapacityResource {
	s.CPU = &v
	return s
}

func (s *CapacityResource) SetGPU(v string) *CapacityResource {
	s.GPU = &v
	return s
}

func (s *CapacityResource) SetMemory(v string) *CapacityResource {
	s.Memory = &v
	return s
}

func (s *CapacityResource) Validate() error {
	return dara.Validate(s)
}
