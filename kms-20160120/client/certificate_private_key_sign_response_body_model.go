// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePrivateKeySignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *CertificatePrivateKeySignResponseBody
	GetCertificateId() *string
	SetRequestId(v string) *CertificatePrivateKeySignResponseBody
	GetRequestId() *string
	SetSignatureValue(v string) *CertificatePrivateKeySignResponseBody
	GetSignatureValue() *string
}

type CertificatePrivateKeySignResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature value.
	//
	// The value is encoded in Base64.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
}

func (s CertificatePrivateKeySignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CertificatePrivateKeySignResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeySignResponseBody) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CertificatePrivateKeySignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CertificatePrivateKeySignResponseBody) GetSignatureValue() *string {
	return s.SignatureValue
}

func (s *CertificatePrivateKeySignResponseBody) SetCertificateId(v string) *CertificatePrivateKeySignResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeySignResponseBody) SetRequestId(v string) *CertificatePrivateKeySignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePrivateKeySignResponseBody) SetSignatureValue(v string) *CertificatePrivateKeySignResponseBody {
	s.SignatureValue = &v
	return s
}

func (s *CertificatePrivateKeySignResponseBody) Validate() error {
	return dara.Validate(s)
}
