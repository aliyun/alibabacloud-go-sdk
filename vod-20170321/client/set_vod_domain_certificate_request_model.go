// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *SetVodDomainCertificateRequest
	GetCertName() *string
	SetDomainName(v string) *SetVodDomainCertificateRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetVodDomainCertificateRequest
	GetOwnerId() *int64
	SetSSLPri(v string) *SetVodDomainCertificateRequest
	GetSSLPri() *string
	SetSSLProtocol(v string) *SetVodDomainCertificateRequest
	GetSSLProtocol() *string
	SetSSLPub(v string) *SetVodDomainCertificateRequest
	GetSSLPub() *string
	SetSecurityToken(v string) *SetVodDomainCertificateRequest
	GetSecurityToken() *string
}

type SetVodDomainCertificateRequest struct {
	// The name of the certificate.
	//
	// example:
	//
	// cert_name
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The domain name that is secured by the certificate. The domain name must use HTTPS acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private key. This parameter is required only if you enable the SSL certificate.
	//
	// example:
	//
	// ****
	SSLPri *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	// Specifies whether to enable the SSL certificate. Default value: off. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The content of the certificate. This parameter is required only if you enable the SSL certificate.
	//
	// example:
	//
	// ****
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetVodDomainCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetVodDomainCertificateRequest) GetCertName() *string {
	return s.CertName
}

func (s *SetVodDomainCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetVodDomainCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetVodDomainCertificateRequest) GetSSLPri() *string {
	return s.SSLPri
}

func (s *SetVodDomainCertificateRequest) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *SetVodDomainCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *SetVodDomainCertificateRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetVodDomainCertificateRequest) SetCertName(v string) *SetVodDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetDomainName(v string) *SetVodDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetOwnerId(v int64) *SetVodDomainCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetSSLPri(v string) *SetVodDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetSSLProtocol(v string) *SetVodDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetSSLPub(v string) *SetVodDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetSecurityToken(v string) *SetVodDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetVodDomainCertificateRequest) Validate() error {
	return dara.Validate(s)
}
