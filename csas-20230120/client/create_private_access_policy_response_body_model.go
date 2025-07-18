// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *CreatePrivateAccessPolicyResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreatePrivateAccessPolicyResponseBody
	GetRequestId() *string
}

type CreatePrivateAccessPolicyResponseBody struct {
	// The ID of the private access policy.
	//
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the current request.
	//
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrivateAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreatePrivateAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrivateAccessPolicyResponseBody) SetPolicyId(v string) *CreatePrivateAccessPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreatePrivateAccessPolicyResponseBody) SetRequestId(v string) *CreatePrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrivateAccessPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
