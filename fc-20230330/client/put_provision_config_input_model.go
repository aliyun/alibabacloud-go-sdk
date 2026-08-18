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
	// Specifies whether to always allocate CPU. Default value: true.
	//
	// example:
	//
	// true
	AlwaysAllocateCPU *bool `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	// Specifies whether to always allocate GPU. Default value: true.
	//
	// example:
	//
	// true
	AlwaysAllocateGPU *bool `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	// The default minimum number of provisioned instances. Valid values: 0 to 10000.
	//
	// > - If no metric-based auto elastic policy or scheduled elastic policy is configured, the current minimum number of instances equals the minimum number of instances you configured.
	//
	// > - If you configured multiple elastic policies for the minimum number of instances, the system calculates the minimum number of instances triggered by each policy and uses the maximum value among the elastic policies that are effective at the current time as the current minimum number of instances.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5
	DefaultTarget *int64 `json:"defaultTarget,omitempty" xml:"defaultTarget,omitempty"`
	// The scheduled scaling configuration.
	ScheduledActions []*ScheduledAction `json:"scheduledActions" xml:"scheduledActions" type:"Repeated"`
	// Deprecated
	//
	// 	Notice: This parameter is no longer recommended. Use the defaultTarget parameter instead.</notice>
	//
	// The target number of provisioned resources. Valid values: 0 to 10000.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// The metric-based scaling policy configuration.
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
