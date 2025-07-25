// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidatePdnsUdpIpSegmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *ValidatePdnsUdpIpSegmentRequest
	GetIp() *string
	SetIpToken(v string) *ValidatePdnsUdpIpSegmentRequest
	GetIpToken() *string
	SetLang(v string) *ValidatePdnsUdpIpSegmentRequest
	GetLang() *string
}

type ValidatePdnsUdpIpSegmentRequest struct {
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	IpToken *string `json:"IpToken,omitempty" xml:"IpToken,omitempty"`
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ValidatePdnsUdpIpSegmentRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidatePdnsUdpIpSegmentRequest) GoString() string {
	return s.String()
}

func (s *ValidatePdnsUdpIpSegmentRequest) GetIp() *string {
	return s.Ip
}

func (s *ValidatePdnsUdpIpSegmentRequest) GetIpToken() *string {
	return s.IpToken
}

func (s *ValidatePdnsUdpIpSegmentRequest) GetLang() *string {
	return s.Lang
}

func (s *ValidatePdnsUdpIpSegmentRequest) SetIp(v string) *ValidatePdnsUdpIpSegmentRequest {
	s.Ip = &v
	return s
}

func (s *ValidatePdnsUdpIpSegmentRequest) SetIpToken(v string) *ValidatePdnsUdpIpSegmentRequest {
	s.IpToken = &v
	return s
}

func (s *ValidatePdnsUdpIpSegmentRequest) SetLang(v string) *ValidatePdnsUdpIpSegmentRequest {
	s.Lang = &v
	return s
}

func (s *ValidatePdnsUdpIpSegmentRequest) Validate() error {
	return dara.Validate(s)
}
