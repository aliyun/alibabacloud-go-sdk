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
	SetUid(v string) *RtcSipInviteMemberRequest
	GetUid() *string
}

type RtcSipInviteMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 000qaaasas……b
	AppToken *string `json:"AppToken,omitempty" xml:"AppToken,omitempty"`
	// example:
	//
	// 055112345678
	CallNumber *string `json:"CallNumber,omitempty" xml:"CallNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mcu
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Registered *bool `json:"Registered,omitempty" xml:"Registered,omitempty"`
	// example:
	//
	// 47.116.80.180
	ServerAddress *string `json:"ServerAddress,omitempty" xml:"ServerAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ceo
	SipDisplayName *string `json:"SipDisplayName,omitempty" xml:"SipDisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	SipRoomId *string `json:"SipRoomId,omitempty" xml:"SipRoomId,omitempty"`
	// example:
	//
	// 30.240.160.66:5060;transport=tcp
	SipUri *string `json:"SipUri,omitempty" xml:"SipUri,omitempty"`
	// example:
	//
	// pstn
	SipUserAgent *string `json:"SipUserAgent,omitempty" xml:"SipUserAgent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	SipUserId *string `json:"SipUserId,omitempty" xml:"SipUserId,omitempty"`
	// example:
	//
	// 123
	SipUserPassword *string `json:"SipUserPassword,omitempty" xml:"SipUserPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
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

func (s *RtcSipInviteMemberRequest) SetUid(v string) *RtcSipInviteMemberRequest {
	s.Uid = &v
	return s
}

func (s *RtcSipInviteMemberRequest) Validate() error {
	return dara.Validate(s)
}
