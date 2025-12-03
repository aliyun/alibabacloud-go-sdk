// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTLSCipherPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTLSCipherPolicyResponseBody
	GetRequestId() *string
	SetTLSCipherPolicyId(v string) *CreateTLSCipherPolicyResponseBody
	GetTLSCipherPolicyId() *string
}

type CreateTLSCipherPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D7A8875F-373A-5F48-8484-25B07A61F2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// tls-bp14bb1e7dll4f****
	TLSCipherPolicyId *string `json:"TLSCipherPolicyId,omitempty" xml:"TLSCipherPolicyId,omitempty"`
}

func (s CreateTLSCipherPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTLSCipherPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTLSCipherPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTLSCipherPolicyResponseBody) GetTLSCipherPolicyId() *string {
	return s.TLSCipherPolicyId
}

func (s *CreateTLSCipherPolicyResponseBody) SetRequestId(v string) *CreateTLSCipherPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTLSCipherPolicyResponseBody) SetTLSCipherPolicyId(v string) *CreateTLSCipherPolicyResponseBody {
	s.TLSCipherPolicyId = &v
	return s
}

func (s *CreateTLSCipherPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
