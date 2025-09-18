// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeGroupConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalSecurityGroupIds(v []*string) *CreateNodeGroupConfig
	GetAdditionalSecurityGroupIds() []*string
	SetAutoScalingPolicy(v *AutoScalingPolicy) *CreateNodeGroupConfig
	GetAutoScalingPolicy() *AutoScalingPolicy
	SetCompensateWithOnDemand(v bool) *CreateNodeGroupConfig
	GetCompensateWithOnDemand() *bool
	SetComponentTags(v []*string) *CreateNodeGroupConfig
	GetComponentTags() []*string
	SetCostOptimizedConfig(v *CostOptimizedConfig) *CreateNodeGroupConfig
	GetCostOptimizedConfig() *CostOptimizedConfig
	SetDataDisks(v []*DataDisk) *CreateNodeGroupConfig
	GetDataDisks() []*DataDisk
	SetDeploymentSetStrategy(v string) *CreateNodeGroupConfig
	GetDeploymentSetStrategy() *string
	SetGracefulShutdown(v bool) *CreateNodeGroupConfig
	GetGracefulShutdown() *bool
	SetInstanceTypes(v []*string) *CreateNodeGroupConfig
	GetInstanceTypes() []*string
	SetNodeCount(v int32) *CreateNodeGroupConfig
	GetNodeCount() *int32
	SetNodeGroupName(v string) *CreateNodeGroupConfig
	GetNodeGroupName() *string
	SetNodeGroupType(v string) *CreateNodeGroupConfig
	GetNodeGroupType() *string
	SetNodeResizeStrategy(v string) *CreateNodeGroupConfig
	GetNodeResizeStrategy() *string
	SetPaymentType(v string) *CreateNodeGroupConfig
	GetPaymentType() *string
	SetPrivatePoolOptions(v *PrivatePoolOptions) *CreateNodeGroupConfig
	GetPrivatePoolOptions() *PrivatePoolOptions
	SetSpotBidPrices(v []*CreateNodeGroupConfigSpotBidPrices) *CreateNodeGroupConfig
	GetSpotBidPrices() []*CreateNodeGroupConfigSpotBidPrices
	SetSpotInstanceRemedy(v bool) *CreateNodeGroupConfig
	GetSpotInstanceRemedy() *bool
	SetSpotStrategy(v string) *CreateNodeGroupConfig
	GetSpotStrategy() *string
	SetSubscriptionConfig(v *SubscriptionConfig) *CreateNodeGroupConfig
	GetSubscriptionConfig() *SubscriptionConfig
	SetSystemDisk(v *SystemDisk) *CreateNodeGroupConfig
	GetSystemDisk() *SystemDisk
	SetVSwitchIds(v []*string) *CreateNodeGroupConfig
	GetVSwitchIds() []*string
	SetWithPublicIp(v bool) *CreateNodeGroupConfig
	GetWithPublicIp() *bool
}

