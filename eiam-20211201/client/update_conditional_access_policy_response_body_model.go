// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConditionalAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateConditionalAccessPolicyResponseBody
	GetRequestId() *string
}

type UpdateConditionalAccessPolicyResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConditionalAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConditionalAccessPolicyResponseBody) SetRequestId(v string) *UpdateConditionalAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConditionalAccessPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
