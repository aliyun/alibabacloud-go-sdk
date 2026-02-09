// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubDomainInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *SubDomainInfo
	GetDomainId() *string
	SetName(v string) *SubDomainInfo
	GetName() *string
	SetNetworkType(v string) *SubDomainInfo
	GetNetworkType() *string
	SetProtocol(v string) *SubDomainInfo
	GetProtocol() *string
}

type SubDomainInfo struct {
	// The domain name ID.
	//
	// example:
	//
	// d-cpudb0llhtgl2djvq2sg
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// abc.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network type. Valid values:
	//
	// Valid values:
	//
	// 	- Intranet
	//
	// 	- Internet
	//
	// example:
	//
	// Intranet
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The protocol.
	//
	// Valid values:
	//
	// 	- HTTPS
	//
	// 	- HTTP
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s SubDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s SubDomainInfo) GoString() string {
	return s.String()
}

func (s *SubDomainInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *SubDomainInfo) GetName() *string {
	return s.Name
}

func (s *SubDomainInfo) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SubDomainInfo) GetProtocol() *string {
	return s.Protocol
}

func (s *SubDomainInfo) SetDomainId(v string) *SubDomainInfo {
	s.DomainId = &v
	return s
}

func (s *SubDomainInfo) SetName(v string) *SubDomainInfo {
	s.Name = &v
	return s
}

func (s *SubDomainInfo) SetNetworkType(v string) *SubDomainInfo {
	s.NetworkType = &v
	return s
}

func (s *SubDomainInfo) SetProtocol(v string) *SubDomainInfo {
	s.Protocol = &v
	return s
}

func (s *SubDomainInfo) Validate() error {
	return dara.Validate(s)
}
