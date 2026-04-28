// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *CreateUserRequest
	GetAvatar() *string
	SetDescription(v string) *CreateUserRequest
	GetDescription() *string
	SetEmail(v string) *CreateUserRequest
	GetEmail() *string
	SetGroupInfoList(v []*CreateUserRequestGroupInfoList) *CreateUserRequest
	GetGroupInfoList() []*CreateUserRequestGroupInfoList
	SetNickName(v string) *CreateUserRequest
	GetNickName() *string
	SetPhone(v string) *CreateUserRequest
	GetPhone() *string
	SetRole(v string) *CreateUserRequest
	GetRole() *string
	SetStatus(v string) *CreateUserRequest
	GetStatus() *string
	SetUserData(v map[string]interface{}) *CreateUserRequest
	GetUserData() map[string]interface{}
	SetUserId(v string) *CreateUserRequest
	GetUserId() *string
	SetUserName(v string) *CreateUserRequest
	GetUserName() *string
}

type CreateUserRequest struct {
	// The URL of the profile picture.
	//
	// If you specify the parameter in the HTTP URL format, the URL must start with http:// or https:// and can be up to 4 KB in size.
	//
	// If you specify the parameter in the data URL format, the URL must start with data:// and be encoded in Base64. The URL can be up to 300 KB in size.
	//
	// example:
	//
	// http://a.b.c/pds.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The description of the user. The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// The VIP user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The email address.
	//
	// example:
	//
	// 123@pds.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The information about the group.
	GroupInfoList []*CreateUserRequestGroupInfoList `json:"group_info_list,omitempty" xml:"group_info_list,omitempty" type:"Repeated"`
	// The nickname of the user. The nickname can be up to 128 characters in length.
	//
	// example:
	//
	// pdsuer
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 13900001111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The role of the user. Default value: user. Valid values:
	//
	// 	- superadmin
	//
	// 	- admin
	//
	// 	- user
	//
	// If the domain can be divided into subdomains, the subdomain_super_admin and subdomain_admin roles are also supported.
	//
	// Valid values:
	//
	// 	- subdomain_super_admin
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- subdomain_admin
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- superadmin
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- admin
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- user
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The state of the user. Default value: enabled. Valid values:
	//
	// 	- enabled: The user is in a normal state.
	//
	// 	- disabled: The user is prohibited from logon.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The custom data. The data can be up to 1,024 characters in length.
	//
	// example:
	//
	// md
	UserData map[string]interface{} `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The user ID. The ID can be up to 64 characters in length and cannot contain number signs (#).
	//
	// This parameter is required.
	//
	// example:
	//
	// pdsuserid1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// The username. The username can be up to 128 characters in length.
	//
	// example:
	//
	// pdsusername
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *CreateUserRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequest) GetGroupInfoList() []*CreateUserRequestGroupInfoList {
	return s.GroupInfoList
}

func (s *CreateUserRequest) GetNickName() *string {
	return s.NickName
}

func (s *CreateUserRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateUserRequest) GetRole() *string {
	return s.Role
}

func (s *CreateUserRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateUserRequest) GetUserData() map[string]interface{} {
	return s.UserData
}

func (s *CreateUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserRequest) SetAvatar(v string) *CreateUserRequest {
	s.Avatar = &v
	return s
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetGroupInfoList(v []*CreateUserRequestGroupInfoList) *CreateUserRequest {
	s.GroupInfoList = v
	return s
}

func (s *CreateUserRequest) SetNickName(v string) *CreateUserRequest {
	s.NickName = &v
	return s
}

func (s *CreateUserRequest) SetPhone(v string) *CreateUserRequest {
	s.Phone = &v
	return s
}

func (s *CreateUserRequest) SetRole(v string) *CreateUserRequest {
	s.Role = &v
	return s
}

func (s *CreateUserRequest) SetStatus(v string) *CreateUserRequest {
	s.Status = &v
	return s
}

func (s *CreateUserRequest) SetUserData(v map[string]interface{}) *CreateUserRequest {
	s.UserData = v
	return s
}

func (s *CreateUserRequest) SetUserId(v string) *CreateUserRequest {
	s.UserId = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
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

type CreateUserRequestGroupInfoList struct {
	// The group ID.
	//
	// example:
	//
	// g123
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s CreateUserRequestGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequestGroupInfoList) GoString() string {
	return s.String()
}

func (s *CreateUserRequestGroupInfoList) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateUserRequestGroupInfoList) SetGroupId(v string) *CreateUserRequestGroupInfoList {
	s.GroupId = &v
	return s
}

func (s *CreateUserRequestGroupInfoList) Validate() error {
	return dara.Validate(s)
}
