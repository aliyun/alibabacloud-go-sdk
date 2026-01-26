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
	// 云盾证书ID，使用云盾证书时必填
	//
	// example:
	//
	// 30000478
	CasId *int64 `json:"CasId,omitempty" xml:"CasId,omitempty"`
	// 云盾证书所在地域，使用云盾证书时必填
	//
	// example:
	//
	// cn-hangzhou
	CasRegion *string `json:"CasRegion,omitempty" xml:"CasRegion,omitempty"`
	// 证书类型，SSL 开启时必填
	//
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// 证书公钥，使用上传证书时必填
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// 自定义主机名
	//
	// This parameter is required.
	//
	// example:
	//
	// custom.site.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// 证书私钥，使用上传证书时必填
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// 绑定的源站记录ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// 关联站点ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// SSL开关
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
