// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePdnsUdpIpSegmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *RemovePdnsUdpIpSegmentRequest
	GetIp() *string
	SetLang(v string) *RemovePdnsUdpIpSegmentRequest
	GetLang() *string
}

type RemovePdnsUdpIpSegmentRequest struct {
	Ip   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s RemovePdnsUdpIpSegmentRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePdnsUdpIpSegmentRequest) GoString() string {
	return s.String()
}

func (s *RemovePdnsUdpIpSegmentRequest) GetIp() *string {
	return s.Ip
}

func (s *RemovePdnsUdpIpSegmentRequest) GetLang() *string {
	return s.Lang
}

func (s *RemovePdnsUdpIpSegmentRequest) SetIp(v string) *RemovePdnsUdpIpSegmentRequest {
	s.Ip = &v
	return s
}

func (s *RemovePdnsUdpIpSegmentRequest) SetLang(v string) *RemovePdnsUdpIpSegmentRequest {
	s.Lang = &v
	return s
}

func (s *RemovePdnsUdpIpSegmentRequest) Validate() error {
	return dara.Validate(s)
}
