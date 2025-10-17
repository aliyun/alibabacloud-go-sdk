// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMembersResponseBody
	GetMaxResults() *int32
	SetMembers(v []*ListMembersResponseBodyMembers) *ListMembersResponseBody
	GetMembers() []*ListMembersResponseBodyMembers
	SetNextToken(v string) *ListMembersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMembersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMembersResponseBody
	GetTotalCount() *int32
}

type ListMembersResponseBody struct {
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Members    []*ListMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 下一页TOKEN。
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 记录总数。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMembersResponseBody) GetMembers() []*ListMembersResponseBodyMembers {
	return s.Members
}

func (s *ListMembersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMembersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMembersResponseBody) SetMaxResults(v int32) *ListMembersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMembersResponseBody) SetMembers(v []*ListMembersResponseBodyMembers) *ListMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListMembersResponseBody) SetNextToken(v string) *ListMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMembersResponseBody) SetRequestId(v string) *ListMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMembersResponseBody) SetTotalCount(v int32) *ListMembersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMembersResponseBody) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMembersResponseBodyMembers struct {
	// 针对此用户允许的操作列表。
	Actions []*ListMembersResponseBodyMembersActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// example:
	//
	// 1753412502000
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 用户展示名称。
	//
	// example:
	//
	// jia***test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 用户 arn。
	//
	// example:
	//
	// acs:emr::w-1234****abcd:member/202265*****276
	MemberArn *string `json:"memberArn,omitempty" xml:"memberArn,omitempty"`
	// 用户角色列表。
	Roles []*ListMembersResponseBodyMembersRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
	// 用户名称。
	//
	// example:
	//
	// jia***test@195*****7311.onaliyun.com
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 用户类型。
	//
	// example:
	//
	// MEMBER
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
	// example:
	//
	// true
	Visible *bool `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s ListMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMembers) GetActions() []*ListMembersResponseBodyMembersActions {
	return s.Actions
}

func (s *ListMembersResponseBodyMembers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMembersResponseBodyMembers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMembersResponseBodyMembers) GetMemberArn() *string {
	return s.MemberArn
}

func (s *ListMembersResponseBodyMembers) GetRoles() []*ListMembersResponseBodyMembersRoles {
	return s.Roles
}

func (s *ListMembersResponseBodyMembers) GetUserName() *string {
	return s.UserName
}

func (s *ListMembersResponseBodyMembers) GetUserType() *string {
	return s.UserType
}

func (s *ListMembersResponseBodyMembers) GetVisible() *bool {
	return s.Visible
}

func (s *ListMembersResponseBodyMembers) SetActions(v []*ListMembersResponseBodyMembersActions) *ListMembersResponseBodyMembers {
	s.Actions = v
	return s
}

func (s *ListMembersResponseBodyMembers) SetCreateTime(v string) *ListMembersResponseBodyMembers {
	s.CreateTime = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetDisplayName(v string) *ListMembersResponseBodyMembers {
	s.DisplayName = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetMemberArn(v string) *ListMembersResponseBodyMembers {
	s.MemberArn = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetRoles(v []*ListMembersResponseBodyMembersRoles) *ListMembersResponseBodyMembers {
	s.Roles = v
	return s
}

func (s *ListMembersResponseBodyMembers) SetUserName(v string) *ListMembersResponseBodyMembers {
	s.UserName = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetUserType(v string) *ListMembersResponseBodyMembers {
	s.UserType = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetVisible(v bool) *ListMembersResponseBodyMembers {
	s.Visible = &v
	return s
}

func (s *ListMembersResponseBodyMembers) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMembersResponseBodyMembersActions struct {
	// 行为 arn。
	//
	// example:
	//
	// acs:emr::w-1234****abcd:action/add_MEMBER
	ActionArn *string `json:"actionArn,omitempty" xml:"actionArn,omitempty"`
	// 权限名称。
	//
	// example:
	//
	// add
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// action 依赖列表。
	//
	// example:
	//
	// ["view"]
	Dependencies []*string `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	// action 描述。
	//
	// example:
	//
	// add members
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 权限展示名称。
	//
	// example:
	//
	// add members
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s ListMembersResponseBodyMembersActions) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponseBodyMembersActions) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMembersActions) GetActionArn() *string {
	return s.ActionArn
}

func (s *ListMembersResponseBodyMembersActions) GetActionName() *string {
	return s.ActionName
}

func (s *ListMembersResponseBodyMembersActions) GetDependencies() []*string {
	return s.Dependencies
}

func (s *ListMembersResponseBodyMembersActions) GetDescription() *string {
	return s.Description
}

func (s *ListMembersResponseBodyMembersActions) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMembersResponseBodyMembersActions) SetActionArn(v string) *ListMembersResponseBodyMembersActions {
	s.ActionArn = &v
	return s
}

