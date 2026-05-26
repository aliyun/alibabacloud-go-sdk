// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *DeletePolicyRequest
	GetPolicyName() *string
	SetPolicySetName(v string) *DeletePolicyRequest
	GetPolicySetName() *string
}

type DeletePolicyRequest struct {
	// example:
	//
	// rate-limit-policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DeletePolicyRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *DeletePolicyRequest) SetPolicyName(v string) *DeletePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *DeletePolicyRequest) SetPolicySetName(v string) *DeletePolicyRequest {
	s.PolicySetName = &v
	return s
}

func (s *DeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
