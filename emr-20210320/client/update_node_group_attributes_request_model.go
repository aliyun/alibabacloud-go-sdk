// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAckConfig(v *AckConfig) *UpdateNodeGroupAttributesRequest
	GetAckConfig() *AckConfig
	SetAdditionalSecurityGroupIds(v []*string) *UpdateNodeGroupAttributesRequest
	GetAdditionalSecurityGroupIds() []*string
	SetAutoCompensateState(v bool) *UpdateNodeGroupAttributesRequest
	GetAutoCompensateState() *bool
	SetClusterId(v string) *UpdateNodeGroupAttributesRequest
	GetClusterId() *string
	SetCostOptimizedConfig(v *CostOptimizedConfig) *UpdateNodeGroupAttributesRequest
	GetCostOptimizedConfig() *CostOptimizedConfig
	SetDescription(v string) *UpdateNodeGroupAttributesRequest
	GetDescription() *string
	SetEcsSpotStrategy(v string) *UpdateNodeGroupAttributesRequest
	GetEcsSpotStrategy() *string
	SetEnableGracefulDecommission(v bool) *UpdateNodeGroupAttributesRequest
	GetEnableGracefulDecommission() *bool
	SetInstanceTypeList(v []*string) *UpdateNodeGroupAttributesRequest
	GetInstanceTypeList() []*string
	SetKeyPairName(v string) *UpdateNodeGroupAttributesRequest
	GetKeyPairName() *string
	SetMaxSize(v int32) *UpdateNodeGroupAttributesRequest
	GetMaxSize() *int32
	SetMinSize(v int32) *UpdateNodeGroupAttributesRequest
	GetMinSize() *int32
	SetNodeCount(v int32) *UpdateNodeGroupAttributesRequest
	GetNodeCount() *int32
	SetNodeGroupId(v string) *UpdateNodeGroupAttributesRequest
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *UpdateNodeGroupAttributesRequest
	GetNodeGroupName() *string
	SetNodeResizeStrategy(v string) *UpdateNodeGroupAttributesRequest
	GetNodeResizeStrategy() *string
	SetRegionId(v string) *UpdateNodeGroupAttributesRequest
	GetRegionId() *string
	SetSpotBidPrices(v []*SpotBidPrice) *UpdateNodeGroupAttributesRequest
	GetSpotBidPrices() []*SpotBidPrice
	SetSpotInstanceRemedy(v bool) *UpdateNodeGroupAttributesRequest
	GetSpotInstanceRemedy() *bool
	SetVSwitchId(v string) *UpdateNodeGroupAttributesRequest
	GetVSwitchId() *string
}

