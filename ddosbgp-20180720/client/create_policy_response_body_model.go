// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreatePolicyResponseBody
	GetId() *string
	SetRequestId(v string) *CreatePolicyResponseBody
	GetRequestId() *string
}

type CreatePolicyResponseBody struct {
	// The ID of the policy.
	//
	// example:
	//
	// 83967609-7ea5-4f6d-a6ea-380b09e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 864FE2F4-CB2E-4024-B9EF-D59FD08A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) GetId() *string {
	return s.Id
}

func (s *CreatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyResponseBody) SetId(v string) *CreatePolicyResponseBody {
	s.Id = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
