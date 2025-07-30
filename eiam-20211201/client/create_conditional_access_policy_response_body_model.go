// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConditionalAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicyId(v string) *CreateConditionalAccessPolicyResponseBody
	GetConditionalAccessPolicyId() *string
	SetRequestId(v string) *CreateConditionalAccessPolicyResponseBody
	GetRequestId() *string
}

type CreateConditionalAccessPolicyResponseBody struct {
	// Conditional Access Policy ID
	//
	// example:
	//
	// cp_xxxxx
	ConditionalAccessPolicyId *string `json:"ConditionalAccessPolicyId,omitempty" xml:"ConditionalAccessPolicyId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConditionalAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConditionalAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConditionalAccessPolicyResponseBody) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *CreateConditionalAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConditionalAccessPolicyResponseBody) SetConditionalAccessPolicyId(v string) *CreateConditionalAccessPolicyResponseBody {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *CreateConditionalAccessPolicyResponseBody) SetRequestId(v string) *CreateConditionalAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConditionalAccessPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
