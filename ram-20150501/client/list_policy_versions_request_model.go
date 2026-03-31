// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *ListPolicyVersionsRequest
	GetPolicyName() *string
	SetPolicyType(v string) *ListPolicyVersionsRequest
	GetPolicyType() *string
}

type ListPolicyVersionsRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPolicyVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyVersionsRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPolicyVersionsRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPolicyVersionsRequest) SetPolicyName(v string) *ListPolicyVersionsRequest {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyVersionsRequest) SetPolicyType(v string) *ListPolicyVersionsRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyVersionsRequest) Validate() error {
	return dara.Validate(s)
}
