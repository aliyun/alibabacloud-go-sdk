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
	// The user groups.
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
	if s.UserGroups != nil {
		for _, item := range s.UserGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DsgUserGroupAddOrUpdateRequestUserGroups struct {
	// The users in the user group.
	//
	// - For user groups created based on Alibaba Cloud accounts and Resource Access Management (RAM) roles, call the [DsgUserGroupQueryUserList](https://help.aliyun.com/document_detail/2786445.html) API to query the user list.
	//
	// - For user groups created based on MaxCompute roles, call the [DsgUserGroupGetOdpsRoleGroups](https://help.aliyun.com/document_detail/2785695.html) API to query the user list.
	Accounts []*string `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The ID of the user group.
	//
	// - If you do not specify this parameter, a new user group is created.
	//
	// - If you specify this parameter, the specified user group is modified. You can call the [DsgUserGroupQueryList](https://help.aliyun.com/document_detail/2786441.html) operation to query the ID of the user group.
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
	// The name of the MaxCompute project. Set this parameter when you create a MaxCompute user group.
	//
	// example:
	//
	// dev_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The type of the user group. The following values are valid:
	//
	// - 1: Alibaba Cloud user
	//
	// - 2: RAM Role
	//
	// - 3: MaxCompute Role
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserGroupType *int32  `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Projects      *string `json:"projects,omitempty" xml:"projects,omitempty"`
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

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) GetProjects() *string {
	return s.Projects
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

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) SetProjects(v string) *DsgUserGroupAddOrUpdateRequestUserGroups {
	s.Projects = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateRequestUserGroups) Validate() error {
	return dara.Validate(s)
}
