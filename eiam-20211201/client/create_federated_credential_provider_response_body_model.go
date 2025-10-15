// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFederatedCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProviderId(v string) *CreateFederatedCredentialProviderResponseBody
	GetFederatedCredentialProviderId() *string
	SetRequestId(v string) *CreateFederatedCredentialProviderResponseBody
	GetRequestId() *string
}

type CreateFederatedCredentialProviderResponseBody struct {
	// example:
	//
	// fcp_sada123XXXX
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFederatedCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderResponseBody) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *CreateFederatedCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFederatedCredentialProviderResponseBody) SetFederatedCredentialProviderId(v string) *CreateFederatedCredentialProviderResponseBody {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *CreateFederatedCredentialProviderResponseBody) SetRequestId(v string) *CreateFederatedCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFederatedCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
