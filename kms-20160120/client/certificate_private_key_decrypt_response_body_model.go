// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertificatePrivateKeyDecryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *CertificatePrivateKeyDecryptResponseBody
	GetCertificateId() *string
	SetPlaintext(v string) *CertificatePrivateKeyDecryptResponseBody
	GetPlaintext() *string
	SetRequestId(v string) *CertificatePrivateKeyDecryptResponseBody
	GetRequestId() *string
}

type CertificatePrivateKeyDecryptResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The plaintext after data is decrypted.
	//
	// The value is encoded in Base64.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CertificatePrivateKeyDecryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CertificatePrivateKeyDecryptResponseBody) GoString() string {
	return s.String()
}

func (s *CertificatePrivateKeyDecryptResponseBody) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CertificatePrivateKeyDecryptResponseBody) GetPlaintext() *string {
	return s.Plaintext
}

func (s *CertificatePrivateKeyDecryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetCertificateId(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetPlaintext(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponseBody) SetRequestId(v string) *CertificatePrivateKeyDecryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *CertificatePrivateKeyDecryptResponseBody) Validate() error {
	return dara.Validate(s)
}
