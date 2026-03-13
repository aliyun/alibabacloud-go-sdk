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
	CPUUsageRate      *string   `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	DiskReadRate      *string   `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string   `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUUsageRate      *string   `json:"GPUUsageRate,omitempty" xml:"GPUUsageRate,omitempty"`
	JobId             *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType           *string   `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MemoryUsageRate   *string   `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string   `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string   `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	NodeNames         []*string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	RequestCPU        *int32    `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU        *int32    `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory     *int64    `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// example:
	//
	// rg17tmvwiokhzaxg
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
	TotalCPU        *int32  `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU        *int32  `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory     *int64  `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
