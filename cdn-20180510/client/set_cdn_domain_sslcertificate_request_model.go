// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainSSLCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v int64) *SetCdnDomainSSLCertificateRequest
	GetCertId() *int64
	SetCertName(v string) *SetCdnDomainSSLCertificateRequest
	GetCertName() *string
	SetCertRegion(v string) *SetCdnDomainSSLCertificateRequest
	GetCertRegion() *string
	SetCertType(v string) *SetCdnDomainSSLCertificateRequest
	GetCertType() *string
	SetDomainName(v string) *SetCdnDomainSSLCertificateRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetCdnDomainSSLCertificateRequest
	GetOwnerId() *int64
	SetSSLPri(v string) *SetCdnDomainSSLCertificateRequest
	GetSSLPri() *string
	SetSSLProtocol(v string) *SetCdnDomainSSLCertificateRequest
	GetSSLProtocol() *string
	SetSSLPub(v string) *SetCdnDomainSSLCertificateRequest
	GetSSLPub() *string
	SetSecurityToken(v string) *SetCdnDomainSSLCertificateRequest
	GetSecurityToken() *string
}

type SetCdnDomainSSLCertificateRequest struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 8089870
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the SSL certificate. You can specify only one certificate name.
	//
	// example:
	//
	// yourCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The region ID of the certificate. Valid values: **cn-hangzhou*	- and **ap-southeast-1**. Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// The type of the certificate.
	//
	// 	- **upload**: a user-uploaded SSL certificate.
	//
	// 	- **cas**: a certificate that is acquired through Certificate Management Service.
	//
	// example:
	//
	// upload
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The accelerated domain name for which you want to configure the SSL certificate. The type of request supported by the domain name must be HTTPS. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private key. Specify the private key only if you want to enable the SSL certificate.
	//
	// example:
	//
	// y****
	SSLPri *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	// Specifies whether to enable the SSL certificate.
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
	// The content of the SSL certificate. Specify the content of the SSL certificate only if you want to enable the SSL certificate.
	//
	// example:
	//
	// xxx
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetCdnDomainSSLCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainSSLCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetCdnDomainSSLCertificateRequest) GetCertId() *int64 {
	return s.CertId
}

func (s *SetCdnDomainSSLCertificateRequest) GetCertName() *string {
	return s.CertName
}

func (s *SetCdnDomainSSLCertificateRequest) GetCertRegion() *string {
	return s.CertRegion
}

func (s *SetCdnDomainSSLCertificateRequest) GetCertType() *string {
	return s.CertType
}

func (s *SetCdnDomainSSLCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetCdnDomainSSLCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCdnDomainSSLCertificateRequest) GetSSLPri() *string {
	return s.SSLPri
}

func (s *SetCdnDomainSSLCertificateRequest) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *SetCdnDomainSSLCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *SetCdnDomainSSLCertificateRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetCdnDomainSSLCertificateRequest) SetCertId(v int64) *SetCdnDomainSSLCertificateRequest {
	s.CertId = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetCertName(v string) *SetCdnDomainSSLCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetCertRegion(v string) *SetCdnDomainSSLCertificateRequest {
	s.CertRegion = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetCertType(v string) *SetCdnDomainSSLCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetDomainName(v string) *SetCdnDomainSSLCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetOwnerId(v int64) *SetCdnDomainSSLCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetSSLPri(v string) *SetCdnDomainSSLCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetSSLProtocol(v string) *SetCdnDomainSSLCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetSSLPub(v string) *SetCdnDomainSSLCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) SetSecurityToken(v string) *SetCdnDomainSSLCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetCdnDomainSSLCertificateRequest) Validate() error {
	return dara.Validate(s)
}
