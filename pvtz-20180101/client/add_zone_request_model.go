// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddZoneRequest
	GetClientToken() *string
	SetDnsGroup(v string) *AddZoneRequest
	GetDnsGroup() *string
	SetLang(v string) *AddZoneRequest
	GetLang() *string
	SetProxyPattern(v string) *AddZoneRequest
	GetProxyPattern() *string
	SetResourceGroupId(v string) *AddZoneRequest
	GetResourceGroupId() *string
	SetZoneName(v string) *AddZoneRequest
	GetZoneName() *string
	SetZoneTag(v string) *AddZoneRequest
	GetZoneTag() *string
	SetZoneType(v string) *AddZoneRequest
	GetZoneType() *string
}

type AddZoneRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The location of the built-in authoritative zone. Valid values:
	//
	// - **NORMAL_ZONE**: Standard zone. DNS responses are cached. If a cache miss occurs, the query is sent to the built-in authoritative standard zone. The time to live (TTL) value affects the time when a DNS record change takes effect. You cannot use custom DNS lines or weighted round-robin.
	//
	// - **FAST_ZONE**: Accelerated zone (recommended). DNS queries are directly responded to with the lowest latency. DNS record changes take effect in real time. You can use custom DNS lines and weighted round-robin.
	//
	// Default value: **NORMAL_ZONE**.
	//
	// > The built-in authoritative accelerated zone is located before the cache module. DNS responses are not cached. This may increase the number of DNS queries and your costs.
	//
	// <props="china">
	//
	// > Starting from April 30, 2025 (UTC+8), when new users activate Alibaba Cloud DNS PrivateZone, added zones are set as accelerated zones by default.
	//
	// example:
	//
	// FAST_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to enable subdomain recursive proxy. Valid values:
	//
	// - **ZONE**: Disables the feature. If a DNS query for a subdomain that does not exist under the current domain name is received, an NXDOMAIN error is returned.
	//
	// - **RECORD**: Enables the feature. If a DNS query for a subdomain that does not exist under the current domain name is received, the query is processed by the forwarding and recursion modules in sequence. The final result is used to respond to the DNS query.
	//
	// Default value: **ZONE**.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmykd63gt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the zone to add.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// This parameter is not available to users. You do not need to specify this parameter.
	//
	// example:
	//
	// BLINK
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// This parameter is not available to users. You do not need to specify this parameter.
	//
	// example:
	//
	// CLOUD_PRODUCT_ZONE
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s AddZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s AddZoneRequest) GoString() string {
	return s.String()
}

func (s *AddZoneRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddZoneRequest) GetDnsGroup() *string {
	return s.DnsGroup
}

func (s *AddZoneRequest) GetLang() *string {
	return s.Lang
}

func (s *AddZoneRequest) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *AddZoneRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddZoneRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *AddZoneRequest) GetZoneTag() *string {
	return s.ZoneTag
}

func (s *AddZoneRequest) GetZoneType() *string {
	return s.ZoneType
}

func (s *AddZoneRequest) SetClientToken(v string) *AddZoneRequest {
	s.ClientToken = &v
	return s
}

func (s *AddZoneRequest) SetDnsGroup(v string) *AddZoneRequest {
	s.DnsGroup = &v
	return s
}

func (s *AddZoneRequest) SetLang(v string) *AddZoneRequest {
	s.Lang = &v
	return s
}

func (s *AddZoneRequest) SetProxyPattern(v string) *AddZoneRequest {
	s.ProxyPattern = &v
	return s
}

func (s *AddZoneRequest) SetResourceGroupId(v string) *AddZoneRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddZoneRequest) SetZoneName(v string) *AddZoneRequest {
	s.ZoneName = &v
	return s
}

func (s *AddZoneRequest) SetZoneTag(v string) *AddZoneRequest {
	s.ZoneTag = &v
	return s
}

func (s *AddZoneRequest) SetZoneType(v string) *AddZoneRequest {
	s.ZoneType = &v
	return s
}

func (s *AddZoneRequest) Validate() error {
	return dara.Validate(s)
}
