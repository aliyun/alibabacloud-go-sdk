// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaDetail interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v string) *QuotaDetail
	GetCPU() *string
	SetGPU(v string) *QuotaDetail
	GetGPU() *string
	SetGPUDetails(v []*GPUDetail) *QuotaDetail
	GetGPUDetails() []*GPUDetail
	SetGPUType(v string) *QuotaDetail
	GetGPUType() *string
	SetGPUTypeFullName(v string) *QuotaDetail
	GetGPUTypeFullName() *string
	SetMemory(v string) *QuotaDetail
	GetMemory() *string
}

type QuotaDetail struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 5
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// The GPU details.
	GPUDetails []*GPUDetail `json:"GPUDetails,omitempty" xml:"GPUDetails,omitempty" type:"Repeated"`
	// The GPU type.
	//
	// example:
	//
	// Tesla-V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The full name of the GPU type.
	//
	// example:
	//
	// nvidia.com/gpu
	GPUTypeFullName *string `json:"GPUTypeFullName,omitempty" xml:"GPUTypeFullName,omitempty"`
	// The memory capacity.
	//
	// example:
	//
	// 10Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s QuotaDetail) String() string {
	return dara.Prettify(s)
}

func (s QuotaDetail) GoString() string {
	return s.String()
}

func (s *QuotaDetail) GetCPU() *string {
	return s.CPU
}

func (s *QuotaDetail) GetGPU() *string {
	return s.GPU
}

func (s *QuotaDetail) GetGPUDetails() []*GPUDetail {
	return s.GPUDetails
}

func (s *QuotaDetail) GetGPUType() *string {
	return s.GPUType
}

func (s *QuotaDetail) GetGPUTypeFullName() *string {
	return s.GPUTypeFullName
}

func (s *QuotaDetail) GetMemory() *string {
	return s.Memory
}

func (s *QuotaDetail) SetCPU(v string) *QuotaDetail {
	s.CPU = &v
	return s
}

func (s *QuotaDetail) SetGPU(v string) *QuotaDetail {
	s.GPU = &v
	return s
}

func (s *QuotaDetail) SetGPUDetails(v []*GPUDetail) *QuotaDetail {
	s.GPUDetails = v
	return s
}

func (s *QuotaDetail) SetGPUType(v string) *QuotaDetail {
	s.GPUType = &v
	return s
}

func (s *QuotaDetail) SetGPUTypeFullName(v string) *QuotaDetail {
	s.GPUTypeFullName = &v
	return s
}

func (s *QuotaDetail) SetMemory(v string) *QuotaDetail {
	s.Memory = &v
	return s
}

func (s *QuotaDetail) Validate() error {
	if s.GPUDetails != nil {
		for _, item := range s.GPUDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
