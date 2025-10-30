// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomPrivacyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCustomPrivacyPolicyResponseBody
	GetRequestId() *string
}

type DisableCustomPrivacyPolicyResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCustomPrivacyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomPrivacyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCustomPrivacyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCustomPrivacyPolicyResponseBody) SetRequestId(v string) *DisableCustomPrivacyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCustomPrivacyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
