// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserViewMetric interface {
	dara.Model
	String() string
	GoString() string
	SetCPUNodeNumber(v int32) *UserViewMetric
	GetCPUNodeNumber() *int32
	SetCPUUsageRate(v string) *UserViewMetric
	GetCPUUsageRate() *string
	SetCpuJobNames(v []*string) *UserViewMetric
	GetCpuJobNames() []*string
	SetCpuNodeNames(v []*string) *UserViewMetric
	GetCpuNodeNames() []*string
	SetDiskReadRate(v string) *UserViewMetric
	GetDiskReadRate() *string
	SetDiskWriteRate(v string) *UserViewMetric
	GetDiskWriteRate() *string
	SetGPUNodeNumber(v int32) *UserViewMetric
	GetGPUNodeNumber() *int32
	SetGPUUsageRate(v string) *UserViewMetric
	GetGPUUsageRate() *string
	SetGpuJobNames(v []*string) *UserViewMetric
	GetGpuJobNames() []*string
	SetGpuNodeNames(v []*string) *UserViewMetric
	GetGpuNodeNames() []*string
	SetJobType(v string) *UserViewMetric
	GetJobType() *string
	SetMemoryUsageRate(v string) *UserViewMetric
	GetMemoryUsageRate() *string
	SetNetworkInputRate(v string) *UserViewMetric
	GetNetworkInputRate() *string
	SetNetworkOutputRate(v string) *UserViewMetric
	GetNetworkOutputRate() *string
	SetNodeNames(v []*string) *UserViewMetric
	GetNodeNames() []*string
	SetRequestCPU(v int32) *UserViewMetric
	GetRequestCPU() *int32
	SetRequestGPU(v int32) *UserViewMetric
	GetRequestGPU() *int32
	SetRequestMemory(v int64) *UserViewMetric
	GetRequestMemory() *int64
	SetResourceGroupId(v string) *UserViewMetric
	GetResourceGroupId() *string
	SetTotalCPU(v int32) *UserViewMetric
	GetTotalCPU() *int32
	SetTotalGPU(v int32) *UserViewMetric
	GetTotalGPU() *int32
	SetTotalMemory(v int64) *UserViewMetric
	GetTotalMemory() *int64
	SetUserId(v string) *UserViewMetric
	GetUserId() *string
}

type UserViewMetric struct {
	CPUNodeNumber     *int32    `json:"CPUNodeNumber,omitempty" xml:"CPUNodeNumber,omitempty"`
	CPUUsageRate      *string   `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	CpuJobNames       []*string `json:"CpuJobNames,omitempty" xml:"CpuJobNames,omitempty" type:"Repeated"`
	CpuNodeNames      []*string `json:"CpuNodeNames,omitempty" xml:"CpuNodeNames,omitempty" type:"Repeated"`
	DiskReadRate      *string   `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string   `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUNodeNumber     *int32    `json:"GPUNodeNumber,omitempty" xml:"GPUNodeNumber,omitempty"`
	GPUUsageRate      *string   `json:"GPUUsageRate,omitempty" xml:"GPUUsageRate,omitempty"`
	GpuJobNames       []*string `json:"GpuJobNames,omitempty" xml:"GpuJobNames,omitempty" type:"Repeated"`
	GpuNodeNames      []*string `json:"GpuNodeNames,omitempty" xml:"GpuNodeNames,omitempty" type:"Repeated"`
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
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TotalCPU        *int32  `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU        *int32  `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory     *int64  `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UserViewMetric) String() string {
	return dara.Prettify(s)
}

func (s UserViewMetric) GoString() string {
	return s.String()
}

func (s *UserViewMetric) GetCPUNodeNumber() *int32 {
	return s.CPUNodeNumber
}

func (s *UserViewMetric) GetCPUUsageRate() *string {
	return s.CPUUsageRate
}

func (s *UserViewMetric) GetCpuJobNames() []*string {
	return s.CpuJobNames
}

func (s *UserViewMetric) GetCpuNodeNames() []*string {
	return s.CpuNodeNames
}

func (s *UserViewMetric) GetDiskReadRate() *string {
	return s.DiskReadRate
}

