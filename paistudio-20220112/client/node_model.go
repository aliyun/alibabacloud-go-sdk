// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNode interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *Node
	GetAcceleratorType() *string
	SetAllocatableCPU(v string) *Node
	GetAllocatableCPU() *string
	SetAllocatableMemory(v string) *Node
	GetAllocatableMemory() *string
	SetAncestorQuotaWorkloadNum(v int64) *Node
	GetAncestorQuotaWorkloadNum() *int64
	SetAvailabilityZone(v string) *Node
	GetAvailabilityZone() *string
	SetBoundQuotas(v []*QuotaIdName) *Node
	GetBoundQuotas() []*QuotaIdName
	SetCPU(v string) *Node
	GetCPU() *string
	SetCreatorId(v string) *Node
	GetCreatorId() *string
	SetDescendantQuotaWorkloadNum(v int64) *Node
	GetDescendantQuotaWorkloadNum() *int64
	SetGPU(v string) *Node
	GetGPU() *string
	SetGPUMemory(v string) *Node
	GetGPUMemory() *string
	SetGPUType(v string) *Node
	GetGPUType() *string
	SetGmtCreateTime(v string) *Node
	GetGmtCreateTime() *string
	SetGmtExpiredTime(v string) *Node
	GetGmtExpiredTime() *string
	SetGmtModifiedTime(v string) *Node
	GetGmtModifiedTime() *string
	SetHyperZone(v string) *Node
	GetHyperZone() *string
	SetIsBound(v bool) *Node
	GetIsBound() *bool
	SetLimitCPU(v string) *Node
	GetLimitCPU() *string
	SetLimitGPU(v string) *Node
	GetLimitGPU() *string
	SetLimitMemory(v string) *Node
	GetLimitMemory() *string
	SetMachineGroupId(v string) *Node
	GetMachineGroupId() *string
	SetMemory(v string) *Node
	GetMemory() *string
	SetNodeName(v string) *Node
	GetNodeName() *string
	SetNodeStatus(v string) *Node
	GetNodeStatus() *string
	SetNodeType(v string) *Node
	GetNodeType() *string
	SetOrderStatus(v string) *Node
	GetOrderStatus() *string
	SetPodNum(v int64) *Node
	GetPodNum() *int64
	SetReasonCode(v string) *Node
	GetReasonCode() *string
	SetReasonMessage(v string) *Node
	GetReasonMessage() *string
	SetRequestCPU(v string) *Node
	GetRequestCPU() *string
	SetRequestGPU(v string) *Node
	GetRequestGPU() *string
	SetRequestMemory(v string) *Node
	GetRequestMemory() *string
	SetResourceGroupId(v string) *Node
	GetResourceGroupId() *string
	SetResourceGroupName(v string) *Node
	GetResourceGroupName() *string
	SetSelfQuotaWorkloadNum(v int64) *Node
	GetSelfQuotaWorkloadNum() *int64
	SetSystemReservedCPU(v string) *Node
	GetSystemReservedCPU() *string
	SetSystemReservedMemory(v string) *Node
	GetSystemReservedMemory() *string
	SetUsers(v []*UserInfo) *Node
	GetUsers() []*UserInfo
	SetWorkloadNum(v int64) *Node
	GetWorkloadNum() *int64
}

