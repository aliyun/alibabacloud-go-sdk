// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeGroup interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalSecurityGroupIds(v []*string) *NodeGroup
	GetAdditionalSecurityGroupIds() []*string
	SetCompensateWithOnDemand(v bool) *NodeGroup
	GetCompensateWithOnDemand() *bool
	SetCostOptimizedConfig(v *CostOptimizedConfig) *NodeGroup
	GetCostOptimizedConfig() *CostOptimizedConfig
	SetDataDisks(v []*DataDisk) *NodeGroup
	GetDataDisks() []*DataDisk
	SetDeploymentSetStrategy(v string) *NodeGroup
	GetDeploymentSetStrategy() *string
	SetGracefulShutdown(v bool) *NodeGroup
	GetGracefulShutdown() *bool
	SetInstanceTypes(v []*string) *NodeGroup
	GetInstanceTypes() []*string
	SetNodeGroupId(v string) *NodeGroup
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *NodeGroup
	GetNodeGroupName() *string
	SetNodeGroupState(v string) *NodeGroup
	GetNodeGroupState() *string
	SetNodeGroupType(v string) *NodeGroup
	GetNodeGroupType() *string
	SetNodeResizeStrategy(v string) *NodeGroup
	GetNodeResizeStrategy() *string
	SetPaymentType(v string) *NodeGroup
	GetPaymentType() *string
	SetPrivatePoolOptions(v *PrivatePoolOptions) *NodeGroup
	GetPrivatePoolOptions() *PrivatePoolOptions
	SetRunningNodeCount(v int32) *NodeGroup
	GetRunningNodeCount() *int32
	SetSpotBidPrices(v []*SpotBidPrice) *NodeGroup
	GetSpotBidPrices() []*SpotBidPrice
	SetSpotInstanceRemedy(v bool) *NodeGroup
	GetSpotInstanceRemedy() *bool
	SetSpotStrategy(v string) *NodeGroup
	GetSpotStrategy() *string
	SetStateChangeReason(v *NodeGroupStateChangeReason) *NodeGroup
	GetStateChangeReason() *NodeGroupStateChangeReason
	SetStatus(v string) *NodeGroup
	GetStatus() *string
	SetSystemDisk(v *SystemDisk) *NodeGroup
	GetSystemDisk() *SystemDisk
	SetVSwitchIds(v []*string) *NodeGroup
	GetVSwitchIds() []*string
	SetWithPublicIp(v bool) *NodeGroup
	GetWithPublicIp() *bool
	SetZoneId(v string) *NodeGroup
	GetZoneId() *string
}

type NodeGroup struct {
	// 安全组ID。
	//
	// example:
	//
	// ["sg-hp3abbae8lb6lmb1****"]
	AdditionalSecurityGroupIds []*string `json:"AdditionalSecurityGroupIds,omitempty" xml:"AdditionalSecurityGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// 成本优化模式配置。
	CostOptimizedConfig *CostOptimizedConfig `json:"CostOptimizedConfig,omitempty" xml:"CostOptimizedConfig,omitempty"`
	// 数据盘列表。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// 部署集策略。取值范围：
	//
	// - NONE：不适用部署集。
	//
	// - CLUSTER：使用集群级别部署集。
	//
	// - NODE_GROUP：使用节点组级别部署集。
	//
	// 默认值：NONE。
	//
	// example:
	//
	// NONE
	DeploymentSetStrategy *string `json:"DeploymentSetStrategy,omitempty" xml:"DeploymentSetStrategy,omitempty"`
	// 节点组上部署的组件是否开启优雅下线。取值范围：
	//
	// - true：开启优雅下线。
	//
	// - false：不开启优雅下线。
	//
	// example:
	//
	// false
	GracefulShutdown *bool `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	// 实例类型列表。
	//
	// example:
	//
	// ["ecs.g6.4xlarge"]
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// 节点组ID。
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// 节点组名称。最大长度128个字符。
	//
	// example:
	//
	// core-1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// 节点组状态。
	//
	// example:
	//
	// CREATED
	NodeGroupState *string `json:"NodeGroupState,omitempty" xml:"NodeGroupState,omitempty"`
	// 节点组类型。取值范围：
	//
	// - MASTER：管理类型节点组。
	//
	// - CORE：存储类型节点组。
	//
	// - TASK：计算类型节点组。
	//
	// example:
	//
	// CORE
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// - COST_OPTIMIZED：成本优化策略。
	//
	// - PRIORITY：优先级策略。
	//
	// example:
	//
	// PRIORITY
	NodeResizeStrategy *string `json:"NodeResizeStrategy,omitempty" xml:"NodeResizeStrategy,omitempty"`
	// 节点组付费类型。取值范围：
	//
	// - PayAsYouGo：后付费，按量付费。
	//
	// - Subscription：预付费，包年包月。
	//
	// example:
	//
	// PayAsYouGo
	PaymentType        *string             `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	PrivatePoolOptions *PrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty"`
	// 存活节点数量。
	//
	// example:
	//
	// 3
	RunningNodeCount *int32          `json:"RunningNodeCount,omitempty" xml:"RunningNodeCount,omitempty"`
	SpotBidPrices    []*SpotBidPrice `json:"SpotBidPrices,omitempty" xml:"SpotBidPrices,omitempty" type:"Repeated"`
	// 开启补齐抢占式实例后，当收到抢占式实例将被回收的系统消息时，伸缩组将尝试创建新的实例，替换掉将被回收的抢占式实例。取值范围：
	//
	// - true：开启补齐抢占式实例。
	//
	// - false：不开启补齐抢占式实例。
	//
	// 默认值：false。
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// 是否支持竞价实例。
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// 状态变化原因。
	StateChangeReason *NodeGroupStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// 虚拟机交换机ID列表。
	//
	// example:
	//
	// ["vsw-hp35g7ya5ymw68mmg****"]
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// 是否开公网IP。取值范围：
	//
	// - true：开公网。
	//
	// - false：不开公网。
	//
	// example:
	//
	// false
	WithPublicIp *bool `json:"WithPublicIp,omitempty" xml:"WithPublicIp,omitempty"`
	// 可用区ID。
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s NodeGroup) String() string {
	return dara.Prettify(s)
}

