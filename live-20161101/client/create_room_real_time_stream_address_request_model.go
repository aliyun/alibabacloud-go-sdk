// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoomRealTimeStreamAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateRoomRealTimeStreamAddressRequest
	GetAppId() *string
	SetChannelId(v string) *CreateRoomRealTimeStreamAddressRequest
	GetChannelId() *string
	SetDisplayName(v string) *CreateRoomRealTimeStreamAddressRequest
	GetDisplayName() *string
	SetExpireTime(v int32) *CreateRoomRealTimeStreamAddressRequest
	GetExpireTime() *int32
	SetUserId(v string) *CreateRoomRealTimeStreamAddressRequest
	GetUserId() *string
}

type CreateRoomRealTimeStreamAddressRequest struct {
	// The ID of the ARTC application. You can specify only one application ID. The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel. You can specify only one ID. The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The name of the RTMP stream. The name can be up to 40 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp-dname
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The validity period of the RTMP URL. Unit: seconds. The default value is 36,000 seconds, which is 10 hours.
	//
	// example:
	//
	// 43200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the user who ingests the stream over RTMP. The user ID must be different from IDs of other users in the channel. The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp-uuid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateRoomRealTimeStreamAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoomRealTimeStreamAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomRealTimeStreamAddressRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateRoomRealTimeStreamAddressRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateRoomRealTimeStreamAddressRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateRoomRealTimeStreamAddressRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *CreateRoomRealTimeStreamAddressRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateRoomRealTimeStreamAddressRequest) SetAppId(v string) *CreateRoomRealTimeStreamAddressRequest {
	s.AppId = &v
	return s
}

func (s *CreateRoomRealTimeStreamAddressRequest) SetChannelId(v string) *CreateRoomRealTimeStreamAddressRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateRoomRealTimeStreamAddressRequest) SetDisplayName(v string) *CreateRoomRealTimeStreamAddressRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateRoomRealTimeStreamAddressRequest) SetExpireTime(v int32) *CreateRoomRealTimeStreamAddressRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateRoomRealTimeStreamAddressRequest) SetUserId(v string) *CreateRoomRealTimeStreamAddressRequest {
	s.UserId = &v
	return s
}

func (s *CreateRoomRealTimeStreamAddressRequest) Validate() error {
	return dara.Validate(s)
}
