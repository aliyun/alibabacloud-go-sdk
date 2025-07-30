// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConditionalAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableConditionalAccessPolicyResponseBody
	GetRequestId() *string
}

type DisableConditionalAccessPolicyResponseBody struct {
	// 请求ID。
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableConditionalAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableConditionalAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableConditionalAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableConditionalAccessPolicyResponseBody) SetRequestId(v string) *DisableConditionalAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableConditionalAccessPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