func (s NodeGroup) GoString() string {
	return s.String()
}

func (s *NodeGroup) GetAdditionalSecurityGroupIds() []*string {
	return s.AdditionalSecurityGroupIds
}

func (s *NodeGroup) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *NodeGroup) GetCostOptimizedConfig() *CostOptimizedConfig {
	return s.CostOptimizedConfig
}

func (s *NodeGroup) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *NodeGroup) GetDeploymentSetStrategy() *string {
	return s.DeploymentSetStrategy
}

func (s *NodeGroup) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *NodeGroup) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *NodeGroup) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *NodeGroup) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *NodeGroup) GetNodeGroupState() *string {
	return s.NodeGroupState
}

func (s *NodeGroup) GetNodeGroupType() *string {
	return s.NodeGroupType
}

func (s *NodeGroup) GetNodeResizeStrategy() *string {
	return s.NodeResizeStrategy
}

func (s *NodeGroup) GetPaymentType() *string {
	return s.PaymentType
}

func (s *NodeGroup) GetPrivatePoolOptions() *PrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *NodeGroup) GetRunningNodeCount() *int32 {
	return s.RunningNodeCount
}

func (s *NodeGroup) GetSpotBidPrices() []*SpotBidPrice {
	return s.SpotBidPrices
}

func (s *NodeGroup) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *NodeGroup) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *NodeGroup) GetStateChangeReason() *NodeGroupStateChangeReason {
	return s.StateChangeReason
}

func (s *NodeGroup) GetStatus() *string {
	return s.Status
}

func (s *NodeGroup) GetSystemDisk() *SystemDisk {
	return s.SystemDisk
}

func (s *NodeGroup) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *NodeGroup) GetWithPublicIp() *bool {
	return s.WithPublicIp
}

func (s *NodeGroup) GetZoneId() *string {
	return s.ZoneId
}

func (s *NodeGroup) SetAdditionalSecurityGroupIds(v []*string) *NodeGroup {
	s.AdditionalSecurityGroupIds = v
	return s
}

func (s *NodeGroup) SetCompensateWithOnDemand(v bool) *NodeGroup {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *NodeGroup) SetCostOptimizedConfig(v *CostOptimizedConfig) *NodeGroup {
	s.CostOptimizedConfig = v
	return s
}

func (s *NodeGroup) SetDataDisks(v []*DataDisk) *NodeGroup {
	s.DataDisks = v
	return s
}

func (s *NodeGroup) SetDeploymentSetStrategy(v string) *NodeGroup {
	s.DeploymentSetStrategy = &v
	return s
}

func (s *NodeGroup) SetGracefulShutdown(v bool) *NodeGroup {
	s.GracefulShutdown = &v
	return s
}

func (s *NodeGroup) SetInstanceTypes(v []*string) *NodeGroup {
	s.InstanceTypes = v
	return s
}

func (s *NodeGroup) SetNodeGroupId(v string) *NodeGroup {
	s.NodeGroupId = &v
	return s
}

func (s *NodeGroup) SetNodeGroupName(v string) *NodeGroup {
	s.NodeGroupName = &v
	return s
}

func (s *NodeGroup) SetNodeGroupState(v string) *NodeGroup {
	s.NodeGroupState = &v
	return s
}

func (s *NodeGroup) SetNodeGroupType(v string) *NodeGroup {
	s.NodeGroupType = &v
	return s
}

func (s *NodeGroup) SetNodeResizeStrategy(v string) *NodeGroup {
	s.NodeResizeStrategy = &v
	return s
}

func (s *NodeGroup) SetPaymentType(v string) *NodeGroup {
	s.PaymentType = &v
	return s
}

func (s *NodeGroup) SetPrivatePoolOptions(v *PrivatePoolOptions) *NodeGroup {
	s.PrivatePoolOptions = v
	return s
}

func (s *NodeGroup) SetRunningNodeCount(v int32) *NodeGroup {
	s.RunningNodeCount = &v
	return s
}

func (s *NodeGroup) SetSpotBidPrices(v []*SpotBidPrice) *NodeGroup {
	s.SpotBidPrices = v
	return s
}

func (s *NodeGroup) SetSpotInstanceRemedy(v bool) *NodeGroup {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *NodeGroup) SetSpotStrategy(v string) *NodeGroup {
	s.SpotStrategy = &v
	return s
}

func (s *NodeGroup) SetStateChangeReason(v *NodeGroupStateChangeReason) *NodeGroup {
	s.StateChangeReason = v
	return s
}

func (s *NodeGroup) SetStatus(v string) *NodeGroup {
	s.Status = &v
	return s
}

func (s *NodeGroup) SetSystemDisk(v *SystemDisk) *NodeGroup {
	s.SystemDisk = v
	return s
}

func (s *NodeGroup) SetVSwitchIds(v []*string) *NodeGroup {
	s.VSwitchIds = v
	return s
}

func (s *NodeGroup) SetWithPublicIp(v bool) *NodeGroup {
	s.WithPublicIp = &v
	return s
}

func (s *NodeGroup) SetZoneId(v string) *NodeGroup {
	s.ZoneId = &v
	return s
}

func (s *NodeGroup) Validate() error {
	return dara.Validate(s)
}
