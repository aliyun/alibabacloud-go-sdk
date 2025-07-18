// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePrivateAccessPolicyResponseBody
	GetRequestId() *string
}

type UpdatePrivateAccessPolicyResponseBody struct {
	// The ID of this request.
	//
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrivateAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrivateAccessPolicyResponseBody) SetRequestId(v string) *UpdatePrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
