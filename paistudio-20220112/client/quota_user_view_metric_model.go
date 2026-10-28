// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaUserViewMetric interface {
	dara.Model
	String() string
	GoString() string
	SetCPUNodeNumber(v int32) *QuotaUserViewMetric
	GetCPUNodeNumber() *int32
	SetCPUUsageRate(v string) *QuotaUserViewMetric
	GetCPUUsageRate() *string
	SetCpuJobNames(v []*string) *QuotaUserViewMetric
	GetCpuJobNames() []*string
	SetCpuNodeNames(v []*string) *QuotaUserViewMetric
	GetCpuNodeNames() []*string
	SetDiskReadRate(v string) *QuotaUserViewMetric
	GetDiskReadRate() *string
	SetDiskWriteRate(v string) *QuotaUserViewMetric
	GetDiskWriteRate() *string
	SetGPUNodeNumber(v int32) *QuotaUserViewMetric
	GetGPUNodeNumber() *int32
	SetGPUUsageRate(v string) *QuotaUserViewMetric
	GetGPUUsageRate() *string
	SetGpuJobNames(v []*string) *QuotaUserViewMetric
	GetGpuJobNames() []*string
	SetGpuNodeNames(v []*string) *QuotaUserViewMetric
	GetGpuNodeNames() []*string
	SetJobType(v string) *QuotaUserViewMetric
	GetJobType() *string
	SetMemoryUsageRate(v string) *QuotaUserViewMetric
	GetMemoryUsageRate() *string
	SetNetworkInputRate(v string) *QuotaUserViewMetric
	GetNetworkInputRate() *string
	SetNetworkOutputRate(v string) *QuotaUserViewMetric
	GetNetworkOutputRate() *string
	SetNodeNames(v []*string) *QuotaUserViewMetric
	GetNodeNames() []*string
	SetRequestCPU(v int32) *QuotaUserViewMetric
	GetRequestCPU() *int32
	SetRequestGPU(v int32) *QuotaUserViewMetric
	GetRequestGPU() *int32
	SetRequestMemory(v int64) *QuotaUserViewMetric
	GetRequestMemory() *int64
	SetTotalCPU(v int32) *QuotaUserViewMetric
	GetTotalCPU() *int32
	SetTotalGPU(v int32) *QuotaUserViewMetric
	GetTotalGPU() *int32
	SetTotalMemory(v int64) *QuotaUserViewMetric
	GetTotalMemory() *int64
	SetUserId(v string) *QuotaUserViewMetric
	GetUserId() *string
}

type QuotaUserViewMetric struct {
	// The number of CPU nodes.
	//
	// example:
	//
	// 2
	CPUNodeNumber *int32 `json:"CPUNodeNumber,omitempty" xml:"CPUNodeNumber,omitempty"`
	// The CPU utilization.
	//
	// example:
	//
	// 59
	CPUUsageRate *string `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	// The CPU jobs.
	CpuJobNames []*string `json:"CpuJobNames,omitempty" xml:"CpuJobNames,omitempty" type:"Repeated"`
	// The CPU node list.
	CpuNodeNames []*string `json:"CpuNodeNames,omitempty" xml:"CpuNodeNames,omitempty" type:"Repeated"`
	// The disk read rate.
	//
	// example:
	//
	// 22
	DiskReadRate *string `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	// The disk write rate.
	//
	// example:
	//
	// 22
	DiskWriteRate *string `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	// The number of GPU nodes.
	//
	// example:
	//
	// 1
	GPUNodeNumber *int32 `json:"GPUNodeNumber,omitempty" xml:"GPUNodeNumber,omitempty"`
	// The GPU utilization.
	//
	// example:
	//
	// 10
	GPUUsageRate *string `json:"GPUUsageRate,omitempty" xml:"GPUUsageRate,omitempty"`
	// The GPU jobs.
	GpuJobNames []*string `json:"GpuJobNames,omitempty" xml:"GpuJobNames,omitempty" type:"Repeated"`
	// The GPU node list.
	GpuNodeNames []*string `json:"GpuNodeNames,omitempty" xml:"GpuNodeNames,omitempty" type:"Repeated"`
	// The job type.
	//
	// example:
	//
	// PyTorch
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The memory utilization.
	//
	// example:
	//
	// 20
	MemoryUsageRate *string `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	// The network input rate.
	//
	// example:
	//
	// 1
	NetworkInputRate *string `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	// The network output rate.
	//
	// example:
	//
	// 1
	NetworkOutputRate *string `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	// The node list.
	NodeNames []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	// The number of requested CPU cores.
	//
	// example:
	//
	// 100
	RequestCPU *int32 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The number of requested GPU cores.
	//
	// example:
	//
	// 10
	RequestGPU *int32 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// The requested memory in KB.
	//
	// example:
	//
	// 102400000
	RequestMemory *int64 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// The total number of CPU cores.
	//
	// example:
	//
	// 1000
	TotalCPU *int32 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// The total number of GPUs.
	//
	// example:
	//
	// 1
	TotalGPU *int32 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	// The total memory in KB.
	//
	// example:
	//
	// 10240
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuotaUserViewMetric) String() string {
	return dara.Prettify(s)
}

