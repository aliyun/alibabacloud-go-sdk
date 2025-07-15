// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *SetLiveDomainCertificateRequest
	GetCertName() *string
	SetCertType(v string) *SetLiveDomainCertificateRequest
	GetCertType() *string
	SetDomainName(v string) *SetLiveDomainCertificateRequest
	GetDomainName() *string
	SetForceSet(v string) *SetLiveDomainCertificateRequest
	GetForceSet() *string
	SetOwnerId(v int64) *SetLiveDomainCertificateRequest
	GetOwnerId() *int64
	SetSSLPri(v string) *SetLiveDomainCertificateRequest
	GetSSLPri() *string
	SetSSLProtocol(v string) *SetLiveDomainCertificateRequest
	GetSSLProtocol() *string
	SetSSLPub(v string) *SetLiveDomainCertificateRequest
	GetSSLPub() *string
	SetSecurityToken(v string) *SetLiveDomainCertificateRequest
	GetSecurityToken() *string
}

type SetLiveDomainCertificateRequest struct {
	// The certificate name.
	//
	// example:
	//
	// Cert-****
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The certificate type. Valid values:
	//
	// 	- **upload**: a custom certificate
	//
	// 	- **cas**: a certificate that is purchased from Certificate Management Service
	//
	// 	- **free**: a free certificate (for testing)
	//
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The domain name that is secured by the certificate. The domain name uses `HTTPS`-based acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// developer.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to check the certificate name for duplicates. A value of 1 indicates that the system does not perform the check and overwrites the information about the certificate that has the same name. Set the value to **1**.
	//
	// example:
	//
	// 1
	ForceSet *string `json:"ForceSet,omitempty" xml:"ForceSet,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private key.
	//
	// >  This parameter is required only if you set the SSLProtocol parameter to on.
	//
	// example:
	//
	// ****
	SSLPri *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	// Specifies whether to enable the HTTPS certificate. Valid values:
	//
	// 	- **on**. If you set this parameter to **on**, you must also specify the SSLPub and SSLPri parameters.
	//
	// 	- **off**. This is the default value.
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The public key.
	//
	// >  This parameter is required only if you set the SSLProtocol parameter to on.
	//
	// example:
	//
	// ****
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetLiveDomainCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetLiveDomainCertificateRequest) GetCertName() *string {
	return s.CertName
}

func (s *SetLiveDomainCertificateRequest) GetCertType() *string {
	return s.CertType
}

func (s *SetLiveDomainCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveDomainCertificateRequest) GetForceSet() *string {
	return s.ForceSet
}

func (s *SetLiveDomainCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveDomainCertificateRequest) GetSSLPri() *string {
	return s.SSLPri
}

func (s *SetLiveDomainCertificateRequest) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *SetLiveDomainCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *SetLiveDomainCertificateRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetLiveDomainCertificateRequest) SetCertName(v string) *SetLiveDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetCertType(v string) *SetLiveDomainCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetDomainName(v string) *SetLiveDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetForceSet(v string) *SetLiveDomainCertificateRequest {
	s.ForceSet = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetOwnerId(v int64) *SetLiveDomainCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetSSLPri(v string) *SetLiveDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetSSLProtocol(v string) *SetLiveDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetSSLPub(v string) *SetLiveDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) SetSecurityToken(v string) *SetLiveDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetLiveDomainCertificateRequest) Validate() error {
	return dara.Validate(s)
}
