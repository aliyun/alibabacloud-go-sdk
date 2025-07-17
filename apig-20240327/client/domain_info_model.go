// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *DomainInfo
	GetCertIdentifier() *string
	SetClientCACert(v string) *DomainInfo
	GetClientCACert() *string
	SetCreateFrom(v string) *DomainInfo
	GetCreateFrom() *string
	SetCreateTimestamp(v int64) *DomainInfo
	GetCreateTimestamp() *int64
	SetDomainId(v string) *DomainInfo
	GetDomainId() *string
	SetForceHttps(v bool) *DomainInfo
	GetForceHttps() *bool
	SetMTLSEnabled(v bool) *DomainInfo
	GetMTLSEnabled() *bool
	SetName(v string) *DomainInfo
	GetName() *string
	SetProtocol(v string) *DomainInfo
	GetProtocol() *string
	SetResourceGroupId(v string) *DomainInfo
	GetResourceGroupId() *string
	SetStatus(v string) *DomainInfo
	GetStatus() *string
	SetUpdateTimestamp(v int64) *DomainInfo
	GetUpdateTimestamp() *int64
}

type DomainInfo struct {
	CertIdentifier  *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	ClientCACert    *string `json:"clientCACert,omitempty" xml:"clientCACert,omitempty"`
	CreateFrom      *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	CreateTimestamp *int64  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	DomainId        *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	ForceHttps      *bool   `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	MTLSEnabled     *bool   `json:"mTLSEnabled,omitempty" xml:"mTLSEnabled,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	Protocol        *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTimestamp *int64  `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s DomainInfo) String() string {
	return dara.Prettify(s)
}

func (s DomainInfo) GoString() string {
	return s.String()
}

func (s *DomainInfo) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DomainInfo) GetClientCACert() *string {
	return s.ClientCACert
}

func (s *DomainInfo) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *DomainInfo) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DomainInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *DomainInfo) GetForceHttps() *bool {
	return s.ForceHttps
}

func (s *DomainInfo) GetMTLSEnabled() *bool {
	return s.MTLSEnabled
}

func (s *DomainInfo) GetName() *string {
	return s.Name
}

func (s *DomainInfo) GetProtocol() *string {
	return s.Protocol
}

func (s *DomainInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DomainInfo) GetStatus() *string {
	return s.Status
}

func (s *DomainInfo) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DomainInfo) SetCertIdentifier(v string) *DomainInfo {
	s.CertIdentifier = &v
	return s
}

func (s *DomainInfo) SetClientCACert(v string) *DomainInfo {
	s.ClientCACert = &v
	return s
}

func (s *DomainInfo) SetCreateFrom(v string) *DomainInfo {
	s.CreateFrom = &v
	return s
}

func (s *DomainInfo) SetCreateTimestamp(v int64) *DomainInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *DomainInfo) SetDomainId(v string) *DomainInfo {
	s.DomainId = &v
	return s
}

func (s *DomainInfo) SetForceHttps(v bool) *DomainInfo {
	s.ForceHttps = &v
	return s
}

func (s *DomainInfo) SetMTLSEnabled(v bool) *DomainInfo {
	s.MTLSEnabled = &v
	return s
}

func (s *DomainInfo) SetName(v string) *DomainInfo {
	s.Name = &v
	return s
}

func (s *DomainInfo) SetProtocol(v string) *DomainInfo {
	s.Protocol = &v
	return s
}

func (s *DomainInfo) SetResourceGroupId(v string) *DomainInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *DomainInfo) SetStatus(v string) *DomainInfo {
	s.Status = &v
	return s
}

func (s *DomainInfo) SetUpdateTimestamp(v int64) *DomainInfo {
	s.UpdateTimestamp = &v
	return s
}

func (s *DomainInfo) Validate() error {
	return dara.Validate(s)
}