type Node struct {
	AcceleratorType            *string        `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	AllocatableCPU             *string        `json:"AllocatableCPU,omitempty" xml:"AllocatableCPU,omitempty"`
	AllocatableMemory          *string        `json:"AllocatableMemory,omitempty" xml:"AllocatableMemory,omitempty"`
	AncestorQuotaWorkloadNum   *int64         `json:"AncestorQuotaWorkloadNum,omitempty" xml:"AncestorQuotaWorkloadNum,omitempty"`
	AvailabilityZone           *string        `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	BoundQuotas                []*QuotaIdName `json:"BoundQuotas,omitempty" xml:"BoundQuotas,omitempty" type:"Repeated"`
	CPU                        *string        `json:"CPU,omitempty" xml:"CPU,omitempty"`
	CreatorId                  *string        `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DescendantQuotaWorkloadNum *int64         `json:"DescendantQuotaWorkloadNum,omitempty" xml:"DescendantQuotaWorkloadNum,omitempty"`
	GPU                        *string        `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUMemory                  *string        `json:"GPUMemory,omitempty" xml:"GPUMemory,omitempty"`
	GPUType                    *string        `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	GmtCreateTime              *string        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtExpiredTime             *string        `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
	GmtModifiedTime            *string        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HyperZone                  *string        `json:"HyperZone,omitempty" xml:"HyperZone,omitempty"`
	IsBound                    *bool          `json:"IsBound,omitempty" xml:"IsBound,omitempty"`
	LimitCPU                   *string        `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	LimitGPU                   *string        `json:"LimitGPU,omitempty" xml:"LimitGPU,omitempty"`
	LimitMemory                *string        `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	MachineGroupId             *string        `json:"MachineGroupId,omitempty" xml:"MachineGroupId,omitempty"`
	Memory                     *string        `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeName                   *string        `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeStatus                 *string        `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType                   *string        `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OrderStatus                *string        `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	PodNum                     *int64         `json:"PodNum,omitempty" xml:"PodNum,omitempty"`
	ReasonCode                 *string        `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage              *string        `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestCPU                 *string        `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	RequestGPU                 *string        `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	RequestMemory              *string        `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	ResourceGroupId            *string        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceGroupName          *string        `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	SelfQuotaWorkloadNum       *int64         `json:"SelfQuotaWorkloadNum,omitempty" xml:"SelfQuotaWorkloadNum,omitempty"`
	SystemReservedCPU          *string        `json:"SystemReservedCPU,omitempty" xml:"SystemReservedCPU,omitempty"`
	SystemReservedMemory       *string        `json:"SystemReservedMemory,omitempty" xml:"SystemReservedMemory,omitempty"`
	Users                      []*UserInfo    `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	WorkloadNum                *int64         `json:"WorkloadNum,omitempty" xml:"WorkloadNum,omitempty"`
}

func (s Node) String() string {
	return dara.Prettify(s)
}

func (s Node) GoString() string {
	return s.String()
}

func (s *Node) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *Node) GetAllocatableCPU() *string {
	return s.AllocatableCPU
}

func (s *Node) GetAllocatableMemory() *string {
	return s.AllocatableMemory
}

func (s *Node) GetAncestorQuotaWorkloadNum() *int64 {
	return s.AncestorQuotaWorkloadNum
}

func (s *Node) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *Node) GetBoundQuotas() []*QuotaIdName {
	return s.BoundQuotas
}

func (s *Node) GetCPU() *string {
	return s.CPU
}

func (s *Node) GetCreatorId() *string {
	return s.CreatorId
}

func (s *Node) GetDescendantQuotaWorkloadNum() *int64 {
	return s.DescendantQuotaWorkloadNum
}

func (s *Node) GetGPU() *string {
	return s.GPU
}

func (s *Node) GetGPUMemory() *string {
	return s.GPUMemory
}

func (s *Node) GetGPUType() *string {
	return s.GPUType
}

func (s *Node) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Node) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *Node) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Node) GetHyperZone() *string {
	return s.HyperZone
}

func (s *Node) GetIsBound() *bool {
	return s.IsBound
}

func (s *Node) GetLimitCPU() *string {
	return s.LimitCPU
}

func (s *Node) GetLimitGPU() *string {
	return s.LimitGPU
}

func (s *Node) GetLimitMemory() *string {
	return s.LimitMemory
}

func (s *Node) GetMachineGroupId() *string {
	return s.MachineGroupId
}

func (s *Node) GetMemory() *string {
	return s.Memory
}

func (s *Node) GetNodeName() *string {
	return s.NodeName
}

func (s *Node) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *Node) GetNodeType() *string {
	return s.NodeType
}

func (s *Node) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *Node) GetPodNum() *int64 {
	return s.PodNum
}

func (s *Node) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *Node) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *Node) GetRequestCPU() *string {
	return s.RequestCPU
}

func (s *Node) GetRequestGPU() *string {
	return s.RequestGPU
}

func (s *Node) GetRequestMemory() *string {
	return s.RequestMemory
}

func (s *Node) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *Node) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *Node) GetSelfQuotaWorkloadNum() *int64 {
	return s.SelfQuotaWorkloadNum
}

func (s *Node) GetSystemReservedCPU() *string {
	return s.SystemReservedCPU
}

func (s *Node) GetSystemReservedMemory() *string {
	return s.SystemReservedMemory
}

func (s *Node) GetUsers() []*UserInfo {
	return s.Users
}

func (s *Node) GetWorkloadNum() *int64 {
	return s.WorkloadNum
}

func (s *Node) SetAcceleratorType(v string) *Node {
	s.AcceleratorType = &v
	return s
}

