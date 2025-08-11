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
	// If `ActivityType` is set to `CapacityChange`, only the expected number of instances is changed during the scaling activity specified by ScalingActivityId and no scale-out is triggered.
	//
	// This parameter is applicable to only scaling groups that have an expected number of instances.
	//
	// example:
	//
	// CapacityChange
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// The elastic planning result returned when the ExecutionMode is set to PlanOnly.
	PlanResult *ScaleWithAdjustmentResponseBodyPlanResult `json:"PlanResult,omitempty" xml:"PlanResult,omitempty" type:"Struct"`
	// The ID of the request.
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
	return dara.Validate(s)
}

type ScaleWithAdjustmentResponseBodyPlanResult struct {
	// The resource allocation information in the elastic planning result.
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
	return dara.Validate(s)
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
	// 	- **Prepaid**: subscription.
	//
	// 	- **Postpaid**: pay-as-you-go.
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
	// The spot policy of instances. Valid values:
	//
	// 	- NoSpot: The instances are created as pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are created as spot instances for which you can specify the maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are spot instances for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the zone.
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
