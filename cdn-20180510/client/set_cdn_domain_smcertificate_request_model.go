// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainSMCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *SetCdnDomainSMCertificateRequest
	GetCertIdentifier() *string
	SetDomainName(v string) *SetCdnDomainSMCertificateRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetCdnDomainSMCertificateRequest
	GetOwnerId() *int64
	SetSSLProtocol(v string) *SetCdnDomainSMCertificateRequest
	GetSSLProtocol() *string
	SetSecurityToken(v string) *SetCdnDomainSMCertificateRequest
	GetSecurityToken() *string
}

type SetCdnDomainSMCertificateRequest struct {
	// The ID of the SM certificate that you want to configure. The identifier of the certificate. The value is Certificate ID-cn-hangzhou. For example, if the certificate ID is 123, set the value of this parameter to 123-cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234****-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The accelerated domain name for which you want to configure the SM certificate.
	//
	// > The domain name must use HTTPS acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable the SSL certificate. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	SSLProtocol   *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetCdnDomainSMCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainSMCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetCdnDomainSMCertificateRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *SetCdnDomainSMCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetCdnDomainSMCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCdnDomainSMCertificateRequest) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *SetCdnDomainSMCertificateRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetCdnDomainSMCertificateRequest) SetCertIdentifier(v string) *SetCdnDomainSMCertificateRequest {
	s.CertIdentifier = &v
	return s
}

func (s *SetCdnDomainSMCertificateRequest) SetDomainName(v string) *SetCdnDomainSMCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetCdnDomainSMCertificateRequest) SetOwnerId(v int64) *SetCdnDomainSMCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCdnDomainSMCertificateRequest) SetSSLProtocol(v string) *SetCdnDomainSMCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetCdnDomainSMCertificateRequest) SetSecurityToken(v string) *SetCdnDomainSMCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetCdnDomainSMCertificateRequest) Validate() error {
	return dara.Validate(s)
}
