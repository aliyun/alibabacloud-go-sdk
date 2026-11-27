// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSipInviteMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RtcSipInviteMemberRequest
	GetAppId() *string
	SetAppToken(v string) *RtcSipInviteMemberRequest
	GetAppToken() *string
	SetCallNumber(v string) *RtcSipInviteMemberRequest
	GetCallNumber() *string
	SetChannelId(v string) *RtcSipInviteMemberRequest
	GetChannelId() *string
	SetDeviceType(v string) *RtcSipInviteMemberRequest
	GetDeviceType() *string
	SetRegistered(v bool) *RtcSipInviteMemberRequest
	GetRegistered() *bool
	SetServerAddress(v string) *RtcSipInviteMemberRequest
	GetServerAddress() *string
	SetSipDisplayName(v string) *RtcSipInviteMemberRequest
	GetSipDisplayName() *string
	SetSipRoomId(v string) *RtcSipInviteMemberRequest
	GetSipRoomId() *string
	SetSipUri(v string) *RtcSipInviteMemberRequest
	GetSipUri() *string
	SetSipUserAgent(v string) *RtcSipInviteMemberRequest
	GetSipUserAgent() *string
	SetSipUserId(v string) *RtcSipInviteMemberRequest
	GetSipUserId() *string
	SetSipUserPassword(v string) *RtcSipInviteMemberRequest
	GetSipUserPassword() *string
	SetTaskId(v string) *RtcSipInviteMemberRequest
	GetTaskId() *string
	SetUid(v string) *RtcSipInviteMemberRequest
	GetUid() *string
}

type RtcSipInviteMemberRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The meeting token. For code generation, refer to https://help.aliyun.com/document_detail/2689025.html.
	//
	// example:
	//
	// 000qaaasas……b
	AppToken *string `json:"AppToken,omitempty" xml:"AppToken,omitempty"`
	// The custom line number in PSTN scenarios.
	//
	// example:
	//
	// 055112345678
	CallNumber *string `json:"CallNumber,omitempty" xml:"CallNumber,omitempty"`
	// The channel ID. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// $rtc_channelid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The type of the device used to call the third-party SIP endpoint. Valid values:
	//
	// - sip_device
	//
	// - mcu
	//
	// - phone
	//
	// Note: Currently, only the phone type is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// phone
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// Specifies whether the third-party meeting system requires registration for inbound calls.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Registered *bool `json:"Registered,omitempty" xml:"Registered,omitempty"`
	// 1. The IP address of the nearest RTC SIP server.
	//
	// 2. You can resolve the domain name rtcsip.dingtalk.com locally to obtain the nearest access address.
	//
	// For example, run dig rtcsip.dingtalk.com to obtain the address.
	//
	// example:
	//
	// 39.101.232.13
	ServerAddress *string `json:"ServerAddress,omitempty" xml:"ServerAddress,omitempty"`
	// The username displayed on the third-party system.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdevice
	SipDisplayName *string `json:"SipDisplayName,omitempty" xml:"SipDisplayName,omitempty"`
	// 1. When the device is mcu, it indicates the channel number of the third-party meeting system to be called.
	//
	// 2. When the device is sip_device, it indicates the registration number of the registered device to be invited.
	//
	// 3. When the device is phone, it indicates the called phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18512514528
	SipRoomId *string `json:"SipRoomId,omitempty" xml:"SipRoomId,omitempty"`
	// 1. When the device is mcu, it indicates the address information of the third-party meeting system in the format of ip:port;transport=tcp/udp/tls.
	//
	// Example: "30.240.160.66:5060;transport=tcp"
	//
	// 2. When the device is sip_device, it indicates the registration address of the registered device to be invited.
	//
	// Example: 39.101.232.13
	//
	// 3. When the device is phone, it indicates the address information of the telephone gateway.
	//
	// Example: "30.240.160.66:5060;transport=tcp"
	//
	// example:
	//
	// 30.240.160.66:5060;transport=tcp
	SipUri *string `json:"SipUri,omitempty" xml:"SipUri,omitempty"`
	// The brand and model of the third-party SIP endpoint being called, such as polycom, cisco, tencent, zoom, or pstn.
	//
	// example:
	//
	// pstn
	SipUserAgent *string `json:"SipUserAgent,omitempty" xml:"SipUserAgent,omitempty"`
	// 1. When DeviceType is mcu, it represents the third-party meeting system.
	//
	// 2. When DeviceType is sip_device, it represents the registered device.
	//
	// 3. If the third-party meeting system requires registration for inbound calls, it is the account provided by the third-party system to create an account.
	//
	// 4. If registration is not required, it is the SIP username or RTC user ID for calling the third-party system.
	//
	// This parameter is required.
	//
	// example:
	//
	// $reg_userid
	SipUserId *string `json:"SipUserId,omitempty" xml:"SipUserId,omitempty"`
	// The registration password that matches the SipUserId if the third-party system requires registration for inbound calls.
	//
	// example:
	//
	// 654321
	SipUserPassword *string `json:"SipUserPassword,omitempty" xml:"SipUserPassword,omitempty"`
	// The task ID.
	//
	// example:
	//
	// taskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// $rtc_userid
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s RtcSipInviteMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s RtcSipInviteMemberRequest) GoString() string {
	return s.String()
}

