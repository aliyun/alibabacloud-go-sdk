// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupAddOrUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroups(v []*DsgUserGroupAddOrUpdateRequestUserGroups) *DsgUserGroupAddOrUpdateRequest
	GetUserGroups() []*DsgUserGroupAddOrUpdateRequestUserGroups
}

type DsgUserGroupAddOrUpdateRequest struct {
	// The information about the user group.
	//
	// This parameter is required.
	UserGroups []*DsgUserGroupAddOrUpdateRequestUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s DsgUserGroupAddOrUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupAddOrUpdateRequest) GoString() string {
	return s.String()
}

func (s *DsgUserGroupAddOrUpdateRequest) GetUserGroups() []*DsgUserGroupAddOrUpdateRequestUserGroups {
	return s.UserGroups
}

func (s *DsgUserGroupAddOrUpdateRequest) SetUserGroups(v []*DsgUserGroupAddOrUpdateRequestUserGroups) *DsgUserGroupAddOrUpdateRequest {
	s.UserGroups = v
	return s
}

func (s *DsgUserGroupAddOrUpdateRequest) Validate() error {
	return dara.Validate(s)
}

type DsgUserGroupAddOrUpdateRequestUserGroups struct {
	// The users in the group.
	//
	// 	- If a user group is created by using an Alibaba Cloud account and a RAM role, you can call the [DsgUserGroupQueryUserList](https://help.aliyun.com/document_detail/2786445.html) operation to query the users in the group.
	//
	// 	- If a user group is created by using a MaxCompute role, you can call the [DsgUserGroupQueryUserList](https://help.aliyun.com/document_detail/2785695.html) operation to query the users in the group.
	//
	// This parameter is required.
	Accounts []*string `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The user group ID.
	//
	// 	- If you do not configure this parameter, the current operation is to add a user group.
	//
	// 	- If you configure this parameter, the current operation is to modify a user group. You can call the [DsgUserGroupQueryList](https://help.aliyun.com/document_detail/2786441.html) operation to query the user group ID.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// yun_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// user1
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the MaxCompute project. You must configure this parameter when you create a MaxCompute user group.
	//
	// example:
	//
	// dev_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The type of the user group. Valid values:
	//
	// 	- 1: Alibaba Cloud account
	//
	// 	- 2: RAM role
	//
	// 	- 3: MaxCompute role
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserGroupType *int32 `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
}

func (s DsgUserGroupAddOrUpdateRequestUserGroups) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupAddOrUpdateRequestUserGroups) GoString() string {
	return s.String()
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) GetAccounts() []*string {
	return s.Accounts
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) GetId() *int64 {
	return s.Id
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) GetName() *string {
	return s.Name
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) GetOwner() *string {
	return s.Owner
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) GetProjectName() *string {
	return s.ProjectName
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) GetUserGroupType() *int32 {
	return s.UserGroupType
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) SetAccounts(v []*string) *DsgUserGroupAddOrUpdateRequestUserGroups {
	s.Accounts = v
	return s
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) SetId(v int64) *DsgUserGroupAddOrUpdateRequestUserGroups {
	s.Id = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) SetName(v string) *DsgUserGroupAddOrUpdateRequestUserGroups {
	s.Name = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) SetOwner(v string) *DsgUserGroupAddOrUpdateRequestUserGroups {
	s.Owner = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) SetProjectName(v string) *DsgUserGroupAddOrUpdateRequestUserGroups {
	s.ProjectName = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) SetUserGroupType(v int32) *DsgUserGroupAddOrUpdateRequestUserGroups {
	s.UserGroupType = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) Validate() error {
	return dara.Validate(s)
}
