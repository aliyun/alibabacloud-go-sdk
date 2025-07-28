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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logical location type of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- **NORMAL_ZONE**: the regular module. DNS results are stored in the cache module and DNS requests are sent to the regular module if the DNS requests do not match the DNS records in the cache module. DNS record updates take effect based on the time to live (TTL) value. The regular module does not support DNS resolution over user-defined lines or based on weight values.
	//
	// 	- **FAST_ZONE**: the acceleration module. It directly responds to DNS requests with the lowest latency and updates DNS records in real time. The acceleration module supports DNS resolution over user-defined lines or based on weight values.
	//
	// Default value: **NORMAL_ZONE**.
	//
	// >  The DNS results returned by the built-in authoritative acceleration module are not stored in the cache module because the built-in authoritative acceleration module is located before the cache module. As a result, you are charged more for DNS requests.
	//
	// example:
	//
	// FAST_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to enable the recursive resolution proxy for subdomain names. Valid values:
	//
	// 	- **ZONE**: disables the recursive resolution proxy for subdomain names. In this case, NXDOMAIN is returned if the queried subdomain name does not exist in the zone.
	//
	// 	- **RECORD**: enables the recursive resolution proxy for subdomain names. In this case, if the queried subdomain name does not exist in the zone, DNS requests are recursively forwarded to the forward module and then to the recursion module until DNS results are returned.
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
	// The name of the zone to be added.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// This parameter is not available. You can ignore it.
	//
	// example:
	//
	// BLINK
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// This parameter is not available. You can ignore it.
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
