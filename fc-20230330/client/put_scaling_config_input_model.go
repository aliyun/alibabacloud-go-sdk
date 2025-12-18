// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutScalingConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetEnableMixMode(v bool) *PutScalingConfigInput
	GetEnableMixMode() *bool
	SetEnableOnDemandScaling(v bool) *PutScalingConfigInput
	GetEnableOnDemandScaling() *bool
	SetHorizontalScalingPolicies(v []*ScalingPolicy) *PutScalingConfigInput
	GetHorizontalScalingPolicies() []*ScalingPolicy
	SetMinInstances(v int64) *PutScalingConfigInput
	GetMinInstances() *int64
	SetRequestDispatchPolicy(v string) *PutScalingConfigInput
	GetRequestDispatchPolicy() *string
	SetResidentPoolId(v string) *PutScalingConfigInput
	GetResidentPoolId() *string
	SetScheduledPolicies(v []*ScheduledPolicy) *PutScalingConfigInput
	GetScheduledPolicies() []*ScheduledPolicy
}

type PutScalingConfigInput struct {
	EnableMixMode             *bool              `json:"enableMixMode,omitempty" xml:"enableMixMode,omitempty"`
	EnableOnDemandScaling     *bool              `json:"enableOnDemandScaling,omitempty" xml:"enableOnDemandScaling,omitempty"`
	HorizontalScalingPolicies []*ScalingPolicy   `json:"horizontalScalingPolicies" xml:"horizontalScalingPolicies" type:"Repeated"`
	MinInstances              *int64             `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	RequestDispatchPolicy     *string            `json:"requestDispatchPolicy,omitempty" xml:"requestDispatchPolicy,omitempty"`
	ResidentPoolId            *string            `json:"residentPoolId,omitempty" xml:"residentPoolId,omitempty"`
	ScheduledPolicies         []*ScheduledPolicy `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
}

func (s PutScalingConfigInput) String() string {
	return dara.Prettify(s)
}

func (s PutScalingConfigInput) GoString() string {
	return s.String()
}

func (s *PutScalingConfigInput) GetEnableMixMode() *bool {
	return s.EnableMixMode
}

func (s *PutScalingConfigInput) GetEnableOnDemandScaling() *bool {
	return s.EnableOnDemandScaling
}

func (s *PutScalingConfigInput) GetHorizontalScalingPolicies() []*ScalingPolicy {
	return s.HorizontalScalingPolicies
}

func (s *PutScalingConfigInput) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *PutScalingConfigInput) GetRequestDispatchPolicy() *string {
	return s.RequestDispatchPolicy
}

func (s *PutScalingConfigInput) GetResidentPoolId() *string {
	return s.ResidentPoolId
}

func (s *PutScalingConfigInput) GetScheduledPolicies() []*ScheduledPolicy {
	return s.ScheduledPolicies
}

func (s *PutScalingConfigInput) SetEnableMixMode(v bool) *PutScalingConfigInput {
	s.EnableMixMode = &v
	return s
}

func (s *PutScalingConfigInput) SetEnableOnDemandScaling(v bool) *PutScalingConfigInput {
	s.EnableOnDemandScaling = &v
	return s
}

func (s *PutScalingConfigInput) SetHorizontalScalingPolicies(v []*ScalingPolicy) *PutScalingConfigInput {
	s.HorizontalScalingPolicies = v
	return s
}

func (s *PutScalingConfigInput) SetMinInstances(v int64) *PutScalingConfigInput {
	s.MinInstances = &v
	return s
}

func (s *PutScalingConfigInput) SetRequestDispatchPolicy(v string) *PutScalingConfigInput {
	s.RequestDispatchPolicy = &v
	return s
}

func (s *PutScalingConfigInput) SetResidentPoolId(v string) *PutScalingConfigInput {
	s.ResidentPoolId = &v
	return s
}

func (s *PutScalingConfigInput) SetScheduledPolicies(v []*ScheduledPolicy) *PutScalingConfigInput {
	s.ScheduledPolicies = v
	return s
}

func (s *PutScalingConfigInput) Validate() error {
	if s.HorizontalScalingPolicies != nil {
		for _, item := range s.HorizontalScalingPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduledPolicies != nil {
		for _, item := range s.ScheduledPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
