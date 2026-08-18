// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAlwaysAllocateCPU(v bool) *ProvisionConfig
	GetAlwaysAllocateCPU() *bool
	SetAlwaysAllocateGPU(v bool) *ProvisionConfig
	GetAlwaysAllocateGPU() *bool
	SetCurrent(v int64) *ProvisionConfig
	GetCurrent() *int64
	SetCurrentError(v string) *ProvisionConfig
	GetCurrentError() *string
	SetDefaultTarget(v int64) *ProvisionConfig
	GetDefaultTarget() *int64
	SetFunctionArn(v string) *ProvisionConfig
	GetFunctionArn() *string
	SetScheduledActions(v []*ScheduledAction) *ProvisionConfig
	GetScheduledActions() []*ScheduledAction
	SetTarget(v int64) *ProvisionConfig
	GetTarget() *int64
	SetTargetTrackingPolicies(v []*TargetTrackingPolicy) *ProvisionConfig
	GetTargetTrackingPolicies() []*TargetTrackingPolicy
}

type ProvisionConfig struct {
	// Specifies whether to always allocate CPU to function instances.
	//
	// example:
	//
	// true
	AlwaysAllocateCPU *bool `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	// Specifies whether to always allocate GPU to function instances.
	//
	// example:
	//
	// true
	AlwaysAllocateGPU *bool `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	// The actual number of resources.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// The error message when provisioned instance creation fails.
	//
	// example:
	//
	// image not found
	CurrentError *string `json:"currentError,omitempty" xml:"currentError,omitempty"`
	// The default number of resources when all metric-based scaling policies and scheduled scaling policies are inactive.
	//
	// example:
	//
	// 5
	DefaultTarget *int64 `json:"defaultTarget,omitempty" xml:"defaultTarget,omitempty"`
	// The resource descriptor of the function.
	//
	// example:
	//
	// acs:fc:cn-shanghai:124:functions/myFunction/prod
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// The scheduled scaling policy configurations.
	ScheduledActions []*ScheduledAction `json:"scheduledActions" xml:"scheduledActions" type:"Repeated"`
	// The current target number of resources. If a metric-based scaling policy or scheduled scaling policy exists, this value is the number of resources calculated by the policy. Otherwise, it is the default number of provisioned instances.
	//
	//
	// > What is the difference between target and defaultTarget?\\
	//
	// > Assume that the number of provisioned instances is configured as 1, and then a scheduled scaling policy is added to set the number of provisioned instances to 5 during a specific time period.
	//
	// > - During the **active period*	- of the scheduled scaling policy, target and defaultTarget are 5 and 1, respectively.
	//
	// >-  During the **inactive period*	- of the scheduled scaling policy, both target and defaultTarget are 1.
	//
	// example:
	//
	// 5
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// The metric-based scaling policy configurations.
	TargetTrackingPolicies []*TargetTrackingPolicy `json:"targetTrackingPolicies" xml:"targetTrackingPolicies" type:"Repeated"`
}

func (s ProvisionConfig) String() string {
	return dara.Prettify(s)
}

func (s ProvisionConfig) GoString() string {
	return s.String()
}

func (s *ProvisionConfig) GetAlwaysAllocateCPU() *bool {
	return s.AlwaysAllocateCPU
}

func (s *ProvisionConfig) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *ProvisionConfig) GetCurrent() *int64 {
	return s.Current
}

func (s *ProvisionConfig) GetCurrentError() *string {
	return s.CurrentError
}

func (s *ProvisionConfig) GetDefaultTarget() *int64 {
	return s.DefaultTarget
}

func (s *ProvisionConfig) GetFunctionArn() *string {
	return s.FunctionArn
}

func (s *ProvisionConfig) GetScheduledActions() []*ScheduledAction {
	return s.ScheduledActions
}

func (s *ProvisionConfig) GetTarget() *int64 {
	return s.Target
}

func (s *ProvisionConfig) GetTargetTrackingPolicies() []*TargetTrackingPolicy {
	return s.TargetTrackingPolicies
}

func (s *ProvisionConfig) SetAlwaysAllocateCPU(v bool) *ProvisionConfig {
	s.AlwaysAllocateCPU = &v
	return s
}

func (s *ProvisionConfig) SetAlwaysAllocateGPU(v bool) *ProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *ProvisionConfig) SetCurrent(v int64) *ProvisionConfig {
	s.Current = &v
	return s
}

func (s *ProvisionConfig) SetCurrentError(v string) *ProvisionConfig {
	s.CurrentError = &v
	return s
}

func (s *ProvisionConfig) SetDefaultTarget(v int64) *ProvisionConfig {
	s.DefaultTarget = &v
	return s
}

func (s *ProvisionConfig) SetFunctionArn(v string) *ProvisionConfig {
	s.FunctionArn = &v
	return s
}

func (s *ProvisionConfig) SetScheduledActions(v []*ScheduledAction) *ProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *ProvisionConfig) SetTarget(v int64) *ProvisionConfig {
	s.Target = &v
	return s
}

func (s *ProvisionConfig) SetTargetTrackingPolicies(v []*TargetTrackingPolicy) *ProvisionConfig {
	s.TargetTrackingPolicies = v
	return s
}

func (s *ProvisionConfig) Validate() error {
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
