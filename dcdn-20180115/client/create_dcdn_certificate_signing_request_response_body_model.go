// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnCertificateSigningRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonName(v string) *CreateDcdnCertificateSigningRequestResponseBody
	GetCommonName() *string
	SetCsr(v string) *CreateDcdnCertificateSigningRequestResponseBody
	GetCsr() *string
	SetPubMd5(v string) *CreateDcdnCertificateSigningRequestResponseBody
	GetPubMd5() *string
	SetRequestId(v string) *CreateDcdnCertificateSigningRequestResponseBody
	GetRequestId() *string
}

type CreateDcdnCertificateSigningRequestResponseBody struct {
	// The Common Name of the certificate.
	//
	// example:
	//
	// example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The content of the CSR file.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----\\nMIIC/zCCAecCAQAwZTELMAkGA1UEBhMCQ04xCzAJBgNVBAgTAlpKMQswCQYDVQQH
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The MD5 hash value of the certificate public key.
	//
	// example:
	//
	// 629bf4fd8104eda171135bcb0f77a10b
	PubMd5 *string `json:"PubMd5,omitempty" xml:"PubMd5,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDcdnCertificateSigningRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnCertificateSigningRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) GetCsr() *string {
	return s.Csr
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) GetPubMd5() *string {
	return s.PubMd5
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) SetCommonName(v string) *CreateDcdnCertificateSigningRequestResponseBody {
	s.CommonName = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) SetCsr(v string) *CreateDcdnCertificateSigningRequestResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) SetPubMd5(v string) *CreateDcdnCertificateSigningRequestResponseBody {
	s.PubMd5 = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) SetRequestId(v string) *CreateDcdnCertificateSigningRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
