// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomPrivacyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomPrivacyPolicyResponseBody
	GetRequestId() *string
}

type UpdateCustomPrivacyPolicyResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomPrivacyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomPrivacyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomPrivacyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomPrivacyPolicyResponseBody) SetRequestId(v string) *UpdateCustomPrivacyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
