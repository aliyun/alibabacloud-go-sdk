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
	CasId         *int64  `json:"CasId,omitempty" xml:"CasId,omitempty"`
	Certificate   *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	KeyServerId   *string `json:"KeyServerId,omitempty" xml:"KeyServerId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PrivateKey    *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
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
