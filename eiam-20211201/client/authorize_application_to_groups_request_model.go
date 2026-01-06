// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AuthorizeApplicationToGroupsRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *AuthorizeApplicationToGroupsRequest
	GetApplicationRoleId() *string
	SetGroupIds(v []*string) *AuthorizeApplicationToGroupsRequest
	GetGroupIds() []*string
	SetInstanceId(v string) *AuthorizeApplicationToGroupsRequest
	GetInstanceId() *string
}

type AuthorizeApplicationToGroupsRequest struct {
	// The application ID.
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
	// The group IDs. You can specify up to 100 group IDs at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AuthorizeApplicationToGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToGroupsRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToGroupsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AuthorizeApplicationToGroupsRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *AuthorizeApplicationToGroupsRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *AuthorizeApplicationToGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthorizeApplicationToGroupsRequest) SetApplicationId(v string) *AuthorizeApplicationToGroupsRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeApplicationToGroupsRequest) SetApplicationRoleId(v string) *AuthorizeApplicationToGroupsRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *AuthorizeApplicationToGroupsRequest) SetGroupIds(v []*string) *AuthorizeApplicationToGroupsRequest {
	s.GroupIds = v
	return s
}

func (s *AuthorizeApplicationToGroupsRequest) SetInstanceId(v string) *AuthorizeApplicationToGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeApplicationToGroupsRequest) Validate() error {
	return dara.Validate(s)
}
