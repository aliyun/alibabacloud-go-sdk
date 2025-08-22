// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainSMCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *SetDcdnDomainSMCertificateRequest
	GetCertIdentifier() *string
	SetDomainName(v string) *SetDcdnDomainSMCertificateRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetDcdnDomainSMCertificateRequest
	GetOwnerId() *int64
	SetSSLProtocol(v string) *SetDcdnDomainSMCertificateRequest
	GetSSLProtocol() *string
	SetSecurityToken(v string) *SetDcdnDomainSMCertificateRequest
	GetSecurityToken() *string
}

type SetDcdnDomainSMCertificateRequest struct {
	// The identifier of the certificate. The value is Certificate ID-cn-hangzhou. If the ID of the certificate is 123, CertIdentifier is set to 123-cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The accelerated domain name for which the SM certificate is configured.
	//
	// > The domain name must have HTTPS secure acceleration enabled.
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

func (s SetDcdnDomainSMCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainSMCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSMCertificateRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *SetDcdnDomainSMCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetDcdnDomainSMCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDcdnDomainSMCertificateRequest) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *SetDcdnDomainSMCertificateRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetDcdnDomainSMCertificateRequest) SetCertIdentifier(v string) *SetDcdnDomainSMCertificateRequest {
	s.CertIdentifier = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) SetDomainName(v string) *SetDcdnDomainSMCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) SetOwnerId(v int64) *SetDcdnDomainSMCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) SetSSLProtocol(v string) *SetDcdnDomainSMCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) SetSecurityToken(v string) *SetDcdnDomainSMCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetDcdnDomainSMCertificateRequest) Validate() error {
	return dara.Validate(s)
}
