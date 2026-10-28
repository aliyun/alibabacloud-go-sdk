// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobViewMetric interface {
	dara.Model
	String() string
	GoString() string
	SetCPUUsageRate(v string) *JobViewMetric
	GetCPUUsageRate() *string
	SetDiskReadRate(v string) *JobViewMetric
	GetDiskReadRate() *string
	SetDiskWriteRate(v string) *JobViewMetric
	GetDiskWriteRate() *string
	SetGPUUsageRate(v string) *JobViewMetric
	GetGPUUsageRate() *string
	SetJobId(v string) *JobViewMetric
	GetJobId() *string
	SetJobType(v string) *JobViewMetric
	GetJobType() *string
	SetMemoryUsageRate(v string) *JobViewMetric
	GetMemoryUsageRate() *string
	SetNetworkInputRate(v string) *JobViewMetric
	GetNetworkInputRate() *string
	SetNetworkOutputRate(v string) *JobViewMetric
	GetNetworkOutputRate() *string
	SetNodeNames(v []*string) *JobViewMetric
	GetNodeNames() []*string
	SetRequestCPU(v int32) *JobViewMetric
	GetRequestCPU() *int32
	SetRequestGPU(v int32) *JobViewMetric
	GetRequestGPU() *int32
	SetRequestMemory(v int64) *JobViewMetric
	GetRequestMemory() *int64
	SetResourceGroupID(v string) *JobViewMetric
	GetResourceGroupID() *string
	SetTotalCPU(v int32) *JobViewMetric
	GetTotalCPU() *int32
	SetTotalGPU(v int32) *JobViewMetric
	GetTotalGPU() *int32
	SetTotalMemory(v int64) *JobViewMetric
	GetTotalMemory() *int64
	SetUserId(v string) *JobViewMetric
	GetUserId() *string
}

type JobViewMetric struct {
	// The CPU utilization.
	//
	// example:
	//
	// 59
	CPUUsageRate *string `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
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
	// The GPU utilization.
	//
	// example:
	//
	// 10
	GPUUsageRate *string `json:"GPUUsageRate,omitempty" xml:"GPUUsageRate,omitempty"`
	// The job ID.
	//
	// example:
	//
	// dsw-12580
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job type, such as DLC or DSW.
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The memory utilization.
	//
	// example:
	//
	// 20
	MemoryUsageRate *string `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	// The network inbound rate.
	//
	// example:
	//
	// 1
	NetworkInputRate *string `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	// The network outbound rate.
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
	// The resource group ID.
	//
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	// The total number of CPU cores.
	//
	// example:
	//
	// 1000
	TotalCPU *int32 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// The total number of GPU cards.
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

func (s JobViewMetric) String() string {
	return dara.Prettify(s)
}

func (s JobViewMetric) GoString() string {
	return s.String()
}

func (s *JobViewMetric) GetCPUUsageRate() *string {
	return s.CPUUsageRate
}

func (s *JobViewMetric) GetDiskReadRate() *string {
	return s.DiskReadRate
}

func (s *JobViewMetric) GetDiskWriteRate() *string {
	return s.DiskWriteRate
}

func (s *JobViewMetric) GetGPUUsageRate() *string {
	return s.GPUUsageRate
}

func (s *JobViewMetric) GetJobId() *string {
	return s.JobId
}

func (s *JobViewMetric) GetJobType() *string {
	return s.JobType
}

func (s *JobViewMetric) GetMemoryUsageRate() *string {
	return s.MemoryUsageRate
}

func (s *JobViewMetric) GetNetworkInputRate() *string {
	return s.NetworkInputRate
}

func (s *JobViewMetric) GetNetworkOutputRate() *string {
	return s.NetworkOutputRate
}

func (s *JobViewMetric) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *JobViewMetric) GetRequestCPU() *int32 {
	return s.RequestCPU
}

func (s *JobViewMetric) GetRequestGPU() *int32 {
	return s.RequestGPU
}

func (s *JobViewMetric) GetRequestMemory() *int64 {
	return s.RequestMemory
}

func (s *JobViewMetric) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *JobViewMetric) GetTotalCPU() *int32 {
	return s.TotalCPU
}

func (s *JobViewMetric) GetTotalGPU() *int32 {
	return s.TotalGPU
}

func (s *JobViewMetric) GetTotalMemory() *int64 {
	return s.TotalMemory
}

func (s *JobViewMetric) GetUserId() *string {
	return s.UserId
}

func (s *JobViewMetric) SetCPUUsageRate(v string) *JobViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *JobViewMetric) SetDiskReadRate(v string) *JobViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *JobViewMetric) SetDiskWriteRate(v string) *JobViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *JobViewMetric) SetGPUUsageRate(v string) *JobViewMetric {
	s.GPUUsageRate = &v
	return s
}

func (s *JobViewMetric) SetJobId(v string) *JobViewMetric {
	s.JobId = &v
	return s
}

func (s *JobViewMetric) SetJobType(v string) *JobViewMetric {
	s.JobType = &v
	return s
}

func (s *JobViewMetric) SetMemoryUsageRate(v string) *JobViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *JobViewMetric) SetNetworkInputRate(v string) *JobViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *JobViewMetric) SetNetworkOutputRate(v string) *JobViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *JobViewMetric) SetNodeNames(v []*string) *JobViewMetric {
	s.NodeNames = v
	return s
}

func (s *JobViewMetric) SetRequestCPU(v int32) *JobViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *JobViewMetric) SetRequestGPU(v int32) *JobViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *JobViewMetric) SetRequestMemory(v int64) *JobViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *JobViewMetric) SetResourceGroupID(v string) *JobViewMetric {
	s.ResourceGroupID = &v
	return s
}

func (s *JobViewMetric) SetTotalCPU(v int32) *JobViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *JobViewMetric) SetTotalGPU(v int32) *JobViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *JobViewMetric) SetTotalMemory(v int64) *JobViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *JobViewMetric) SetUserId(v string) *JobViewMetric {
	s.UserId = &v
	return s
}

func (s *JobViewMetric) Validate() error {
	return dara.Validate(s)
}
