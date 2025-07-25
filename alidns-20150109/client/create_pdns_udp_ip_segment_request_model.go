// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdnsUdpIpSegmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *CreatePdnsUdpIpSegmentRequest
	GetIp() *string
	SetIpToken(v string) *CreatePdnsUdpIpSegmentRequest
	GetIpToken() *string
	SetLang(v string) *CreatePdnsUdpIpSegmentRequest
	GetLang() *string
	SetName(v string) *CreatePdnsUdpIpSegmentRequest
	GetName() *string
}

type CreatePdnsUdpIpSegmentRequest struct {
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	IpToken *string `json:"IpToken,omitempty" xml:"IpToken,omitempty"`
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePdnsUdpIpSegmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePdnsUdpIpSegmentRequest) GoString() string {
	return s.String()
}

func (s *CreatePdnsUdpIpSegmentRequest) GetIp() *string {
	return s.Ip
}

func (s *CreatePdnsUdpIpSegmentRequest) GetIpToken() *string {
	return s.IpToken
}

func (s *CreatePdnsUdpIpSegmentRequest) GetLang() *string {
	return s.Lang
}

func (s *CreatePdnsUdpIpSegmentRequest) GetName() *string {
	return s.Name
}

func (s *CreatePdnsUdpIpSegmentRequest) SetIp(v string) *CreatePdnsUdpIpSegmentRequest {
	s.Ip = &v
	return s
}

func (s *CreatePdnsUdpIpSegmentRequest) SetIpToken(v string) *CreatePdnsUdpIpSegmentRequest {
	s.IpToken = &v
	return s
}

func (s *CreatePdnsUdpIpSegmentRequest) SetLang(v string) *CreatePdnsUdpIpSegmentRequest {
	s.Lang = &v
	return s
}

func (s *CreatePdnsUdpIpSegmentRequest) SetName(v string) *CreatePdnsUdpIpSegmentRequest {
	s.Name = &v
	return s
}

func (s *CreatePdnsUdpIpSegmentRequest) Validate() error {
	return dara.Validate(s)
}
