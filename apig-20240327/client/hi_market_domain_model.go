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
	Domain      *string `json:"domain,omitempty" xml:"domain,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Port        *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Protocol    *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
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
