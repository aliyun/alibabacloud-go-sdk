// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcCancelSipInviteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RtcCancelSipInviteRequest
	GetAppId() *string
	SetChannelId(v string) *RtcCancelSipInviteRequest
	GetChannelId() *string
	SetDeviceType(v string) *RtcCancelSipInviteRequest
	GetDeviceType() *string
	SetUserId(v string) *RtcCancelSipInviteRequest
	GetUserId() *string
}

type RtcCancelSipInviteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	// 123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RtcCancelSipInviteRequest) String() string {
	return dara.Prettify(s)
}

func (s RtcCancelSipInviteRequest) GoString() string {
	return s.String()
}

func (s *RtcCancelSipInviteRequest) GetAppId() *string {
	return s.AppId
}

func (s *RtcCancelSipInviteRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *RtcCancelSipInviteRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *RtcCancelSipInviteRequest) GetUserId() *string {
	return s.UserId
}

func (s *RtcCancelSipInviteRequest) SetAppId(v string) *RtcCancelSipInviteRequest {
	s.AppId = &v
	return s
}

func (s *RtcCancelSipInviteRequest) SetChannelId(v string) *RtcCancelSipInviteRequest {
	s.ChannelId = &v
	return s
}

func (s *RtcCancelSipInviteRequest) SetDeviceType(v string) *RtcCancelSipInviteRequest {
	s.DeviceType = &v
	return s
}

func (s *RtcCancelSipInviteRequest) SetUserId(v string) *RtcCancelSipInviteRequest {
	s.UserId = &v
	return s
}

func (s *RtcCancelSipInviteRequest) Validate() error {
	return dara.Validate(s)
}
