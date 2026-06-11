// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketDomain interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *HiMarketDomain
	GetDomain() *string
	SetNetworkType(v string) *HiMarketDomain
	GetNetworkType() *string
	SetPort(v int32) *HiMarketDomain
	GetPort() *int32
	SetProtocol(v string) *HiMarketDomain
	GetProtocol() *string
}

type HiMarketDomain struct {
	// The custom domain name. This must be a valid DNS hostname.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The network type of the endpoint. For example, `VPC` for an internal network or `INTERNET` for a public network.
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The port number for the endpoint. For example, `80` for HTTP or `443` for HTTPS.
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The communication protocol. Valid values include `HTTP` and `HTTPS`.
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HiMarketDomain) String() string {
	return dara.Prettify(s)
}

func (s HiMarketDomain) GoString() string {
	return s.String()
}

func (s *HiMarketDomain) GetDomain() *string {
	return s.Domain
}

func (s *HiMarketDomain) GetNetworkType() *string {
	return s.NetworkType
}

func (s *HiMarketDomain) GetPort() *int32 {
	return s.Port
}

func (s *HiMarketDomain) GetProtocol() *string {
	return s.Protocol
}

func (s *HiMarketDomain) SetDomain(v string) *HiMarketDomain {
	s.Domain = &v
	return s
}

func (s *HiMarketDomain) SetNetworkType(v string) *HiMarketDomain {
	s.NetworkType = &v
	return s
}

func (s *HiMarketDomain) SetPort(v int32) *HiMarketDomain {
	s.Port = &v
	return s
}

func (s *HiMarketDomain) SetProtocol(v string) *HiMarketDomain {
	s.Protocol = &v
	return s
}

func (s *HiMarketDomain) Validate() error {
	return dara.Validate(s)
}
