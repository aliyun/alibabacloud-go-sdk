// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeZoneDnsGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ChangeZoneDnsGroupRequest
	GetClientToken() *string
	SetDnsGroup(v string) *ChangeZoneDnsGroupRequest
	GetDnsGroup() *string
	SetZoneId(v string) *ChangeZoneDnsGroupRequest
	GetZoneId() *string
}

type ChangeZoneDnsGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 85456erer657cfgfg3437
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logical location of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- Normal zone: regular module
	//
	// 	- Fast Zone: acceleration module
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// The global ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0cff188756b1d4579b25e54b66cb830
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ChangeZoneDnsGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeZoneDnsGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeZoneDnsGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ChangeZoneDnsGroupRequest) GetDnsGroup() *string {
	return s.DnsGroup
}

func (s *ChangeZoneDnsGroupRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ChangeZoneDnsGroupRequest) SetClientToken(v string) *ChangeZoneDnsGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeZoneDnsGroupRequest) SetDnsGroup(v string) *ChangeZoneDnsGroupRequest {
	s.DnsGroup = &v
	return s
}

func (s *ChangeZoneDnsGroupRequest) SetZoneId(v string) *ChangeZoneDnsGroupRequest {
	s.ZoneId = &v
	return s
}

func (s *ChangeZoneDnsGroupRequest) Validate() error {
	return dara.Validate(s)
}
