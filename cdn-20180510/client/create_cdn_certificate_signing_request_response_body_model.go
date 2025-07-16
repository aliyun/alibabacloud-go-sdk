// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdnCertificateSigningRequestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonName(v string) *CreateCdnCertificateSigningRequestResponseBody
	GetCommonName() *string
	SetCsr(v string) *CreateCdnCertificateSigningRequestResponseBody
	GetCsr() *string
	SetPubMd5(v string) *CreateCdnCertificateSigningRequestResponseBody
	GetPubMd5() *string
	SetRequestId(v string) *CreateCdnCertificateSigningRequestResponseBody
	GetRequestId() *string
}

type CreateCdnCertificateSigningRequestResponseBody struct {
	// The Common Name of the certificate.
	//
	// example:
	//
	// CommonName
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The content of the CSR file.
	//
	// example:
	//
	// CSRName
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The MD5 hash value of the certificate public key.
	//
	// example:
	//
	// 629bf4fd8104eda171135bcb0f77****
	PubMd5 *string `json:"PubMd5,omitempty" xml:"PubMd5,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCdnCertificateSigningRequestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCdnCertificateSigningRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdnCertificateSigningRequestResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateCdnCertificateSigningRequestResponseBody) GetCsr() *string {
	return s.Csr
}

func (s *CreateCdnCertificateSigningRequestResponseBody) GetPubMd5() *string {
	return s.PubMd5
}

func (s *CreateCdnCertificateSigningRequestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCdnCertificateSigningRequestResponseBody) SetCommonName(v string) *CreateCdnCertificateSigningRequestResponseBody {
	s.CommonName = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponseBody) SetCsr(v string) *CreateCdnCertificateSigningRequestResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponseBody) SetPubMd5(v string) *CreateCdnCertificateSigningRequestResponseBody {
	s.PubMd5 = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponseBody) SetRequestId(v string) *CreateCdnCertificateSigningRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponseBody) Validate() error {
	return dara.Validate(s)
}
