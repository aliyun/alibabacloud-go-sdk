// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResidentResourceAllocation interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *ResidentResourceAllocation
	GetFunctionName() *string
	SetInstanceCount(v int32) *ResidentResourceAllocation
	GetInstanceCount() *int32
	SetInstanceType(v string) *ResidentResourceAllocation
	GetInstanceType() *string
	SetQualifier(v string) *ResidentResourceAllocation
	GetQualifier() *string
	SetTotalCpuCores(v float64) *ResidentResourceAllocation
	GetTotalCpuCores() *float64
	SetTotalDiskSize(v float64) *ResidentResourceAllocation
	GetTotalDiskSize() *float64
	SetTotalGpuMemorySize(v float64) *ResidentResourceAllocation
	GetTotalGpuMemorySize() *float64
	SetTotalMemorySize(v float64) *ResidentResourceAllocation
	GetTotalMemorySize() *float64
}

type ResidentResourceAllocation struct {
	// 使用该资源池的函数名
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// 实例数
	InstanceCount *int32  `json:"instanceCount,omitempty" xml:"instanceCount,omitempty"`
	InstanceType  *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// 函数的别名
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// CPU 占用总核数
	TotalCpuCores *float64 `json:"totalCpuCores,omitempty" xml:"totalCpuCores,omitempty"`
	// 占用磁盘大小，单位 GB
	TotalDiskSize *float64 `json:"totalDiskSize,omitempty" xml:"totalDiskSize,omitempty"`
	// 占用显存大小，单位 GB
	TotalGpuMemorySize *float64 `json:"totalGpuMemorySize,omitempty" xml:"totalGpuMemorySize,omitempty"`
	// 内存占用大小，单位 GB
	TotalMemorySize *float64 `json:"totalMemorySize,omitempty" xml:"totalMemorySize,omitempty"`
}

func (s ResidentResourceAllocation) String() string {
	return dara.Prettify(s)
}

func (s ResidentResourceAllocation) GoString() string {
	return s.String()
}

func (s *ResidentResourceAllocation) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ResidentResourceAllocation) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *ResidentResourceAllocation) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ResidentResourceAllocation) GetQualifier() *string {
	return s.Qualifier
}

func (s *ResidentResourceAllocation) GetTotalCpuCores() *float64 {
	return s.TotalCpuCores
}

func (s *ResidentResourceAllocation) GetTotalDiskSize() *float64 {
	return s.TotalDiskSize
}

func (s *ResidentResourceAllocation) GetTotalGpuMemorySize() *float64 {
	return s.TotalGpuMemorySize
}

func (s *ResidentResourceAllocation) GetTotalMemorySize() *float64 {
	return s.TotalMemorySize
}

func (s *ResidentResourceAllocation) SetFunctionName(v string) *ResidentResourceAllocation {
	s.FunctionName = &v
	return s
}

func (s *ResidentResourceAllocation) SetInstanceCount(v int32) *ResidentResourceAllocation {
	s.InstanceCount = &v
	return s
}

func (s *ResidentResourceAllocation) SetInstanceType(v string) *ResidentResourceAllocation {
	s.InstanceType = &v
	return s
}

func (s *ResidentResourceAllocation) SetQualifier(v string) *ResidentResourceAllocation {
	s.Qualifier = &v
	return s
}

func (s *ResidentResourceAllocation) SetTotalCpuCores(v float64) *ResidentResourceAllocation {
	s.TotalCpuCores = &v
	return s
}

func (s *ResidentResourceAllocation) SetTotalDiskSize(v float64) *ResidentResourceAllocation {
	s.TotalDiskSize = &v
	return s
}

func (s *ResidentResourceAllocation) SetTotalGpuMemorySize(v float64) *ResidentResourceAllocation {
	s.TotalGpuMemorySize = &v
	return s
}

func (s *ResidentResourceAllocation) SetTotalMemorySize(v float64) *ResidentResourceAllocation {
	s.TotalMemorySize = &v
	return s
}

func (s *ResidentResourceAllocation) Validate() error {
	return dara.Validate(s)
}
