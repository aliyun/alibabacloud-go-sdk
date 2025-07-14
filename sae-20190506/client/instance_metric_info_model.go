// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceMetricInfo interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationID(v string) *InstanceMetricInfo
	GetApplicationID() *string
	SetCpuPercent(v float32) *InstanceMetricInfo
	GetCpuPercent() *float32
	SetCpuQuotaPercent(v float32) *InstanceMetricInfo
	GetCpuQuotaPercent() *float32
	SetInstanceID(v string) *InstanceMetricInfo
	GetInstanceID() *string
	SetMemoryLimitMB(v float32) *InstanceMetricInfo
	GetMemoryLimitMB() *float32
	SetMemoryUsageMB(v float32) *InstanceMetricInfo
	GetMemoryUsageMB() *float32
	SetTimestamp(v int64) *InstanceMetricInfo
	GetTimestamp() *int64
}

type InstanceMetricInfo struct {
	// example:
	//
	// a03aa9f9-3d32-4655-8394-05fd10dcbd8a
	ApplicationID *string `json:"applicationID,omitempty" xml:"applicationID,omitempty"`
	// example:
	//
	// 1.98
	CpuPercent *float32 `json:"cpuPercent,omitempty" xml:"cpuPercent,omitempty"`
	// example:
	//
	// 35.0
	CpuQuotaPercent *float32 `json:"cpuQuotaPercent,omitempty" xml:"cpuQuotaPercent,omitempty"`
	// example:
	//
	// c-6498f0fe-33bb4f9249b54789a023
	InstanceID *string `json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	// example:
	//
	// 512.0
	MemoryLimitMB *float32 `json:"memoryLimitMB,omitempty" xml:"memoryLimitMB,omitempty"`
	// example:
	//
	// 8.81
	MemoryUsageMB *float32 `json:"memoryUsageMB,omitempty" xml:"memoryUsageMB,omitempty"`
	// example:
	//
	// 1686568800000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s InstanceMetricInfo) String() string {
	return dara.Prettify(s)
}

func (s InstanceMetricInfo) GoString() string {
	return s.String()
}

func (s *InstanceMetricInfo) GetApplicationID() *string {
	return s.ApplicationID
}

func (s *InstanceMetricInfo) GetCpuPercent() *float32 {
	return s.CpuPercent
}

func (s *InstanceMetricInfo) GetCpuQuotaPercent() *float32 {
	return s.CpuQuotaPercent
}

func (s *InstanceMetricInfo) GetInstanceID() *string {
	return s.InstanceID
}

func (s *InstanceMetricInfo) GetMemoryLimitMB() *float32 {
	return s.MemoryLimitMB
}

func (s *InstanceMetricInfo) GetMemoryUsageMB() *float32 {
	return s.MemoryUsageMB
}

func (s *InstanceMetricInfo) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *InstanceMetricInfo) SetApplicationID(v string) *InstanceMetricInfo {
	s.ApplicationID = &v
	return s
}

func (s *InstanceMetricInfo) SetCpuPercent(v float32) *InstanceMetricInfo {
	s.CpuPercent = &v
	return s
}

func (s *InstanceMetricInfo) SetCpuQuotaPercent(v float32) *InstanceMetricInfo {
	s.CpuQuotaPercent = &v
	return s
}

func (s *InstanceMetricInfo) SetInstanceID(v string) *InstanceMetricInfo {
	s.InstanceID = &v
	return s
}

func (s *InstanceMetricInfo) SetMemoryLimitMB(v float32) *InstanceMetricInfo {
	s.MemoryLimitMB = &v
	return s
}

func (s *InstanceMetricInfo) SetMemoryUsageMB(v float32) *InstanceMetricInfo {
	s.MemoryUsageMB = &v
	return s
}

func (s *InstanceMetricInfo) SetTimestamp(v int64) *InstanceMetricInfo {
	s.Timestamp = &v
	return s
}

func (s *InstanceMetricInfo) Validate() error {
	return dara.Validate(s)
}
