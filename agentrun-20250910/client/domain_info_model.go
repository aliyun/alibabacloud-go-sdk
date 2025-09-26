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
	SetDomainId(v string) *DomainInfo
	GetDomainId() *string
	SetName(v string) *DomainInfo
	GetName() *string
	SetProtocol(v string) *DomainInfo
	GetProtocol() *string
}

type DomainInfo struct {
	CertIdentifier *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	DomainId       *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Protocol       *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
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

func (s *DomainInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *DomainInfo) GetName() *string {
	return s.Name
}

func (s *DomainInfo) GetProtocol() *string {
	return s.Protocol
}

func (s *DomainInfo) SetCertIdentifier(v string) *DomainInfo {
	s.CertIdentifier = &v
	return s
}

func (s *DomainInfo) SetDomainId(v string) *DomainInfo {
	s.DomainId = &v
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

func (s *DomainInfo) Validate() error {
	return dara.Validate(s)
}
