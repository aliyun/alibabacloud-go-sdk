// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayDomain interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *GatewayDomain
	GetCertIdentifier() *string
	SetGatewayId(v int64) *GatewayDomain
	GetGatewayId() *int64
	SetGatewayName(v string) *GatewayDomain
	GetGatewayName() *string
	SetGatewayUniqueId(v string) *GatewayDomain
	GetGatewayUniqueId() *string
	SetGmtCreate(v string) *GatewayDomain
	GetGmtCreate() *string
	SetGmtModified(v string) *GatewayDomain
	GetGmtModified() *string
	SetId(v int64) *GatewayDomain
	GetId() *int64
	SetMustHttps(v string) *GatewayDomain
	GetMustHttps() *string
	SetName(v string) *GatewayDomain
	GetName() *string
	SetProtocol(v string) *GatewayDomain
	GetProtocol() *string
}

type GatewayDomain struct {
	CertIdentifier  *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	GatewayId       *int64  `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayName     *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MustHttps       *string `json:"MustHttps,omitempty" xml:"MustHttps,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s GatewayDomain) String() string {
	return dara.Prettify(s)
}

func (s GatewayDomain) GoString() string {
	return s.String()
}

func (s *GatewayDomain) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GatewayDomain) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GatewayDomain) GetGatewayName() *string {
	return s.GatewayName
}

func (s *GatewayDomain) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GatewayDomain) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GatewayDomain) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GatewayDomain) GetId() *int64 {
	return s.Id
}

func (s *GatewayDomain) GetMustHttps() *string {
	return s.MustHttps
}

func (s *GatewayDomain) GetName() *string {
	return s.Name
}

func (s *GatewayDomain) GetProtocol() *string {
	return s.Protocol
}

func (s *GatewayDomain) SetCertIdentifier(v string) *GatewayDomain {
	s.CertIdentifier = &v
	return s
}

func (s *GatewayDomain) SetGatewayId(v int64) *GatewayDomain {
	s.GatewayId = &v
	return s
}

func (s *GatewayDomain) SetGatewayName(v string) *GatewayDomain {
	s.GatewayName = &v
	return s
}

func (s *GatewayDomain) SetGatewayUniqueId(v string) *GatewayDomain {
	s.GatewayUniqueId = &v
	return s
}

func (s *GatewayDomain) SetGmtCreate(v string) *GatewayDomain {
	s.GmtCreate = &v
	return s
}

func (s *GatewayDomain) SetGmtModified(v string) *GatewayDomain {
	s.GmtModified = &v
	return s
}

func (s *GatewayDomain) SetId(v int64) *GatewayDomain {
	s.Id = &v
	return s
}

func (s *GatewayDomain) SetMustHttps(v string) *GatewayDomain {
	s.MustHttps = &v
	return s
}

func (s *GatewayDomain) SetName(v string) *GatewayDomain {
	s.Name = &v
	return s
}

func (s *GatewayDomain) SetProtocol(v string) *GatewayDomain {
	s.Protocol = &v
	return s
}

func (s *GatewayDomain) Validate() error {
	return dara.Validate(s)
}
