// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSendSipDtmfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RtcSendSipDtmfRequest
	GetAppId() *string
	SetChannelId(v string) *RtcSendSipDtmfRequest
	GetChannelId() *string
	SetChannelUserId(v string) *RtcSendSipDtmfRequest
	GetChannelUserId() *string
	SetDeviceType(v string) *RtcSendSipDtmfRequest
	GetDeviceType() *string
	SetDigits(v string) *RtcSendSipDtmfRequest
	GetDigits() *string
}

type RtcSendSipDtmfRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of an existing channel. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The user ID (sipUserId) in the target OneRTC SIP call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ChannelUserId *string `json:"ChannelUserId,omitempty" xml:"ChannelUserId,omitempty"`
	// The device type of the third-party Session Initiation Protocol (SIP) device. Valid values:
	//
	// - sip_device
	//
	// - mcu
	//
	// - phone
	//
	// This parameter is required.
	//
	// example:
	//
	// mcu
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The keys to be sent. Supported keys include 0-9, A-D, a-d, *, #, W or w (wait), and F or f (flash). The keys are sent in the original order and case.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12#W3f
	Digits *string `json:"Digits,omitempty" xml:"Digits,omitempty"`
}

func (s RtcSendSipDtmfRequest) String() string {
	return dara.Prettify(s)
}

func (s RtcSendSipDtmfRequest) GoString() string {
	return s.String()
}

func (s *RtcSendSipDtmfRequest) GetAppId() *string {
	return s.AppId
}

func (s *RtcSendSipDtmfRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *RtcSendSipDtmfRequest) GetChannelUserId() *string {
	return s.ChannelUserId
}

func (s *RtcSendSipDtmfRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *RtcSendSipDtmfRequest) GetDigits() *string {
	return s.Digits
}

func (s *RtcSendSipDtmfRequest) SetAppId(v string) *RtcSendSipDtmfRequest {
	s.AppId = &v
	return s
}

func (s *RtcSendSipDtmfRequest) SetChannelId(v string) *RtcSendSipDtmfRequest {
	s.ChannelId = &v
	return s
}

func (s *RtcSendSipDtmfRequest) SetChannelUserId(v string) *RtcSendSipDtmfRequest {
	s.ChannelUserId = &v
	return s
}

func (s *RtcSendSipDtmfRequest) SetDeviceType(v string) *RtcSendSipDtmfRequest {
	s.DeviceType = &v
	return s
}

func (s *RtcSendSipDtmfRequest) SetDigits(v string) *RtcSendSipDtmfRequest {
	s.Digits = &v
	return s
}

func (s *RtcSendSipDtmfRequest) Validate() error {
	return dara.Validate(s)
}
