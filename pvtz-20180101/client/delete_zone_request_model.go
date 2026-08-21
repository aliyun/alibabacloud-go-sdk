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
	// A client token to ensure the idempotence of the request. Generate a unique value from your client for this parameter. The token must contain only ASCII characters and be no more than 64 characters in length.
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
	// Default: en.
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
	// The unique ID of the zone.
	//
	// > You must dissociate a built-in authoritative domain name from its scope before you delete it.
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
