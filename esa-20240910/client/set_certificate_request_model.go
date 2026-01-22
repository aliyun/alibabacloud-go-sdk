// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasId(v int64) *SetCertificateRequest
	GetCasId() *int64
	SetCertificate(v string) *SetCertificateRequest
	GetCertificate() *string
	SetId(v string) *SetCertificateRequest
	GetId() *string
	SetKeyServerId(v string) *SetCertificateRequest
	GetKeyServerId() *string
	SetName(v string) *SetCertificateRequest
	GetName() *string
	SetOwnerId(v int64) *SetCertificateRequest
	GetOwnerId() *int64
	SetPrivateKey(v string) *SetCertificateRequest
	GetPrivateKey() *string
	SetRegion(v string) *SetCertificateRequest
	GetRegion() *string
	SetSecurityToken(v string) *SetCertificateRequest
	GetSecurityToken() *string
	SetSiteId(v int64) *SetCertificateRequest
	GetSiteId() *int64
	SetType(v string) *SetCertificateRequest
	GetType() *string
}

type SetCertificateRequest struct {
	// The certificate ID on Certificate Management Service.
	//
	// example:
	//
	// 30000478
	CasId *int64 `json:"CasId,omitempty" xml:"CasId,omitempty"`
	// The certificate content.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate ID on ESA.
	//
	// example:
	//
	// 30001303
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	KeyServerId *string `json:"KeyServerId,omitempty" xml:"KeyServerId,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// yourCertName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private key of the certificate.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The certificate type. Valid values:
	//
	// 	- cas: a certificate purchased by using Certificate Management Service.
	//
	// 	- upload: a custom certificate that you upload.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SetCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetCertificateRequest) GetCasId() *int64 {
	return s.CasId
}

func (s *SetCertificateRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *SetCertificateRequest) GetId() *string {
	return s.Id
}

func (s *SetCertificateRequest) GetKeyServerId() *string {
	return s.KeyServerId
}

func (s *SetCertificateRequest) GetName() *string {
	return s.Name
}

func (s *SetCertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCertificateRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *SetCertificateRequest) GetRegion() *string {
	return s.Region
}

func (s *SetCertificateRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetCertificateRequest) GetType() *string {
	return s.Type
}

func (s *SetCertificateRequest) SetCasId(v int64) *SetCertificateRequest {
	s.CasId = &v
	return s
}

func (s *SetCertificateRequest) SetCertificate(v string) *SetCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *SetCertificateRequest) SetId(v string) *SetCertificateRequest {
	s.Id = &v
	return s
}

func (s *SetCertificateRequest) SetKeyServerId(v string) *SetCertificateRequest {
	s.KeyServerId = &v
	return s
}

func (s *SetCertificateRequest) SetName(v string) *SetCertificateRequest {
	s.Name = &v
	return s
}

func (s *SetCertificateRequest) SetOwnerId(v int64) *SetCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCertificateRequest) SetPrivateKey(v string) *SetCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *SetCertificateRequest) SetRegion(v string) *SetCertificateRequest {
	s.Region = &v
	return s
}

func (s *SetCertificateRequest) SetSecurityToken(v string) *SetCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetCertificateRequest) SetSiteId(v int64) *SetCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *SetCertificateRequest) SetType(v string) *SetCertificateRequest {
	s.Type = &v
	return s
}

func (s *SetCertificateRequest) Validate() error {
	return dara.Validate(s)
}
