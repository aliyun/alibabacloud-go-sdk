// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateZoneRemarkRequest
	GetClientToken() *string
	SetLang(v string) *UpdateZoneRemarkRequest
	GetLang() *string
	SetRemark(v string) *UpdateZoneRemarkRequest
	GetRemark() *string
	SetUserClientIp(v string) *UpdateZoneRemarkRequest
	GetUserClientIp() *string
	SetZoneId(v string) *UpdateZoneRemarkRequest
	GetZoneId() *string
}

type UpdateZoneRemarkRequest struct {
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
	// The new description. If you leave Remark empty, the zone has no description.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.XX.XX
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

func (s UpdateZoneRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateZoneRemarkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateZoneRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateZoneRemarkRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *UpdateZoneRemarkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateZoneRemarkRequest) SetClientToken(v string) *UpdateZoneRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetLang(v string) *UpdateZoneRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetRemark(v string) *UpdateZoneRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetUserClientIp(v string) *UpdateZoneRemarkRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetZoneId(v string) *UpdateZoneRemarkRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateZoneRemarkRequest) Validate() error {
	return dara.Validate(s)
}
