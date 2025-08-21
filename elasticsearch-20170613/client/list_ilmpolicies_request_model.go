// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListILMPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *ListILMPoliciesRequest
	GetPolicyName() *string
}

type ListILMPoliciesRequest struct {
	// example:
	//
	// policy-1
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
}

func (s ListILMPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListILMPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListILMPoliciesRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListILMPoliciesRequest) SetPolicyName(v string) *ListILMPoliciesRequest {
	s.PolicyName = &v
	return s
}

func (s *ListILMPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
