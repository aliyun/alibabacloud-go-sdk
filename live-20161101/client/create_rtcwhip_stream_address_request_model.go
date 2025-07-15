// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRTCWhipStreamAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateRTCWhipStreamAddressRequest
	GetAppId() *string
	SetChannelId(v string) *CreateRTCWhipStreamAddressRequest
	GetChannelId() *string
	SetClientToken(v string) *CreateRTCWhipStreamAddressRequest
	GetClientToken() *string
	SetDisplayName(v string) *CreateRTCWhipStreamAddressRequest
	GetDisplayName() *string
	SetExpireTime(v int32) *CreateRTCWhipStreamAddressRequest
	GetExpireTime() *int32
	SetUserId(v string) *CreateRTCWhipStreamAddressRequest
	GetUserId() *string
}

type CreateRTCWhipStreamAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6a0**6dcc-xxxx-xxxx-xxxx-e**e3exxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ch00000****001
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 58E73333-xxxx-xxxx-xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 43200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zb0000****0001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateRTCWhipStreamAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRTCWhipStreamAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateRTCWhipStreamAddressRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateRTCWhipStreamAddressRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateRTCWhipStreamAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRTCWhipStreamAddressRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateRTCWhipStreamAddressRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *CreateRTCWhipStreamAddressRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateRTCWhipStreamAddressRequest) SetAppId(v string) *CreateRTCWhipStreamAddressRequest {
	s.AppId = &v
	return s
}

func (s *CreateRTCWhipStreamAddressRequest) SetChannelId(v string) *CreateRTCWhipStreamAddressRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateRTCWhipStreamAddressRequest) SetClientToken(v string) *CreateRTCWhipStreamAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRTCWhipStreamAddressRequest) SetDisplayName(v string) *CreateRTCWhipStreamAddressRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateRTCWhipStreamAddressRequest) SetExpireTime(v int32) *CreateRTCWhipStreamAddressRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateRTCWhipStreamAddressRequest) SetUserId(v string) *CreateRTCWhipStreamAddressRequest {
	s.UserId = &v
	return s
}

func (s *CreateRTCWhipStreamAddressRequest) Validate() error {
	return dara.Validate(s)
}
