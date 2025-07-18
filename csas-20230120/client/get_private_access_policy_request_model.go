// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrivateAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *GetPrivateAccessPolicyRequest
	GetPolicyId() *string
}

type GetPrivateAccessPolicyRequest struct {
	// Intranet access policy ID. The value can be obtained from:
	//
	// - [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): Query multiple intranet access policies in bulk.
	//
	// - [CreatePrivateAccessPolicy](~~CreatePrivateAccessPolicy~~): Create an intranet access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetPrivateAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPrivateAccessPolicyRequest) SetPolicyId(v string) *GetPrivateAccessPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetPrivateAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
