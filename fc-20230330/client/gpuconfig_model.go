// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGPUConfig interface {
	dara.Model
	String() string
	GoString() string
	SetGpuMemorySize(v int32) *GPUConfig
	GetGpuMemorySize() *int32
	SetGpuType(v string) *GPUConfig
	GetGpuType() *string
}

type GPUConfig struct {
	// The GPU memory size. Unit: MB. The value is a multiple of 1024 MB.
	//
	// example:
	//
	// 2048
	GpuMemorySize *int32 `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	// The type of GPU cards. Valid values: fc.gpu.tesla.1: Tesla T4 fc.gpu.ampere.1: Ampere A10
	//
	// example:
	//
	// fc.gpu.ampere.1
	GpuType *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s GPUConfig) String() string {
	return dara.Prettify(s)
}

func (s GPUConfig) GoString() string {
	return s.String()
}

func (s *GPUConfig) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *GPUConfig) GetGpuType() *string {
	return s.GpuType
}

func (s *GPUConfig) SetGpuMemorySize(v int32) *GPUConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *GPUConfig) SetGpuType(v string) *GPUConfig {
	s.GpuType = &v
	return s
}

func (s *GPUConfig) Validate() error {
	return dara.Validate(s)
}
