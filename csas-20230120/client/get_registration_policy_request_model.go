// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *GetRegistrationPolicyRequest
	GetPolicyId() *string
}

type GetRegistrationPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetRegistrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetRegistrationPolicyRequest) SetPolicyId(v string) *GetRegistrationPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetRegistrationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
