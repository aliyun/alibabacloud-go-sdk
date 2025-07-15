// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarUrl(v string) *CreateUserRequest
	GetAvatarUrl() *string
	SetDisplayId(v string) *CreateUserRequest
	GetDisplayId() *string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *CreateUserRequest
	GetEmail() *string
	SetInstanceId(v string) *CreateUserRequest
	GetInstanceId() *string
	SetLoginName(v string) *CreateUserRequest
	GetLoginName() *string
	SetMobile(v string) *CreateUserRequest
	GetMobile() *string
	SetNickname(v string) *CreateUserRequest
	GetNickname() *string
	SetResetPassword(v bool) *CreateUserRequest
	GetResetPassword() *bool
	SetRoleId(v string) *CreateUserRequest
	GetRoleId() *string
	SetSkillLevelList(v string) *CreateUserRequest
	GetSkillLevelList() *string
	SetWorkMode(v string) *CreateUserRequest
	GetWorkMode() *string
}

type CreateUserRequest struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// 1001
	DisplayId *string `json:"DisplayId,omitempty" xml:"DisplayId,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// example:
	//
	// 1382114****
	Mobile   *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// false
	ResetPassword *bool `json:"ResetPassword,omitempty" xml:"ResetPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Agent@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// [{"skillGroupId":"skillgroup1@ccc-test","skillLevel":1},{"skillGroupId":"skillgroup2@ccc-test","skillLevel":10}]
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateUserRequest) GetDisplayId() *string {
	return s.DisplayId
}

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserRequest) GetLoginName() *string {
	return s.LoginName
}

func (s *CreateUserRequest) GetMobile() *string {
	return s.Mobile
}

func (s *CreateUserRequest) GetNickname() *string {
	return s.Nickname
}

func (s *CreateUserRequest) GetResetPassword() *bool {
	return s.ResetPassword
}

func (s *CreateUserRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *CreateUserRequest) GetSkillLevelList() *string {
	return s.SkillLevelList
}

func (s *CreateUserRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *CreateUserRequest) SetAvatarUrl(v string) *CreateUserRequest {
	s.AvatarUrl = &v
	return s
}

func (s *CreateUserRequest) SetDisplayId(v string) *CreateUserRequest {
	s.DisplayId = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetInstanceId(v string) *CreateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserRequest) SetLoginName(v string) *CreateUserRequest {
	s.LoginName = &v
	return s
}

func (s *CreateUserRequest) SetMobile(v string) *CreateUserRequest {
	s.Mobile = &v
	return s
}

func (s *CreateUserRequest) SetNickname(v string) *CreateUserRequest {
	s.Nickname = &v
	return s
}

func (s *CreateUserRequest) SetResetPassword(v bool) *CreateUserRequest {
	s.ResetPassword = &v
	return s
}

func (s *CreateUserRequest) SetRoleId(v string) *CreateUserRequest {
	s.RoleId = &v
	return s
}

func (s *CreateUserRequest) SetSkillLevelList(v string) *CreateUserRequest {
	s.SkillLevelList = &v
	return s
}

func (s *CreateUserRequest) SetWorkMode(v string) *CreateUserRequest {
	s.WorkMode = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}
