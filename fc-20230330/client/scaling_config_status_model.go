// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingConfigStatus interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentError(v string) *ScalingConfigStatus
	GetCurrentError() *string
	SetCurrentInstances(v int64) *ScalingConfigStatus
	GetCurrentInstances() *int64
	SetEnableOnDemandScaling(v bool) *ScalingConfigStatus
	GetEnableOnDemandScaling() *bool
	SetFunctionArn(v string) *ScalingConfigStatus
	GetFunctionArn() *string
	SetHorizontalScalingPolicies(v []*ScalingPolicy) *ScalingConfigStatus
	GetHorizontalScalingPolicies() []*ScalingPolicy
	SetMinInstances(v int64) *ScalingConfigStatus
	GetMinInstances() *int64
	SetResidentPoolId(v string) *ScalingConfigStatus
	GetResidentPoolId() *string
	SetScheduledPolicies(v []*ScheduledPolicy) *ScalingConfigStatus
	GetScheduledPolicies() []*ScheduledPolicy
	SetTargetInstances(v int64) *ScalingConfigStatus
	GetTargetInstances() *int64
}

type ScalingConfigStatus struct {
	CurrentError              *string            `json:"currentError,omitempty" xml:"currentError,omitempty"`
	CurrentInstances          *int64             `json:"currentInstances,omitempty" xml:"currentInstances,omitempty"`
	EnableOnDemandScaling     *bool              `json:"enableOnDemandScaling,omitempty" xml:"enableOnDemandScaling,omitempty"`
	FunctionArn               *string            `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	HorizontalScalingPolicies []*ScalingPolicy   `json:"horizontalScalingPolicies" xml:"horizontalScalingPolicies" type:"Repeated"`
	MinInstances              *int64             `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	ResidentPoolId            *string            `json:"residentPoolId,omitempty" xml:"residentPoolId,omitempty"`
	ScheduledPolicies         []*ScheduledPolicy `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
	TargetInstances           *int64             `json:"targetInstances,omitempty" xml:"targetInstances,omitempty"`
}

func (s ScalingConfigStatus) String() string {
	return dara.Prettify(s)
}

func (s ScalingConfigStatus) GoString() string {
	return s.String()
}

func (s *ScalingConfigStatus) GetCurrentError() *string {
	return s.CurrentError
}

func (s *ScalingConfigStatus) GetCurrentInstances() *int64 {
	return s.CurrentInstances
}

func (s *ScalingConfigStatus) GetEnableOnDemandScaling() *bool {
	return s.EnableOnDemandScaling
}

func (s *ScalingConfigStatus) GetFunctionArn() *string {
	return s.FunctionArn
}

func (s *ScalingConfigStatus) GetHorizontalScalingPolicies() []*ScalingPolicy {
	return s.HorizontalScalingPolicies
}

func (s *ScalingConfigStatus) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *ScalingConfigStatus) GetResidentPoolId() *string {
	return s.ResidentPoolId
}

func (s *ScalingConfigStatus) GetScheduledPolicies() []*ScheduledPolicy {
	return s.ScheduledPolicies
}

func (s *ScalingConfigStatus) GetTargetInstances() *int64 {
	return s.TargetInstances
}

func (s *ScalingConfigStatus) SetCurrentError(v string) *ScalingConfigStatus {
	s.CurrentError = &v
	return s
}

func (s *ScalingConfigStatus) SetCurrentInstances(v int64) *ScalingConfigStatus {
	s.CurrentInstances = &v
	return s
}

func (s *ScalingConfigStatus) SetEnableOnDemandScaling(v bool) *ScalingConfigStatus {
	s.EnableOnDemandScaling = &v
	return s
}

func (s *ScalingConfigStatus) SetFunctionArn(v string) *ScalingConfigStatus {
	s.FunctionArn = &v
	return s
}

func (s *ScalingConfigStatus) SetHorizontalScalingPolicies(v []*ScalingPolicy) *ScalingConfigStatus {
	s.HorizontalScalingPolicies = v
	return s
}

func (s *ScalingConfigStatus) SetMinInstances(v int64) *ScalingConfigStatus {
	s.MinInstances = &v
	return s
}

func (s *ScalingConfigStatus) SetResidentPoolId(v string) *ScalingConfigStatus {
	s.ResidentPoolId = &v
	return s
}

func (s *ScalingConfigStatus) SetScheduledPolicies(v []*ScheduledPolicy) *ScalingConfigStatus {
	s.ScheduledPolicies = v
	return s
}

func (s *ScalingConfigStatus) SetTargetInstances(v int64) *ScalingConfigStatus {
	s.TargetInstances = &v
	return s
}

func (s *ScalingConfigStatus) Validate() error {
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
