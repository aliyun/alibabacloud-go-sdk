// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomHostnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasId(v int64) *UpdateCustomHostnameRequest
	GetCasId() *int64
	SetCasRegion(v string) *UpdateCustomHostnameRequest
	GetCasRegion() *string
	SetCertType(v string) *UpdateCustomHostnameRequest
	GetCertType() *string
	SetCertificate(v string) *UpdateCustomHostnameRequest
	GetCertificate() *string
	SetHostnameId(v int64) *UpdateCustomHostnameRequest
	GetHostnameId() *int64
	SetPrivateKey(v string) *UpdateCustomHostnameRequest
	GetPrivateKey() *string
	SetRecordId(v int64) *UpdateCustomHostnameRequest
	GetRecordId() *int64
	SetSslFlag(v string) *UpdateCustomHostnameRequest
	GetSslFlag() *string
}

type UpdateCustomHostnameRequest struct {
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
	// cas
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// 证书公钥，使用上传证书时必填
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	HostnameId *int64 `json:"HostnameId,omitempty" xml:"HostnameId,omitempty"`
	// 证书私钥，使用上传证书时必填
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// 绑定的源站记录ID
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// SSL开关
	//
	// example:
	//
	// on
	SslFlag *string `json:"SslFlag,omitempty" xml:"SslFlag,omitempty"`
}

func (s UpdateCustomHostnameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomHostnameRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomHostnameRequest) GetCasId() *int64 {
	return s.CasId
}

func (s *UpdateCustomHostnameRequest) GetCasRegion() *string {
	return s.CasRegion
}

func (s *UpdateCustomHostnameRequest) GetCertType() *string {
	return s.CertType
}

func (s *UpdateCustomHostnameRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *UpdateCustomHostnameRequest) GetHostnameId() *int64 {
	return s.HostnameId
}

func (s *UpdateCustomHostnameRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *UpdateCustomHostnameRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateCustomHostnameRequest) GetSslFlag() *string {
	return s.SslFlag
}

func (s *UpdateCustomHostnameRequest) SetCasId(v int64) *UpdateCustomHostnameRequest {
	s.CasId = &v
	return s
}

func (s *UpdateCustomHostnameRequest) SetCasRegion(v string) *UpdateCustomHostnameRequest {
	s.CasRegion = &v
	return s
}

func (s *UpdateCustomHostnameRequest) SetCertType(v string) *UpdateCustomHostnameRequest {
	s.CertType = &v
	return s
}

func (s *UpdateCustomHostnameRequest) SetCertificate(v string) *UpdateCustomHostnameRequest {
	s.Certificate = &v
	return s
}

func (s *UpdateCustomHostnameRequest) SetHostnameId(v int64) *UpdateCustomHostnameRequest {
	s.HostnameId = &v
	return s
}

func (s *UpdateCustomHostnameRequest) SetPrivateKey(v string) *UpdateCustomHostnameRequest {
	s.PrivateKey = &v
	return s
}

func (s *UpdateCustomHostnameRequest) SetRecordId(v int64) *UpdateCustomHostnameRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateCustomHostnameRequest) SetSslFlag(v string) *UpdateCustomHostnameRequest {
	s.SslFlag = &v
	return s
}

func (s *UpdateCustomHostnameRequest) Validate() error {
	return dara.Validate(s)
}
