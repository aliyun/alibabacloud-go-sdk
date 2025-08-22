// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnWafPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *CreateDcdnWafPolicyRequest
	GetDefenseScene() *string
	SetPolicyName(v string) *CreateDcdnWafPolicyRequest
	GetPolicyName() *string
	SetPolicyStatus(v string) *CreateDcdnWafPolicyRequest
	GetPolicyStatus() *string
	SetPolicyType(v string) *CreateDcdnWafPolicyRequest
	GetPolicyType() *string
}

type CreateDcdnWafPolicyRequest struct {
	// The type of the WAF protection policy. Valid values:
	//
	// 	- waf_group: basic web protection
	//
	// 	- custom_acl: custom protection
	//
	// 	- whitelist: IP address whitelist
	//
	// 	- ip_blacklist: IP address blacklist
	//
	// 	- region_block: region blacklist
	//
	// 	- bot: bot management
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The name of the protection policy. The name can be up to 64 characters in length and can contain letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The status of the protection policy. Valid values:
	//
	// 	- on: The policy is enabled.
	//
	// 	- off: The policy is disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// Specifies whether to set the current policy as the default policy. Valid values:
	//
	// 	- default: sets the current policy as the default policy.
	//
	// 	- custom: does not set the current policy as the default policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateDcdnWafPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnWafPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnWafPolicyRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *CreateDcdnWafPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateDcdnWafPolicyRequest) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *CreateDcdnWafPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateDcdnWafPolicyRequest) SetDefenseScene(v string) *CreateDcdnWafPolicyRequest {
	s.DefenseScene = &v
	return s
}

func (s *CreateDcdnWafPolicyRequest) SetPolicyName(v string) *CreateDcdnWafPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateDcdnWafPolicyRequest) SetPolicyStatus(v string) *CreateDcdnWafPolicyRequest {
	s.PolicyStatus = &v
	return s
}

func (s *CreateDcdnWafPolicyRequest) SetPolicyType(v string) *CreateDcdnWafPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateDcdnWafPolicyRequest) Validate() error {
	return dara.Validate(s)
}
