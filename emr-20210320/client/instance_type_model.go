// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceType interface {
	dara.Model
	String() string
	GoString() string
	SetCpuArchitecture(v string) *InstanceType
	GetCpuArchitecture() *string
	SetCpuCore(v int32) *InstanceType
	GetCpuCore() *int32
	SetInstanceCategory(v string) *InstanceType
	GetInstanceCategory() *string
	SetInstanceType(v string) *InstanceType
	GetInstanceType() *string
	SetInstanceTypeFamily(v string) *InstanceType
	GetInstanceTypeFamily() *string
	SetLocalStorageAmount(v int32) *InstanceType
	GetLocalStorageAmount() *int32
	SetLocalStorageCapacity(v int64) *InstanceType
	GetLocalStorageCapacity() *int64
	SetModifyType(v string) *InstanceType
	GetModifyType() *string
	SetOptimized(v bool) *InstanceType
	GetOptimized() *bool
}

type InstanceType struct {
	// The CPU architecture. Valid values:
	//
	// - `X86`: X86 architecture.
	//
	// - `ARM`: ARM architecture.
	//
	// example:
	//
	// X86
	CpuArchitecture *string `json:"CpuArchitecture,omitempty" xml:"CpuArchitecture,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 4
	CpuCore *int32 `json:"CpuCore,omitempty" xml:"CpuCore,omitempty"`
	// The instance category. Valid values:
	//
	// - `General-purpose`: A general-purpose instance type.
	//
	// - `Compute-optimized`: A compute-optimized instance type.
	//
	// - `Memory-optimized`: A memory-optimized instance type.
	//
	// - `Big data`: A big data instance type.
	//
	// - `Local SSDs`: A local SSD instance type.
	//
	// - `High Clock Speed`: A high clock speed instance type.
	//
	// - `Enhanced`: An enhanced instance type.
	//
	// - `Shared`: A shared instance type.
	//
	// - `Compute-optimized with GPU`: A compute-optimized instance type with GPUs.
	//
	// - `Visual Compute-optimized`: A visual compute-optimized instance type.
	//
	// - `Heterogeneous Service`: A heterogeneous service instance type.
	//
	// - `Compute-optimized with FPGA`: A compute-optimized instance type with FPGAs.
	//
	// - `Compute-optimized with NPU`: A compute-optimized instance type with NPUs.
	//
	// - `ECS Bare Metal`: An ECS Bare Metal instance.
	//
	// - `Super Computing Cluster`: A supercomputing cluster instance type.
	//
	// example:
	//
	// Compute-optimized
	InstanceCategory *string `json:"InstanceCategory,omitempty" xml:"InstanceCategory,omitempty"`
	// The ECS instance type. For more information, see [Instance type families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance type family. For valid values, see the ECS documentation for [DescribeInstanceTypeFamilies](https://help.aliyun.com/document_detail/25621.html).
	//
	// example:
	//
	// ecs.g6
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The number of local disks attached to the instance.
	//
	// example:
	//
	// 8
	LocalStorageAmount *int32 `json:"LocalStorageAmount,omitempty" xml:"LocalStorageAmount,omitempty"`
	// The capacity of each local disk attached to the instance, in GiB.
	//
	// example:
	//
	// 40
	LocalStorageCapacity *int64  `json:"LocalStorageCapacity,omitempty" xml:"LocalStorageCapacity,omitempty"`
	ModifyType           *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// Specifies whether the instance type is I/O optimized. Valid values:
	//
	// - `true`: The instance type is I/O optimized.
	//
	// - `false`: The instance type is not I/O optimized.
	//
	// example:
	//
	// true
	Optimized *bool `json:"Optimized,omitempty" xml:"Optimized,omitempty"`
}

func (s InstanceType) String() string {
	return dara.Prettify(s)
}

func (s InstanceType) GoString() string {
	return s.String()
}

func (s *InstanceType) GetCpuArchitecture() *string {
	return s.CpuArchitecture
}

func (s *InstanceType) GetCpuCore() *int32 {
	return s.CpuCore
}

func (s *InstanceType) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *InstanceType) GetInstanceType() *string {
	return s.InstanceType
}

func (s *InstanceType) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *InstanceType) GetLocalStorageAmount() *int32 {
	return s.LocalStorageAmount
}

func (s *InstanceType) GetLocalStorageCapacity() *int64 {
	return s.LocalStorageCapacity
}

func (s *InstanceType) GetModifyType() *string {
	return s.ModifyType
}

func (s *InstanceType) GetOptimized() *bool {
	return s.Optimized
}

func (s *InstanceType) SetCpuArchitecture(v string) *InstanceType {
	s.CpuArchitecture = &v
	return s
}

func (s *InstanceType) SetCpuCore(v int32) *InstanceType {
	s.CpuCore = &v
	return s
}

func (s *InstanceType) SetInstanceCategory(v string) *InstanceType {
	s.InstanceCategory = &v
	return s
}

func (s *InstanceType) SetInstanceType(v string) *InstanceType {
	s.InstanceType = &v
	return s
}

func (s *InstanceType) SetInstanceTypeFamily(v string) *InstanceType {
	s.InstanceTypeFamily = &v
	return s
}

func (s *InstanceType) SetLocalStorageAmount(v int32) *InstanceType {
	s.LocalStorageAmount = &v
	return s
}

func (s *InstanceType) SetLocalStorageCapacity(v int64) *InstanceType {
	s.LocalStorageCapacity = &v
	return s
}

func (s *InstanceType) SetModifyType(v string) *InstanceType {
	s.ModifyType = &v
	return s
}

func (s *InstanceType) SetOptimized(v bool) *InstanceType {
	s.Optimized = &v
	return s
}

func (s *InstanceType) Validate() error {
	return dara.Validate(s)
}
