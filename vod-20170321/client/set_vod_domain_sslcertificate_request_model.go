// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainSSLCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v int64) *SetVodDomainSSLCertificateRequest
	GetCertId() *int64
	SetCertName(v string) *SetVodDomainSSLCertificateRequest
	GetCertName() *string
	SetCertRegion(v string) *SetVodDomainSSLCertificateRequest
	GetCertRegion() *string
	SetCertType(v string) *SetVodDomainSSLCertificateRequest
	GetCertType() *string
	SetDomainName(v string) *SetVodDomainSSLCertificateRequest
	GetDomainName() *string
	SetEnv(v string) *SetVodDomainSSLCertificateRequest
	GetEnv() *string
	SetOwnerId(v int64) *SetVodDomainSSLCertificateRequest
	GetOwnerId() *int64
	SetSSLPri(v string) *SetVodDomainSSLCertificateRequest
	GetSSLPri() *string
	SetSSLProtocol(v string) *SetVodDomainSSLCertificateRequest
	GetSSLProtocol() *string
	SetSSLPub(v string) *SetVodDomainSSLCertificateRequest
	GetSSLPub() *string
	SetSecurityToken(v string) *SetVodDomainSSLCertificateRequest
	GetSecurityToken() *string
}

type SetVodDomainSSLCertificateRequest struct {
	// The certificate ID.
	//
	// example:
	//
	// 12342707
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// cert_name
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The certificate region. Valid values:
	//
	// - **ap-southeast-1*	- (Singapore)
	//
	// - **cn-hangzhou*	- (Hangzhou)
	//
	// Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// The certificate type. Valid values:
	//
	// - **upload**: an uploaded certificate.
	//
	// - **cas**: a certificate from SSL Certificates Service.
	//
	// example:
	//
	// cas
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The accelerated domain name for ApsaraVideo VOD.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to distribute the certificate in a canary release environment. Valid values:
	//
	// - **staging**: distributes the certificate in a canary release environment.
	//
	// If this parameter is not specified or set to any other value, the certificate is formally distributed.
	//
	// example:
	//
	// staging
	Env     *string `json:"Env,omitempty" xml:"Env,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The content of the private key. If you do not enable the certificate, you do not need to specify this parameter. If you configure a certificate, enter the private key content.
	//
	// example:
	//
	// ****
	SSLPri *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	// Specifies whether to enable the HTTPS certificate. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The content of the security certificate. If you do not enable the certificate, you do not need to specify this parameter. If you configure a certificate, enter the certificate content.
	//
	// example:
	//
	// ****
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetVodDomainSSLCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainSSLCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetVodDomainSSLCertificateRequest) GetCertId() *int64 {
	return s.CertId
}

func (s *SetVodDomainSSLCertificateRequest) GetCertName() *string {
	return s.CertName
}

func (s *SetVodDomainSSLCertificateRequest) GetCertRegion() *string {
	return s.CertRegion
}

func (s *SetVodDomainSSLCertificateRequest) GetCertType() *string {
	return s.CertType
}

func (s *SetVodDomainSSLCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetVodDomainSSLCertificateRequest) GetEnv() *string {
	return s.Env
}

func (s *SetVodDomainSSLCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetVodDomainSSLCertificateRequest) GetSSLPri() *string {
	return s.SSLPri
}

func (s *SetVodDomainSSLCertificateRequest) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *SetVodDomainSSLCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *SetVodDomainSSLCertificateRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetVodDomainSSLCertificateRequest) SetCertId(v int64) *SetVodDomainSSLCertificateRequest {
	s.CertId = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetCertName(v string) *SetVodDomainSSLCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetCertRegion(v string) *SetVodDomainSSLCertificateRequest {
	s.CertRegion = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetCertType(v string) *SetVodDomainSSLCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetDomainName(v string) *SetVodDomainSSLCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetEnv(v string) *SetVodDomainSSLCertificateRequest {
	s.Env = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetOwnerId(v int64) *SetVodDomainSSLCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetSSLPri(v string) *SetVodDomainSSLCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetSSLProtocol(v string) *SetVodDomainSSLCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetSSLPub(v string) *SetVodDomainSSLCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) SetSecurityToken(v string) *SetVodDomainSSLCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetVodDomainSSLCertificateRequest) Validate() error {
	return dara.Validate(s)
}
