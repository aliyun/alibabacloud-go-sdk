// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNode interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *Node
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *Node
	GetAutoRenewDuration() *int32
	SetAutoRenewDurationUnit(v string) *Node
	GetAutoRenewDurationUnit() *string
	SetCreateTime(v int64) *Node
	GetCreateTime() *int64
	SetExpireTime(v int64) *Node
	GetExpireTime() *int64
	SetInstanceType(v string) *Node
	GetInstanceType() *string
	SetMaintenanceStatus(v string) *Node
	GetMaintenanceStatus() *string
	SetNodeGroupId(v string) *Node
	GetNodeGroupId() *string
	SetNodeGroupType(v string) *Node
	GetNodeGroupType() *string
	SetNodeId(v string) *Node
	GetNodeId() *string
	SetNodeName(v string) *Node
	GetNodeName() *string
	SetNodeState(v string) *Node
	GetNodeState() *string
	SetPrivateIp(v string) *Node
	GetPrivateIp() *string
	SetPublicIp(v string) *Node
	GetPublicIp() *string
	SetZoneId(v string) *Node
	GetZoneId() *string
}

type Node struct {
	// 节点是否自动续费。
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// 节点自动续费时长。
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// 节点自动续费时长单位。
	//
	// example:
	//
	// Month
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 节点过期时间。
	//
	// example:
	//
	// 1603728394857
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// 实例类型。
	//
	// example:
	//
	// ecs.g6e.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 运维模式状态。取值范围：
	//
	// - ON：处于运维模式。
	//
	// - OFF：处于非运维模式。
	//
	// 为空表示处于非运维模式。
	MaintenanceStatus *string `json:"MaintenanceStatus,omitempty" xml:"MaintenanceStatus,omitempty"`
	// 节点组ID。
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// 节点组类型。
	//
	// example:
	//
	// CORE
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// 节点ID。
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// 节点名称。
	//
	// example:
	//
	// core1-1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// 节点状态。取值范围：
	//
	// - Pending：创建中。
	//
	// - Starting：启动中。
	//
	// - Running：运行中。
	//
	// - Stopping：停止中。
	//
	// - Stopped：已停止。
	//
	// - Terminated：已终止。
	//
	// example:
	//
	// Running
	NodeState *string `json:"NodeState,omitempty" xml:"NodeState,omitempty"`
	// 私网IP。
	//
	// example:
	//
	// 10.10.10.1
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// 公网IP。
	//
	// example:
	//
	// 42.120.75.***
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// 可用区ID。
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s Node) String() string {
	return dara.Prettify(s)
}

func (s Node) GoString() string {
	return s.String()
}

func (s *Node) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *Node) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *Node) GetAutoRenewDurationUnit() *string {
	return s.AutoRenewDurationUnit
}

func (s *Node) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Node) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *Node) GetInstanceType() *string {
	return s.InstanceType
}

func (s *Node) GetMaintenanceStatus() *string {
	return s.MaintenanceStatus
}

func (s *Node) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *Node) GetNodeGroupType() *string {
	return s.NodeGroupType
}

func (s *Node) GetNodeId() *string {
	return s.NodeId
}

func (s *Node) GetNodeName() *string {
	return s.NodeName
}

func (s *Node) GetNodeState() *string {
	return s.NodeState
}

func (s *Node) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *Node) GetPublicIp() *string {
	return s.PublicIp
}

func (s *Node) GetZoneId() *string {
	return s.ZoneId
}

func (s *Node) SetAutoRenew(v bool) *Node {
	s.AutoRenew = &v
	return s
}

func (s *Node) SetAutoRenewDuration(v int32) *Node {
	s.AutoRenewDuration = &v
	return s
}

func (s *Node) SetAutoRenewDurationUnit(v string) *Node {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *Node) SetCreateTime(v int64) *Node {
	s.CreateTime = &v
	return s
}

func (s *Node) SetExpireTime(v int64) *Node {
	s.ExpireTime = &v
	return s
}

func (s *Node) SetInstanceType(v string) *Node {
	s.InstanceType = &v
	return s
}

func (s *Node) SetMaintenanceStatus(v string) *Node {
	s.MaintenanceStatus = &v
	return s
}

func (s *Node) SetNodeGroupId(v string) *Node {
	s.NodeGroupId = &v
	return s
}

func (s *Node) SetNodeGroupType(v string) *Node {
	s.NodeGroupType = &v
	return s
}

func (s *Node) SetNodeId(v string) *Node {
	s.NodeId = &v
	return s
}

func (s *Node) SetNodeName(v string) *Node {
	s.NodeName = &v
	return s
}

func (s *Node) SetNodeState(v string) *Node {
	s.NodeState = &v
	return s
}

func (s *Node) SetPrivateIp(v string) *Node {
	s.PrivateIp = &v
	return s
}

func (s *Node) SetPublicIp(v string) *Node {
	s.PublicIp = &v
	return s
}

func (s *Node) SetZoneId(v string) *Node {
	s.ZoneId = &v
	return s
}

func (s *Node) Validate() error {
	return dara.Validate(s)
}
