// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSAMLIdentityProviderCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddSAMLIdentityProviderCertificateResponseBody
	GetRequestId() *string
}

type AddSAMLIdentityProviderCertificateResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSAMLIdentityProviderCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSAMLIdentityProviderCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *AddSAMLIdentityProviderCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSAMLIdentityProviderCertificateResponseBody) SetRequestId(v string) *AddSAMLIdentityProviderCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSAMLIdentityProviderCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
