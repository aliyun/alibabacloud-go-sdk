// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAvatarSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelToken(v string) *StartAvatarSessionResponseBody
	GetChannelToken() *string
	SetRequestId(v string) *StartAvatarSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *StartAvatarSessionResponseBody
	GetSessionId() *string
	SetToken(v string) *StartAvatarSessionResponseBody
	GetToken() *string
	SetWebSocketUrl(v string) *StartAvatarSessionResponseBody
	GetWebSocketUrl() *string
}

type StartAvatarSessionResponseBody struct {
	ChannelToken *string `json:"channelToken,omitempty" xml:"channelToken,omitempty"`
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 121dlsga4o7golrl1hoja
	SessionId    *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
	WebSocketUrl *string `json:"webSocketUrl,omitempty" xml:"webSocketUrl,omitempty"`
}

func (s StartAvatarSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAvatarSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StartAvatarSessionResponseBody) GetChannelToken() *string {
	return s.ChannelToken
}

func (s *StartAvatarSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAvatarSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *StartAvatarSessionResponseBody) GetToken() *string {
	return s.Token
}

func (s *StartAvatarSessionResponseBody) GetWebSocketUrl() *string {
	return s.WebSocketUrl
}

func (s *StartAvatarSessionResponseBody) SetChannelToken(v string) *StartAvatarSessionResponseBody {
	s.ChannelToken = &v
	return s
}

func (s *StartAvatarSessionResponseBody) SetRequestId(v string) *StartAvatarSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAvatarSessionResponseBody) SetSessionId(v string) *StartAvatarSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *StartAvatarSessionResponseBody) SetToken(v string) *StartAvatarSessionResponseBody {
	s.Token = &v
	return s
}

func (s *StartAvatarSessionResponseBody) SetWebSocketUrl(v string) *StartAvatarSessionResponseBody {
	s.WebSocketUrl = &v
	return s
}

func (s *StartAvatarSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
