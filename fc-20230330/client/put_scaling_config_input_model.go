// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutScalingConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetHorizontalScalingPolicies(v []*ScalingPolicy) *PutScalingConfigInput
	GetHorizontalScalingPolicies() []*ScalingPolicy
	SetMinInstances(v int64) *PutScalingConfigInput
	GetMinInstances() *int64
	SetResidentPoolId(v string) *PutScalingConfigInput
	GetResidentPoolId() *string
	SetScheduledPolicies(v []*ScheduledPolicy) *PutScalingConfigInput
	GetScheduledPolicies() []*ScheduledPolicy
}

type PutScalingConfigInput struct {
	HorizontalScalingPolicies []*ScalingPolicy   `json:"horizontalScalingPolicies" xml:"horizontalScalingPolicies" type:"Repeated"`
	MinInstances              *int64             `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	ResidentPoolId            *string            `json:"residentPoolId,omitempty" xml:"residentPoolId,omitempty"`
	ScheduledPolicies         []*ScheduledPolicy `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
}

func (s PutScalingConfigInput) String() string {
	return dara.Prettify(s)
}

func (s PutScalingConfigInput) GoString() string {
	return s.String()
}

func (s *PutScalingConfigInput) GetHorizontalScalingPolicies() []*ScalingPolicy {
	return s.HorizontalScalingPolicies
}

func (s *PutScalingConfigInput) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *PutScalingConfigInput) GetResidentPoolId() *string {
	return s.ResidentPoolId
}

func (s *PutScalingConfigInput) GetScheduledPolicies() []*ScheduledPolicy {
	return s.ScheduledPolicies
}

func (s *PutScalingConfigInput) SetHorizontalScalingPolicies(v []*ScalingPolicy) *PutScalingConfigInput {
	s.HorizontalScalingPolicies = v
	return s
}

func (s *PutScalingConfigInput) SetMinInstances(v int64) *PutScalingConfigInput {
	s.MinInstances = &v
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
	return dara.Validate(s)
}
