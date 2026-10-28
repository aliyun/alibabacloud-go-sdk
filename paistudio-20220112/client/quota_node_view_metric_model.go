// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaNodeViewMetric interface {
	dara.Model
	String() string
	GoString() string
	SetCPUUsageRate(v string) *QuotaNodeViewMetric
	GetCPUUsageRate() *string
	SetCreatedTime(v string) *QuotaNodeViewMetric
	GetCreatedTime() *string
	SetDiskReadRate(v string) *QuotaNodeViewMetric
	GetDiskReadRate() *string
	SetDiskWriteRate(v string) *QuotaNodeViewMetric
	GetDiskWriteRate() *string
	SetGPUType(v string) *QuotaNodeViewMetric
	GetGPUType() *string
	SetMemoryUsageRate(v string) *QuotaNodeViewMetric
	GetMemoryUsageRate() *string
	SetNetworkInputRate(v string) *QuotaNodeViewMetric
	GetNetworkInputRate() *string
	SetNetworkOutputRate(v string) *QuotaNodeViewMetric
	GetNetworkOutputRate() *string
	SetNodeID(v string) *QuotaNodeViewMetric
	GetNodeID() *string
	SetNodeStatus(v string) *QuotaNodeViewMetric
	GetNodeStatus() *string
	SetNodeType(v string) *QuotaNodeViewMetric
	GetNodeType() *string
	SetQuotaId(v string) *QuotaNodeViewMetric
	GetQuotaId() *string
	SetRequestCPU(v int64) *QuotaNodeViewMetric
	GetRequestCPU() *int64
	SetRequestGPU(v int64) *QuotaNodeViewMetric
	GetRequestGPU() *int64
	SetRequestMemory(v int64) *QuotaNodeViewMetric
	GetRequestMemory() *int64
	SetTaskIdMap(v map[string]interface{}) *QuotaNodeViewMetric
	GetTaskIdMap() map[string]interface{}
	SetTotalCPU(v int64) *QuotaNodeViewMetric
	GetTotalCPU() *int64
	SetTotalGPU(v int64) *QuotaNodeViewMetric
	GetTotalGPU() *int64
	SetTotalMemory(v int64) *QuotaNodeViewMetric
	GetTotalMemory() *int64
	SetTotalTasks(v int64) *QuotaNodeViewMetric
	GetTotalTasks() *int64
	SetUserIDs(v []*string) *QuotaNodeViewMetric
	GetUserIDs() []*string
	SetUserNumber(v string) *QuotaNodeViewMetric
	GetUserNumber() *string
}

