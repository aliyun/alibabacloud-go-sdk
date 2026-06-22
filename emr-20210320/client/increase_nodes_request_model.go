// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v []*ApplicationConfig) *IncreaseNodesRequest
	GetApplicationConfigs() []*ApplicationConfig
	SetAutoPayOrder(v bool) *IncreaseNodesRequest
	GetAutoPayOrder() *bool
	SetAutoRenew(v bool) *IncreaseNodesRequest
	GetAutoRenew() *bool
	SetClusterId(v string) *IncreaseNodesRequest
	GetClusterId() *string
	SetIncreaseNodeCount(v int32) *IncreaseNodesRequest
	GetIncreaseNodeCount() *int32
	SetMinIncreaseNodeCount(v int32) *IncreaseNodesRequest
	GetMinIncreaseNodeCount() *int32
	SetNodeGroupId(v string) *IncreaseNodesRequest
	GetNodeGroupId() *string
	SetPaymentDuration(v int32) *IncreaseNodesRequest
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *IncreaseNodesRequest
	GetPaymentDurationUnit() *string
	SetPromotions(v []*Promotion) *IncreaseNodesRequest
	GetPromotions() []*Promotion
	SetRegionId(v string) *IncreaseNodesRequest
	GetRegionId() *string
}

type IncreaseNodesRequest struct {
	// The application configurations. The number of array elements can range from 1 to 1,000.
	//
	// example:
	//
	// Month
	ApplicationConfigs []*ApplicationConfig `json:"ApplicationConfigs,omitempty" xml:"ApplicationConfigs,omitempty" type:"Repeated"`
	// Specifies whether to automatically pay for the order. This parameter is effective only when the PaymentType of the node group is Subscription. Valid values:
	//
	// - true: Automatically pays for the order.
	//
	// - false: Does not automatically pay for the order.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoPayOrder *bool `json:"AutoPayOrder,omitempty" xml:"AutoPayOrder,omitempty"`
	// Specifies whether to enable auto-renewal for the added nodes. The default value is false. Valid values:
	//
	// - true: Enables auto-renewal.
	//
	// - false: Disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of nodes to add. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	IncreaseNodeCount *int32 `json:"IncreaseNodeCount,omitempty" xml:"IncreaseNodeCount,omitempty"`
	// The minimum number of nodes to add in this scale-out operation. The value must be between 1 and 500.
	//
	// - If you set this parameter and the available ECS inventory is less than IncreaseNodeCount, the system attempts to create at least the number of nodes specified by MinIncreaseNodeCount. The scale-out flow is then marked as `PARTIAL_COMPLETED`.
	//
	// - If you do not set this parameter and the available ECS inventory is less than IncreaseNodeCount, the scale-out flow fails and is marked as `FAILED`.
	MinIncreaseNodeCount *int32 `json:"MinIncreaseNodeCount,omitempty" xml:"MinIncreaseNodeCount,omitempty"`
	// The ID of the node group to scale out.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The subscription duration. If PaymentDurationUnit is set to Month, valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 12
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// - Month: The unit is month.
	//
	// example:
	//
	// Month
	PaymentDurationUnit *string      `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	Promotions          []*Promotion `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s IncreaseNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s IncreaseNodesRequest) GoString() string {
	return s.String()
}

func (s *IncreaseNodesRequest) GetApplicationConfigs() []*ApplicationConfig {
	return s.ApplicationConfigs
}

func (s *IncreaseNodesRequest) GetAutoPayOrder() *bool {
	return s.AutoPayOrder
}

func (s *IncreaseNodesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *IncreaseNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *IncreaseNodesRequest) GetIncreaseNodeCount() *int32 {
	return s.IncreaseNodeCount
}

func (s *IncreaseNodesRequest) GetMinIncreaseNodeCount() *int32 {
	return s.MinIncreaseNodeCount
}

func (s *IncreaseNodesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *IncreaseNodesRequest) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *IncreaseNodesRequest) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *IncreaseNodesRequest) GetPromotions() []*Promotion {
	return s.Promotions
}

func (s *IncreaseNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *IncreaseNodesRequest) SetApplicationConfigs(v []*ApplicationConfig) *IncreaseNodesRequest {
	s.ApplicationConfigs = v
	return s
}

func (s *IncreaseNodesRequest) SetAutoPayOrder(v bool) *IncreaseNodesRequest {
	s.AutoPayOrder = &v
	return s
}

func (s *IncreaseNodesRequest) SetAutoRenew(v bool) *IncreaseNodesRequest {
	s.AutoRenew = &v
	return s
}

func (s *IncreaseNodesRequest) SetClusterId(v string) *IncreaseNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *IncreaseNodesRequest) SetIncreaseNodeCount(v int32) *IncreaseNodesRequest {
	s.IncreaseNodeCount = &v
	return s
}

func (s *IncreaseNodesRequest) SetMinIncreaseNodeCount(v int32) *IncreaseNodesRequest {
	s.MinIncreaseNodeCount = &v
	return s
}

func (s *IncreaseNodesRequest) SetNodeGroupId(v string) *IncreaseNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *IncreaseNodesRequest) SetPaymentDuration(v int32) *IncreaseNodesRequest {
	s.PaymentDuration = &v
	return s
}

func (s *IncreaseNodesRequest) SetPaymentDurationUnit(v string) *IncreaseNodesRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *IncreaseNodesRequest) SetPromotions(v []*Promotion) *IncreaseNodesRequest {
	s.Promotions = v
	return s
}

func (s *IncreaseNodesRequest) SetRegionId(v string) *IncreaseNodesRequest {
	s.RegionId = &v
	return s
}

func (s *IncreaseNodesRequest) Validate() error {
	if s.ApplicationConfigs != nil {
		for _, item := range s.ApplicationConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Promotions != nil {
		for _, item := range s.Promotions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
