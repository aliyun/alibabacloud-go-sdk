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
	// Whether auto-renewal is enabled for the node. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration for the node.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// The unit of the auto-renewal duration.
	//
	// example:
	//
	// Month
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	// The creation time of the node.
	//
	// example:
	//
	// 1603728394857
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the node.
	//
	// example:
	//
	// 1603728394857
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance type of the node. This corresponds to an ECS instance type. You can call the ECS [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the available instance types.
	//
	// example:
	//
	// ecs.g7.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maintenance status of the node. Valid values:
	//
	// - ON: The node is in maintenance mode.
	//
	// - OFF: The node is not in maintenance mode.
	//
	// If this parameter is empty, the node is not in maintenance mode.
	MaintenanceStatus *string `json:"MaintenanceStatus,omitempty" xml:"MaintenanceStatus,omitempty"`
	// The ID of the node group.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The type of the node group. Valid values:
	//
	// - MASTER: A master node group.
	//
	// - CORE: A core node group.
	//
	// - TASK: A task node group.
	//
	// example:
	//
	// CORE
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// core1-1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The state of the node. Valid values:
	//
	// - Pending: The node is being created.
	//
	// - Starting: The node is starting up.
	//
	// - Running: The node is operational and running services.
	//
	// - Stopping: The node is shutting down.
	//
	// - Stopped: The node is powered off.
	//
	// - Terminated: The node has been permanently deleted.
	//
	// example:
	//
	// Running
	NodeState *string `json:"NodeState,omitempty" xml:"NodeState,omitempty"`
	// The private IP address of the node.
	//
	// example:
	//
	// ``10.10.**.**``
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The public IP address of the node.
	//
	// example:
	//
	// 42.1.1.**
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The ID of the zone.
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
