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
	// The ID of the ApsaraVideo Real-time Communication application. Only a single ID is supported. The value can contain uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters. You can view your application IDs by navigating to **ApsaraVideo Live > Live+ > Real-time Communication > Application Management**. If no application exists, create one by clicking **Create Application**.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel to join. Only a single ID is supported. The value can contain uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The display name of the RTMP stream in the channel. Maximum length: 40 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp-dname
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The validity period of the RTMP URL. Unit: seconds. Default value: 36000 (10 hours).
	//
	// example:
	//
	// 43200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The user ID for the RTMP stream ingest. This value must not duplicate any other user ID in the channel. The value can contain uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters.
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
