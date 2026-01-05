// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityReservationState(v []*DescribeCapacityReservationResponseBodyCapacityReservationState) *DescribeCapacityReservationResponseBody
	GetCapacityReservationState() []*DescribeCapacityReservationResponseBodyCapacityReservationState
	SetDecreaseRequestsRemaining(v int32) *DescribeCapacityReservationResponseBody
	GetDecreaseRequestsRemaining() *int32
	SetLastModifiedTime(v string) *DescribeCapacityReservationResponseBody
	GetLastModifiedTime() *string
	SetMinimumLoadBalancerCapacity(v *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity) *DescribeCapacityReservationResponseBody
	GetMinimumLoadBalancerCapacity() *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity
	SetRequestId(v string) *DescribeCapacityReservationResponseBody
	GetRequestId() *string
}

type DescribeCapacityReservationResponseBody struct {
	CapacityReservationState []*DescribeCapacityReservationResponseBodyCapacityReservationState `json:"CapacityReservationState,omitempty" xml:"CapacityReservationState,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	DecreaseRequestsRemaining *int32 `json:"DecreaseRequestsRemaining,omitempty" xml:"DecreaseRequestsRemaining,omitempty"`
	// example:
	//
	// 2025-08-22 05:06:58
	LastModifiedTime            *string                                                             `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	MinimumLoadBalancerCapacity *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity `json:"MinimumLoadBalancerCapacity,omitempty" xml:"MinimumLoadBalancerCapacity,omitempty" type:"Struct"`
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCapacityReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationResponseBody) GetCapacityReservationState() []*DescribeCapacityReservationResponseBodyCapacityReservationState {
	return s.CapacityReservationState
}

func (s *DescribeCapacityReservationResponseBody) GetDecreaseRequestsRemaining() *int32 {
	return s.DecreaseRequestsRemaining
}

func (s *DescribeCapacityReservationResponseBody) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *DescribeCapacityReservationResponseBody) GetMinimumLoadBalancerCapacity() *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity {
	return s.MinimumLoadBalancerCapacity
}

func (s *DescribeCapacityReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCapacityReservationResponseBody) SetCapacityReservationState(v []*DescribeCapacityReservationResponseBodyCapacityReservationState) *DescribeCapacityReservationResponseBody {
	s.CapacityReservationState = v
	return s
}

func (s *DescribeCapacityReservationResponseBody) SetDecreaseRequestsRemaining(v int32) *DescribeCapacityReservationResponseBody {
	s.DecreaseRequestsRemaining = &v
	return s
}

func (s *DescribeCapacityReservationResponseBody) SetLastModifiedTime(v string) *DescribeCapacityReservationResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeCapacityReservationResponseBody) SetMinimumLoadBalancerCapacity(v *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity) *DescribeCapacityReservationResponseBody {
	s.MinimumLoadBalancerCapacity = v
	return s
}

func (s *DescribeCapacityReservationResponseBody) SetRequestId(v string) *DescribeCapacityReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCapacityReservationResponseBody) Validate() error {
	if s.CapacityReservationState != nil {
		for _, item := range s.CapacityReservationState {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MinimumLoadBalancerCapacity != nil {
		if err := s.MinimumLoadBalancerCapacity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCapacityReservationResponseBodyCapacityReservationState struct {
	// example:
	//
	// cn-hangzhou-k
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// example:
	//
	// 50.0
	EffectiveCapacityUnits *float64 `json:"EffectiveCapacityUnits,omitempty" xml:"EffectiveCapacityUnits,omitempty"`
	// example:
	//
	// INSUFFICIENT_CAPACITY
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// Provisioned
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCapacityReservationResponseBodyCapacityReservationState) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationResponseBodyCapacityReservationState) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) GetEffectiveCapacityUnits() *float64 {
	return s.EffectiveCapacityUnits
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) GetReason() *string {
	return s.Reason
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) GetStatus() *string {
	return s.Status
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) SetAvailabilityZone(v string) *DescribeCapacityReservationResponseBodyCapacityReservationState {
	s.AvailabilityZone = &v
	return s
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) SetEffectiveCapacityUnits(v float64) *DescribeCapacityReservationResponseBodyCapacityReservationState {
	s.EffectiveCapacityUnits = &v
	return s
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) SetReason(v string) *DescribeCapacityReservationResponseBodyCapacityReservationState {
	s.Reason = &v
	return s
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) SetStatus(v string) *DescribeCapacityReservationResponseBodyCapacityReservationState {
	s.Status = &v
	return s
}

func (s *DescribeCapacityReservationResponseBodyCapacityReservationState) Validate() error {
	return dara.Validate(s)
}

type DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity struct {
	// example:
	//
	// 100
	CapacityUnits *int32 `json:"CapacityUnits,omitempty" xml:"CapacityUnits,omitempty"`
}

func (s DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity) GetCapacityUnits() *int32 {
	return s.CapacityUnits
}

func (s *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity) SetCapacityUnits(v int32) *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity {
	s.CapacityUnits = &v
	return s
}

func (s *DescribeCapacityReservationResponseBodyMinimumLoadBalancerCapacity) Validate() error {
	return dara.Validate(s)
}
