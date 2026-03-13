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
	CPUUsageRate      *string `json:"CPUUsageRate,omitempty" xml:"CPUUsageRate,omitempty"`
	CreatedTime       *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DiskReadRate      *string `json:"DiskReadRate,omitempty" xml:"DiskReadRate,omitempty"`
	DiskWriteRate     *string `json:"DiskWriteRate,omitempty" xml:"DiskWriteRate,omitempty"`
	GPUType           *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	MachineGroupID    *string `json:"MachineGroupID,omitempty" xml:"MachineGroupID,omitempty"`
	MemoryUsageRate   *string `json:"MemoryUsageRate,omitempty" xml:"MemoryUsageRate,omitempty"`
	NetworkInputRate  *string `json:"NetworkInputRate,omitempty" xml:"NetworkInputRate,omitempty"`
	NetworkOutputRate *string `json:"NetworkOutputRate,omitempty" xml:"NetworkOutputRate,omitempty"`
	// example:
	//
	// -i121212node
	NodeID        *string                `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
	NodeStatus    *string                `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType      *string                `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	RequestCPU    *int64                 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU    *int64                 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory *int64                 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	TaskIdMap     map[string]interface{} `json:"TaskIdMap,omitempty" xml:"TaskIdMap,omitempty"`
	TotalCPU      *int64                 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	TotalGPU      *int64                 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	TotalMemory   *int64                 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	TotalTasks    *int64                 `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	UserIDs       []*string              `json:"UserIDs,omitempty" xml:"UserIDs,omitempty" type:"Repeated"`
	UserNumber    *string                `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
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
