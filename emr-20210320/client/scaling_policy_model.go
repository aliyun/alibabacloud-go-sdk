// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ScalingPolicy
	GetClusterId() *string
	SetConstraints(v *ManagedScalingConstraints) *ScalingPolicy
	GetConstraints() *ManagedScalingConstraints
	SetDisabled(v bool) *ScalingPolicy
	GetDisabled() *bool
	SetNodeGroupId(v string) *ScalingPolicy
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *ScalingPolicy
	GetNodeGroupName() *string
	SetScalingPolicyId(v string) *ScalingPolicy
	GetScalingPolicyId() *string
	SetScalingPolicyType(v string) *ScalingPolicy
	GetScalingPolicyType() *string
	SetScalingRules(v []*ScalingRule) *ScalingPolicy
	GetScalingRules() []*ScalingRule
}

type ScalingPolicy struct {
	ClusterId       *string                    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Constraints     *ManagedScalingConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	Disabled        *bool                      `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	NodeGroupId     *string                    `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName   *string                    `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	ScalingPolicyId *string                    `json:"ScalingPolicyId,omitempty" xml:"ScalingPolicyId,omitempty"`
	// example:
	//
	// AUTO / MANAGED
	ScalingPolicyType *string        `json:"ScalingPolicyType,omitempty" xml:"ScalingPolicyType,omitempty"`
	ScalingRules      []*ScalingRule `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Repeated"`
}

func (s ScalingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ScalingPolicy) GoString() string {
	return s.String()
}

func (s *ScalingPolicy) GetClusterId() *string {
	return s.ClusterId
}

func (s *ScalingPolicy) GetConstraints() *ManagedScalingConstraints {
	return s.Constraints
}

func (s *ScalingPolicy) GetDisabled() *bool {
	return s.Disabled
}

func (s *ScalingPolicy) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ScalingPolicy) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ScalingPolicy) GetScalingPolicyId() *string {
	return s.ScalingPolicyId
}

func (s *ScalingPolicy) GetScalingPolicyType() *string {
	return s.ScalingPolicyType
}

func (s *ScalingPolicy) GetScalingRules() []*ScalingRule {
	return s.ScalingRules
}

func (s *ScalingPolicy) SetClusterId(v string) *ScalingPolicy {
	s.ClusterId = &v
	return s
}

func (s *ScalingPolicy) SetConstraints(v *ManagedScalingConstraints) *ScalingPolicy {
	s.Constraints = v
	return s
}

func (s *ScalingPolicy) SetDisabled(v bool) *ScalingPolicy {
	s.Disabled = &v
	return s
}

func (s *ScalingPolicy) SetNodeGroupId(v string) *ScalingPolicy {
	s.NodeGroupId = &v
	return s
}

func (s *ScalingPolicy) SetNodeGroupName(v string) *ScalingPolicy {
	s.NodeGroupName = &v
	return s
}

func (s *ScalingPolicy) SetScalingPolicyId(v string) *ScalingPolicy {
	s.ScalingPolicyId = &v
	return s
}

func (s *ScalingPolicy) SetScalingPolicyType(v string) *ScalingPolicy {
	s.ScalingPolicyType = &v
	return s
}

func (s *ScalingPolicy) SetScalingRules(v []*ScalingRule) *ScalingPolicy {
	s.ScalingRules = v
	return s
}

func (s *ScalingPolicy) Validate() error {
	return dara.Validate(s)
}
