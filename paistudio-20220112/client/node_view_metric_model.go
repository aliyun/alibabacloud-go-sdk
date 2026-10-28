// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeViewMetric interface {
	dara.Model
	String() string
	GoString() string
	SetCPUUsageRate(v string) *NodeViewMetric
	GetCPUUsageRate() *string
	SetCreatedTime(v string) *NodeViewMetric
	GetCreatedTime() *string
	SetDiskReadRate(v string) *NodeViewMetric
	GetDiskReadRate() *string
	SetDiskWriteRate(v string) *NodeViewMetric
	GetDiskWriteRate() *string
	SetGPUType(v string) *NodeViewMetric
	GetGPUType() *string
	SetMachineGroupID(v string) *NodeViewMetric
	GetMachineGroupID() *string
	SetMemoryUsageRate(v string) *NodeViewMetric
	GetMemoryUsageRate() *string
	SetNetworkInputRate(v string) *NodeViewMetric
	GetNetworkInputRate() *string
	SetNetworkOutputRate(v string) *NodeViewMetric
	GetNetworkOutputRate() *string
	SetNodeID(v string) *NodeViewMetric
	GetNodeID() *string
	SetNodeStatus(v string) *NodeViewMetric
	GetNodeStatus() *string
	SetNodeType(v string) *NodeViewMetric
	GetNodeType() *string
	SetRequestCPU(v int64) *NodeViewMetric
	GetRequestCPU() *int64
	SetRequestGPU(v int64) *NodeViewMetric
	GetRequestGPU() *int64
	SetRequestMemory(v int64) *NodeViewMetric
	GetRequestMemory() *int64
	SetTaskIdMap(v map[string]interface{}) *NodeViewMetric
	GetTaskIdMap() map[string]interface{}
	SetTotalCPU(v int64) *NodeViewMetric
	GetTotalCPU() *int64
	SetTotalGPU(v int64) *NodeViewMetric
	GetTotalGPU() *int64
	SetTotalMemory(v int64) *NodeViewMetric
	GetTotalMemory() *int64
	SetTotalTasks(v int64) *NodeViewMetric
	GetTotalTasks() *int64
	SetUserIDs(v []*string) *NodeViewMetric
	GetUserIDs() []*string
	SetUserNumber(v string) *NodeViewMetric
	GetUserNumber() *string
}

type NodeViewMetric struct {
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
	// The machine group ID.
	//
	// example:
	//
	// mg1
	MachineGroupID *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
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

func (s NodeViewMetric) String() string {
	return dara.Prettify(s)
}

func (s NodeViewMetric) GoString() string {
	return s.String()
}

func (s *NodeViewMetric) GetCPUUsageRate() *string {
	return s.CPUUsageRate
}

func (s *NodeViewMetric) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *NodeViewMetric) GetDiskReadRate() *string {
	return s.DiskReadRate
}

func (s *NodeViewMetric) GetDiskWriteRate() *string {
	return s.DiskWriteRate
}

func (s *NodeViewMetric) GetGPUType() *string {
	return s.GPUType
}

func (s *NodeViewMetric) GetMachineGroupID() *string {
	return s.MachineGroupID
}

func (s *NodeViewMetric) GetMemoryUsageRate() *string {
	return s.MemoryUsageRate
}

func (s *NodeViewMetric) GetNetworkInputRate() *string {
	return s.NetworkInputRate
}

func (s *NodeViewMetric) GetNetworkOutputRate() *string {
	return s.NetworkOutputRate
}

func (s *NodeViewMetric) GetNodeID() *string {
	return s.NodeID
}

func (s *NodeViewMetric) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *NodeViewMetric) GetNodeType() *string {
	return s.NodeType
}

func (s *NodeViewMetric) GetRequestCPU() *int64 {
	return s.RequestCPU
}

func (s *NodeViewMetric) GetRequestGPU() *int64 {
	return s.RequestGPU
}

func (s *NodeViewMetric) GetRequestMemory() *int64 {
	return s.RequestMemory
}

