// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeGroupConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalSecurityGroupIds(v []*string) *NodeGroupConfig
	GetAdditionalSecurityGroupIds() []*string
	SetAutoScalingPolicy(v *AutoScalingPolicy) *NodeGroupConfig
	GetAutoScalingPolicy() *AutoScalingPolicy
	SetCompensateWithOnDemand(v bool) *NodeGroupConfig
	GetCompensateWithOnDemand() *bool
	SetComponentTags(v []*string) *NodeGroupConfig
	GetComponentTags() []*string
	SetCostOptimizedConfig(v *CostOptimizedConfig) *NodeGroupConfig
	GetCostOptimizedConfig() *CostOptimizedConfig
	SetDataDisks(v []*DataDisk) *NodeGroupConfig
	GetDataDisks() []*DataDisk
	SetDeploymentSetStrategy(v string) *NodeGroupConfig
	GetDeploymentSetStrategy() *string
	SetGracefulShutdown(v bool) *NodeGroupConfig
	GetGracefulShutdown() *bool
	SetInstanceTypes(v []*string) *NodeGroupConfig
	GetInstanceTypes() []*string
	SetNodeCount(v int32) *NodeGroupConfig
	GetNodeCount() *int32
	SetNodeGroupName(v string) *NodeGroupConfig
	GetNodeGroupName() *string
	SetNodeGroupType(v string) *NodeGroupConfig
	GetNodeGroupType() *string
	SetNodeResizeStrategy(v string) *NodeGroupConfig
	GetNodeResizeStrategy() *string
	SetPaymentType(v string) *NodeGroupConfig
	GetPaymentType() *string
	SetPrivatePoolOptions(v *PrivatePoolOptions) *NodeGroupConfig
	GetPrivatePoolOptions() *PrivatePoolOptions
	SetSpotBidPrices(v []*SpotBidPrice) *NodeGroupConfig
	GetSpotBidPrices() []*SpotBidPrice
	SetSpotInstanceRemedy(v bool) *NodeGroupConfig
	GetSpotInstanceRemedy() *bool
	SetSpotStrategy(v string) *NodeGroupConfig
	GetSpotStrategy() *string
	SetSubscriptionConfig(v *SubscriptionConfig) *NodeGroupConfig
	GetSubscriptionConfig() *SubscriptionConfig
	SetSystemDisk(v *SystemDisk) *NodeGroupConfig
	GetSystemDisk() *SystemDisk
	SetVSwitchIds(v []*string) *NodeGroupConfig
	GetVSwitchIds() []*string
	SetWithPublicIp(v bool) *NodeGroupConfig
	GetWithPublicIp() *bool
}

type NodeGroupConfig struct {
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
	SpotBidPrices []*SpotBidPrice `json:"SpotBidPrices,omitempty" xml:"SpotBidPrices,omitempty" type:"Repeated"`
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

func (s NodeGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s NodeGroupConfig) GoString() string {
	return s.String()
}

func (s *NodeGroupConfig) GetAdditionalSecurityGroupIds() []*string {
	return s.AdditionalSecurityGroupIds
}

func (s *NodeGroupConfig) GetAutoScalingPolicy() *AutoScalingPolicy {
	return s.AutoScalingPolicy
}

func (s *NodeGroupConfig) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *NodeGroupConfig) GetComponentTags() []*string {
	return s.ComponentTags
}

func (s *NodeGroupConfig) GetCostOptimizedConfig() *CostOptimizedConfig {
	return s.CostOptimizedConfig
}

func (s *NodeGroupConfig) GetDataDisks() []*DataDisk {
	return s.DataDisks
}

func (s *NodeGroupConfig) GetDeploymentSetStrategy() *string {
	return s.DeploymentSetStrategy
}

func (s *NodeGroupConfig) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *NodeGroupConfig) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *NodeGroupConfig) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *NodeGroupConfig) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *NodeGroupConfig) GetNodeGroupType() *string {
	return s.NodeGroupType
}

func (s *NodeGroupConfig) GetNodeResizeStrategy() *string {
	return s.NodeResizeStrategy
}

func (s *NodeGroupConfig) GetPaymentType() *string {
	return s.PaymentType
}

func (s *NodeGroupConfig) GetPrivatePoolOptions() *PrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *NodeGroupConfig) GetSpotBidPrices() []*SpotBidPrice {
	return s.SpotBidPrices
}

