// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckZoneNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckZoneNameRequest
	GetLang() *string
	SetUserClientIp(v string) *CheckZoneNameRequest
	GetUserClientIp() *string
	SetZoneName(v string) *CheckZoneNameRequest
	GetZoneName() *string
}

type CheckZoneNameRequest struct {
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
	// The name of the zone. This parameter is required.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s CheckZoneNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckZoneNameRequest) GoString() string {
	return s.String()
}

func (s *CheckZoneNameRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckZoneNameRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CheckZoneNameRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *CheckZoneNameRequest) SetLang(v string) *CheckZoneNameRequest {
	s.Lang = &v
	return s
}

func (s *CheckZoneNameRequest) SetUserClientIp(v string) *CheckZoneNameRequest {
	s.UserClientIp = &v
	return s
}

func (s *CheckZoneNameRequest) SetZoneName(v string) *CheckZoneNameRequest {
	s.ZoneName = &v
	return s
}

func (s *CheckZoneNameRequest) Validate() error {
	return dara.Validate(s)
}
