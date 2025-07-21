// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePublicKeyEncryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *CertificatePublicKeyEncryptResponseBody
	GetCertificateId() *string
	SetCiphertextBlob(v string) *CertificatePublicKeyEncryptResponseBody
	GetCiphertextBlob() *string
	SetRequestId(v string) *CertificatePublicKeyEncryptResponseBody
	GetRequestId() *string
}

type CertificatePublicKeyEncryptResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ciphertext.
	//
	// The value is encoded in Base64.
	//
	// example:
	//
	// ZOyIygCyaOW6Gj****MlNKiuyjfzw=
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CertificatePublicKeyEncryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CertificatePublicKeyEncryptResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePublicKeyEncryptResponseBody) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CertificatePublicKeyEncryptResponseBody) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *CertificatePublicKeyEncryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CertificatePublicKeyEncryptResponseBody) SetCertificateId(v string) *CertificatePublicKeyEncryptResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponseBody) SetCiphertextBlob(v string) *CertificatePublicKeyEncryptResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponseBody) SetRequestId(v string) *CertificatePublicKeyEncryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePublicKeyEncryptResponseBody) Validate() error {
	return dara.Validate(s)
}
