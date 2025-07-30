// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConditionalAccessPolicyDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateConditionalAccessPolicyDescriptionResponseBody
	GetRequestId() *string
}

type UpdateConditionalAccessPolicyDescriptionResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConditionalAccessPolicyDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConditionalAccessPolicyDescriptionResponseBody) SetRequestId(v string) *UpdateConditionalAccessPolicyDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConditionalAccessPolicyDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
