// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStaticsInfo interface {
	dara.Model
	String() string
	GoString() string
	SetActiveCPUUsage(v int64) *StaticsInfo
	GetActiveCPUUsage() *int64
	SetCost(v float32) *StaticsInfo
	GetCost() *float32
	SetDiskUsage(v int64) *StaticsInfo
	GetDiskUsage() *int64
	SetFunctionName(v string) *StaticsInfo
	GetFunctionName() *string
	SetGpuUsage(v int64) *StaticsInfo
	GetGpuUsage() *int64
	SetIdleCPUUsage(v int64) *StaticsInfo
	GetIdleCPUUsage() *int64
	SetInstanceTrafficOut(v int64) *StaticsInfo
	GetInstanceTrafficOut() *int64
	SetInvocations(v int64) *StaticsInfo
	GetInvocations() *int64
	SetInvokeCDNOut(v int64) *StaticsInfo
	GetInvokeCDNOut() *int64
	SetInvokeInternetOut(v int64) *StaticsInfo
	GetInvokeInternetOut() *int64
	SetMemoryUsage(v int64) *StaticsInfo
	GetMemoryUsage() *int64
	SetRegion(v string) *StaticsInfo
	GetRegion() *string
	SetServiceName(v string) *StaticsInfo
	GetServiceName() *string
}

type StaticsInfo struct {
	ActiveCPUUsage     *int64   `json:"activeCPUUsage,omitempty" xml:"activeCPUUsage,omitempty"`
	Cost               *float32 `json:"cost,omitempty" xml:"cost,omitempty"`
	DiskUsage          *int64   `json:"diskUsage,omitempty" xml:"diskUsage,omitempty"`
	FunctionName       *string  `json:"functionName,omitempty" xml:"functionName,omitempty"`
	GpuUsage           *int64   `json:"gpuUsage,omitempty" xml:"gpuUsage,omitempty"`
	IdleCPUUsage       *int64   `json:"idleCPUUsage,omitempty" xml:"idleCPUUsage,omitempty"`
	InstanceTrafficOut *int64   `json:"instanceTrafficOut,omitempty" xml:"instanceTrafficOut,omitempty"`
	Invocations        *int64   `json:"invocations,omitempty" xml:"invocations,omitempty"`
	InvokeCDNOut       *int64   `json:"invokeCDNOut,omitempty" xml:"invokeCDNOut,omitempty"`
	InvokeInternetOut  *int64   `json:"invokeInternetOut,omitempty" xml:"invokeInternetOut,omitempty"`
	MemoryUsage        *int64   `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	Region             *string  `json:"region,omitempty" xml:"region,omitempty"`
	ServiceName        *string  `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s StaticsInfo) String() string {
	return dara.Prettify(s)
}

func (s StaticsInfo) GoString() string {
	return s.String()
}

func (s *StaticsInfo) GetActiveCPUUsage() *int64 {
	return s.ActiveCPUUsage
}

func (s *StaticsInfo) GetCost() *float32 {
	return s.Cost
}

func (s *StaticsInfo) GetDiskUsage() *int64 {
	return s.DiskUsage
}

func (s *StaticsInfo) GetFunctionName() *string {
	return s.FunctionName
}

func (s *StaticsInfo) GetGpuUsage() *int64 {
	return s.GpuUsage
}

func (s *StaticsInfo) GetIdleCPUUsage() *int64 {
	return s.IdleCPUUsage
}

func (s *StaticsInfo) GetInstanceTrafficOut() *int64 {
	return s.InstanceTrafficOut
}

func (s *StaticsInfo) GetInvocations() *int64 {
	return s.Invocations
}

func (s *StaticsInfo) GetInvokeCDNOut() *int64 {
	return s.InvokeCDNOut
}

func (s *StaticsInfo) GetInvokeInternetOut() *int64 {
	return s.InvokeInternetOut
}

func (s *StaticsInfo) GetMemoryUsage() *int64 {
	return s.MemoryUsage
}

func (s *StaticsInfo) GetRegion() *string {
	return s.Region
}

func (s *StaticsInfo) GetServiceName() *string {
	return s.ServiceName
}

func (s *StaticsInfo) SetActiveCPUUsage(v int64) *StaticsInfo {
	s.ActiveCPUUsage = &v
	return s
}

func (s *StaticsInfo) SetCost(v float32) *StaticsInfo {
	s.Cost = &v
	return s
}

func (s *StaticsInfo) SetDiskUsage(v int64) *StaticsInfo {
	s.DiskUsage = &v
	return s
}

func (s *StaticsInfo) SetFunctionName(v string) *StaticsInfo {
	s.FunctionName = &v
	return s
}

func (s *StaticsInfo) SetGpuUsage(v int64) *StaticsInfo {
	s.GpuUsage = &v
	return s
}

func (s *StaticsInfo) SetIdleCPUUsage(v int64) *StaticsInfo {
	s.IdleCPUUsage = &v
	return s
}

func (s *StaticsInfo) SetInstanceTrafficOut(v int64) *StaticsInfo {
	s.InstanceTrafficOut = &v
	return s
}

func (s *StaticsInfo) SetInvocations(v int64) *StaticsInfo {
	s.Invocations = &v
	return s
}

func (s *StaticsInfo) SetInvokeCDNOut(v int64) *StaticsInfo {
	s.InvokeCDNOut = &v
	return s
}

func (s *StaticsInfo) SetInvokeInternetOut(v int64) *StaticsInfo {
	s.InvokeInternetOut = &v
	return s
}

func (s *StaticsInfo) SetMemoryUsage(v int64) *StaticsInfo {
	s.MemoryUsage = &v
	return s
}

func (s *StaticsInfo) SetRegion(v string) *StaticsInfo {
	s.Region = &v
	return s
}

func (s *StaticsInfo) SetServiceName(v string) *StaticsInfo {
	s.ServiceName = &v
	return s
}

func (s *StaticsInfo) Validate() error {
	return dara.Validate(s)
}
