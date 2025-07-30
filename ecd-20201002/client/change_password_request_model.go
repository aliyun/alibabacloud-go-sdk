// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangePasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *ChangePasswordRequest
	GetClientId() *string
	SetEndUserId(v string) *ChangePasswordRequest
	GetEndUserId() *string
	SetLoginToken(v string) *ChangePasswordRequest
	GetLoginToken() *string
	SetNewPassword(v string) *ChangePasswordRequest
	GetNewPassword() *string
	SetOfficeSiteId(v string) *ChangePasswordRequest
	GetOfficeSiteId() *string
	SetOldPassword(v string) *ChangePasswordRequest
	GetOldPassword() *string
	SetRegionId(v string) *ChangePasswordRequest
	GetRegionId() *string
	SetSessionId(v string) *ChangePasswordRequest
	GetSessionId() *string
}

type ChangePasswordRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 42f6645a-9c3c-4772-be2a-cc5f5732****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// liming
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The new password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 67436290
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+dir-227468****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The current password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
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
	// example:
	//
	// 1
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ChangePasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangePasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangePasswordRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ChangePasswordRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ChangePasswordRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ChangePasswordRequest) GetNewPassword() *string {
	return s.NewPassword
}

func (s *ChangePasswordRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ChangePasswordRequest) GetOldPassword() *string {
	return s.OldPassword
}

func (s *ChangePasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangePasswordRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChangePasswordRequest) SetClientId(v string) *ChangePasswordRequest {
	s.ClientId = &v
	return s
}

func (s *ChangePasswordRequest) SetEndUserId(v string) *ChangePasswordRequest {
	s.EndUserId = &v
	return s
}

func (s *ChangePasswordRequest) SetLoginToken(v string) *ChangePasswordRequest {
	s.LoginToken = &v
	return s
}

func (s *ChangePasswordRequest) SetNewPassword(v string) *ChangePasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetOfficeSiteId(v string) *ChangePasswordRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ChangePasswordRequest) SetOldPassword(v string) *ChangePasswordRequest {
	s.OldPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetRegionId(v string) *ChangePasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ChangePasswordRequest) SetSessionId(v string) *ChangePasswordRequest {
	s.SessionId = &v
	return s
}

func (s *ChangePasswordRequest) Validate() error {
	return dara.Validate(s)
}