func (s *Node) SetAllocatableCPU(v string) *Node {
	s.AllocatableCPU = &v
	return s
}

func (s *Node) SetAllocatableMemory(v string) *Node {
	s.AllocatableMemory = &v
	return s
}

func (s *Node) SetAncestorQuotaWorkloadNum(v int64) *Node {
	s.AncestorQuotaWorkloadNum = &v
	return s
}

func (s *Node) SetAvailabilityZone(v string) *Node {
	s.AvailabilityZone = &v
	return s
}

func (s *Node) SetBoundQuotas(v []*QuotaIdName) *Node {
	s.BoundQuotas = v
	return s
}

func (s *Node) SetCPU(v string) *Node {
	s.CPU = &v
	return s
}

func (s *Node) SetCreatorId(v string) *Node {
	s.CreatorId = &v
	return s
}

func (s *Node) SetDescendantQuotaWorkloadNum(v int64) *Node {
	s.DescendantQuotaWorkloadNum = &v
	return s
}

func (s *Node) SetGPU(v string) *Node {
	s.GPU = &v
	return s
}

func (s *Node) SetGPUMemory(v string) *Node {
	s.GPUMemory = &v
	return s
}

func (s *Node) SetGPUType(v string) *Node {
	s.GPUType = &v
	return s
}

func (s *Node) SetGmtCreateTime(v string) *Node {
	s.GmtCreateTime = &v
	return s
}

func (s *Node) SetGmtExpiredTime(v string) *Node {
	s.GmtExpiredTime = &v
	return s
}

func (s *Node) SetGmtModifiedTime(v string) *Node {
	s.GmtModifiedTime = &v
	return s
}

func (s *Node) SetHyperZone(v string) *Node {
	s.HyperZone = &v
	return s
}

func (s *Node) SetIsBound(v bool) *Node {
	s.IsBound = &v
	return s
}

func (s *Node) SetLimitCPU(v string) *Node {
	s.LimitCPU = &v
	return s
}

func (s *Node) SetLimitGPU(v string) *Node {
	s.LimitGPU = &v
	return s
}

func (s *Node) SetLimitMemory(v string) *Node {
	s.LimitMemory = &v
	return s
}

func (s *Node) SetMachineGroupId(v string) *Node {
	s.MachineGroupId = &v
	return s
}

func (s *Node) SetMemory(v string) *Node {
	s.Memory = &v
	return s
}

func (s *Node) SetNodeName(v string) *Node {
	s.NodeName = &v
	return s
}

func (s *Node) SetNodeStatus(v string) *Node {
	s.NodeStatus = &v
	return s
}

func (s *Node) SetNodeType(v string) *Node {
	s.NodeType = &v
	return s
}

func (s *Node) SetOrderStatus(v string) *Node {
	s.OrderStatus = &v
	return s
}

func (s *Node) SetPodNum(v int64) *Node {
	s.PodNum = &v
	return s
}

func (s *Node) SetReasonCode(v string) *Node {
	s.ReasonCode = &v
	return s
}

func (s *Node) SetReasonMessage(v string) *Node {
	s.ReasonMessage = &v
	return s
}

func (s *Node) SetRequestCPU(v string) *Node {
	s.RequestCPU = &v
	return s
}

func (s *Node) SetRequestGPU(v string) *Node {
	s.RequestGPU = &v
	return s
}

func (s *Node) SetRequestMemory(v string) *Node {
	s.RequestMemory = &v
	return s
}

func (s *Node) SetResourceGroupId(v string) *Node {
	s.ResourceGroupId = &v
	return s
}

func (s *Node) SetResourceGroupName(v string) *Node {
	s.ResourceGroupName = &v
	return s
}

func (s *Node) SetSelfQuotaWorkloadNum(v int64) *Node {
	s.SelfQuotaWorkloadNum = &v
	return s
}

func (s *Node) SetSystemReservedCPU(v string) *Node {
	s.SystemReservedCPU = &v
	return s
}

func (s *Node) SetSystemReservedMemory(v string) *Node {
	s.SystemReservedMemory = &v
	return s
}

func (s *Node) SetUsers(v []*UserInfo) *Node {
	s.Users = v
	return s
}

func (s *Node) SetWorkloadNum(v int64) *Node {
	s.WorkloadNum = &v
	return s
}

func (s *Node) Validate() error {
	if s.BoundQuotas != nil {
		for _, item := range s.BoundQuotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
