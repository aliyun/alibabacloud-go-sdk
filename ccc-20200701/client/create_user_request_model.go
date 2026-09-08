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
	SetNeedEmailNotification(v string) *CreateUserRequest
	GetNeedEmailNotification() *string
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
	// The URL of the agent\\"s profile picture.
	//
	// example:
	//
	// http://abc.com/sam.jpg
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// The agent\\"s ID number. Set this as needed.
	//
	// example:
	//
	// 1001
	DisplayId *string `json:"DisplayId,omitempty" xml:"DisplayId,omitempty"`
	// The display name of the agent. It must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 坐席小王
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the agent. After the agent is created, an email is sent to this address. The email contains the logon URL for Cloud Contact Center, and the username and password for the RAM account. Keep this information secure.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The logon name of the agent. It must be 4 to 64 characters in length and can contain uppercase letters, lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The personal phone number of the agent. This number is used in OFF_SITE mode. The agent can use this number to answer calls in OFF_SITE mode.
	//
	// example:
	//
	// 1382114****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Specifies whether to send an email notification.
	//
	// - true: Send
	//
	// - false: Do not send
	//
	// example:
	//
	// true
	NeedEmailNotification *string `json:"NeedEmailNotification,omitempty" xml:"NeedEmailNotification,omitempty"`
	// The agent\\"s nickname.
	//
	// example:
	//
	// 老王
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// Specifies whether the agent must reset the password upon the first logon. If set to true, the agent is prompted to reset the password when they first log on to the RAM account. Otherwise, they are not prompted. The default value is false.
	//
	// example:
	//
	// false
	ResetPassword *bool `json:"ResetPassword,omitempty" xml:"ResetPassword,omitempty"`
	// The role ID. The format is Role\\@InstanceID. The following roles are supported: Admin (administrator), Manager (skill group leader), and Agent (agent).
	//
	// This parameter is required.
	//
	// example:
	//
	// Agent@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// A list of skill levels for skill groups. This is a string in the format of a JSON array. The array can contain up to 100 elements. Each element is an object that contains two fields: skillGroupId and skillLevel. For skillGroupId, enter the ID of the skill group to add. For skillLevel, enter the skill level to add. The value can range from 1 to 10. A smaller value indicates a higher skill level, meaning the agent can handle more calls per unit of time.
	//
	// example:
	//
	// [{"skillGroupId":"skillgroup1@ccc-test","skillLevel":1},{"skillGroupId":"skillgroup2@ccc-test","skillLevel":10}]
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
	// The work mode.
	//
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

func (s *CreateUserRequest) GetNeedEmailNotification() *string {
	return s.NeedEmailNotification
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

func (s *CreateUserRequest) SetNeedEmailNotification(v string) *CreateUserRequest {
	s.NeedEmailNotification = &v
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
