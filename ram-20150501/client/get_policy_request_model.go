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
	SetPolicyType(v string) *GetPolicyRequest
	GetPolicyType() *string
}

type GetPolicyRequest struct {
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

func (s GetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetPolicyRequest) SetPolicyName(v string) *GetPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyRequest) SetPolicyType(v string) *GetPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *GetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
