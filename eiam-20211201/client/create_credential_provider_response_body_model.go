// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialProviderId(v string) *CreateCredentialProviderResponseBody
	GetCredentialProviderId() *string
	SetRequestId(v string) *CreateCredentialProviderResponseBody
	GetRequestId() *string
}

type CreateCredentialProviderResponseBody struct {
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"CredentialProviderId,omitempty" xml:"CredentialProviderId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialProviderResponseBody) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *CreateCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCredentialProviderResponseBody) SetCredentialProviderId(v string) *CreateCredentialProviderResponseBody {
	s.CredentialProviderId = &v
	return s
}

func (s *CreateCredentialProviderResponseBody) SetRequestId(v string) *CreateCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
