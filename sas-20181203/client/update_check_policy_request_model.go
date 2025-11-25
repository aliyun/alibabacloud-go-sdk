// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDependentPolicyId(v int64) *UpdateCheckPolicyRequest
	GetDependentPolicyId() *int64
	SetPolicyId(v int64) *UpdateCheckPolicyRequest
	GetPolicyId() *int64
	SetPolicyShowName(v string) *UpdateCheckPolicyRequest
	GetPolicyShowName() *string
	SetPolicyType(v string) *UpdateCheckPolicyRequest
	GetPolicyType() *string
	SetType(v string) *UpdateCheckPolicyRequest
	GetType() *string
}

type UpdateCheckPolicyRequest struct {
	// ID of the associated parent policy.
	//
	// (The specific dependency hierarchy, from low to high, is Section -> Requirement -> Standard).
	//
	// example:
	//
	// 1000000000002
	DependentPolicyId *int64 `json:"DependentPolicyId,omitempty" xml:"DependentPolicyId,omitempty"`
	// ID of the custom policy.
	//
	// > You can obtain this parameter by calling the [ListCheckPolicies](~~ListCheckPolicies~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Name of the custom classification.
	//
	// example:
	//
	// testPolicyName
	PolicyShowName *string `json:"PolicyShowName,omitempty" xml:"PolicyShowName,omitempty"`
	// Classification type of the custom check item rule:
	//
	// - **STANDARD**: Add to standard.
	//
	// - **REQUIREMENT**: Add to requirement.
	//
	// - **SECTION**: Add to section.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// Name of the associated major policy category (required when PolicyType is STANDARD):
	//
	// - **AISPM**: AI Configuration Management (AISPM).
	//
	// - **RISK**: Security Risk.
	//
	// - **COMPLIANCE**: Compliance Risk.
	//
	// example:
	//
	// AISPM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateCheckPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckPolicyRequest) GetDependentPolicyId() *int64 {
	return s.DependentPolicyId
}

func (s *UpdateCheckPolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *UpdateCheckPolicyRequest) GetPolicyShowName() *string {
	return s.PolicyShowName
}

func (s *UpdateCheckPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateCheckPolicyRequest) GetType() *string {
	return s.Type
}

func (s *UpdateCheckPolicyRequest) SetDependentPolicyId(v int64) *UpdateCheckPolicyRequest {
	s.DependentPolicyId = &v
	return s
}

func (s *UpdateCheckPolicyRequest) SetPolicyId(v int64) *UpdateCheckPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateCheckPolicyRequest) SetPolicyShowName(v string) *UpdateCheckPolicyRequest {
	s.PolicyShowName = &v
	return s
}

func (s *UpdateCheckPolicyRequest) SetPolicyType(v string) *UpdateCheckPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *UpdateCheckPolicyRequest) SetType(v string) *UpdateCheckPolicyRequest {
	s.Type = &v
	return s
}

func (s *UpdateCheckPolicyRequest) Validate() error {
	return dara.Validate(s)
}
