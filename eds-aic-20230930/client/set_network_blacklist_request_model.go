// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetNetworkBlacklistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainBlacklist(v []*string) *SetNetworkBlacklistRequest
	GetDomainBlacklist() []*string
	SetIpBlacklist(v []*string) *SetNetworkBlacklistRequest
	GetIpBlacklist() []*string
}

type SetNetworkBlacklistRequest struct {
	DomainBlacklist []*string `json:"DomainBlacklist,omitempty" xml:"DomainBlacklist,omitempty" type:"Repeated"`
	IpBlacklist     []*string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty" type:"Repeated"`
}

func (s SetNetworkBlacklistRequest) String() string {
	return dara.Prettify(s)
}

func (s SetNetworkBlacklistRequest) GoString() string {
	return s.String()
}

func (s *SetNetworkBlacklistRequest) GetDomainBlacklist() []*string {
	return s.DomainBlacklist
}

func (s *SetNetworkBlacklistRequest) GetIpBlacklist() []*string {
	return s.IpBlacklist
}

func (s *SetNetworkBlacklistRequest) SetDomainBlacklist(v []*string) *SetNetworkBlacklistRequest {
	s.DomainBlacklist = v
	return s
}

func (s *SetNetworkBlacklistRequest) SetIpBlacklist(v []*string) *SetNetworkBlacklistRequest {
	s.IpBlacklist = v
	return s
}

func (s *SetNetworkBlacklistRequest) Validate() error {
	return dara.Validate(s)
}
