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
	// A client token that is used to ensure the idempotence of the request. Generate a unique value for this parameter on your client. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 85456erer657cfgfg3437
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The built-in authoritative DNS group.
	//
	// - Standard zone group: NORMAL_ZONE
	//
	// - Acceleration zone group: FAST_ZONE
	//
	// <props="china">
	//
	// > Starting from April 30, 2025 (UTC+8), when new users of Alibaba Cloud DNS PrivateZone create a zone, the zone is set to an acceleration zone by default.
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
