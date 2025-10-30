// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomPrivacyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomPrivacyPolicyResponseBody
	GetRequestId() *string
}

type DeleteCustomPrivacyPolicyResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomPrivacyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomPrivacyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomPrivacyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomPrivacyPolicyResponseBody) SetRequestId(v string) *DeleteCustomPrivacyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomPrivacyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
