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
	// The additional security group IDs.
	//
	// example:
	//
	// ["sg-hp3abbae8lb6lmb1****"]
	AdditionalSecurityGroupIds []*string `json:"AdditionalSecurityGroupIds,omitempty" xml:"AdditionalSecurityGroupIds,omitempty" type:"Repeated"`
	// Applies only when `NodeResizeStrategy` is set to `COST_OPTIMIZED`. If set to `true`, the system creates Pay-As-You-Go instances to meet the target capacity if it fails to create enough spot instances due to price or inventory constraints.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The configurations of the cost-optimized mode.
	CostOptimizedConfig *CostOptimizedConfig `json:"CostOptimizedConfig,omitempty" xml:"CostOptimizedConfig,omitempty"`
	// The data disks.
	DataDisks []*DataDisk `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The Deployment Set strategy. Valid values:
	//
	// - NONE: Does not use a Deployment Set.
	//
	// - CLUSTER: Uses a cluster-level Deployment Set.
	//
	// - NODE_GROUP: Uses a node group-level Deployment Set.
	//
	// Default: `NONE`.
	//
	// example:
	//
	// NONE
	DeploymentSetStrategy *string `json:"DeploymentSetStrategy,omitempty" xml:"DeploymentSetStrategy,omitempty"`
	// Specifies whether to enable graceful shutdown for components deployed in the node group. Valid values:
	//
	// - true: Enables graceful shutdown.
	//
	// - false: Disables graceful shutdown.
	//
	// example:
	//
	// false
	GracefulShutdown *bool `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	// The instance types.
	//
	// example:
	//
	// ["ecs.g6.4xlarge"]
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The node group ID.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node group name.
	//
	// example:
	//
	// core-1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The state of the node group.
	//
	// example:
	//
	// RESIZING
	NodeGroupState *string `json:"NodeGroupState,omitempty" xml:"NodeGroupState,omitempty"`
	// The type of the node group. Valid values:
	//
	// - MASTER: A master node.
	//
	// - CORE: A core node.
	//
	// - TASK: A task node.
	//
	// - GATEWAY: A gateway node. This value is not applicable to DATALAKE, OLAP, or DATASERVING clusters.
	//
	// example:
	//
	// MASTER
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// - COST_OPTIMIZED: The cost-optimized strategy.
	//
	// - PRIORITY: The priority-based strategy.
	//
	// example:
	//
	// PRIORITY
	NodeResizeStrategy *string `json:"NodeResizeStrategy,omitempty" xml:"NodeResizeStrategy,omitempty"`
	// The payment type. Valid values are `Subscription` for the subscription billing method and `PayAsYouGo` for the Pay-As-You-Go billing method.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The private pool options.
	PrivatePoolOptions *PrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty"`
	// The number of running nodes.
	//
	// example:
	//
	// 3
	RunningNodeCount *int32 `json:"RunningNodeCount,omitempty" xml:"RunningNodeCount,omitempty"`
	// The bid prices for the spot instances. This parameter is effective only when `SpotStrategy` is set to `SpotWithPriceLimit`. The array can contain 0 to 100 elements.
	SpotBidPrices []*SpotBidPrice `json:"SpotBidPrices,omitempty" xml:"SpotBidPrices,omitempty" type:"Repeated"`
	// Specifies whether to enable spot instance remedy. If enabled, the scaling group attempts to create a new instance to replace a spot instance that is about to be reclaimed. Valid values:
	//
	// - true: Enables spot instance remedy.
	//
	// - false: Disables spot instance remedy.
	//
	// Default: `false`.
	//
	// example:
	//
	// false
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The policy for using spot instances. Valid values:
	//
	// - NoSpot: No spot instances are used.
	//
	// - SpotWithPriceLimit: Spot instances are created with a user-defined maximum bid price.
	//
	// - SpotAsPriceGo: The system automatically bids for spot instances. The bid price does not exceed the price of a Pay-As-You-Go instance.
	//
	// Default: `NoSpot`.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The reason for the state change.
	//
	// example:
	//
	// 手动重启
	StateChangeReason *NodeGroupStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
	// The node group state.
	//
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The system disk.
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// The VSwitch IDs.
	//
	// example:
	//
	// ["vsw-hp35g7ya5ymw68mmg****"]
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// Specifies whether to assign a public IP address.
	//
	// example:
	//
	// true
	WithPublicIp *bool `json:"WithPublicIp,omitempty" xml:"WithPublicIp,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou
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
	if s.StateChangeReason != nil {
		if err := s.StateChangeReason.Validate(); err != nil {
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
