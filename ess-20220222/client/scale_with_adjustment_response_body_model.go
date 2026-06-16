// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleWithAdjustmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivityType(v string) *ScaleWithAdjustmentResponseBody
	GetActivityType() *string
	SetPlanResult(v *ScaleWithAdjustmentResponseBodyPlanResult) *ScaleWithAdjustmentResponseBody
	GetPlanResult() *ScaleWithAdjustmentResponseBodyPlanResult
	SetRequestId(v string) *ScaleWithAdjustmentResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *ScaleWithAdjustmentResponseBody
	GetScalingActivityId() *string
}

type ScaleWithAdjustmentResponseBody struct {
	// The type of the scaling activity.
	//
	// If this parameter is set to `CapacityChange`, the scaling activity only adjusts the desired capacity of the scaling group without immediately adding or removing instances.
	//
	// This setting only affects scaling groups with a configured desired capacity.
	//
	// example:
	//
	// CapacityChange
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// The scaling plan result returned when ExecutionMode is set to PlanOnly.
	PlanResult *ScaleWithAdjustmentResponseBodyPlanResult `json:"PlanResult,omitempty" xml:"PlanResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-bp175o6f6ego3r2j****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s ScaleWithAdjustmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentResponseBody) GetActivityType() *string {
	return s.ActivityType
}

func (s *ScaleWithAdjustmentResponseBody) GetPlanResult() *ScaleWithAdjustmentResponseBodyPlanResult {
	return s.PlanResult
}

func (s *ScaleWithAdjustmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleWithAdjustmentResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *ScaleWithAdjustmentResponseBody) SetActivityType(v string) *ScaleWithAdjustmentResponseBody {
	s.ActivityType = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBody) SetPlanResult(v *ScaleWithAdjustmentResponseBodyPlanResult) *ScaleWithAdjustmentResponseBody {
	s.PlanResult = v
	return s
}

func (s *ScaleWithAdjustmentResponseBody) SetRequestId(v string) *ScaleWithAdjustmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBody) SetScalingActivityId(v string) *ScaleWithAdjustmentResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBody) Validate() error {
	if s.PlanResult != nil {
		if err := s.PlanResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScaleWithAdjustmentResponseBodyPlanResult struct {
	// The resource allocation details in the scaling plan result.
	ResourceAllocations []*ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations `json:"ResourceAllocations,omitempty" xml:"ResourceAllocations,omitempty" type:"Repeated"`
}

func (s ScaleWithAdjustmentResponseBodyPlanResult) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentResponseBodyPlanResult) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentResponseBodyPlanResult) GetResourceAllocations() []*ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations {
	return s.ResourceAllocations
}

func (s *ScaleWithAdjustmentResponseBodyPlanResult) SetResourceAllocations(v []*ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) *ScaleWithAdjustmentResponseBodyPlanResult {
	s.ResourceAllocations = v
	return s
}

func (s *ScaleWithAdjustmentResponseBodyPlanResult) Validate() error {
	if s.ResourceAllocations != nil {
		for _, item := range s.ResourceAllocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations struct {
	// The number of instances.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Prepaid**: subscription
	//
	// - **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.u1-c1m8.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The spot strategy of the instance. Valid values:
	//
	// - NoSpot: A pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: A spot instance with a user-specified price limit.
	//
	// - SpotAsPriceGo: A spot instance where the system automatically bids based on the current market price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The availability zone ID.
	//
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) GetAmount() *int32 {
	return s.Amount
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) GetZoneId() *string {
	return s.ZoneId
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) SetAmount(v int32) *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations {
	s.Amount = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) SetInstanceChargeType(v string) *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations {
	s.InstanceChargeType = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) SetInstanceType(v string) *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations {
	s.InstanceType = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) SetSpotStrategy(v string) *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations {
	s.SpotStrategy = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) SetZoneId(v string) *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations {
	s.ZoneId = &v
	return s
}

func (s *ScaleWithAdjustmentResponseBodyPlanResultResourceAllocations) Validate() error {
	return dara.Validate(s)
}
