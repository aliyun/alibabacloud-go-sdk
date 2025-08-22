// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainCSRCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetDcdnDomainCSRCertificateRequest
	GetDomainName() *string
	SetServerCertificate(v string) *SetDcdnDomainCSRCertificateRequest
	GetServerCertificate() *string
}

type SetDcdnDomainCSRCertificateRequest struct {
	// The domain name that is secured by the certificate. The domain name uses HTTPS acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The content of the certificate. The certificate must match the certificate signing request (CSR) created by calling the [CreateDcdnCertificateSigningRequest](https://help.aliyun.com/document_detail/144478.html) operation. Make sure that the certificate is in PEM format and its content is Base64-encoded and then encoded by encodeURIComponent.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
}

func (s SetDcdnDomainCSRCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainCSRCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainCSRCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetDcdnDomainCSRCertificateRequest) GetServerCertificate() *string {
	return s.ServerCertificate
}

func (s *SetDcdnDomainCSRCertificateRequest) SetDomainName(v string) *SetDcdnDomainCSRCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetDcdnDomainCSRCertificateRequest) SetServerCertificate(v string) *SetDcdnDomainCSRCertificateRequest {
	s.ServerCertificate = &v
	return s
}

func (s *SetDcdnDomainCSRCertificateRequest) Validate() error {
	return dara.Validate(s)
}
