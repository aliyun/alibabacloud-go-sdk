// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainCSRCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetCdnDomainCSRCertificateRequest
	GetDomainName() *string
	SetServerCertificate(v string) *SetCdnDomainCSRCertificateRequest
	GetServerCertificate() *string
}

type SetCdnDomainCSRCertificateRequest struct {
	// The accelerated domain name for which you want to configure an SSL certificate. The domain name must have HTTPS secure acceleration enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The content of the certificate. The certificate must match the certificate signing request (CSR) created by calling the [CreateCdnCertificateSigningRequest](https://help.aliyun.com/document_detail/144478.html) operation. Make sure that the content of the certificate is encoded in Base64 and then encoded by encodeURIComponent.
	//
	// This parameter is required.
	//
	// example:
	//
	// ----BEGIN CERTIFICATE----- MIIFz****-----END CERTIFICATE-----
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
}

func (s SetCdnDomainCSRCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainCSRCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetCdnDomainCSRCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetCdnDomainCSRCertificateRequest) GetServerCertificate() *string {
	return s.ServerCertificate
}

func (s *SetCdnDomainCSRCertificateRequest) SetDomainName(v string) *SetCdnDomainCSRCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetCdnDomainCSRCertificateRequest) SetServerCertificate(v string) *SetCdnDomainCSRCertificateRequest {
	s.ServerCertificate = &v
	return s
}

func (s *SetCdnDomainCSRCertificateRequest) Validate() error {
	return dara.Validate(s)
}
