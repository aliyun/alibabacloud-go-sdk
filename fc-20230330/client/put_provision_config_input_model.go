// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutProvisionConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetAlwaysAllocateCPU(v bool) *PutProvisionConfigInput
	GetAlwaysAllocateCPU() *bool
	SetAlwaysAllocateGPU(v bool) *PutProvisionConfigInput
	GetAlwaysAllocateGPU() *bool
	SetDefaultTarget(v int64) *PutProvisionConfigInput
	GetDefaultTarget() *int64
	SetScheduledActions(v []*ScheduledAction) *PutProvisionConfigInput
	GetScheduledActions() []*ScheduledAction
	SetTarget(v int64) *PutProvisionConfigInput
	GetTarget() *int64
	SetTargetTrackingPolicies(v []*TargetTrackingPolicy) *PutProvisionConfigInput
	GetTargetTrackingPolicies() []*TargetTrackingPolicy
}

type PutProvisionConfigInput struct {
	// example:
	//
	// true
	AlwaysAllocateCPU *bool `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	// example:
	//
	// true
	AlwaysAllocateGPU *bool `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	// if can be null:
	// true
	DefaultTarget    *int64             `json:"defaultTarget,omitempty" xml:"defaultTarget,omitempty"`
	ScheduledActions []*ScheduledAction `json:"scheduledActions" xml:"scheduledActions" type:"Repeated"`
	// Deprecated
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	Target                 *int64                  `json:"target,omitempty" xml:"target,omitempty"`
	TargetTrackingPolicies []*TargetTrackingPolicy `json:"targetTrackingPolicies" xml:"targetTrackingPolicies" type:"Repeated"`
}

func (s PutProvisionConfigInput) String() string {
	return dara.Prettify(s)
}

func (s PutProvisionConfigInput) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigInput) GetAlwaysAllocateCPU() *bool {
	return s.AlwaysAllocateCPU
}

func (s *PutProvisionConfigInput) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *PutProvisionConfigInput) GetDefaultTarget() *int64 {
	return s.DefaultTarget
}

func (s *PutProvisionConfigInput) GetScheduledActions() []*ScheduledAction {
	return s.ScheduledActions
}

func (s *PutProvisionConfigInput) GetTarget() *int64 {
	return s.Target
}

func (s *PutProvisionConfigInput) GetTargetTrackingPolicies() []*TargetTrackingPolicy {
	return s.TargetTrackingPolicies
}

func (s *PutProvisionConfigInput) SetAlwaysAllocateCPU(v bool) *PutProvisionConfigInput {
	s.AlwaysAllocateCPU = &v
	return s
}

func (s *PutProvisionConfigInput) SetAlwaysAllocateGPU(v bool) *PutProvisionConfigInput {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *PutProvisionConfigInput) SetDefaultTarget(v int64) *PutProvisionConfigInput {
	s.DefaultTarget = &v
	return s
}

func (s *PutProvisionConfigInput) SetScheduledActions(v []*ScheduledAction) *PutProvisionConfigInput {
	s.ScheduledActions = v
	return s
}

func (s *PutProvisionConfigInput) SetTarget(v int64) *PutProvisionConfigInput {
	s.Target = &v
	return s
}

func (s *PutProvisionConfigInput) SetTargetTrackingPolicies(v []*TargetTrackingPolicy) *PutProvisionConfigInput {
	s.TargetTrackingPolicies = v
	return s
}

func (s *PutProvisionConfigInput) Validate() error {
	if s.ScheduledActions != nil {
		for _, item := range s.ScheduledActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TargetTrackingPolicies != nil {
		for _, item := range s.TargetTrackingPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
