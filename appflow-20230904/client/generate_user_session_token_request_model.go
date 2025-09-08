// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUserSessionTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatbotId(v string) *GenerateUserSessionTokenRequest
	GetChatbotId() *string
	SetExpireSecond(v int64) *GenerateUserSessionTokenRequest
	GetExpireSecond() *int64
	SetExtraInfo(v string) *GenerateUserSessionTokenRequest
	GetExtraInfo() *string
	SetIntegrateId(v string) *GenerateUserSessionTokenRequest
	GetIntegrateId() *string
	SetUserAvatar(v string) *GenerateUserSessionTokenRequest
	GetUserAvatar() *string
	SetUserId(v string) *GenerateUserSessionTokenRequest
	GetUserId() *string
	SetUserName(v string) *GenerateUserSessionTokenRequest
	GetUserName() *string
}

type GenerateUserSessionTokenRequest struct {
	// AI Assistant ID
	//
	// example:
	//
	// cb-069d508f9ab341b1****
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Expiration Time, in seconds, default 24 hours
	//
	// example:
	//
	// 6000
	ExpireSecond *int64  `json:"ExpireSecond,omitempty" xml:"ExpireSecond,omitempty"`
	ExtraInfo    *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// Integration ID
	//
	// example:
	//
	// cit-960f499au184m7****
	IntegrateId *string `json:"IntegrateId,omitempty" xml:"IntegrateId,omitempty"`
	// User Avatar (URL)
	//
	// example:
	//
	// https://xxxx.com/xxx
	UserAvatar *string `json:"UserAvatar,omitempty" xml:"UserAvatar,omitempty"`
	// User ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 929293312213****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// User Nickname
	//
	// example:
	//
	// testxxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GenerateUserSessionTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateUserSessionTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateUserSessionTokenRequest) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *GenerateUserSessionTokenRequest) GetExpireSecond() *int64 {
	return s.ExpireSecond
}

func (s *GenerateUserSessionTokenRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GenerateUserSessionTokenRequest) GetIntegrateId() *string {
	return s.IntegrateId
}

func (s *GenerateUserSessionTokenRequest) GetUserAvatar() *string {
	return s.UserAvatar
}

func (s *GenerateUserSessionTokenRequest) GetUserId() *string {
	return s.UserId
}

func (s *GenerateUserSessionTokenRequest) GetUserName() *string {
	return s.UserName
}

func (s *GenerateUserSessionTokenRequest) SetChatbotId(v string) *GenerateUserSessionTokenRequest {
	s.ChatbotId = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetExpireSecond(v int64) *GenerateUserSessionTokenRequest {
	s.ExpireSecond = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetExtraInfo(v string) *GenerateUserSessionTokenRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetIntegrateId(v string) *GenerateUserSessionTokenRequest {
	s.IntegrateId = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetUserAvatar(v string) *GenerateUserSessionTokenRequest {
	s.UserAvatar = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetUserId(v string) *GenerateUserSessionTokenRequest {
	s.UserId = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) SetUserName(v string) *GenerateUserSessionTokenRequest {
	s.UserName = &v
	return s
}

func (s *GenerateUserSessionTokenRequest) Validate() error {
	return dara.Validate(s)
}
