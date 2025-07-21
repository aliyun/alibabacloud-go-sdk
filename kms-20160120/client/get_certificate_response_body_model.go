// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetCertificateResponseBody
	GetCertificate() *string
	SetCertificateChain(v string) *GetCertificateResponseBody
	GetCertificateChain() *string
	SetCertificateId(v string) *GetCertificateResponseBody
	GetCertificateId() *string
	SetCsr(v string) *GetCertificateResponseBody
	GetCsr() *string
	SetRequestId(v string) *GetCertificateResponseBody
	GetRequestId() *string
}

type GetCertificateResponseBody struct {
	// The certificate in the Privacy Enhanced Mail (PEM) format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  (X.509 Certificate PEM Content)  -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  (Sub CA Certificate PEM Content)  -----END CERTIFICATE-----  -----BEGIN CERTIFICATE-----  (Sub CA Certificate PEM Content)  -----END CERTIFICATE-----  -----BEGIN CERTIFICATE-----  (Root CA Certificate PEM Content)  -----END CERTIFICATE-----
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The CSR in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----MIICxjCCAa4CAQAwPzELMAkGA1UEBhMCQ04xDzANBgNVBAoTBmFsaXl1bjEMMAoGA1UECxMDa21zMREwDwY****-----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b3e104b4-0319-4a20-ab7f-9fef6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetCertificateResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *GetCertificateResponseBody) GetCertificateId() *string {
	return s.CertificateId
}

func (s *GetCertificateResponseBody) GetCsr() *string {
	return s.Csr
}

func (s *GetCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertificateResponseBody) SetCertificate(v string) *GetCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetCertificateResponseBody) SetCertificateChain(v string) *GetCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *GetCertificateResponseBody) SetCertificateId(v string) *GetCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateResponseBody) SetCsr(v string) *GetCertificateResponseBody {
	s.Csr = &v
	return s
}

func (s *GetCertificateResponseBody) SetRequestId(v string) *GetCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
