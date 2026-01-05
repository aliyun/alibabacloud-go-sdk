// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCapacityReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyCapacityReservationRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyCapacityReservationRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *ModifyCapacityReservationRequest
	GetLoadBalancerId() *string
	SetMinimumLoadBalancerCapacity(v *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity) *ModifyCapacityReservationRequest
	GetMinimumLoadBalancerCapacity() *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity
	SetResetCapacityReservation(v bool) *ModifyCapacityReservationRequest
	GetResetCapacityReservation() *bool
}

type ModifyCapacityReservationRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alb-iv9gj3spak6fbj****
	LoadBalancerId              *string                                                      `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	MinimumLoadBalancerCapacity *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity `json:"MinimumLoadBalancerCapacity,omitempty" xml:"MinimumLoadBalancerCapacity,omitempty" type:"Struct"`
	// example:
	//
	// false
	ResetCapacityReservation *bool `json:"ResetCapacityReservation,omitempty" xml:"ResetCapacityReservation,omitempty"`
}

func (s ModifyCapacityReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *ModifyCapacityReservationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyCapacityReservationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyCapacityReservationRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ModifyCapacityReservationRequest) GetMinimumLoadBalancerCapacity() *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity {
	return s.MinimumLoadBalancerCapacity
}

func (s *ModifyCapacityReservationRequest) GetResetCapacityReservation() *bool {
	return s.ResetCapacityReservation
}

func (s *ModifyCapacityReservationRequest) SetClientToken(v string) *ModifyCapacityReservationRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetDryRun(v bool) *ModifyCapacityReservationRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetLoadBalancerId(v string) *ModifyCapacityReservationRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetMinimumLoadBalancerCapacity(v *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity) *ModifyCapacityReservationRequest {
	s.MinimumLoadBalancerCapacity = v
	return s
}

func (s *ModifyCapacityReservationRequest) SetResetCapacityReservation(v bool) *ModifyCapacityReservationRequest {
	s.ResetCapacityReservation = &v
	return s
}

func (s *ModifyCapacityReservationRequest) Validate() error {
	if s.MinimumLoadBalancerCapacity != nil {
		if err := s.MinimumLoadBalancerCapacity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCapacityReservationRequestMinimumLoadBalancerCapacity struct {
	// example:
	//
	// 100
	CapacityUnits *int32 `json:"CapacityUnits,omitempty" xml:"CapacityUnits,omitempty"`
}

func (s ModifyCapacityReservationRequestMinimumLoadBalancerCapacity) String() string {
	return dara.Prettify(s)
}

func (s ModifyCapacityReservationRequestMinimumLoadBalancerCapacity) GoString() string {
	return s.String()
}

func (s *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity) GetCapacityUnits() *int32 {
	return s.CapacityUnits
}

func (s *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity) SetCapacityUnits(v int32) *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity {
	s.CapacityUnits = &v
	return s
}

func (s *ModifyCapacityReservationRequestMinimumLoadBalancerCapacity) Validate() error {
	return dara.Validate(s)
}
