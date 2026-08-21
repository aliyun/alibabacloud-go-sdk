// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetZoneLbaStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SetZoneLbaStatusRequest
	GetClientToken() *string
	SetLang(v string) *SetZoneLbaStatusRequest
	GetLang() *string
	SetLine(v string) *SetZoneLbaStatusRequest
	GetLine() *string
	SetOpen(v bool) *SetZoneLbaStatusRequest
	GetOpen() *bool
	SetRr(v string) *SetZoneLbaStatusRequest
	GetRr() *string
	SetType(v string) *SetZoneLbaStatusRequest
	GetType() *string
	SetUserClientIp(v string) *SetZoneLbaStatusRequest
	GetUserClientIp() *string
	SetZoneId(v string) *SetZoneLbaStatusRequest
	GetZoneId() *string
}

type SetZoneLbaStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 210bc45716943908285687176dcf0a
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The resolution line.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// Specifies whether to enable the weight configuration. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Open *bool `json:"Open,omitempty" xml:"Open,omitempty"`
	// The hostname record.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The DNS record type. Currently, only **A*	- and **AAAA*	- record types support the weight toggle.
	//
	// This parameter is required.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IP address of the user.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// Zone ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 8fd507b3eec6bba982060561f5624ea6
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SetZoneLbaStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetZoneLbaStatusRequest) GoString() string {
	return s.String()
}

func (s *SetZoneLbaStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetZoneLbaStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *SetZoneLbaStatusRequest) GetLine() *string {
	return s.Line
}

func (s *SetZoneLbaStatusRequest) GetOpen() *bool {
	return s.Open
}

func (s *SetZoneLbaStatusRequest) GetRr() *string {
	return s.Rr
}

func (s *SetZoneLbaStatusRequest) GetType() *string {
	return s.Type
}

func (s *SetZoneLbaStatusRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SetZoneLbaStatusRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *SetZoneLbaStatusRequest) SetClientToken(v string) *SetZoneLbaStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *SetZoneLbaStatusRequest) SetLang(v string) *SetZoneLbaStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetZoneLbaStatusRequest) SetLine(v string) *SetZoneLbaStatusRequest {
	s.Line = &v
	return s
}

func (s *SetZoneLbaStatusRequest) SetOpen(v bool) *SetZoneLbaStatusRequest {
	s.Open = &v
	return s
}

func (s *SetZoneLbaStatusRequest) SetRr(v string) *SetZoneLbaStatusRequest {
	s.Rr = &v
	return s
}

func (s *SetZoneLbaStatusRequest) SetType(v string) *SetZoneLbaStatusRequest {
	s.Type = &v
	return s
}

func (s *SetZoneLbaStatusRequest) SetUserClientIp(v string) *SetZoneLbaStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetZoneLbaStatusRequest) SetZoneId(v string) *SetZoneLbaStatusRequest {
	s.ZoneId = &v
	return s
}

func (s *SetZoneLbaStatusRequest) Validate() error {
	return dara.Validate(s)
}
