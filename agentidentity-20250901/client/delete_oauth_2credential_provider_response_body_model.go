// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOAuth2CredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOAuth2CredentialProviderResponseBody
	GetRequestId() *string
}

type DeleteOAuth2CredentialProviderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOAuth2CredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOAuth2CredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOAuth2CredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOAuth2CredentialProviderResponseBody) SetRequestId(v string) *DeleteOAuth2CredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOAuth2CredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