func (s QuotaUserViewMetric) GoString() string {
	return s.String()
}

func (s *QuotaUserViewMetric) GetCPUNodeNumber() *int32 {
	return s.CPUNodeNumber
}

func (s *QuotaUserViewMetric) GetCPUUsageRate() *string {
	return s.CPUUsageRate
}

func (s *QuotaUserViewMetric) GetCpuJobNames() []*string {
	return s.CpuJobNames
}

func (s *QuotaUserViewMetric) GetCpuNodeNames() []*string {
	return s.CpuNodeNames
}

func (s *QuotaUserViewMetric) GetDiskReadRate() *string {
	return s.DiskReadRate
}

func (s *QuotaUserViewMetric) GetDiskWriteRate() *string {
	return s.DiskWriteRate
}

func (s *QuotaUserViewMetric) GetGPUNodeNumber() *int32 {
	return s.GPUNodeNumber
}

func (s *QuotaUserViewMetric) GetGPUUsageRate() *string {
	return s.GPUUsageRate
}

func (s *QuotaUserViewMetric) GetGpuJobNames() []*string {
	return s.GpuJobNames
}

func (s *QuotaUserViewMetric) GetGpuNodeNames() []*string {
	return s.GpuNodeNames
}

func (s *QuotaUserViewMetric) GetJobType() *string {
	return s.JobType
}

func (s *QuotaUserViewMetric) GetMemoryUsageRate() *string {
	return s.MemoryUsageRate
}

func (s *QuotaUserViewMetric) GetNetworkInputRate() *string {
	return s.NetworkInputRate
}

func (s *QuotaUserViewMetric) GetNetworkOutputRate() *string {
	return s.NetworkOutputRate
}

func (s *QuotaUserViewMetric) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *QuotaUserViewMetric) GetRequestCPU() *int32 {
	return s.RequestCPU
}

func (s *QuotaUserViewMetric) GetRequestGPU() *int32 {
	return s.RequestGPU
}

func (s *QuotaUserViewMetric) GetRequestMemory() *int64 {
	return s.RequestMemory
}

func (s *QuotaUserViewMetric) GetTotalCPU() *int32 {
	return s.TotalCPU
}

func (s *QuotaUserViewMetric) GetTotalGPU() *int32 {
	return s.TotalGPU
}

func (s *QuotaUserViewMetric) GetTotalMemory() *int64 {
	return s.TotalMemory
}

func (s *QuotaUserViewMetric) GetUserId() *string {
	return s.UserId
}

func (s *QuotaUserViewMetric) SetCPUNodeNumber(v int32) *QuotaUserViewMetric {
	s.CPUNodeNumber = &v
	return s
}

func (s *QuotaUserViewMetric) SetCPUUsageRate(v string) *QuotaUserViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetCpuJobNames(v []*string) *QuotaUserViewMetric {
	s.CpuJobNames = v
	return s
}

func (s *QuotaUserViewMetric) SetCpuNodeNames(v []*string) *QuotaUserViewMetric {
	s.CpuNodeNames = v
	return s
}

func (s *QuotaUserViewMetric) SetDiskReadRate(v string) *QuotaUserViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetDiskWriteRate(v string) *QuotaUserViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetGPUNodeNumber(v int32) *QuotaUserViewMetric {
	s.GPUNodeNumber = &v
	return s
}

func (s *QuotaUserViewMetric) SetGPUUsageRate(v string) *QuotaUserViewMetric {
	s.GPUUsageRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetGpuJobNames(v []*string) *QuotaUserViewMetric {
	s.GpuJobNames = v
	return s
}

func (s *QuotaUserViewMetric) SetGpuNodeNames(v []*string) *QuotaUserViewMetric {
	s.GpuNodeNames = v
	return s
}

func (s *QuotaUserViewMetric) SetJobType(v string) *QuotaUserViewMetric {
	s.JobType = &v
	return s
}

func (s *QuotaUserViewMetric) SetMemoryUsageRate(v string) *QuotaUserViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetNetworkInputRate(v string) *QuotaUserViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetNetworkOutputRate(v string) *QuotaUserViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *QuotaUserViewMetric) SetNodeNames(v []*string) *QuotaUserViewMetric {
	s.NodeNames = v
	return s
}

func (s *QuotaUserViewMetric) SetRequestCPU(v int32) *QuotaUserViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *QuotaUserViewMetric) SetRequestGPU(v int32) *QuotaUserViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *QuotaUserViewMetric) SetRequestMemory(v int64) *QuotaUserViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *QuotaUserViewMetric) SetTotalCPU(v int32) *QuotaUserViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *QuotaUserViewMetric) SetTotalGPU(v int32) *QuotaUserViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *QuotaUserViewMetric) SetTotalMemory(v int64) *QuotaUserViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *QuotaUserViewMetric) SetUserId(v string) *QuotaUserViewMetric {
	s.UserId = &v
	return s
}

func (s *QuotaUserViewMetric) Validate() error {
	return dara.Validate(s)
}
