// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketPortalDomainConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *HiMarketPortalDomainConfig
	GetDomain() *string
	SetProtocol(v string) *HiMarketPortalDomainConfig
	GetProtocol() *string
	SetType(v string) *HiMarketPortalDomainConfig
	GetType() *string
}

type HiMarketPortalDomainConfig struct {
	// The domain name address.
	//
	// example:
	//
	// portal.example.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The request protocol.
	//
	// example:
	//
	// HTTPS
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The domain name type.
	//
	// example:
	//
	// custom
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HiMarketPortalDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketPortalDomainConfig) GoString() string {
	return s.String()
}

func (s *HiMarketPortalDomainConfig) GetDomain() *string {
	return s.Domain
}

func (s *HiMarketPortalDomainConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *HiMarketPortalDomainConfig) GetType() *string {
	return s.Type
}

func (s *HiMarketPortalDomainConfig) SetDomain(v string) *HiMarketPortalDomainConfig {
	s.Domain = &v
	return s
}

func (s *HiMarketPortalDomainConfig) SetProtocol(v string) *HiMarketPortalDomainConfig {
	s.Protocol = &v
	return s
}

func (s *HiMarketPortalDomainConfig) SetType(v string) *HiMarketPortalDomainConfig {
	s.Type = &v
	return s
}

func (s *HiMarketPortalDomainConfig) Validate() error {
	return dara.Validate(s)
}
