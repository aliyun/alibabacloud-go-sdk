// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *UpdateUserRequest
	GetAvatar() *string
	SetDescription(v string) *UpdateUserRequest
	GetDescription() *string
	SetEmail(v string) *UpdateUserRequest
	GetEmail() *string
	SetGroupInfoList(v []*UpdateUserRequestGroupInfoList) *UpdateUserRequest
	GetGroupInfoList() []*UpdateUserRequestGroupInfoList
	SetNickName(v string) *UpdateUserRequest
	GetNickName() *string
	SetPhone(v string) *UpdateUserRequest
	GetPhone() *string
	SetRole(v string) *UpdateUserRequest
	GetRole() *string
	SetStatus(v string) *UpdateUserRequest
	GetStatus() *string
	SetUserData(v map[string]*string) *UpdateUserRequest
	GetUserData() map[string]*string
	SetUserId(v string) *UpdateUserRequest
	GetUserId() *string
}

type UpdateUserRequest struct {
	// The URL of the profile picture.
	//
	// If you specify the parameter in the HTTP URL format, the URL must start with http:// or https:// and can be up to 4 KB in size.
	//
	// If you specify the parameter in the DATA URL format, the URL must start with data:// and be encoded in Base64. The URL can be up to 300 KB in size.
	//
	// example:
	//
	// http://a.b.c/pds.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The description of the user. The description can be up to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// a@aliyunpds.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The information about the group.
	GroupInfoList []*UpdateUserRequestGroupInfoList `json:"group_info_list,omitempty" xml:"group_info_list,omitempty" type:"Repeated"`
	// The nickname of the user. The nickname can be up to 128 characters in length.
	//
	// example:
	//
	// pdsuer
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The mobile number of the user.
	//
	// example:
	//
	// 13900001111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The role of the user. Valid values:
	//
	// 	- superadmin
	//
	// 	- admin
	//
	// 	- user
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The state of the user. Valid values:
	//
	// 	- disabled: The user is prohibited from logon.
	//
	// 	- enabled: The user is in a normal state.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The custom data. The data can be up to 1,024 characters in length.
	UserData map[string]*string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The user ID. The ID can be up to 64 characters in length and cannot contain a number sign (#).
	//
	// This parameter is required.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateUserRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserRequest) GetGroupInfoList() []*UpdateUserRequestGroupInfoList {
	return s.GroupInfoList
}

func (s *UpdateUserRequest) GetNickName() *string {
	return s.NickName
}

func (s *UpdateUserRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateUserRequest) GetRole() *string {
	return s.Role
}

func (s *UpdateUserRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserRequest) GetUserData() map[string]*string {
	return s.UserData
}

func (s *UpdateUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserRequest) SetAvatar(v string) *UpdateUserRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateUserRequest) SetDescription(v string) *UpdateUserRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetGroupInfoList(v []*UpdateUserRequestGroupInfoList) *UpdateUserRequest {
	s.GroupInfoList = v
	return s
}

func (s *UpdateUserRequest) SetNickName(v string) *UpdateUserRequest {
	s.NickName = &v
	return s
}

func (s *UpdateUserRequest) SetPhone(v string) *UpdateUserRequest {
	s.Phone = &v
	return s
}

func (s *UpdateUserRequest) SetRole(v string) *UpdateUserRequest {
	s.Role = &v
	return s
}

func (s *UpdateUserRequest) SetStatus(v string) *UpdateUserRequest {
	s.Status = &v
	return s
}

func (s *UpdateUserRequest) SetUserData(v map[string]*string) *UpdateUserRequest {
	s.UserData = v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	if s.GroupInfoList != nil {
		for _, item := range s.GroupInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateUserRequestGroupInfoList struct {
	// The group ID.
	//
	// example:
	//
	// g123
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s UpdateUserRequestGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequestGroupInfoList) GoString() string {
	return s.String()
}

func (s *UpdateUserRequestGroupInfoList) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateUserRequestGroupInfoList) SetGroupId(v string) *UpdateUserRequestGroupInfoList {
	s.GroupId = &v
	return s
}

func (s *UpdateUserRequestGroupInfoList) Validate() error {
	return dara.Validate(s)
}