type CreateNodeGroupConfig struct {
	// 附加安全组。除集群设置的安全组外，为节点组单独设置的附加安全组。数组元数个数N的取值范围：0~2。
	//
	// example:
	//
	// ["sg-hp3abbae8lb6lmb1****"]
	AdditionalSecurityGroupIds []*string          `json:"AdditionalSecurityGroupIds,omitempty" xml:"AdditionalSecurityGroupIds,omitempty" type:"Repeated"`
	AutoScalingPolicy          *AutoScalingPolicy `json:"AutoScalingPolicy,omitempty" xml:"AutoScalingPolicy,omitempty"`
	// example:
	//
	// true
	CompensateWithOnDemand *bool     `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	ComponentTags          []*string `json:"ComponentTags,omitempty" xml:"ComponentTags,omitempty" type:"Repeated"`
	// 成本优化模式配置。
	CostOptimizedConfig *CostOptimizedConfig `json:"CostOptimizedConfig,omitempty" xml:"CostOptimizedConfig,omitempty"`
	// 数据盘。当前数据盘只支持一种磁盘类型，即数组元数个数N的取值范围：1~1。
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
	// 默认值：false。
	//
	// example:
	//
	// false
	GracefulShutdown *bool `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	// 节点实例类型列表。数组元数个数N的取值范围：1~100。
	//
	// example:
	//
	// ["ecs.g6.xlarge"]
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// 节点数量。取值范围：1~1000。
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// 节点组名称。最大长度128个字符。集群内要求节点组名称唯一。
	//
	// example:
	//
	// core-1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// 节点组类型。取值范围：
	//
	// - MASTER：管理类型节点组。
	//
	// - CORE：存储类型节点组。
	//
	// - TASK：计算类型节点组。
	//
	// This parameter is required.
	//
	// example:
	//
	// CORE
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// 节点扩容策略。取值范围：
	//
	// - COST_OPTIMIZED：成本优化策略。
	//
	// - PRIORITY：优先级策略。
	//
	// 默认值：PRIORITY。
	//
	// example:
	//
	// PRIORITY
	NodeResizeStrategy *string `json:"NodeResizeStrategy,omitempty" xml:"NodeResizeStrategy,omitempty"`
	// 节点组付费类型。不传入时默认和集群付费类型一致。取值范围：
	//
	// - PayAsYouGo：后付费，按量付费。
	//
	// - Subscription：预付费，包年包月。
	//
	// 默认值：PayAsYouGo。
	//
	// example:
	//
	// PayAsYouGo
	PaymentType        *string             `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	PrivatePoolOptions *PrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty"`
	// 抢占式Spot实例出价价格。参数SpotStrategy取值为SpotWithPriceLimit时生效。数组元数个数N的取值范围：0~100。
	SpotBidPrices []*CreateNodeGroupConfigSpotBidPrices `json:"SpotBidPrices,omitempty" xml:"SpotBidPrices,omitempty" type:"Repeated"`
	// 开启后，当收到抢占式实例将被回收的系统消息时，伸缩组将尝试创建新的实例，替换掉将被回收的抢占式实例。取值范围：
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
	// 抢占式Spot实例策略。取值范围：
	//
	// - NoSpot：正常按量付费实例。
	//
	// - SpotWithPriceLimit：设置最高出价的抢占式实例。
	//
	// - SpotAsPriceGo：系统自动出价，最高按量付费价格的抢占式实例。
	//
	// 默认值：NoSpot。
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// 节点组预付费配置。不传入时默认和集群预付费配置一致。
	SubscriptionConfig *SubscriptionConfig `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	// 系统盘。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// 虚拟机交换机ID列表。数组元数个数N的取值范围：1~20。
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
	// 默认值：false。
	//
	// example:
	//
	// false
	WithPublicIp *bool `json:"WithPublicIp,omitempty" xml:"WithPublicIp,omitempty"`
}

func (s CreateNodeGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupConfig) GetAdditionalSecurityGroupIds() []*string {
	return s.AdditionalSecurityGroupIds
}

func (s *CreateNodeGroupConfig) GetAutoScalingPolicy() *AutoScalingPolicy {
	return s.AutoScalingPolicy
}

func (s *CreateNodeGroupConfig) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *CreateNodeGroupConfig) GetComponentTags() []*string {
	return s.ComponentTags
}

func (s *CreateNodeGroupConfig) GetCostOptimizedConfig() *CostOptimizedConfig {
	return s.CostOptimizedConfig
}

func (s *CreateNodeGroupConfig) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *CreateNodeGroupConfig) GetDeploymentSetStrategy() *string {
	return s.DeploymentSetStrategy
}

func (s *CreateNodeGroupConfig) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *CreateNodeGroupConfig) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateNodeGroupConfig) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CreateNodeGroupConfig) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *CreateNodeGroupConfig) GetNodeGroupType() *string {
	return s.NodeGroupType
}

func (s *CreateNodeGroupConfig) GetNodeResizeStrategy() *string {
	return s.NodeResizeStrategy
}

func (s *CreateNodeGroupConfig) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateNodeGroupConfig) GetPrivatePoolOptions() *PrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateNodeGroupConfig) GetSpotBidPrices() []*CreateNodeGroupConfigSpotBidPrices {
	return s.SpotBidPrices
}

func (s *CreateNodeGroupConfig) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *CreateNodeGroupConfig) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateNodeGroupConfig) GetSubscriptionConfig() *SubscriptionConfig {
	return s.SubscriptionConfig
}

func (s *CreateNodeGroupConfig) GetSystemDisk() *SystemDisk {
	return s.SystemDisk
}

func (s *CreateNodeGroupConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateNodeGroupConfig) GetWithPublicIp() *bool {
	return s.WithPublicIp
}

func (s *CreateNodeGroupConfig) SetAdditionalSecurityGroupIds(v []*string) *CreateNodeGroupConfig {
	s.AdditionalSecurityGroupIds = v
	return s
}

func (s *CreateNodeGroupConfig) SetAutoScalingPolicy(v *AutoScalingPolicy) *CreateNodeGroupConfig {
	s.AutoScalingPolicy = v
	return s
}

func (s *CreateNodeGroupConfig) SetCompensateWithOnDemand(v bool) *CreateNodeGroupConfig {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *CreateNodeGroupConfig) SetComponentTags(v []*string) *CreateNodeGroupConfig {
	s.ComponentTags = v
	return s
}

func (s *CreateNodeGroupConfig) SetCostOptimizedConfig(v *CostOptimizedConfig) *CreateNodeGroupConfig {
	s.CostOptimizedConfig = v
	return s
}

func (s *CreateNodeGroupConfig) SetDataDisks(v []*DataDisk) *CreateNodeGroupConfig {
	s.DataDisks = v
	return s
}

func (s *CreateNodeGroupConfig) SetDeploymentSetStrategy(v string) *CreateNodeGroupConfig {
	s.DeploymentSetStrategy = &v
	return s
}

func (s *CreateNodeGroupConfig) SetGracefulShutdown(v bool) *CreateNodeGroupConfig {
	s.GracefulShutdown = &v
	return s
}

func (s *CreateNodeGroupConfig) SetInstanceTypes(v []*string) *CreateNodeGroupConfig {
	s.InstanceTypes = v
	return s
}

func (s *CreateNodeGroupConfig) SetNodeCount(v int32) *CreateNodeGroupConfig {
	s.NodeCount = &v
	return s
}

func (s *CreateNodeGroupConfig) SetNodeGroupName(v string) *CreateNodeGroupConfig {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupConfig) SetNodeGroupType(v string) *CreateNodeGroupConfig {
	s.NodeGroupType = &v
	return s
}

func (s *CreateNodeGroupConfig) SetNodeResizeStrategy(v string) *CreateNodeGroupConfig {
	s.NodeResizeStrategy = &v
	return s
}

func (s *CreateNodeGroupConfig) SetPaymentType(v string) *CreateNodeGroupConfig {
	s.PaymentType = &v
	return s
}

func (s *CreateNodeGroupConfig) SetPrivatePoolOptions(v *PrivatePoolOptions) *CreateNodeGroupConfig {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateNodeGroupConfig) SetSpotBidPrices(v []*CreateNodeGroupConfigSpotBidPrices) *CreateNodeGroupConfig {
	s.SpotBidPrices = v
	return s
}

func (s *CreateNodeGroupConfig) SetSpotInstanceRemedy(v bool) *CreateNodeGroupConfig {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *CreateNodeGroupConfig) SetSpotStrategy(v string) *CreateNodeGroupConfig {
	s.SpotStrategy = &v
	return s
}

func (s *CreateNodeGroupConfig) SetSubscriptionConfig(v *SubscriptionConfig) *CreateNodeGroupConfig {
	s.SubscriptionConfig = v
	return s
}

func (s *CreateNodeGroupConfig) SetSystemDisk(v *SystemDisk) *CreateNodeGroupConfig {
	s.SystemDisk = v
	return s
}

func (s *CreateNodeGroupConfig) SetVSwitchIds(v []*string) *CreateNodeGroupConfig {
	s.VSwitchIds = v
	return s
}

func (s *CreateNodeGroupConfig) SetWithPublicIp(v bool) *CreateNodeGroupConfig {
	s.WithPublicIp = &v
	return s
}

func (s *CreateNodeGroupConfig) Validate() error {
	return dara.Validate(s)
}

type CreateNodeGroupConfigSpotBidPrices struct {
	// 实例的每小时最高出价。支持最大3位小数，参数SpotStrategy=SpotWithPriceLimit时，该参数生效。
	//
	// example:
	//
	// 1000.0
	BidPrice *float64 `json:"BidPrice,omitempty" xml:"BidPrice,omitempty"`
	// 实例类型。
	//
	// example:
	//
	// ecs.g7.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateNodeGroupConfigSpotBidPrices) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupConfigSpotBidPrices) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupConfigSpotBidPrices) GetBidPrice() *float64 {
	return s.BidPrice
}

func (s *CreateNodeGroupConfigSpotBidPrices) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateNodeGroupConfigSpotBidPrices) SetBidPrice(v float64) *CreateNodeGroupConfigSpotBidPrices {
	s.BidPrice = &v
	return s
}

func (s *CreateNodeGroupConfigSpotBidPrices) SetInstanceType(v string) *CreateNodeGroupConfigSpotBidPrices {
	s.InstanceType = &v
	return s
}

func (s *CreateNodeGroupConfigSpotBidPrices) Validate() error {
	return dara.Validate(s)
}