func (s *UserViewMetric) GetDiskWriteRate() *string {
	return s.DiskWriteRate
}

func (s *UserViewMetric) GetGPUNodeNumber() *int32 {
	return s.GPUNodeNumber
}

func (s *UserViewMetric) GetGPUUsageRate() *string {
	return s.GPUUsageRate
}

func (s *UserViewMetric) GetGpuJobNames() []*string {
	return s.GpuJobNames
}

func (s *UserViewMetric) GetGpuNodeNames() []*string {
	return s.GpuNodeNames
}

func (s *UserViewMetric) GetJobType() *string {
	return s.JobType
}

func (s *UserViewMetric) GetMemoryUsageRate() *string {
	return s.MemoryUsageRate
}

func (s *UserViewMetric) GetNetworkInputRate() *string {
	return s.NetworkInputRate
}

func (s *UserViewMetric) GetNetworkOutputRate() *string {
	return s.NetworkOutputRate
}

func (s *UserViewMetric) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *UserViewMetric) GetRequestCPU() *int32 {
	return s.RequestCPU
}

func (s *UserViewMetric) GetRequestGPU() *int32 {
	return s.RequestGPU
}

func (s *UserViewMetric) GetRequestMemory() *int64 {
	return s.RequestMemory
}

func (s *UserViewMetric) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UserViewMetric) GetTotalCPU() *int32 {
	return s.TotalCPU
}

func (s *UserViewMetric) GetTotalGPU() *int32 {
	return s.TotalGPU
}

func (s *UserViewMetric) GetTotalMemory() *int64 {
	return s.TotalMemory
}

func (s *UserViewMetric) GetUserId() *string {
	return s.UserId
}

func (s *UserViewMetric) SetCPUNodeNumber(v int32) *UserViewMetric {
	s.CPUNodeNumber = &v
	return s
}

func (s *UserViewMetric) SetCPUUsageRate(v string) *UserViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *UserViewMetric) SetCpuJobNames(v []*string) *UserViewMetric {
	s.CpuJobNames = v
	return s
}

func (s *UserViewMetric) SetCpuNodeNames(v []*string) *UserViewMetric {
	s.CpuNodeNames = v
	return s
}

func (s *UserViewMetric) SetDiskReadRate(v string) *UserViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *UserViewMetric) SetDiskWriteRate(v string) *UserViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *UserViewMetric) SetGPUNodeNumber(v int32) *UserViewMetric {
	s.GPUNodeNumber = &v
	return s
}

func (s *UserViewMetric) SetGPUUsageRate(v string) *UserViewMetric {
	s.GPUUsageRate = &v
	return s
}

func (s *UserViewMetric) SetGpuJobNames(v []*string) *UserViewMetric {
	s.GpuJobNames = v
	return s
}

func (s *UserViewMetric) SetGpuNodeNames(v []*string) *UserViewMetric {
	s.GpuNodeNames = v
	return s
}

func (s *UserViewMetric) SetJobType(v string) *UserViewMetric {
	s.JobType = &v
	return s
}

func (s *UserViewMetric) SetMemoryUsageRate(v string) *UserViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *UserViewMetric) SetNetworkInputRate(v string) *UserViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *UserViewMetric) SetNetworkOutputRate(v string) *UserViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *UserViewMetric) SetNodeNames(v []*string) *UserViewMetric {
	s.NodeNames = v
	return s
}

func (s *UserViewMetric) SetRequestCPU(v int32) *UserViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *UserViewMetric) SetRequestGPU(v int32) *UserViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *UserViewMetric) SetRequestMemory(v int64) *UserViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *UserViewMetric) SetResourceGroupId(v string) *UserViewMetric {
	s.ResourceGroupId = &v
	return s
}

func (s *UserViewMetric) SetTotalCPU(v int32) *UserViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *UserViewMetric) SetTotalGPU(v int32) *UserViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *UserViewMetric) SetTotalMemory(v int64) *UserViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *UserViewMetric) SetUserId(v string) *UserViewMetric {
	s.UserId = &v
	return s
}

func (s *UserViewMetric) Validate() error {
	return dara.Validate(s)
}
