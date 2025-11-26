// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoScalingPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetConstraints(v *AutoScalingPolicyConstraints) *AutoScalingPolicy
	GetConstraints() *AutoScalingPolicyConstraints
	SetScalingRules(v []*ScalingRule) *AutoScalingPolicy
	GetScalingRules() []*ScalingRule
}

type AutoScalingPolicy struct {
	Constraints  *AutoScalingPolicyConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty" type:"Struct"`
	ScalingRules []*ScalingRule                `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Repeated"`
}

func (s AutoScalingPolicy) String() string {
	return dara.Prettify(s)
}

func (s AutoScalingPolicy) GoString() string {
	return s.String()
}

func (s *AutoScalingPolicy) GetConstraints() *AutoScalingPolicyConstraints {
	return s.Constraints
}

func (s *AutoScalingPolicy) GetScalingRules() []*ScalingRule {
	return s.ScalingRules
}

func (s *AutoScalingPolicy) SetConstraints(v *AutoScalingPolicyConstraints) *AutoScalingPolicy {
	s.Constraints = v
	return s
}

func (s *AutoScalingPolicy) SetScalingRules(v []*ScalingRule) *AutoScalingPolicy {
	s.ScalingRules = v
	return s
}

func (s *AutoScalingPolicy) Validate() error {
	if s.Constraints != nil {
		if err := s.Constraints.Validate(); err != nil {
			return err
		}
	}
	if s.ScalingRules != nil {
		for _, item := range s.ScalingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AutoScalingPolicyConstraints struct {
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s AutoScalingPolicyConstraints) String() string {
	return dara.Prettify(s)
}

func (s AutoScalingPolicyConstraints) GoString() string {
	return s.String()
}

func (s *AutoScalingPolicyConstraints) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *AutoScalingPolicyConstraints) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *AutoScalingPolicyConstraints) SetMaxCapacity(v int32) *AutoScalingPolicyConstraints {
	s.MaxCapacity = &v
	return s
}

func (s *AutoScalingPolicyConstraints) SetMinCapacity(v int32) *AutoScalingPolicyConstraints {
	s.MinCapacity = &v
	return s
}

func (s *AutoScalingPolicyConstraints) Validate() error {
	return dara.Validate(s)
}
