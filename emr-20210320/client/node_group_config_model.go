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
	// The IDs of the additional security groups. In addition to the security group of the cluster, you can specify additional security groups for the node group. You can specify up to two security group IDs.
	//
	// example:
	//
	// ["sg-hp3abbae8lb6lmb1****"]
	AdditionalSecurityGroupIds []*string `json:"AdditionalSecurityGroupIds,omitempty" xml:"AdditionalSecurityGroupIds,omitempty" type:"Repeated"`
	// The auto scaling policy.
	AutoScalingPolicy *AutoScalingPolicy `json:"AutoScalingPolicy,omitempty" xml:"AutoScalingPolicy,omitempty"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the required capacity when the number of preemptible instances is insufficient. This parameter is effective only when `nodeResizeStrategy` is set to `COST_OPTIMIZED`.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// A list of custom component tags.
	ComponentTags []*string `json:"ComponentTags,omitempty" xml:"ComponentTags,omitempty" type:"Repeated"`
	// The cost optimization settings.
	CostOptimizedConfig *CostOptimizedConfig `json:"CostOptimizedConfig,omitempty" xml:"CostOptimizedConfig,omitempty"`
	// The data disks. Currently, the array can contain only one data disk.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The deployment set strategy. Valid values:
	//
	// - `NONE`: No deployment sets are used.
	//
	// - `CLUSTER`: The cluster-level deployment set is used.
	//
	// - `NODE_GROUP`: The node group-level deployment set is used.
	//
	// Default value: `NONE`.
	//
	// example:
	//
	// NONE
	DeploymentSetStrategy *string `json:"DeploymentSetStrategy,omitempty" xml:"DeploymentSetStrategy,omitempty"`
	// Specifies whether to enable graceful shutdown for the components in the node group. Valid values:
	//
	// - `true`: Enables graceful shutdown.
	//
	// - `false`: Disables graceful shutdown.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	GracefulShutdown *bool `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	// The instance types. You can specify 1 to 100 instance types.
	//
	// example:
	//
	// ["ecs.g6.xlarge"]
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The number of nodes in the node group. Valid values: 1 to 1,000.
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The name of the node group. The name can be up to 128 characters in length and must be unique within the cluster.
	//
	// example:
	//
	// core-1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// - `MASTER`: a master node group.
	//
	// - `CORE`: a core node group.
	//
	// - `TASK`: a task node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// CORE
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// The node scale-out strategy. Valid values:
	//
	// - `COST_OPTIMIZED`: The cost-optimized strategy.
	//
	// - `PRIORITY`: The priority-based strategy.
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// PRIORITY
	NodeResizeStrategy *string `json:"NodeResizeStrategy,omitempty" xml:"NodeResizeStrategy,omitempty"`
	// The billing method of the node group. If you do not specify this parameter, the billing method of the cluster is used. Valid values:
	//
	// - `PayAsYouGo`: pay-as-you-go.
	//
	// - `Subscription`: subscription.
	//
	// Default value: `PayAsYouGo`.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The private pool options. This parameter is effective only when you create pay-as-you-go instances.
	PrivatePoolOptions *PrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty"`
	// The bid prices for the preemptible instances. This parameter is effective only when `SpotStrategy` is set to `SpotWithPriceLimit`. You can specify up to 100 bid prices.
	SpotBidPrices []*SpotBidPrice `json:"SpotBidPrices,omitempty" xml:"SpotBidPrices,omitempty" type:"Repeated"`
	// If enabled, the auto scaling group attempts to create a new instance to replace a spot instance that is about to be reclaimed. This process is triggered when the auto scaling group receives a system message about the impending reclamation. Valid values:
	//
	// - `true`: The auto scaling group attempts to replace a spot instance that is about to be reclaimed.
	//
	// - `false`: The auto scaling group does not attempt to replace a spot instance that is about to be reclaimed.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The preemption strategy for preemptible instances. Valid values:
	//
	// - `NoSpot`: pay-as-you-go instances.
	//
	// - `SpotWithPriceLimit`: preemptible instances with a user-defined maximum hourly price.
	//
	// - `SpotAsPriceGo`: preemptible instances that are automatically bid at the pay-as-you-go price.
	//
	// Default value: `NoSpot`.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The subscription settings of the node group. If you do not specify this parameter, the subscription settings of the cluster are used.
	SubscriptionConfig *SubscriptionConfig `json:"SubscriptionConfig,omitempty" xml:"SubscriptionConfig,omitempty"`
	// The system disk.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// The vSwitch IDs. You can specify 1 to 20 vSwitch IDs.
	//
	// example:
	//
	// ["vsw-hp35g7ya5ymw68mmg****"]
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// Specifies whether to assign a public IP address to the instances. Valid values:
	//
	// - `true`: Assigns a public IP address.
	//
	// - `false`: Does not assign a public IP address.
	//
	// Default value: `false`.
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
