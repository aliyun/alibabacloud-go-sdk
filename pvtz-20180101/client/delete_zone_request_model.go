// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteZoneRequest
	GetClientToken() *string
	SetLang(v string) *DeleteZoneRequest
	GetLang() *string
	SetUserClientIp(v string) *DeleteZoneRequest
	GetUserClientIp() *string
	SetZoneId(v string) *DeleteZoneRequest
	GetZoneId() *string
}

type DeleteZoneRequest struct {
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
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// >  If you want to delete a built-in authoritative zone whose effective scope is configured, you must disassociate the zone from the effective scope first.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0e41496f12da01311d314f17b801****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteZoneRequest) GoString() string {
	return s.String()
}

func (s *DeleteZoneRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteZoneRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteZoneRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteZoneRequest) SetClientToken(v string) *DeleteZoneRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteZoneRequest) SetLang(v string) *DeleteZoneRequest {
	s.Lang = &v
	return s
}

func (s *DeleteZoneRequest) SetUserClientIp(v string) *DeleteZoneRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteZoneRequest) SetZoneId(v string) *DeleteZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteZoneRequest) Validate() error {
	return dara.Validate(s)
}
