// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomHostnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasId(v int64) *CreateCustomHostnameRequest
	GetCasId() *int64
	SetCasRegion(v string) *CreateCustomHostnameRequest
	GetCasRegion() *string
	SetCertType(v string) *CreateCustomHostnameRequest
	GetCertType() *string
	SetCertificate(v string) *CreateCustomHostnameRequest
	GetCertificate() *string
	SetHostname(v string) *CreateCustomHostnameRequest
	GetHostname() *string
	SetPrivateKey(v string) *CreateCustomHostnameRequest
	GetPrivateKey() *string
	SetRecordId(v int64) *CreateCustomHostnameRequest
	GetRecordId() *int64
	SetSiteId(v int64) *CreateCustomHostnameRequest
	GetSiteId() *int64
	SetSslFlag(v string) *CreateCustomHostnameRequest
	GetSslFlag() *string
}

type CreateCustomHostnameRequest struct {
	// The ID of the CAS certificate. This parameter is required if `CertType` is set to `cas`.
	//
	// example:
	//
	// 30000478
	CasId *int64 `json:"CasId,omitempty" xml:"CasId,omitempty"`
	// The region of the CAS certificate. This parameter is required if `CertType` is set to `cas`.
	//
	// - For accounts on the China site, set this parameter to `cn-hangzhou`.
	//
	// - For accounts on the International site, set this parameter to `ap-southeast-1`.
	//
	// example:
	//
	// cn-hangzhou
	CasRegion *string `json:"CasRegion,omitempty" xml:"CasRegion,omitempty"`
	// The certificate type. This parameter is required if `SslFlag` is set to `on`. Valid values:
	//
	// - **free**: A free certificate.
	//
	// - **upload**: A user-uploaded certificate.
	//
	// - **cas**: A CAS certificate.
	//
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The content of the certificate. This parameter is required if `CertType` is set to `upload`.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The custom hostname.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom.site.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The private key of the certificate. This parameter is required if `CertType` is set to `upload`.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The ID of the record to bind. Call the [ListRecords](https://help.aliyun.com/document_detail/2850265.html) operation to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The site ID. Call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to get this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Specifies whether to enable SSL. Valid values:
	//
	// - **on**: Enable SSL.
	//
	// - **off**: Disable SSL.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	SslFlag *string `json:"SslFlag,omitempty" xml:"SslFlag,omitempty"`
}

func (s CreateCustomHostnameRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomHostnameRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomHostnameRequest) GetCasId() *int64 {
	return s.CasId
}

func (s *CreateCustomHostnameRequest) GetCasRegion() *string {
	return s.CasRegion
}

func (s *CreateCustomHostnameRequest) GetCertType() *string {
	return s.CertType
}

func (s *CreateCustomHostnameRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *CreateCustomHostnameRequest) GetHostname() *string {
	return s.Hostname
}

func (s *CreateCustomHostnameRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CreateCustomHostnameRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *CreateCustomHostnameRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateCustomHostnameRequest) GetSslFlag() *string {
	return s.SslFlag
}

func (s *CreateCustomHostnameRequest) SetCasId(v int64) *CreateCustomHostnameRequest {
	s.CasId = &v
	return s
}

func (s *CreateCustomHostnameRequest) SetCasRegion(v string) *CreateCustomHostnameRequest {
	s.CasRegion = &v
	return s
}

func (s *CreateCustomHostnameRequest) SetCertType(v string) *CreateCustomHostnameRequest {
	s.CertType = &v
	return s
}

func (s *CreateCustomHostnameRequest) SetCertificate(v string) *CreateCustomHostnameRequest {
	s.Certificate = &v
	return s
}

func (s *CreateCustomHostnameRequest) SetHostname(v string) *CreateCustomHostnameRequest {
	s.Hostname = &v
	return s
}

func (s *CreateCustomHostnameRequest) SetPrivateKey(v string) *CreateCustomHostnameRequest {
	s.PrivateKey = &v
	return s
}

func (s *CreateCustomHostnameRequest) SetRecordId(v int64) *CreateCustomHostnameRequest {
	s.RecordId = &v
	return s
}

func (s *CreateCustomHostnameRequest) SetSiteId(v int64) *CreateCustomHostnameRequest {
	s.SiteId = &v
	return s
}

func (s *CreateCustomHostnameRequest) SetSslFlag(v string) *CreateCustomHostnameRequest {
	s.SslFlag = &v
	return s
}

func (s *CreateCustomHostnameRequest) Validate() error {
	return dara.Validate(s)
}
