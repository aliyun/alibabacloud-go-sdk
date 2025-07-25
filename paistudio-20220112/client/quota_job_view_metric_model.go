// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaJobViewMetric interface {
	dara.Model
	String() string
	GoString() string
	SetCPUUsageRate(v string) *QuotaJobViewMetric
	GetCPUUsageRate() *string
	SetDiskReadRate(v string) *QuotaJobViewMetric
	GetDiskReadRate() *string
	SetDiskWriteRate(v string) *QuotaJobViewMetric
	GetDiskWriteRate() *string
	SetGPUUsageRate(v string) *QuotaJobViewMetric
	GetGPUUsageRate() *string
	SetJobId(v string) *QuotaJobViewMetric
	GetJobId() *string
	SetJobType(v string) *QuotaJobViewMetric
	GetJobType() *string
	SetMemoryUsageRate(v string) *QuotaJobViewMetric
	GetMemoryUsageRate() *string
	SetNetworkInputRate(v string) *QuotaJobViewMetric
	GetNetworkInputRate() *string
	SetNetworkOutputRate(v string) *QuotaJobViewMetric
	GetNetworkOutputRate() *string
	SetNodeNames(v []*string) *QuotaJobViewMetric
	GetNodeNames() []*string
	SetRequestCPU(v int32) *QuotaJobViewMetric
	GetRequestCPU() *int32
	SetRequestGPU(v int32) *QuotaJobViewMetric
	GetRequestGPU() *int32
	SetRequestMemory(v int64) *QuotaJobViewMetric
	GetRequestMemory() *int64
	SetTotalCPU(v int32) *QuotaJobViewMetric
	GetTotalCPU() *int32
	SetTotalGPU(v int32) *QuotaJobViewMetric
	GetTotalGPU() *int32
	SetTotalMemory(v int64) *QuotaJobViewMetric
	GetTotalMemory() *int64
	SetUserId(v string) *QuotaJobViewMetric
	GetUserId() *string
}

type QuotaJobViewMetric struct {
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
	TotalCPU          *int32    `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU          *int32    `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory       *int64    `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuotaJobViewMetric) String() string {
	return dara.Prettify(s)
}

func (s QuotaJobViewMetric) GoString() string {
	return s.String()
}

func (s *QuotaJobViewMetric) GetCPUUsageRate() *string {
	return s.CPUUsageRate
}

func (s *QuotaJobViewMetric) GetDiskReadRate() *string {
	return s.DiskReadRate
}

func (s *QuotaJobViewMetric) GetDiskWriteRate() *string {
	return s.DiskWriteRate
}

func (s *QuotaJobViewMetric) GetGPUUsageRate() *string {
	return s.GPUUsageRate
}

func (s *QuotaJobViewMetric) GetJobId() *string {
	return s.JobId
}

func (s *QuotaJobViewMetric) GetJobType() *string {
	return s.JobType
}

func (s *QuotaJobViewMetric) GetMemoryUsageRate() *string {
	return s.MemoryUsageRate
}

func (s *QuotaJobViewMetric) GetNetworkInputRate() *string {
	return s.NetworkInputRate
}

func (s *QuotaJobViewMetric) GetNetworkOutputRate() *string {
	return s.NetworkOutputRate
}

func (s *QuotaJobViewMetric) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *QuotaJobViewMetric) GetRequestCPU() *int32 {
	return s.RequestCPU
}

func (s *QuotaJobViewMetric) GetRequestGPU() *int32 {
	return s.RequestGPU
}

func (s *QuotaJobViewMetric) GetRequestMemory() *int64 {
	return s.RequestMemory
}

func (s *QuotaJobViewMetric) GetTotalCPU() *int32 {
	return s.TotalCPU
}

func (s *QuotaJobViewMetric) GetTotalGPU() *int32 {
	return s.TotalGPU
}

func (s *QuotaJobViewMetric) GetTotalMemory() *int64 {
	return s.TotalMemory
}

func (s *QuotaJobViewMetric) GetUserId() *string {
	return s.UserId
}

func (s *QuotaJobViewMetric) SetCPUUsageRate(v string) *QuotaJobViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetDiskReadRate(v string) *QuotaJobViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetDiskWriteRate(v string) *QuotaJobViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetGPUUsageRate(v string) *QuotaJobViewMetric {
	s.GPUUsageRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetJobId(v string) *QuotaJobViewMetric {
	s.JobId = &v
	return s
}

func (s *QuotaJobViewMetric) SetJobType(v string) *QuotaJobViewMetric {
	s.JobType = &v
	return s
}

func (s *QuotaJobViewMetric) SetMemoryUsageRate(v string) *QuotaJobViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetNetworkInputRate(v string) *QuotaJobViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetNetworkOutputRate(v string) *QuotaJobViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *QuotaJobViewMetric) SetNodeNames(v []*string) *QuotaJobViewMetric {
	s.NodeNames = v
	return s
}

func (s *QuotaJobViewMetric) SetRequestCPU(v int32) *QuotaJobViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *QuotaJobViewMetric) SetRequestGPU(v int32) *QuotaJobViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *QuotaJobViewMetric) SetRequestMemory(v int64) *QuotaJobViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *QuotaJobViewMetric) SetTotalCPU(v int32) *QuotaJobViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *QuotaJobViewMetric) SetTotalGPU(v int32) *QuotaJobViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *QuotaJobViewMetric) SetTotalMemory(v int64) *QuotaJobViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *QuotaJobViewMetric) SetUserId(v string) *QuotaJobViewMetric {
	s.UserId = &v
	return s
}

func (s *QuotaJobViewMetric) Validate() error {
	return dara.Validate(s)
}
