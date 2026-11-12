// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTerminalSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartTerminalSessionResponseBody
	GetRequestId() *string
	SetSecurityToken(v string) *StartTerminalSessionResponseBody
	GetSecurityToken() *string
	SetSessionId(v string) *StartTerminalSessionResponseBody
	GetSessionId() *string
	SetWebSocketUrl(v string) *StartTerminalSessionResponseBody
	GetWebSocketUrl() *string
}

type StartTerminalSessionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security token that is included in the request header of the WebSocket request for authentication.
	//
	// example:
	//
	// token-xxxaaz
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The session ID.
	//
	// example:
	//
	// ffb90b6e-b18a-4a33-88cf-86fb88****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The URL of the WebSocket session that is used to remotely connect to the Simple Application Server instance. The URL contains the session ID (`SessionId`) and the security token (`SecurityToken`) that is used for authentication.
	//
	// example:
	//
	// wss://xxxx
	WebSocketUrl *string `json:"WebSocketUrl,omitempty" xml:"WebSocketUrl,omitempty"`
}

func (s StartTerminalSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTerminalSessionResponseBody) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StartTerminalSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *StartTerminalSessionResponseBody) GetWebSocketUrl() *string {
	return s.WebSocketUrl
}

func (s *StartTerminalSessionResponseBody) SetRequestId(v string) *StartTerminalSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetSecurityToken(v string) *StartTerminalSessionResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetSessionId(v string) *StartTerminalSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetWebSocketUrl(v string) *StartTerminalSessionResponseBody {
	s.WebSocketUrl = &v
	return s
}

func (s *StartTerminalSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
