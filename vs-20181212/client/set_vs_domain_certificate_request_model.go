// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVsDomainCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *SetVsDomainCertificateRequest
	GetCertName() *string
	SetCertType(v string) *SetVsDomainCertificateRequest
	GetCertType() *string
	SetDomainName(v string) *SetVsDomainCertificateRequest
	GetDomainName() *string
	SetForceSet(v string) *SetVsDomainCertificateRequest
	GetForceSet() *string
	SetOwnerId(v int64) *SetVsDomainCertificateRequest
	GetOwnerId() *int64
	SetRegion(v string) *SetVsDomainCertificateRequest
	GetRegion() *string
	SetSSLPri(v string) *SetVsDomainCertificateRequest
	GetSSLPri() *string
	SetSSLProtocol(v string) *SetVsDomainCertificateRequest
	GetSSLProtocol() *string
	SetSSLPub(v string) *SetVsDomainCertificateRequest
	GetSSLPub() *string
}

type SetVsDomainCertificateRequest struct {
	// example:
	//
	// Cert-77****7
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1
	ForceSet *string `json:"ForceSet,omitempty" xml:"ForceSet,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// xxxxxxx
	SSLPri *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// example:
	//
	// 328uiuii28****82dsada81
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
}

func (s SetVsDomainCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetVsDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetVsDomainCertificateRequest) GetCertName() *string {
	return s.CertName
}

func (s *SetVsDomainCertificateRequest) GetCertType() *string {
	return s.CertType
}

func (s *SetVsDomainCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetVsDomainCertificateRequest) GetForceSet() *string {
	return s.ForceSet
}

func (s *SetVsDomainCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetVsDomainCertificateRequest) GetRegion() *string {
	return s.Region
}

func (s *SetVsDomainCertificateRequest) GetSSLPri() *string {
	return s.SSLPri
}

func (s *SetVsDomainCertificateRequest) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *SetVsDomainCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *SetVsDomainCertificateRequest) SetCertName(v string) *SetVsDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetCertType(v string) *SetVsDomainCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetDomainName(v string) *SetVsDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetForceSet(v string) *SetVsDomainCertificateRequest {
	s.ForceSet = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetOwnerId(v int64) *SetVsDomainCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetRegion(v string) *SetVsDomainCertificateRequest {
	s.Region = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetSSLPri(v string) *SetVsDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetSSLProtocol(v string) *SetVsDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetVsDomainCertificateRequest) SetSSLPub(v string) *SetVsDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetVsDomainCertificateRequest) Validate() error {
	return dara.Validate(s)
}