func (s *NodeViewMetric) GetTaskIdMap() map[string]interface{} {
	return s.TaskIdMap
}

func (s *NodeViewMetric) GetTotalCPU() *int64 {
	return s.TotalCPU
}

func (s *NodeViewMetric) GetTotalGPU() *int64 {
	return s.TotalGPU
}

func (s *NodeViewMetric) GetTotalMemory() *int64 {
	return s.TotalMemory
}

func (s *NodeViewMetric) GetTotalTasks() *int64 {
	return s.TotalTasks
}

func (s *NodeViewMetric) GetUserIDs() []*string {
	return s.UserIDs
}

func (s *NodeViewMetric) GetUserNumber() *string {
	return s.UserNumber
}

func (s *NodeViewMetric) SetCPUUsageRate(v string) *NodeViewMetric {
	s.CPUUsageRate = &v
	return s
}

func (s *NodeViewMetric) SetCreatedTime(v string) *NodeViewMetric {
	s.CreatedTime = &v
	return s
}

func (s *NodeViewMetric) SetDiskReadRate(v string) *NodeViewMetric {
	s.DiskReadRate = &v
	return s
}

func (s *NodeViewMetric) SetDiskWriteRate(v string) *NodeViewMetric {
	s.DiskWriteRate = &v
	return s
}

func (s *NodeViewMetric) SetGPUType(v string) *NodeViewMetric {
	s.GPUType = &v
	return s
}

func (s *NodeViewMetric) SetMachineGroupID(v string) *NodeViewMetric {
	s.MachineGroupID = &v
	return s
}

func (s *NodeViewMetric) SetMemoryUsageRate(v string) *NodeViewMetric {
	s.MemoryUsageRate = &v
	return s
}

func (s *NodeViewMetric) SetNetworkInputRate(v string) *NodeViewMetric {
	s.NetworkInputRate = &v
	return s
}

func (s *NodeViewMetric) SetNetworkOutputRate(v string) *NodeViewMetric {
	s.NetworkOutputRate = &v
	return s
}

func (s *NodeViewMetric) SetNodeID(v string) *NodeViewMetric {
	s.NodeID = &v
	return s
}

func (s *NodeViewMetric) SetNodeStatus(v string) *NodeViewMetric {
	s.NodeStatus = &v
	return s
}

func (s *NodeViewMetric) SetNodeType(v string) *NodeViewMetric {
	s.NodeType = &v
	return s
}

func (s *NodeViewMetric) SetRequestCPU(v int64) *NodeViewMetric {
	s.RequestCPU = &v
	return s
}

func (s *NodeViewMetric) SetRequestGPU(v int64) *NodeViewMetric {
	s.RequestGPU = &v
	return s
}

func (s *NodeViewMetric) SetRequestMemory(v int64) *NodeViewMetric {
	s.RequestMemory = &v
	return s
}

func (s *NodeViewMetric) SetTaskIdMap(v map[string]interface{}) *NodeViewMetric {
	s.TaskIdMap = v
	return s
}

func (s *NodeViewMetric) SetTotalCPU(v int64) *NodeViewMetric {
	s.TotalCPU = &v
	return s
}

func (s *NodeViewMetric) SetTotalGPU(v int64) *NodeViewMetric {
	s.TotalGPU = &v
	return s
}

func (s *NodeViewMetric) SetTotalMemory(v int64) *NodeViewMetric {
	s.TotalMemory = &v
	return s
}

func (s *NodeViewMetric) SetTotalTasks(v int64) *NodeViewMetric {
	s.TotalTasks = &v
	return s
}

func (s *NodeViewMetric) SetUserIDs(v []*string) *NodeViewMetric {
	s.UserIDs = v
	return s
}

func (s *NodeViewMetric) SetUserNumber(v string) *NodeViewMetric {
	s.UserNumber = &v
	return s
}

func (s *NodeViewMetric) Validate() error {
	return dara.Validate(s)
}
