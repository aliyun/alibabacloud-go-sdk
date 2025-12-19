// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOAuth2CredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOAuth2CredentialProviderResponseBody
	GetRequestId() *string
}

type UpdateOAuth2CredentialProviderResponseBody struct {
	// example:
	//
	// D9A9DC39-61BB-53FD-9ADC-B14884F21038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOAuth2CredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOAuth2CredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOAuth2CredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOAuth2CredentialProviderResponseBody) SetRequestId(v string) *UpdateOAuth2CredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