func (s *ListMembersResponseBodyMembersActions) SetActionName(v string) *ListMembersResponseBodyMembersActions {
	s.ActionName = &v
	return s
}

func (s *ListMembersResponseBodyMembersActions) SetDependencies(v []*string) *ListMembersResponseBodyMembersActions {
	s.Dependencies = v
	return s
}

func (s *ListMembersResponseBodyMembersActions) SetDescription(v string) *ListMembersResponseBodyMembersActions {
	s.Description = &v
	return s
}

func (s *ListMembersResponseBodyMembersActions) SetDisplayName(v string) *ListMembersResponseBodyMembersActions {
	s.DisplayName = &v
	return s
}

func (s *ListMembersResponseBodyMembersActions) Validate() error {
	return dara.Validate(s)
}

type ListMembersResponseBodyMembersRoles struct {
	// 权限列表。
	Actions []*ListMembersResponseBodyMembersRolesActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// 创建时间。
	//
	// example:
	//
	// 1753412502000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 描述。
	//
	// example:
	//
	// DataScience
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 角色 arn。
	//
	// example:
	//
	// acs:emr::w-1234****abcd:role/DataScience
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// 角色名称。
	//
	// example:
	//
	// DataScience
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s ListMembersResponseBodyMembersRoles) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponseBodyMembersRoles) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMembersRoles) GetActions() []*ListMembersResponseBodyMembersRolesActions {
	return s.Actions
}

func (s *ListMembersResponseBodyMembersRoles) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMembersResponseBodyMembersRoles) GetDescription() *string {
	return s.Description
}

func (s *ListMembersResponseBodyMembersRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListMembersResponseBodyMembersRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *ListMembersResponseBodyMembersRoles) SetActions(v []*ListMembersResponseBodyMembersRolesActions) *ListMembersResponseBodyMembersRoles {
	s.Actions = v
	return s
}

func (s *ListMembersResponseBodyMembersRoles) SetCreateTime(v int64) *ListMembersResponseBodyMembersRoles {
	s.CreateTime = &v
	return s
}

func (s *ListMembersResponseBodyMembersRoles) SetDescription(v string) *ListMembersResponseBodyMembersRoles {
	s.Description = &v
	return s
}

func (s *ListMembersResponseBodyMembersRoles) SetRoleArn(v string) *ListMembersResponseBodyMembersRoles {
	s.RoleArn = &v
	return s
}

func (s *ListMembersResponseBodyMembersRoles) SetRoleName(v string) *ListMembersResponseBodyMembersRoles {
	s.RoleName = &v
	return s
}

func (s *ListMembersResponseBodyMembersRoles) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMembersResponseBodyMembersRolesActions struct {
	// 行为 arn。
	//
	// example:
	//
	// acs:emr::w-1234****abcd:action/add_MEMBER
	ActionArn *string `json:"actionArn,omitempty" xml:"actionArn,omitempty"`
	// 权限名称。
	//
	// example:
	//
	// add
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// action 依赖列表。
	//
	// example:
	//
	// ["view"]
	Dependencies []*string `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	// action 描述。
	//
	// example:
	//
	// add members
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 权限展示名称。
	//
	// example:
	//
	// add members
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s ListMembersResponseBodyMembersRolesActions) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponseBodyMembersRolesActions) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMembersRolesActions) GetActionArn() *string {
	return s.ActionArn
}

func (s *ListMembersResponseBodyMembersRolesActions) GetActionName() *string {
	return s.ActionName
}

func (s *ListMembersResponseBodyMembersRolesActions) GetDependencies() []*string {
	return s.Dependencies
}

func (s *ListMembersResponseBodyMembersRolesActions) GetDescription() *string {
	return s.Description
}

func (s *ListMembersResponseBodyMembersRolesActions) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMembersResponseBodyMembersRolesActions) SetActionArn(v string) *ListMembersResponseBodyMembersRolesActions {
	s.ActionArn = &v
	return s
}

func (s *ListMembersResponseBodyMembersRolesActions) SetActionName(v string) *ListMembersResponseBodyMembersRolesActions {
	s.ActionName = &v
	return s
}

func (s *ListMembersResponseBodyMembersRolesActions) SetDependencies(v []*string) *ListMembersResponseBodyMembersRolesActions {
	s.Dependencies = v
	return s
}

func (s *ListMembersResponseBodyMembersRolesActions) SetDescription(v string) *ListMembersResponseBodyMembersRolesActions {
	s.Description = &v
	return s
}

func (s *ListMembersResponseBodyMembersRolesActions) SetDisplayName(v string) *ListMembersResponseBodyMembersRolesActions {
	s.DisplayName = &v
	return s
}

func (s *ListMembersResponseBodyMembersRolesActions) Validate() error {
	return dara.Validate(s)
}
