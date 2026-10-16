// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *GetPolicyRequest
	GetPolicyName() *string
	SetPolicySetName(v string) *GetPolicyRequest
	GetPolicySetName() *string
}

type GetPolicyRequest struct {
	// The name of the policy. The name must be unique within the associated policy set.
	//
	// example:
	//
	// rate-limit-policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The name of the policy set. The name must be unique within the current Alibaba Cloud account.
	//
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s GetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *GetPolicyRequest) SetPolicyName(v string) *GetPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyRequest) SetPolicySetName(v string) *GetPolicyRequest {
	s.PolicySetName = &v
	return s
}

func (s *GetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