type QuotaNodeViewMetric struct {
	// The CPU utilization.
	//
	// example:
	//
	// 59
	CPUUsageRate *string `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	// The node import time.
	//
	// example:
	//
	// 2022-08-25T15:20:56Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
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
	// The GPU model.
	//
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
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
	// The node ID.
	//
	// example:
	//
	// -i121212node
	NodeID *string `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	// The node status.
	//
	// example:
	//
	// Ready
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The GPU type.
	//
	// example:
	//
	// ecs.g6.large
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The resource quota ID.
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The used CPU.
	//
	// example:
	//
	// 100
	RequestCPU *int64 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The used GPU.
	//
	// example:
	//
	// 10
	RequestGPU *int64 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// The used memory.
	//
	// example:
	//
	// 102400000
	RequestMemory *int64 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// The node tasks.
	TaskIdMap map[string]interface{} `json:"TaskIdMap,omitempty" xml:"TaskIdMap,omitempty"`
	// The total CPU.
	//
	// example:
	//
	// 1000
	TotalCPU *int64 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// The total GPU.
	//
	// example:
	//
	// 1
	TotalGPU *int64 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	// The total memory.
	//
	// example:
	//
	// 10240
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	// The total number of node tasks.
	//
	// example:
	//
	// 1
	TotalTasks *int64 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	// The user list.
	UserIDs []*string `json:"UserIDs,omitempty" xml:"UserIDs,omitempty" type:"Repeated"`
	// The total number of users.
	//
	// example:
	//
	// 1
	UserNumber *string `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s QuotaNodeViewMetric) String() string {
	return dara.Prettify(s)
}

func (s QuotaNodeViewMetric) GoString() string {
	return s.String()
}

func (s *QuotaNodeViewMetric) GetCPUUsageRate() *string {
	return s.CPUUsageRate
}

func (s *QuotaNodeViewMetric) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *QuotaNodeViewMetric) GetDiskReadRate() *string {
	return s.DiskReadRate
}

func (s *QuotaNodeViewMetric) GetDiskWriteRate() *string {
	return s.DiskWriteRate
}

func (s *QuotaNodeViewMetric) GetGPUType() *string {
	return s.GPUType
}

func (s *QuotaNodeViewMetric) GetMemoryUsageRate() *string {
	return s.MemoryUsageRate
}

func (s *QuotaNodeViewMetric) GetNetworkInputRate() *string {
	return s.NetworkInputRate
}

func (s *QuotaNodeViewMetric) GetNetworkOutputRate() *string {
	return s.NetworkOutputRate
}

func (s *QuotaNodeViewMetric) GetNodeID() *string {
	return s.NodeID
}

func (s *QuotaNodeViewMetric) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *QuotaNodeViewMetric) GetNodeType() *string {
	return s.NodeType
}

func (s *QuotaNodeViewMetric) GetQuotaId() *string {
	return s.QuotaId
}

func (s *QuotaNodeViewMetric) GetRequestCPU() *int64 {
	return s.RequestCPU
}

func (s *QuotaNodeViewMetric) GetRequestGPU() *int64 {
	return s.RequestGPU
}

func (s *QuotaNodeViewMetric) GetRequestMemory() *int64 {
	return s.RequestMemory
}

func (s *QuotaNodeViewMetric) GetTaskIdMap() map[string]interface{} {
	return s.TaskIdMap
}

func (s *QuotaNodeViewMetric) GetTotalCPU() *int64 {
	return s.TotalCPU
}

func (s *QuotaNodeViewMetric) GetTotalGPU() *int64 {
	return s.TotalGPU
}

func (s *QuotaNodeViewMetric) GetTotalMemory() *int64 {
	return s.TotalMemory
}

func (s *QuotaNodeViewMetric) GetTotalTasks() *int64 {
	return s.TotalTasks
}

func (s *QuotaNodeViewMetric) GetUserIDs() []*string {
	return s.UserIDs
}

func (s *QuotaNodeViewMetric) GetUserNumber() *string {
	return s.UserNumber
}

func (s *QuotaNodeViewMetric) SetCPUUsageRate(v string) *QuotaNodeViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetCreatedTime(v string) *QuotaNodeViewMetric {
	s.CreatedTime = &v
	return s
}

func (s *QuotaNodeViewMetric) SetDiskReadRate(v string) *QuotaNodeViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetDiskWriteRate(v string) *QuotaNodeViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetGPUType(v string) *QuotaNodeViewMetric {
	s.GPUType = &v
	return s
}

func (s *QuotaNodeViewMetric) SetMemoryUsageRate(v string) *QuotaNodeViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNetworkInputRate(v string) *QuotaNodeViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNetworkOutputRate(v string) *QuotaNodeViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNodeID(v string) *QuotaNodeViewMetric {
	s.NodeID = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNodeStatus(v string) *QuotaNodeViewMetric {
	s.NodeStatus = &v
	return s
}

func (s *QuotaNodeViewMetric) SetNodeType(v string) *QuotaNodeViewMetric {
	s.NodeType = &v
	return s
}

func (s *QuotaNodeViewMetric) SetQuotaId(v string) *QuotaNodeViewMetric {
	s.QuotaId = &v
	return s
}

func (s *QuotaNodeViewMetric) SetRequestCPU(v int64) *QuotaNodeViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *QuotaNodeViewMetric) SetRequestGPU(v int64) *QuotaNodeViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *QuotaNodeViewMetric) SetRequestMemory(v int64) *QuotaNodeViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *QuotaNodeViewMetric) SetTaskIdMap(v map[string]interface{}) *QuotaNodeViewMetric {
	s.TaskIdMap = v
	return s
}

func (s *QuotaNodeViewMetric) SetTotalCPU(v int64) *QuotaNodeViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *QuotaNodeViewMetric) SetTotalGPU(v int64) *QuotaNodeViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *QuotaNodeViewMetric) SetTotalMemory(v int64) *QuotaNodeViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *QuotaNodeViewMetric) SetTotalTasks(v int64) *QuotaNodeViewMetric {
	s.TotalTasks = &v
	return s
}

func (s *QuotaNodeViewMetric) SetUserIDs(v []*string) *QuotaNodeViewMetric {
	s.UserIDs = v
	return s
}

func (s *QuotaNodeViewMetric) SetUserNumber(v string) *QuotaNodeViewMetric {
	s.UserNumber = &v
	return s
}

func (s *QuotaNodeViewMetric) Validate() error {
	return dara.Validate(s)
}
