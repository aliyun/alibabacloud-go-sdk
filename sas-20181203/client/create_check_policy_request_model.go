// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDependentPolicyId(v int64) *CreateCheckPolicyRequest
	GetDependentPolicyId() *int64
	SetPolicyShowName(v string) *CreateCheckPolicyRequest
	GetPolicyShowName() *string
	SetPolicyType(v string) *CreateCheckPolicyRequest
	GetPolicyType() *string
	SetType(v string) *CreateCheckPolicyRequest
	GetType() *string
}

type CreateCheckPolicyRequest struct {
	// The ID of the parent policy.
	//
	//
	//
	// (The specific dependency order from low to high is Section -> Requirement -> Standard)
	//
	// example:
	//
	// 123
	DependentPolicyId *int64 `json:"DependentPolicyId,omitempty" xml:"DependentPolicyId,omitempty"`
	// The name of the custom policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestPolicyName
	PolicyShowName *string `json:"PolicyShowName,omitempty" xml:"PolicyShowName,omitempty"`
	// The policy category type for custom check rules:
	//
	// - **STANDARD**: Add to a standard.
	//
	// - **REQUIREMENT**: Add to a requirement.
	//
	// - **SECTION**: Add to a section.
	//
	// This parameter is required.
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the major policy category (required when PolicyType is STANDARD):
	//
	// - **AISPM**: AI Configuration Management (AISPM).
	//
	// - **IDENTITY_PERMISSION**: Identity and Permission Management (CIEM).
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

func (s CreateCheckPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateCheckPolicyRequest) GetDependentPolicyId() *int64 {
	return s.DependentPolicyId
}

func (s *CreateCheckPolicyRequest) GetPolicyShowName() *string {
	return s.PolicyShowName
}

func (s *CreateCheckPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateCheckPolicyRequest) GetType() *string {
	return s.Type
}

func (s *CreateCheckPolicyRequest) SetDependentPolicyId(v int64) *CreateCheckPolicyRequest {
	s.DependentPolicyId = &v
	return s
}

func (s *CreateCheckPolicyRequest) SetPolicyShowName(v string) *CreateCheckPolicyRequest {
	s.PolicyShowName = &v
	return s
}

func (s *CreateCheckPolicyRequest) SetPolicyType(v string) *CreateCheckPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateCheckPolicyRequest) SetType(v string) *CreateCheckPolicyRequest {
	s.Type = &v
	return s
}

func (s *CreateCheckPolicyRequest) Validate() error {
	return dara.Validate(s)
}
