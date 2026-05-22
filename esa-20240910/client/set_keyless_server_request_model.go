// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKeylessServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertificate(v string) *SetKeylessServerRequest
	GetCaCertificate() *string
	SetClientCertificate(v string) *SetKeylessServerRequest
	GetClientCertificate() *string
	SetClientPrivateKey(v string) *SetKeylessServerRequest
	GetClientPrivateKey() *string
	SetHost(v string) *SetKeylessServerRequest
	GetHost() *string
	SetId(v string) *SetKeylessServerRequest
	GetId() *string
	SetName(v string) *SetKeylessServerRequest
	GetName() *string
	SetPort(v int64) *SetKeylessServerRequest
	GetPort() *int64
	SetSiteId(v int64) *SetKeylessServerRequest
	GetSiteId() *int64
	SetVerify(v bool) *SetKeylessServerRequest
	GetVerify() *bool
}

type SetKeylessServerRequest struct {
	// example:
	//
	// -----BEGIN CERTIFICATE-----****
	CaCertificate *string `json:"CaCertificate,omitempty" xml:"CaCertificate,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----****
	ClientCertificate *string `json:"ClientCertificate,omitempty" xml:"ClientCertificate,omitempty"`
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----****
	ClientPrivateKey *string `json:"ClientPrivateKey,omitempty" xml:"ClientPrivateKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// keyless.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Keyless server ID。
	//
	// example:
	//
	// baba39055622c008b90285a8838e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 443
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// true
	Verify *bool `json:"Verify,omitempty" xml:"Verify,omitempty"`
}

func (s SetKeylessServerRequest) String() string {
	return dara.Prettify(s)
}

func (s SetKeylessServerRequest) GoString() string {
	return s.String()
}

func (s *SetKeylessServerRequest) GetCaCertificate() *string {
	return s.CaCertificate
}

func (s *SetKeylessServerRequest) GetClientCertificate() *string {
	return s.ClientCertificate
}

func (s *SetKeylessServerRequest) GetClientPrivateKey() *string {
	return s.ClientPrivateKey
}

func (s *SetKeylessServerRequest) GetHost() *string {
	return s.Host
}

func (s *SetKeylessServerRequest) GetId() *string {
	return s.Id
}

func (s *SetKeylessServerRequest) GetName() *string {
	return s.Name
}

func (s *SetKeylessServerRequest) GetPort() *int64 {
	return s.Port
}

func (s *SetKeylessServerRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetKeylessServerRequest) GetVerify() *bool {
	return s.Verify
}

func (s *SetKeylessServerRequest) SetCaCertificate(v string) *SetKeylessServerRequest {
	s.CaCertificate = &v
	return s
}

func (s *SetKeylessServerRequest) SetClientCertificate(v string) *SetKeylessServerRequest {
	s.ClientCertificate = &v
	return s
}

func (s *SetKeylessServerRequest) SetClientPrivateKey(v string) *SetKeylessServerRequest {
	s.ClientPrivateKey = &v
	return s
}

func (s *SetKeylessServerRequest) SetHost(v string) *SetKeylessServerRequest {
	s.Host = &v
	return s
}

func (s *SetKeylessServerRequest) SetId(v string) *SetKeylessServerRequest {
	s.Id = &v
	return s
}

func (s *SetKeylessServerRequest) SetName(v string) *SetKeylessServerRequest {
	s.Name = &v
	return s
}

func (s *SetKeylessServerRequest) SetPort(v int64) *SetKeylessServerRequest {
	s.Port = &v
	return s
}

func (s *SetKeylessServerRequest) SetSiteId(v int64) *SetKeylessServerRequest {
	s.SiteId = &v
	return s
}

func (s *SetKeylessServerRequest) SetVerify(v bool) *SetKeylessServerRequest {
	s.Verify = &v
	return s
}

func (s *SetKeylessServerRequest) Validate() error {
	return dara.Validate(s)
}