func (s *RtcSipInviteMemberRequest) GetAppId() *string {
	return s.AppId
}

func (s *RtcSipInviteMemberRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RtcSipInviteMemberRequest) GetCallNumber() *string {
	return s.CallNumber
}

func (s *RtcSipInviteMemberRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *RtcSipInviteMemberRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *RtcSipInviteMemberRequest) GetRegistered() *bool {
	return s.Registered
}

func (s *RtcSipInviteMemberRequest) GetServerAddress() *string {
	return s.ServerAddress
}

func (s *RtcSipInviteMemberRequest) GetSipDisplayName() *string {
	return s.SipDisplayName
}

func (s *RtcSipInviteMemberRequest) GetSipRoomId() *string {
	return s.SipRoomId
}

func (s *RtcSipInviteMemberRequest) GetSipUri() *string {
	return s.SipUri
}

func (s *RtcSipInviteMemberRequest) GetSipUserAgent() *string {
	return s.SipUserAgent
}

func (s *RtcSipInviteMemberRequest) GetSipUserId() *string {
	return s.SipUserId
}

func (s *RtcSipInviteMemberRequest) GetSipUserPassword() *string {
	return s.SipUserPassword
}

func (s *RtcSipInviteMemberRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RtcSipInviteMemberRequest) GetUid() *string {
	return s.Uid
}

func (s *RtcSipInviteMemberRequest) SetAppId(v string) *RtcSipInviteMemberRequest {
	s.AppId = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetAppToken(v string) *RtcSipInviteMemberRequest {
	s.AppToken = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetCallNumber(v string) *RtcSipInviteMemberRequest {
	s.CallNumber = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetChannelId(v string) *RtcSipInviteMemberRequest {
	s.ChannelId = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetDeviceType(v string) *RtcSipInviteMemberRequest {
	s.DeviceType = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetRegistered(v bool) *RtcSipInviteMemberRequest {
	s.Registered = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetServerAddress(v string) *RtcSipInviteMemberRequest {
	s.ServerAddress = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetSipDisplayName(v string) *RtcSipInviteMemberRequest {
	s.SipDisplayName = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetSipRoomId(v string) *RtcSipInviteMemberRequest {
	s.SipRoomId = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetSipUri(v string) *RtcSipInviteMemberRequest {
	s.SipUri = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetSipUserAgent(v string) *RtcSipInviteMemberRequest {
	s.SipUserAgent = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetSipUserId(v string) *RtcSipInviteMemberRequest {
	s.SipUserId = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetSipUserPassword(v string) *RtcSipInviteMemberRequest {
	s.SipUserPassword = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetTaskId(v string) *RtcSipInviteMemberRequest {
	s.TaskId = &v
	return s
}

func (s *RtcSipInviteMemberRequest) SetUid(v string) *RtcSipInviteMemberRequest {
	s.Uid = &v
	return s
}

func (s *RtcSipInviteMemberRequest) Validate() error {
	return dara.Validate(s)
}
