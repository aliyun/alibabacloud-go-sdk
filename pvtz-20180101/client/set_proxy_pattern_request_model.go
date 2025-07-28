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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to enable the recursive resolution proxy for subdomain names. Valid values:
	//
	// 	- **ZONE**: disables the recursive resolution proxy for subdomain names. In this case, NXDOMAIN is returned if the queried subdomain name does not exist in the zone.
	//
	// 	- **RECORD**: enables the recursive resolution proxy for subdomain names. In this case, if the queried domain name does not exist in the zone, Domain Name System (DNS) requests are recursively forwarded to the forward module and then to the recursion module until DNS results are returned.
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
	// The zone ID. This ID uniquely identifies the zone.
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
