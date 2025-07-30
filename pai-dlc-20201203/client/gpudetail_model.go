// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGPUDetail interface {
	dara.Model
	String() string
	GoString() string
	SetGPU(v string) *GPUDetail
	GetGPU() *string
	SetGPUType(v string) *GPUDetail
	GetGPUType() *string
	SetGPUTypeFullName(v string) *GPUDetail
	GetGPUTypeFullName() *string
}

type GPUDetail struct {
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// Tesla-V100-32G
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// nvidia.com/gpu-tesla-v100-sxm2-16gb
	GPUTypeFullName *string `json:"GPUTypeFullName,omitempty" xml:"GPUTypeFullName,omitempty"`
}

func (s GPUDetail) String() string {
	return dara.Prettify(s)
}

func (s GPUDetail) GoString() string {
	return s.String()
}

func (s *GPUDetail) GetGPU() *string {
	return s.GPU
}

func (s *GPUDetail) GetGPUType() *string {
	return s.GPUType
}

func (s *GPUDetail) GetGPUTypeFullName() *string {
	return s.GPUTypeFullName
}

func (s *GPUDetail) SetGPU(v string) *GPUDetail {
	s.GPU = &v
	return s
}

func (s *GPUDetail) SetGPUType(v string) *GPUDetail {
	s.GPUType = &v
	return s
}

func (s *GPUDetail) SetGPUTypeFullName(v string) *GPUDetail {
	s.GPUTypeFullName = &v
	return s
}

func (s *GPUDetail) Validate() error {
	return dara.Validate(s)
}