func (s *NodeGroupConfig) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *NodeGroupConfig) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *NodeGroupConfig) GetSubscriptionConfig() *SubscriptionConfig {
	return s.SubscriptionConfig
}

func (s *NodeGroupConfig) GetSystemDisk() *SystemDisk {
	return s.SystemDisk
}

func (s *NodeGroupConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *NodeGroupConfig) GetWithPublicIp() *bool {
	return s.WithPublicIp
}

func (s *NodeGroupConfig) SetAdditionalSecurityGroupIds(v []*string) *NodeGroupConfig {
	s.AdditionalSecurityGroupIds = v
	return s
}

func (s *NodeGroupConfig) SetAutoScalingPolicy(v *AutoScalingPolicy) *NodeGroupConfig {
	s.AutoScalingPolicy = v
	return s
}

func (s *NodeGroupConfig) SetCompensateWithOnDemand(v bool) *NodeGroupConfig {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *NodeGroupConfig) SetComponentTags(v []*string) *NodeGroupConfig {
	s.ComponentTags = v
	return s
}

func (s *NodeGroupConfig) SetCostOptimizedConfig(v *CostOptimizedConfig) *NodeGroupConfig {
	s.CostOptimizedConfig = v
	return s
}

func (s *NodeGroupConfig) SetDataDisks(v []*DataDisk) *NodeGroupConfig {
	s.DataDisks = v
	return s
}

func (s *NodeGroupConfig) SetDeploymentSetStrategy(v string) *NodeGroupConfig {
	s.DeploymentSetStrategy = &v
	return s
}

func (s *NodeGroupConfig) SetGracefulShutdown(v bool) *NodeGroupConfig {
	s.GracefulShutdown = &v
	return s
}

func (s *NodeGroupConfig) SetInstanceTypes(v []*string) *NodeGroupConfig {
	s.InstanceTypes = v
	return s
}

func (s *NodeGroupConfig) SetNodeCount(v int32) *NodeGroupConfig {
	s.NodeCount = &v
	return s
}

func (s *NodeGroupConfig) SetNodeGroupName(v string) *NodeGroupConfig {
	s.NodeGroupName = &v
	return s
}

func (s *NodeGroupConfig) SetNodeGroupType(v string) *NodeGroupConfig {
	s.NodeGroupType = &v
	return s
}

func (s *NodeGroupConfig) SetNodeResizeStrategy(v string) *NodeGroupConfig {
	s.NodeResizeStrategy = &v
	return s
}

func (s *NodeGroupConfig) SetPaymentType(v string) *NodeGroupConfig {
	s.PaymentType = &v
	return s
}

func (s *NodeGroupConfig) SetPrivatePoolOptions(v *PrivatePoolOptions) *NodeGroupConfig {
	s.PrivatePoolOptions = v
	return s
}

func (s *NodeGroupConfig) SetSpotBidPrices(v []*SpotBidPrice) *NodeGroupConfig {
	s.SpotBidPrices = v
	return s
}

func (s *NodeGroupConfig) SetSpotInstanceRemedy(v bool) *NodeGroupConfig {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *NodeGroupConfig) SetSpotStrategy(v string) *NodeGroupConfig {
	s.SpotStrategy = &v
	return s
}

func (s *NodeGroupConfig) SetSubscriptionConfig(v *SubscriptionConfig) *NodeGroupConfig {
	s.SubscriptionConfig = v
	return s
}

func (s *NodeGroupConfig) SetSystemDisk(v *SystemDisk) *NodeGroupConfig {
	s.SystemDisk = v
	return s
}

func (s *NodeGroupConfig) SetVSwitchIds(v []*string) *NodeGroupConfig {
	s.VSwitchIds = v
	return s
}

func (s *NodeGroupConfig) SetWithPublicIp(v bool) *NodeGroupConfig {
	s.WithPublicIp = &v
	return s
}

func (s *NodeGroupConfig) Validate() error {
	if s.AutoScalingPolicy != nil {
		if err := s.AutoScalingPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.CostOptimizedConfig != nil {
		if err := s.CostOptimizedConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DataDisks != nil {
		for _, item := range s.DataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SpotBidPrices != nil {
		for _, item := range s.SpotBidPrices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubscriptionConfig != nil {
		if err := s.SubscriptionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}
