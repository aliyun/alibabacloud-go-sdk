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
	// cpu架构。
	//
	// example:
	//
	// X86
	CpuArchitecture *string `json:"CpuArchitecture,omitempty" xml:"CpuArchitecture,omitempty"`
	// vCPU内核数目。
	//
	// example:
	//
	// 4
	CpuCore *int32 `json:"CpuCore,omitempty" xml:"CpuCore,omitempty"`
	// 实例规格分类。取值范围：
	//
	// - General-purpose： 通用型。
	//
	// - Compute-optimized：计算型。
	//
	// - Memory-optimized：内存型。
	//
	// - Big data：大数据型。
	//
	// - Local SSDs ：本地SSD型。
	//
	// - High Clock Speed ：高主频型。
	//
	// - Enhanced ：增强型。
	//
	// - Shared：共享型。
	//
	// - Compute-optimized with GPU ：GPU计算型。
	//
	// - Visual Compute-optimized ：视觉计算型。
	//
	// - Heterogeneous Service ：异构服务型。
	//
	// - Compute-optimized with FPGA ：FPGA计算型。
	//
	// - Compute-optimized with NPU ：NPU计算型。
	//
	// - ECS Bare Metal ：弹性裸金属服务器。
	//
	// - Super Computing Cluster：超级计算集群。
	//
	// example:
	//
	// Compute-optimized
	InstanceCategory *string `json:"InstanceCategory,omitempty" xml:"InstanceCategory,omitempty"`
	// 实例规格。
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 实例规格所属的实例规格族。取值请参见DescribeInstanceTypeFamilies。
	//
	// example:
	//
	// ecs.g6
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// 实例挂载的本地盘的数量。
	//
	// example:
	//
	// 8
	LocalStorageAmount *int32 `json:"LocalStorageAmount,omitempty" xml:"LocalStorageAmount,omitempty"`
	// 实例挂载的本地盘的单盘容量。单位：GiB
	//
	// example:
	//
	// 40
	LocalStorageCapacity *int64  `json:"LocalStorageCapacity,omitempty" xml:"LocalStorageCapacity,omitempty"`
	ModifyType           *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// 是否IO优化类型。
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