type UpdateNodeGroupAttributesRequest struct {
	// The ACK cluster configuration.
	AckConfig *AckConfig `json:"AckConfig,omitempty" xml:"AckConfig,omitempty"`
	// Deprecated
	//
	// The list of security group IDs.
	//
	// example:
	//
	// ["sg-hp3abbae8lb6lmb1****"]
	AdditionalSecurityGroupIds []*string `json:"AdditionalSecurityGroupIds,omitempty" xml:"AdditionalSecurityGroupIds,omitempty" type:"Repeated"`
	// The automatic compensation state.
	//
	// example:
	//
	// false
	AutoCompensateState *bool `json:"AutoCompensateState,omitempty" xml:"AutoCompensateState,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cost-optimized mode configuration.
	CostOptimizedConfig *CostOptimizedConfig `json:"CostOptimizedConfig,omitempty" xml:"CostOptimizedConfig,omitempty"`
	// The node group description.
	//
	// example:
	//
	// emr-core-1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The spot instance strategy.
	//
	// example:
	//
	// SpotWithPriceLimit
	EcsSpotStrategy *string `json:"EcsSpotStrategy,omitempty" xml:"EcsSpotStrategy,omitempty"`
	// Specifies whether to enable graceful decommission. When a cluster is created in V303, graceful decommission is disabled by default for task groups.
	//
	// example:
	//
	// true
	EnableGracefulDecommission *bool `json:"EnableGracefulDecommission,omitempty" xml:"EnableGracefulDecommission,omitempty"`
	// The list of ECS instance types.
	//
	// example:
	//
	// null
	InstanceTypeList []*string `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The key pair for ECS logon.
	//
	// example:
	//
	// test_pair
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The maximum number of instances in the node group.
	//
	// example:
	//
	// 200
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The minimum number of instances in the node group.
	//
	// example:
	//
	// 0
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The target number of nodes in the node group.
	//
	// example:
	//
	// 8
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The node group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node group name.
	//
	// example:
	//
	// CORE1
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The node scale-out strategy. Valid values: COST_OPTIMIZED and PRIORITY. Default value: PRIORITY.
	//
	// example:
	//
	// PRIORITY
	NodeResizeStrategy *string `json:"NodeResizeStrategy,omitempty" xml:"NodeResizeStrategy,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum bid prices for spot instances. This parameter takes effect only when spotStrategy is set to SpotWithPriceLimit.
	SpotBidPrices []*SpotBidPrice `json:"SpotBidPrices,omitempty" xml:"SpotBidPrices,omitempty" type:"Repeated"`
	// Specifies whether to enable spot instance supplementation. If this feature is enabled, the scaling group attempts to create new instances to replace spot instances that are about to be reclaimed when the system sends a reclamation notification. Default value: false.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// Deprecated
	//
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-hp35g7ya5ymw68mmg****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateNodeGroupAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupAttributesRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupAttributesRequest) GetAckConfig() *AckConfig {
	return s.AckConfig
}

func (s *UpdateNodeGroupAttributesRequest) GetAdditionalSecurityGroupIds() []*string {
	return s.AdditionalSecurityGroupIds
}

func (s *UpdateNodeGroupAttributesRequest) GetAutoCompensateState() *bool {
	return s.AutoCompensateState
}

func (s *UpdateNodeGroupAttributesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateNodeGroupAttributesRequest) GetCostOptimizedConfig() *CostOptimizedConfig {
	return s.CostOptimizedConfig
}

func (s *UpdateNodeGroupAttributesRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateNodeGroupAttributesRequest) GetEcsSpotStrategy() *string {
	return s.EcsSpotStrategy
}

func (s *UpdateNodeGroupAttributesRequest) GetEnableGracefulDecommission() *bool {
	return s.EnableGracefulDecommission
}

func (s *UpdateNodeGroupAttributesRequest) GetInstanceTypeList() []*string {
	return s.InstanceTypeList
}

func (s *UpdateNodeGroupAttributesRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *UpdateNodeGroupAttributesRequest) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *UpdateNodeGroupAttributesRequest) GetMinSize() *int32 {
	return s.MinSize
}

func (s *UpdateNodeGroupAttributesRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *UpdateNodeGroupAttributesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateNodeGroupAttributesRequest) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *UpdateNodeGroupAttributesRequest) GetNodeResizeStrategy() *string {
	return s.NodeResizeStrategy
}

func (s *UpdateNodeGroupAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNodeGroupAttributesRequest) GetSpotBidPrices() []*SpotBidPrice {
	return s.SpotBidPrices
}

func (s *UpdateNodeGroupAttributesRequest) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *UpdateNodeGroupAttributesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateNodeGroupAttributesRequest) SetAckConfig(v *AckConfig) *UpdateNodeGroupAttributesRequest {
	s.AckConfig = v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetAdditionalSecurityGroupIds(v []*string) *UpdateNodeGroupAttributesRequest {
	s.AdditionalSecurityGroupIds = v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetAutoCompensateState(v bool) *UpdateNodeGroupAttributesRequest {
	s.AutoCompensateState = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetClusterId(v string) *UpdateNodeGroupAttributesRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetCostOptimizedConfig(v *CostOptimizedConfig) *UpdateNodeGroupAttributesRequest {
	s.CostOptimizedConfig = v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetDescription(v string) *UpdateNodeGroupAttributesRequest {
	s.Description = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetEcsSpotStrategy(v string) *UpdateNodeGroupAttributesRequest {
	s.EcsSpotStrategy = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetEnableGracefulDecommission(v bool) *UpdateNodeGroupAttributesRequest {
	s.EnableGracefulDecommission = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetInstanceTypeList(v []*string) *UpdateNodeGroupAttributesRequest {
	s.InstanceTypeList = v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetKeyPairName(v string) *UpdateNodeGroupAttributesRequest {
	s.KeyPairName = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetMaxSize(v int32) *UpdateNodeGroupAttributesRequest {
	s.MaxSize = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetMinSize(v int32) *UpdateNodeGroupAttributesRequest {
	s.MinSize = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetNodeCount(v int32) *UpdateNodeGroupAttributesRequest {
	s.NodeCount = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetNodeGroupId(v string) *UpdateNodeGroupAttributesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetNodeGroupName(v string) *UpdateNodeGroupAttributesRequest {
	s.NodeGroupName = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetNodeResizeStrategy(v string) *UpdateNodeGroupAttributesRequest {
	s.NodeResizeStrategy = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetRegionId(v string) *UpdateNodeGroupAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetSpotBidPrices(v []*SpotBidPrice) *UpdateNodeGroupAttributesRequest {
	s.SpotBidPrices = v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetSpotInstanceRemedy(v bool) *UpdateNodeGroupAttributesRequest {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) SetVSwitchId(v string) *UpdateNodeGroupAttributesRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateNodeGroupAttributesRequest) Validate() error {
	if s.AckConfig != nil {
		if err := s.AckConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CostOptimizedConfig != nil {
		if err := s.CostOptimizedConfig.Validate(); err != nil {
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
	return nil
}
