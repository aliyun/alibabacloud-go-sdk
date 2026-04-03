// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingStatus interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentError(v string) *ScalingStatus
	GetCurrentError() *string
	SetCurrentInstances(v int64) *ScalingStatus
	GetCurrentInstances() *int64
	SetMinInstances(v int64) *ScalingStatus
	GetMinInstances() *int64
	SetScheduledPolicies(v []*ScheduledPolicy) *ScalingStatus
	GetScheduledPolicies() []*ScheduledPolicy
	SetTargetInstances(v int64) *ScalingStatus
	GetTargetInstances() *int64
}

type ScalingStatus struct {
	// example:
	//
	// error
	CurrentError *string `json:"currentError,omitempty" xml:"currentError,omitempty"`
	// example:
	//
	// 2
	CurrentInstances *int64 `json:"currentInstances,omitempty" xml:"currentInstances,omitempty"`
	// example:
	//
	// 1
	MinInstances      *int64             `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	ScheduledPolicies []*ScheduledPolicy `json:"scheduledPolicies" xml:"scheduledPolicies" type:"Repeated"`
	// example:
	//
	// 2
	TargetInstances *int64 `json:"targetInstances,omitempty" xml:"targetInstances,omitempty"`
}

func (s ScalingStatus) String() string {
	return dara.Prettify(s)
}

func (s ScalingStatus) GoString() string {
	return s.String()
}

func (s *ScalingStatus) GetCurrentError() *string {
	return s.CurrentError
}

func (s *ScalingStatus) GetCurrentInstances() *int64 {
	return s.CurrentInstances
}

func (s *ScalingStatus) GetMinInstances() *int64 {
	return s.MinInstances
}

func (s *ScalingStatus) GetScheduledPolicies() []*ScheduledPolicy {
	return s.ScheduledPolicies
}

func (s *ScalingStatus) GetTargetInstances() *int64 {
	return s.TargetInstances
}

func (s *ScalingStatus) SetCurrentError(v string) *ScalingStatus {
	s.CurrentError = &v
	return s
}

func (s *ScalingStatus) SetCurrentInstances(v int64) *ScalingStatus {
	s.CurrentInstances = &v
	return s
}

func (s *ScalingStatus) SetMinInstances(v int64) *ScalingStatus {
	s.MinInstances = &v
	return s
}

func (s *ScalingStatus) SetScheduledPolicies(v []*ScheduledPolicy) *ScalingStatus {
	s.ScheduledPolicies = v
	return s
}

func (s *ScalingStatus) SetTargetInstances(v int64) *ScalingStatus {
	s.TargetInstances = &v
	return s
}

func (s *ScalingStatus) Validate() error {
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
