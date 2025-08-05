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
	// The description of the backup policy.
	//
	// example:
	//
	// Data is backed up at 10:00:00 every day and replicated to the China (Shanghai) region for geo-redundancy.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the backup policy.
	//
	// example:
	//
	// Daily Local Backup + Remote Backup
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The policy type. Valid values:
	//
	// 	- **STANDARD**: the general backup policy. This type of policy applies to backups other than Elastic Compute Service (ECS) instance backup.
	//
	// 	- **UDM_ECS_ONLY**: This type of policy applies only to ECS instance backup.
	//
	// If the policy type is not specified, Cloud Backup automatically sets the policy type based on whether the backup vault is specified in the rules of the policy:
	//
	// 	- If the backup vault is specified, Cloud Backup sets the policy type to **STANDARD**.
	//
	// 	- If the backup vault is not specified, Cloud Backup sets the policy type to **UDM_ECS_ONLY**.
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The rules in the backup policy.
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
