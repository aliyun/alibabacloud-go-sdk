// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListRolesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRolesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRolesResponseBody
	GetRequestId() *string
	SetRoles(v *ListRolesResponseBodyRoles) *ListRolesResponseBody
	GetRoles() *ListRolesResponseBodyRoles
	SetTotalCount(v int32) *ListRolesResponseBody
	GetTotalCount() *int32
}

type ListRolesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the roles.
	Roles *ListRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	// The total number of roles.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRolesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRolesResponseBody) GetRoles() *ListRolesResponseBodyRoles {
	return s.Roles
}

func (s *ListRolesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRolesResponseBody) SetPageNumber(v int32) *ListRolesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRolesResponseBody) SetPageSize(v int32) *ListRolesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetRoles(v *ListRolesResponseBodyRoles) *ListRolesResponseBody {
	s.Roles = v
	return s
}

func (s *ListRolesResponseBody) SetTotalCount(v int32) *ListRolesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRolesResponseBody) Validate() error {
	if s.Roles != nil {
		if err := s.Roles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRolesResponseBodyRoles struct {
	Role []*ListRolesResponseBodyRolesRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyRoles) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRoles) GetRole() []*ListRolesResponseBodyRolesRole {
	return s.Role
}

func (s *ListRolesResponseBodyRoles) SetRole(v []*ListRolesResponseBodyRolesRole) *ListRolesResponseBodyRoles {
	s.Role = v
	return s
}

func (s *ListRolesResponseBodyRoles) Validate() error {
	if s.Role != nil {
		for _, item := range s.Role {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRolesResponseBodyRolesRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The time when the role was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the role is a service-linked role.
	//
	// example:
	//
	// true
	IsServiceLinkedRole *bool `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	// The information of the most recent deletion task.
	LatestDeletionTask *ListRolesResponseBodyRolesRoleLatestDeletionTask `json:"LatestDeletionTask,omitempty" xml:"LatestDeletionTask,omitempty" type:"Struct"`
	// The maximum session duration of the role.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the role.
	//
	// example:
	//
	// 90123456789****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the role after authorization.
	//
	// example:
	//
	// ECSAdmin@role.123456.onaliyunservice.com
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	// The time when the role was updated.
	//
	// example:
	//
	// 2016-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListRolesResponseBodyRolesRole) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyRolesRole) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRole) GetArn() *string {
	return s.Arn
}

func (s *ListRolesResponseBodyRolesRole) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListRolesResponseBodyRolesRole) GetDescription() *string {
	return s.Description
}

func (s *ListRolesResponseBodyRolesRole) GetIsServiceLinkedRole() *bool {
	return s.IsServiceLinkedRole
}

func (s *ListRolesResponseBodyRolesRole) GetLatestDeletionTask() *ListRolesResponseBodyRolesRoleLatestDeletionTask {
	return s.LatestDeletionTask
}

func (s *ListRolesResponseBodyRolesRole) GetMaxSessionDuration() *int64 {
	return s.MaxSessionDuration
}

func (s *ListRolesResponseBodyRolesRole) GetRoleId() *string {
	return s.RoleId
}

func (s *ListRolesResponseBodyRolesRole) GetRoleName() *string {
	return s.RoleName
}

func (s *ListRolesResponseBodyRolesRole) GetRolePrincipalName() *string {
	return s.RolePrincipalName
}

func (s *ListRolesResponseBodyRolesRole) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListRolesResponseBodyRolesRole) SetArn(v string) *ListRolesResponseBodyRolesRole {
	s.Arn = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetCreateDate(v string) *ListRolesResponseBodyRolesRole {
	s.CreateDate = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetDescription(v string) *ListRolesResponseBodyRolesRole {
	s.Description = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetIsServiceLinkedRole(v bool) *ListRolesResponseBodyRolesRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetLatestDeletionTask(v *ListRolesResponseBodyRolesRoleLatestDeletionTask) *ListRolesResponseBodyRolesRole {
	s.LatestDeletionTask = v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetMaxSessionDuration(v int64) *ListRolesResponseBodyRolesRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetRoleId(v string) *ListRolesResponseBodyRolesRole {
	s.RoleId = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetRoleName(v string) *ListRolesResponseBodyRolesRole {
	s.RoleName = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetRolePrincipalName(v string) *ListRolesResponseBodyRolesRole {
	s.RolePrincipalName = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) SetUpdateDate(v string) *ListRolesResponseBodyRolesRole {
	s.UpdateDate = &v
	return s
}

func (s *ListRolesResponseBodyRolesRole) Validate() error {
	if s.LatestDeletionTask != nil {
		if err := s.LatestDeletionTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRolesResponseBodyRolesRoleLatestDeletionTask struct {
	// The time when the deletion task was created.
	//
	// example:
	//
	// 2018-10-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the deletion task.
	//
	// example:
	//
	// ECSAdmin/cc61514b-26eb-4453-ab53-b142eb70****
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
}

func (s ListRolesResponseBodyRolesRoleLatestDeletionTask) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyRolesRoleLatestDeletionTask) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyRolesRoleLatestDeletionTask) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListRolesResponseBodyRolesRoleLatestDeletionTask) GetDeletionTaskId() *string {
	return s.DeletionTaskId
}

func (s *ListRolesResponseBodyRolesRoleLatestDeletionTask) SetCreateDate(v string) *ListRolesResponseBodyRolesRoleLatestDeletionTask {
	s.CreateDate = &v
	return s
}

func (s *ListRolesResponseBodyRolesRoleLatestDeletionTask) SetDeletionTaskId(v string) *ListRolesResponseBodyRolesRoleLatestDeletionTask {
	s.DeletionTaskId = &v
	return s
}

func (s *ListRolesResponseBodyRolesRoleLatestDeletionTask) Validate() error {
	return dara.Validate(s)
}
