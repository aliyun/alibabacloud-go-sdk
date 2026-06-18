// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeylessServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertificate(v string) *GetKeylessServerResponseBody
	GetCaCertificate() *string
	SetClientCertificate(v string) *GetKeylessServerResponseBody
	GetClientCertificate() *string
	SetClientPrivateKey(v string) *GetKeylessServerResponseBody
	GetClientPrivateKey() *string
	SetCreateTime(v string) *GetKeylessServerResponseBody
	GetCreateTime() *string
	SetHost(v string) *GetKeylessServerResponseBody
	GetHost() *string
	SetId(v string) *GetKeylessServerResponseBody
	GetId() *string
	SetName(v string) *GetKeylessServerResponseBody
	GetName() *string
	SetPort(v int64) *GetKeylessServerResponseBody
	GetPort() *int64
	SetRequestId(v string) *GetKeylessServerResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetKeylessServerResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetKeylessServerResponseBody
	GetSiteName() *string
	SetUpdateTime(v string) *GetKeylessServerResponseBody
	GetUpdateTime() *string
	SetVerify(v bool) *GetKeylessServerResponseBody
	GetVerify() *bool
}

type GetKeylessServerResponseBody struct {
	// The CA certificate used to verify the server certificate of the keyless server. This parameter applies only when Verify is set to true.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----****
	CaCertificate *string `json:"CaCertificate,omitempty" xml:"CaCertificate,omitempty"`
	// The client certificate. Must be provided as a pair with the client private key.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----****
	ClientCertificate *string `json:"ClientCertificate,omitempty" xml:"ClientCertificate,omitempty"`
	// The client private key. Must be provided as a pair with the client certificate.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----****
	ClientPrivateKey *string `json:"ClientPrivateKey,omitempty" xml:"ClientPrivateKey,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-03-11T01:23:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The hostname of the keyless server.
	//
	// example:
	//
	// example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The keyless server ID.
	//
	// example:
	//
	// baba39055622c008b90285a8838e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The keyless server name.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port of the keyless server. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 443
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3558df77-8a7a-4060-a900-2d794940****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The site name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-03-13T02:13:28Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Indicates whether to verify the server certificate of the keyless server. The default value is false.
	//
	// example:
	//
	// true
	Verify *bool `json:"Verify,omitempty" xml:"Verify,omitempty"`
}

func (s GetKeylessServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKeylessServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeylessServerResponseBody) GetCaCertificate() *string {
	return s.CaCertificate
}

func (s *GetKeylessServerResponseBody) GetClientCertificate() *string {
	return s.ClientCertificate
}

func (s *GetKeylessServerResponseBody) GetClientPrivateKey() *string {
	return s.ClientPrivateKey
}

func (s *GetKeylessServerResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetKeylessServerResponseBody) GetHost() *string {
	return s.Host
}

func (s *GetKeylessServerResponseBody) GetId() *string {
	return s.Id
}

func (s *GetKeylessServerResponseBody) GetName() *string {
	return s.Name
}

func (s *GetKeylessServerResponseBody) GetPort() *int64 {
	return s.Port
}

func (s *GetKeylessServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKeylessServerResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetKeylessServerResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetKeylessServerResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetKeylessServerResponseBody) GetVerify() *bool {
	return s.Verify
}

func (s *GetKeylessServerResponseBody) SetCaCertificate(v string) *GetKeylessServerResponseBody {
	s.CaCertificate = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetClientCertificate(v string) *GetKeylessServerResponseBody {
	s.ClientCertificate = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetClientPrivateKey(v string) *GetKeylessServerResponseBody {
	s.ClientPrivateKey = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetCreateTime(v string) *GetKeylessServerResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetHost(v string) *GetKeylessServerResponseBody {
	s.Host = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetId(v string) *GetKeylessServerResponseBody {
	s.Id = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetName(v string) *GetKeylessServerResponseBody {
	s.Name = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetPort(v int64) *GetKeylessServerResponseBody {
	s.Port = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetRequestId(v string) *GetKeylessServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetSiteId(v int64) *GetKeylessServerResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetSiteName(v string) *GetKeylessServerResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetUpdateTime(v string) *GetKeylessServerResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetKeylessServerResponseBody) SetVerify(v bool) *GetKeylessServerResponseBody {
	s.Verify = &v
	return s
}

func (s *GetKeylessServerResponseBody) Validate() error {
	return dara.Validate(s)
}
