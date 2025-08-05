// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyDescription(v string) *UpdatePolicyV2ShrinkRequest
	GetPolicyDescription() *string
	SetPolicyId(v string) *UpdatePolicyV2ShrinkRequest
	GetPolicyId() *string
	SetPolicyName(v string) *UpdatePolicyV2ShrinkRequest
	GetPolicyName() *string
	SetRulesShrink(v string) *UpdatePolicyV2ShrinkRequest
	GetRulesShrink() *string
}

type UpdatePolicyV2ShrinkRequest struct {
	// The description of the backup policy.
	//
	// example:
	//
	// Data is backed up at 10:00:00 every day and replicated to the China (Shanghai) region for geo-redundancy.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The ID of the backup policy.
	//
	// example:
	//
	// po-000************viy
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the backup policy.
	//
	// example:
	//
	// Daily Local Backup + Remote Backup
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The rules in the backup policy.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s UpdatePolicyV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2ShrinkRequest) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *UpdatePolicyV2ShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdatePolicyV2ShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdatePolicyV2ShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *UpdatePolicyV2ShrinkRequest) SetPolicyDescription(v string) *UpdatePolicyV2ShrinkRequest {
	s.PolicyDescription = &v
	return s
}

func (s *UpdatePolicyV2ShrinkRequest) SetPolicyId(v string) *UpdatePolicyV2ShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdatePolicyV2ShrinkRequest) SetPolicyName(v string) *UpdatePolicyV2ShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyV2ShrinkRequest) SetRulesShrink(v string) *UpdatePolicyV2ShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *UpdatePolicyV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
