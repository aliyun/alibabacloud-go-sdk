// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AuthorizeApplicationToUsersRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *AuthorizeApplicationToUsersRequest
	GetApplicationRoleId() *string
	SetInstanceId(v string) *AuthorizeApplicationToUsersRequest
	GetInstanceId() *string
	SetUserIds(v []*string) *AuthorizeApplicationToUsersRequest
	GetUserIds() []*string
}

type AuthorizeApplicationToUsersRequest struct {
	// The ID of the application on which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色ID。
	//
	// example:
	//
	// app_role_mkv7rgt4ds8d8v0qtzev2mxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the accounts to which you want to grant permissions. You can grant permissions to a maximum of 100 accounts at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s AuthorizeApplicationToUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToUsersRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToUsersRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AuthorizeApplicationToUsersRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *AuthorizeApplicationToUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthorizeApplicationToUsersRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *AuthorizeApplicationToUsersRequest) SetApplicationId(v string) *AuthorizeApplicationToUsersRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeApplicationToUsersRequest) SetApplicationRoleId(v string) *AuthorizeApplicationToUsersRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *AuthorizeApplicationToUsersRequest) SetInstanceId(v string) *AuthorizeApplicationToUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeApplicationToUsersRequest) SetUserIds(v []*string) *AuthorizeApplicationToUsersRequest {
	s.UserIds = v
	return s
}

func (s *AuthorizeApplicationToUsersRequest) Validate() error {
	return dara.Validate(s)
}
