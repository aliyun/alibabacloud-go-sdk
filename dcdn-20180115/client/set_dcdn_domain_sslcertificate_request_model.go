// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainSSLCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v int64) *SetDcdnDomainSSLCertificateRequest
	GetCertId() *int64
	SetCertName(v string) *SetDcdnDomainSSLCertificateRequest
	GetCertName() *string
	SetCertRegion(v string) *SetDcdnDomainSSLCertificateRequest
	GetCertRegion() *string
	SetCertType(v string) *SetDcdnDomainSSLCertificateRequest
	GetCertType() *string
	SetDomainName(v string) *SetDcdnDomainSSLCertificateRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetDcdnDomainSSLCertificateRequest
	GetOwnerId() *int64
	SetSSLPri(v string) *SetDcdnDomainSSLCertificateRequest
	GetSSLPri() *string
	SetSSLProtocol(v string) *SetDcdnDomainSSLCertificateRequest
	GetSSLProtocol() *string
	SetSSLPub(v string) *SetDcdnDomainSSLCertificateRequest
	GetSSLPub() *string
	SetSecurityToken(v string) *SetDcdnDomainSSLCertificateRequest
	GetSecurityToken() *string
}

type SetDcdnDomainSSLCertificateRequest struct {
	// The certificate ID. This parameter is required and valid only when **CertType*	- is set to **cas**. If you specify this parameter, an existing certificate is used.
	//
	// example:
	//
	// 8089870
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the new certificate. You can specify only one certificate name. This parameter is optional and valid only when **CertType*	- is set to **upload**.
	//
	// example:
	//
	// yourCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The region of the SSL certificate. This parameter takes effect only when **CertType*	- is set to **cas**. Default value: **cn-hangzhou**. Valid values: **cn-hangzhou*	- and **ap-southeast-1**.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **upload**: a user-uploaded SSL certificate.
	//
	// 	- **cas**: a certificate that is acquired through Certificate Management Service.
	//
	// example:
	//
	// upload
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The domain name that is secured by the SSL certificate.
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

func (s SetDcdnDomainSSLCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainSSLCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainSSLCertificateRequest) GetCertId() *int64 {
	return s.CertId
}

func (s *SetDcdnDomainSSLCertificateRequest) GetCertName() *string {
	return s.CertName
}

func (s *SetDcdnDomainSSLCertificateRequest) GetCertRegion() *string {
	return s.CertRegion
}

func (s *SetDcdnDomainSSLCertificateRequest) GetCertType() *string {
	return s.CertType
}

func (s *SetDcdnDomainSSLCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetDcdnDomainSSLCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDcdnDomainSSLCertificateRequest) GetSSLPri() *string {
	return s.SSLPri
}

func (s *SetDcdnDomainSSLCertificateRequest) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *SetDcdnDomainSSLCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *SetDcdnDomainSSLCertificateRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetDcdnDomainSSLCertificateRequest) SetCertId(v int64) *SetDcdnDomainSSLCertificateRequest {
	s.CertId = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetCertName(v string) *SetDcdnDomainSSLCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetCertRegion(v string) *SetDcdnDomainSSLCertificateRequest {
	s.CertRegion = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetCertType(v string) *SetDcdnDomainSSLCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetDomainName(v string) *SetDcdnDomainSSLCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetOwnerId(v int64) *SetDcdnDomainSSLCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetSSLPri(v string) *SetDcdnDomainSSLCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetSSLProtocol(v string) *SetDcdnDomainSSLCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetSSLPub(v string) *SetDcdnDomainSSLCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) SetSecurityToken(v string) *SetDcdnDomainSSLCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetDcdnDomainSSLCertificateRequest) Validate() error {
	return dara.Validate(s)
}
