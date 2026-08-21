// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetProxyPatternRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SetProxyPatternRequest
	GetClientToken() *string
	SetLang(v string) *SetProxyPatternRequest
	GetLang() *string
	SetProxyPattern(v string) *SetProxyPatternRequest
	GetProxyPattern() *string
	SetUserClientIp(v string) *SetProxyPatternRequest
	GetUserClientIp() *string
	SetZoneId(v string) *SetProxyPatternRequest
	GetZoneId() *string
}

type SetProxyPatternRequest struct {
	// A client token to ensure the idempotence of the request. The token must be unique for each request. It can contain only ASCII characters and must be no more than 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// Default value: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The proxy mode for recursive resolution of subdomains. Valid values:
	//
	// - **ZONE**: Disables the proxy. If a subdomain does not exist, an NXDOMAIN response is returned.
	//
	// - **RECORD**: Enables the proxy. If a subdomain does not exist, the system queries the forwarding and recursion modules and returns the final result.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 10.61.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The unique ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SetProxyPatternRequest) String() string {
	return dara.Prettify(s)
}

func (s SetProxyPatternRequest) GoString() string {
	return s.String()
}

func (s *SetProxyPatternRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetProxyPatternRequest) GetLang() *string {
	return s.Lang
}

func (s *SetProxyPatternRequest) GetProxyPattern() *string {
	return s.ProxyPattern
}

func (s *SetProxyPatternRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SetProxyPatternRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *SetProxyPatternRequest) SetClientToken(v string) *SetProxyPatternRequest {
	s.ClientToken = &v
	return s
}

func (s *SetProxyPatternRequest) SetLang(v string) *SetProxyPatternRequest {
	s.Lang = &v
	return s
}

func (s *SetProxyPatternRequest) SetProxyPattern(v string) *SetProxyPatternRequest {
	s.ProxyPattern = &v
	return s
}

func (s *SetProxyPatternRequest) SetUserClientIp(v string) *SetProxyPatternRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetProxyPatternRequest) SetZoneId(v string) *SetProxyPatternRequest {
	s.ZoneId = &v
	return s
}

func (s *SetProxyPatternRequest) Validate() error {
	return dara.Validate(s)
}
