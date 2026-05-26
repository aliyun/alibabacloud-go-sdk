// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicySetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicySetName(v string) *GetPolicySetRequest
	GetPolicySetName() *string
}

type GetPolicySetRequest struct {
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s GetPolicySetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicySetRequest) GoString() string {
	return s.String()
}

func (s *GetPolicySetRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *GetPolicySetRequest) SetPolicySetName(v string) *GetPolicySetRequest {
	s.PolicySetName = &v
	return s
}

func (s *GetPolicySetRequest) Validate() error {
	return dara.Validate(s)
}
