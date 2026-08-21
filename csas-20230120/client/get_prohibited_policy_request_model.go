// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProhibitedPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *GetProhibitedPolicyRequest
	GetPolicyId() *string
}

type GetProhibitedPolicyRequest struct {
	// The software prohibition policy ID. You can obtain this value from the following operations:
	//
	// - [ListProhibitedPolicies](~~ListProhibitedPolicies~~): Lists software prohibition policies.
	//
	// - [CreateProhibitedPolicy](~~CreateProhibitedPolicy~~): Creates a software prohibition policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// pid-7da5ea4192c1****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetProhibitedPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetProhibitedPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetProhibitedPolicyRequest) SetPolicyId(v string) *GetProhibitedPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetProhibitedPolicyRequest) Validate() error {
	return dara.Validate(s)
}
