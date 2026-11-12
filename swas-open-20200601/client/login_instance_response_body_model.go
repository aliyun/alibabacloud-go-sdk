// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProtocol(v string) *LoginInstanceResponseBody
	GetProtocol() *string
	SetRedirectUrl(v string) *LoginInstanceResponseBody
	GetRedirectUrl() *string
	SetRequestId(v string) *LoginInstanceResponseBody
	GetRequestId() *string
	SetSecurityToken(v string) *LoginInstanceResponseBody
	GetSecurityToken() *string
	SetSessionId(v string) *LoginInstanceResponseBody
	GetSessionId() *string
}

type LoginInstanceResponseBody struct {
	// The session protocol.
	//
	// example:
	//
	// SSH
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The URL of the remote logon desktop.
	//
	// example:
	//
	// https://ecs-workbench.aliyun.com/view/instance/single/gbktfz****
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C2DE174B-7196-5778-A00D-6EA2601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security token.
	//
	// example:
	//
	// ik1mb2g69ksagyybrd13be7h3655yuhb
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The session ID.
	//
	// example:
	//
	// s-0rwwbwxymfzobrcz8
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s LoginInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *LoginInstanceResponseBody) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *LoginInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoginInstanceResponseBody) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *LoginInstanceResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *LoginInstanceResponseBody) SetProtocol(v string) *LoginInstanceResponseBody {
	s.Protocol = &v
	return s
}

func (s *LoginInstanceResponseBody) SetRedirectUrl(v string) *LoginInstanceResponseBody {
	s.RedirectUrl = &v
	return s
}

func (s *LoginInstanceResponseBody) SetRequestId(v string) *LoginInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginInstanceResponseBody) SetSecurityToken(v string) *LoginInstanceResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *LoginInstanceResponseBody) SetSessionId(v string) *LoginInstanceResponseBody {
	s.SessionId = &v
	return s
}

func (s *LoginInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
