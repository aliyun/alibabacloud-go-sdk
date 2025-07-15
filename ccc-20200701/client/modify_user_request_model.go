// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarUrl(v string) *ModifyUserRequest
	GetAvatarUrl() *string
	SetDisplayId(v string) *ModifyUserRequest
	GetDisplayId() *string
	SetDisplayName(v string) *ModifyUserRequest
	GetDisplayName() *string
	SetForce(v bool) *ModifyUserRequest
	GetForce() *bool
	SetInstanceId(v string) *ModifyUserRequest
	GetInstanceId() *string
	SetMobile(v string) *ModifyUserRequest
	GetMobile() *string
	SetNickname(v string) *ModifyUserRequest
	GetNickname() *string
	SetRoleId(v string) *ModifyUserRequest
	GetRoleId() *string
	SetUserId(v string) *ModifyUserRequest
	GetUserId() *string
	SetWorkMode(v string) *ModifyUserRequest
	GetWorkMode() *string
}

type ModifyUserRequest struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// 1001
	DisplayId   *string `json:"DisplayId,omitempty" xml:"DisplayId,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Force       *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1382114****
	Mobile   *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-test@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ModifyUserRequest) GetDisplayId() *string {
	return s.DisplayId
}

func (s *ModifyUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ModifyUserRequest) GetForce() *bool {
	return s.Force
}

func (s *ModifyUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserRequest) GetMobile() *string {
	return s.Mobile
}

func (s *ModifyUserRequest) GetNickname() *string {
	return s.Nickname
}

func (s *ModifyUserRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *ModifyUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ModifyUserRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ModifyUserRequest) SetAvatarUrl(v string) *ModifyUserRequest {
	s.AvatarUrl = &v
	return s
}

func (s *ModifyUserRequest) SetDisplayId(v string) *ModifyUserRequest {
	s.DisplayId = &v
	return s
}

func (s *ModifyUserRequest) SetDisplayName(v string) *ModifyUserRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifyUserRequest) SetForce(v bool) *ModifyUserRequest {
	s.Force = &v
	return s
}

func (s *ModifyUserRequest) SetInstanceId(v string) *ModifyUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserRequest) SetMobile(v string) *ModifyUserRequest {
	s.Mobile = &v
	return s
}

func (s *ModifyUserRequest) SetNickname(v string) *ModifyUserRequest {
	s.Nickname = &v
	return s
}

func (s *ModifyUserRequest) SetRoleId(v string) *ModifyUserRequest {
	s.RoleId = &v
	return s
}

func (s *ModifyUserRequest) SetUserId(v string) *ModifyUserRequest {
	s.UserId = &v
	return s
}

func (s *ModifyUserRequest) SetWorkMode(v string) *ModifyUserRequest {
	s.WorkMode = &v
	return s
}

func (s *ModifyUserRequest) Validate() error {
	return dara.Validate(s)
}
