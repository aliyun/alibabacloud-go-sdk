// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaUsage interface {
	dara.Model
	String() string
	GoString() string
	SetCpuCores(v float64) *QuotaUsage
	GetCpuCores() *float64
	SetInstanceCount(v int64) *QuotaUsage
	GetInstanceCount() *int64
	SetMemoryGB(v float64) *QuotaUsage
	GetMemoryGB() *float64
}

type QuotaUsage struct {
	// 已占用的 CPU 配额，单位为 vCPU，不是 CPU 使用率。
	//
	// example:
	//
	// 1
	CpuCores *float64 `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// 尚未释放的实例授权数量。
	//
	// example:
	//
	// 1
	InstanceCount *int64 `json:"instanceCount,omitempty" xml:"instanceCount,omitempty"`
	// 已占用的内存配额，单位为 GB，不是进程实际内存。
	//
	// example:
	//
	// 1.5
	MemoryGB *float64 `json:"memoryGB,omitempty" xml:"memoryGB,omitempty"`
}

func (s QuotaUsage) String() string {
	return dara.Prettify(s)
}

func (s QuotaUsage) GoString() string {
	return s.String()
}

func (s *QuotaUsage) GetCpuCores() *float64 {
	return s.CpuCores
}

func (s *QuotaUsage) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *QuotaUsage) GetMemoryGB() *float64 {
	return s.MemoryGB
}

func (s *QuotaUsage) SetCpuCores(v float64) *QuotaUsage {
	s.CpuCores = &v
	return s
}

func (s *QuotaUsage) SetInstanceCount(v int64) *QuotaUsage {
	s.InstanceCount = &v
	return s
}

func (s *QuotaUsage) SetMemoryGB(v float64) *QuotaUsage {
	s.MemoryGB = &v
	return s
}

func (s *QuotaUsage) Validate() error {
	return dara.Validate(s)
}
