// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomPrivacyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrivacyPolicyId(v string) *CreateCustomPrivacyPolicyResponseBody
	GetCustomPrivacyPolicyId() *string
	SetRequestId(v string) *CreateCustomPrivacyPolicyResponseBody
	GetRequestId() *string
}

type CreateCustomPrivacyPolicyResponseBody struct {
	// example:
	//
	// pp_neagxpoznsjdtxxxxx
	CustomPrivacyPolicyId *string `json:"CustomPrivacyPolicyId,omitempty" xml:"CustomPrivacyPolicyId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomPrivacyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomPrivacyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomPrivacyPolicyResponseBody) GetCustomPrivacyPolicyId() *string {
	return s.CustomPrivacyPolicyId
}

func (s *CreateCustomPrivacyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomPrivacyPolicyResponseBody) SetCustomPrivacyPolicyId(v string) *CreateCustomPrivacyPolicyResponseBody {
	s.CustomPrivacyPolicyId = &v
	return s
}

func (s *CreateCustomPrivacyPolicyResponseBody) SetRequestId(v string) *CreateCustomPrivacyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomPrivacyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
