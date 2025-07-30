// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindUserDesktopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *UnbindUserDesktopRequest
	GetClientId() *string
	SetClientType(v string) *UnbindUserDesktopRequest
	GetClientType() *string
	SetForce(v bool) *UnbindUserDesktopRequest
	GetForce() *bool
	SetLoginToken(v string) *UnbindUserDesktopRequest
	GetLoginToken() *string
	SetRegionId(v string) *UnbindUserDesktopRequest
	GetRegionId() *string
	SetSessionId(v string) *UnbindUserDesktopRequest
	GetSessionId() *string
	SetUserDesktopId(v string) *UnbindUserDesktopRequest
	GetUserDesktopId() *string
}

type UnbindUserDesktopRequest struct {
	// The client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 58f96f67-7944-4f97-9342-****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client type.
	//
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// Specifies whether to enable forced unbinding.
	//
	// Valid values:
	//
	// 	- true: Even when end users connect to cloud computers, the forced unbinding is still enforced.
	//
	// 	- false: Forced unbinding is only enforced when end users are disconnected from cloud computers.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v12307f5e0****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b053331-dc98-43d8-b247-****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The cloud computer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ud-sdfs****
	UserDesktopId *string `json:"UserDesktopId,omitempty" xml:"UserDesktopId,omitempty"`
}

func (s UnbindUserDesktopRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindUserDesktopRequest) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UnbindUserDesktopRequest) GetClientType() *string {
	return s.ClientType
}

func (s *UnbindUserDesktopRequest) GetForce() *bool {
	return s.Force
}

func (s *UnbindUserDesktopRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *UnbindUserDesktopRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnbindUserDesktopRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *UnbindUserDesktopRequest) GetUserDesktopId() *string {
	return s.UserDesktopId
}

func (s *UnbindUserDesktopRequest) SetClientId(v string) *UnbindUserDesktopRequest {
	s.ClientId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetClientType(v string) *UnbindUserDesktopRequest {
	s.ClientType = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetForce(v bool) *UnbindUserDesktopRequest {
	s.Force = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetLoginToken(v string) *UnbindUserDesktopRequest {
	s.LoginToken = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetRegionId(v string) *UnbindUserDesktopRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetSessionId(v string) *UnbindUserDesktopRequest {
	s.SessionId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetUserDesktopId(v string) *UnbindUserDesktopRequest {
	s.UserDesktopId = &v
	return s
}

func (s *UnbindUserDesktopRequest) Validate() error {
	return dara.Validate(s)
}
