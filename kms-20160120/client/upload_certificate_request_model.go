// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *UploadCertificateRequest
	GetCertificate() *string
	SetCertificateChain(v string) *UploadCertificateRequest
	GetCertificateChain() *string
	SetCertificateId(v string) *UploadCertificateRequest
	GetCertificateId() *string
}

type UploadCertificateRequest struct {
	// The certificate issued by the CA, which is in the Privacy Enhanced Mail (PEM) format.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  (X.509 Certificate PEM Content)  -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain issued by the CA, which is in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  (Sub CA Certificate PEM Content)  -----END CERTIFICATE-----  -----BEGIN CERTIFICATE-----  (Sub CA Certificate PEM Content)  -----END CERTIFICATE-----  -----BEGIN CERTIFICATE-----  (Root CA Certificate PEM Content)  -----END CERTIFICATE-----
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s UploadCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadCertificateRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *UploadCertificateRequest) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *UploadCertificateRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *UploadCertificateRequest) SetCertificate(v string) *UploadCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *UploadCertificateRequest) SetCertificateChain(v string) *UploadCertificateRequest {
	s.CertificateChain = &v
	return s
}

func (s *UploadCertificateRequest) SetCertificateId(v string) *UploadCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *UploadCertificateRequest) Validate() error {
	return dara.Validate(s)
}
