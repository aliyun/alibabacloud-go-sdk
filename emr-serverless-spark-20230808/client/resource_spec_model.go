// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceSpec interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v int32) *ResourceSpec
	GetCu() *int32
	SetGpu(v int32) *ResourceSpec
	GetGpu() *int32
	SetGpuMachineNum(v int32) *ResourceSpec
	GetGpuMachineNum() *int32
	SetInstanceTypeId(v string) *ResourceSpec
	GetInstanceTypeId() *string
	SetMaxCu(v int32) *ResourceSpec
	GetMaxCu() *int32
}

type ResourceSpec struct {
	// The number of CUs. Currently, 1 CU is equivalent to 1 vCPU and 4 GB of memory.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	// The number of GPU machines.
	//
	// example:
	//
	// 1
	GpuMachineNum *int32 `json:"gpuMachineNum,omitempty" xml:"gpuMachineNum,omitempty"`
	// The instance type ID.
	//
	// example:
	//
	// ecs.gn7i-c32g1.8xlarge
	InstanceTypeId *string `json:"instanceTypeId,omitempty" xml:"instanceTypeId,omitempty"`
	// The upper limit of the absolute value for the queue quota. If not specified, the backend falls back to the number of CUs.
	//
	// example:
	//
	// 4
	MaxCu *int32 `json:"maxCu,omitempty" xml:"maxCu,omitempty"`
}

func (s ResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ResourceSpec) GoString() string {
	return s.String()
}

func (s *ResourceSpec) GetCu() *int32 {
	return s.Cu
}

func (s *ResourceSpec) GetGpu() *int32 {
	return s.Gpu
}

func (s *ResourceSpec) GetGpuMachineNum() *int32 {
	return s.GpuMachineNum
}

func (s *ResourceSpec) GetInstanceTypeId() *string {
	return s.InstanceTypeId
}

func (s *ResourceSpec) GetMaxCu() *int32 {
	return s.MaxCu
}

func (s *ResourceSpec) SetCu(v int32) *ResourceSpec {
	s.Cu = &v
	return s
}

func (s *ResourceSpec) SetGpu(v int32) *ResourceSpec {
	s.Gpu = &v
	return s
}

func (s *ResourceSpec) SetGpuMachineNum(v int32) *ResourceSpec {
	s.GpuMachineNum = &v
	return s
}

func (s *ResourceSpec) SetInstanceTypeId(v string) *ResourceSpec {
	s.InstanceTypeId = &v
	return s
}

func (s *ResourceSpec) SetMaxCu(v int32) *ResourceSpec {
	s.MaxCu = &v
	return s
}

func (s *ResourceSpec) Validate() error {
	return dara.Validate(s)
}
