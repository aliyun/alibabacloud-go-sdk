// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSAMLIdentityProviderCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSAMLIdentityProviderCertificateResponseBody
	GetRequestId() *string
}

type DeleteSAMLIdentityProviderCertificateResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSAMLIdentityProviderCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSAMLIdentityProviderCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSAMLIdentityProviderCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSAMLIdentityProviderCertificateResponseBody) SetRequestId(v string) *DeleteSAMLIdentityProviderCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSAMLIdentityProviderCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
