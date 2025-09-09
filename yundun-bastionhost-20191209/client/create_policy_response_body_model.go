// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *CreatePolicyResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreatePolicyResponseBody
	GetRequestId() *string
}

type CreatePolicyResponseBody struct {
	// The control policy ID.
	//
	// example:
	//
	// 1
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7E68165E-1191-5CC2-B54B-5EF7390A5400
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyResponseBody) SetPolicyId(v string) *CreatePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
