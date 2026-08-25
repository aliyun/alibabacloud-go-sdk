// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyDescription(v string) *CreatePolicyV2ShrinkRequest
	GetPolicyDescription() *string
	SetPolicyName(v string) *CreatePolicyV2ShrinkRequest
	GetPolicyName() *string
	SetPolicyType(v string) *CreatePolicyV2ShrinkRequest
	GetPolicyType() *string
	SetRulesShrink(v string) *CreatePolicyV2ShrinkRequest
	GetRulesShrink() *string
}

type CreatePolicyV2ShrinkRequest struct {
	// The policy description.
	//
	// example:
	//
	// Backup once every day at 10:00 AM, with cross-region backup to Shanghai.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Daily local backup + geo-redundancy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The policy type. Valid values:
	//
	// - **STANDARD**: general backup policy. Supports backing up data sources other than ECS instances.
	//
	// - **UDM_ECS_ONLY**: ECS instance backup policy. Supports backing up only ECS instances.
	//
	// If you do not specify the policy type, Cloud Backup automatically sets the policy type based on whether a backup vault is specified in the policy rules:
	//
	// - A backup vault is specified in the policy rules: **STANDARD**
	//
	// - No backup vault is specified in the policy rules: **UDM_ECS_ONLY**
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The list of policy rules.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreatePolicyV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2ShrinkRequest) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *CreatePolicyV2ShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyV2ShrinkRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePolicyV2ShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreatePolicyV2ShrinkRequest) SetPolicyDescription(v string) *CreatePolicyV2ShrinkRequest {
	s.PolicyDescription = &v
	return s
}

func (s *CreatePolicyV2ShrinkRequest) SetPolicyName(v string) *CreatePolicyV2ShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyV2ShrinkRequest) SetPolicyType(v string) *CreatePolicyV2ShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *CreatePolicyV2ShrinkRequest) SetRulesShrink(v string) *CreatePolicyV2ShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreatePolicyV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
