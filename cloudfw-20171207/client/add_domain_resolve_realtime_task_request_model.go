// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainResolveRealtimeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddDomainResolveRealtimeTaskRequest
	GetDomainName() *string
	SetFirewallType(v string) *AddDomainResolveRealtimeTaskRequest
	GetFirewallType() *string
	SetRegionNo(v string) *AddDomainResolveRealtimeTaskRequest
	GetRegionNo() *string
}

type AddDomainResolveRealtimeTaskRequest struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// internet
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s AddDomainResolveRealtimeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDomainResolveRealtimeTaskRequest) GoString() string {
	return s.String()
}

func (s *AddDomainResolveRealtimeTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDomainResolveRealtimeTaskRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *AddDomainResolveRealtimeTaskRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *AddDomainResolveRealtimeTaskRequest) SetDomainName(v string) *AddDomainResolveRealtimeTaskRequest {
	s.DomainName = &v
	return s
}

func (s *AddDomainResolveRealtimeTaskRequest) SetFirewallType(v string) *AddDomainResolveRealtimeTaskRequest {
	s.FirewallType = &v
	return s
}

func (s *AddDomainResolveRealtimeTaskRequest) SetRegionNo(v string) *AddDomainResolveRealtimeTaskRequest {
	s.RegionNo = &v
	return s
}

func (s *AddDomainResolveRealtimeTaskRequest) Validate() error {
	return dara.Validate(s)
}
