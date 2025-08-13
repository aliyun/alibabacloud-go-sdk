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
	SetPolicyName(v string) *CreatePolicyResponseBody
	GetPolicyName() *string
	SetRequestId(v string) *CreatePolicyResponseBody
	GetRequestId() *string
}

type CreatePolicyResponseBody struct {
	// The ID of the tag policy.
	//
	// example:
	//
	// p-5732750813924f90****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the tag policy.
	//
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FDBE270D-C491-5EEC-A5CD-98245422D3F7
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

func (s *CreatePolicyResponseBody) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyResponseBody) SetPolicyId(v string) *CreatePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyResponseBody) SetPolicyName(v string) *CreatePolicyResponseBody {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
