// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppDomainCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SetAppDomainCertificateRequest
	GetBizId() *string
	SetCertificateName(v string) *SetAppDomainCertificateRequest
	GetCertificateName() *string
	SetCertificateType(v string) *SetAppDomainCertificateRequest
	GetCertificateType() *string
	SetDomainName(v string) *SetAppDomainCertificateRequest
	GetDomainName() *string
	SetPrivateKey(v string) *SetAppDomainCertificateRequest
	GetPrivateKey() *string
	SetPublicKey(v string) *SetAppDomainCertificateRequest
	GetPublicKey() *string
}

type SetAppDomainCertificateRequest struct {
	// example:
	//
	// WD20250821161210000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 2024
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// example:
	//
	// Server
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// example:
	//
	// kaibaidu.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// ***
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// example:
	//
	// c3NoLWVkMjU1MTkgQUFBQUMzTnphQzFsWkRJMU5URTVBQUFBSUxGQnQxUUpyT3IxK2hTTGRkbERMZUx4WGRIZ3hBalBxWHJIbWNFNWxqSk8gbm93Y29kZXJAbm93Y29kZXJkZU1hY0Jvb2stUHJvLmxvY2Fs
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
}

func (s SetAppDomainCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAppDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetAppDomainCertificateRequest) GetBizId() *string {
	return s.BizId
}

func (s *SetAppDomainCertificateRequest) GetCertificateName() *string {
	return s.CertificateName
}

func (s *SetAppDomainCertificateRequest) GetCertificateType() *string {
	return s.CertificateType
}

func (s *SetAppDomainCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetAppDomainCertificateRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *SetAppDomainCertificateRequest) GetPublicKey() *string {
	return s.PublicKey
}

func (s *SetAppDomainCertificateRequest) SetBizId(v string) *SetAppDomainCertificateRequest {
	s.BizId = &v
	return s
}

func (s *SetAppDomainCertificateRequest) SetCertificateName(v string) *SetAppDomainCertificateRequest {
	s.CertificateName = &v
	return s
}

func (s *SetAppDomainCertificateRequest) SetCertificateType(v string) *SetAppDomainCertificateRequest {
	s.CertificateType = &v
	return s
}

func (s *SetAppDomainCertificateRequest) SetDomainName(v string) *SetAppDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetAppDomainCertificateRequest) SetPrivateKey(v string) *SetAppDomainCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *SetAppDomainCertificateRequest) SetPublicKey(v string) *SetAppDomainCertificateRequest {
	s.PublicKey = &v
	return s
}

func (s *SetAppDomainCertificateRequest) Validate() error {
	return dara.Validate(s)
}
