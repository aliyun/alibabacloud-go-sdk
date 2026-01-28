// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceCustomDomain interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *GrafanaWorkspaceCustomDomain
	GetCert() *string
	SetDate(v int64) *GrafanaWorkspaceCustomDomain
	GetDate() *int64
	SetDomain(v string) *GrafanaWorkspaceCustomDomain
	GetDomain() *string
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceCustomDomain
	GetGrafanaWorkspaceId() *string
	SetId(v int64) *GrafanaWorkspaceCustomDomain
	GetId() *int64
	SetKey(v string) *GrafanaWorkspaceCustomDomain
	GetKey() *string
	SetPrivateZone(v string) *GrafanaWorkspaceCustomDomain
	GetPrivateZone() *string
	SetProtocol(v string) *GrafanaWorkspaceCustomDomain
	GetProtocol() *string
	SetStatus(v string) *GrafanaWorkspaceCustomDomain
	GetStatus() *string
	SetUri(v string) *GrafanaWorkspaceCustomDomain
	GetUri() *string
}

type GrafanaWorkspaceCustomDomain struct {
	// example:
	//
	// protocol为http时无需填写;
	Cert *string `json:"cert,omitempty" xml:"cert,omitempty"`
	// example:
	//
	// 1688627798017
	Date *int64 `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// mydomain.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// grafana-cn-***********
	GrafanaWorkspaceId *string `json:"grafanaWorkspaceId,omitempty" xml:"grafanaWorkspaceId,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// protocol为http时无需填写;
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// true
	PrivateZone *string `json:"privateZone,omitempty" xml:"privateZone,omitempty"`
	// example:
	//
	// https
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// CreateSucceed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// /
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s GrafanaWorkspaceCustomDomain) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceCustomDomain) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceCustomDomain) GetCert() *string {
	return s.Cert
}

func (s *GrafanaWorkspaceCustomDomain) GetDate() *int64 {
	return s.Date
}

func (s *GrafanaWorkspaceCustomDomain) GetDomain() *string {
	return s.Domain
}

func (s *GrafanaWorkspaceCustomDomain) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspaceCustomDomain) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceCustomDomain) GetKey() *string {
	return s.Key
}

func (s *GrafanaWorkspaceCustomDomain) GetPrivateZone() *string {
	return s.PrivateZone
}

func (s *GrafanaWorkspaceCustomDomain) GetProtocol() *string {
	return s.Protocol
}

func (s *GrafanaWorkspaceCustomDomain) GetStatus() *string {
	return s.Status
}

func (s *GrafanaWorkspaceCustomDomain) GetUri() *string {
	return s.Uri
}

func (s *GrafanaWorkspaceCustomDomain) SetCert(v string) *GrafanaWorkspaceCustomDomain {
	s.Cert = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetDate(v int64) *GrafanaWorkspaceCustomDomain {
	s.Date = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetDomain(v string) *GrafanaWorkspaceCustomDomain {
	s.Domain = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceCustomDomain {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetId(v int64) *GrafanaWorkspaceCustomDomain {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetKey(v string) *GrafanaWorkspaceCustomDomain {
	s.Key = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetPrivateZone(v string) *GrafanaWorkspaceCustomDomain {
	s.PrivateZone = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetProtocol(v string) *GrafanaWorkspaceCustomDomain {
	s.Protocol = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetStatus(v string) *GrafanaWorkspaceCustomDomain {
	s.Status = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) SetUri(v string) *GrafanaWorkspaceCustomDomain {
	s.Uri = &v
	return s
}

func (s *GrafanaWorkspaceCustomDomain) Validate() error {
	return dara.Validate(s)
}
