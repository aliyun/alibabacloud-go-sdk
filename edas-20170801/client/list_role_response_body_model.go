// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListRoleResponseBody
	GetCode() *int32
	SetMessage(v string) *ListRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRoleResponseBody
	GetRequestId() *string
	SetRoleList(v *ListRoleResponseBodyRoleList) *ListRoleResponseBody
	GetRoleList() *ListRoleResponseBodyRoleList
}

type ListRoleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57609587-DFA2-41EC-****-*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The roles.
	RoleList *ListRoleResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Struct"`
}

func (s ListRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoleResponseBody) GetRoleList() *ListRoleResponseBodyRoleList {
	return s.RoleList
}

func (s *ListRoleResponseBody) SetCode(v int32) *ListRoleResponseBody {
	s.Code = &v
	return s
}

func (s *ListRoleResponseBody) SetMessage(v string) *ListRoleResponseBody {
	s.Message = &v
	return s
}

func (s *ListRoleResponseBody) SetRequestId(v string) *ListRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoleResponseBody) SetRoleList(v *ListRoleResponseBodyRoleList) *ListRoleResponseBody {
	s.RoleList = v
	return s
}

func (s *ListRoleResponseBody) Validate() error {
	if s.RoleList != nil {
		if err := s.RoleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRoleResponseBodyRoleList struct {
	RoleItem []*ListRoleResponseBodyRoleListRoleItem `json:"RoleItem,omitempty" xml:"RoleItem,omitempty" type:"Repeated"`
}

func (s ListRoleResponseBodyRoleList) String() string {
	return dara.Prettify(s)
}

func (s ListRoleResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleList) GetRoleItem() []*ListRoleResponseBodyRoleListRoleItem {
	return s.RoleItem
}

func (s *ListRoleResponseBodyRoleList) SetRoleItem(v []*ListRoleResponseBodyRoleListRoleItem) *ListRoleResponseBodyRoleList {
	s.RoleItem = v
	return s
}

func (s *ListRoleResponseBodyRoleList) Validate() error {
	if s.RoleItem != nil {
		for _, item := range s.RoleItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoleResponseBodyRoleListRoleItem struct {
	// The set of permissions to be granted to the role.
	ActionList *ListRoleResponseBodyRoleListRoleItemActionList `json:"ActionList,omitempty" xml:"ActionList,omitempty" type:"Struct"`
	// The roles.
	Role *ListRoleResponseBodyRoleListRoleItemRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s ListRoleResponseBodyRoleListRoleItem) String() string {
	return dara.Prettify(s)
}

func (s ListRoleResponseBodyRoleListRoleItem) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleListRoleItem) GetActionList() *ListRoleResponseBodyRoleListRoleItemActionList {
	return s.ActionList
}

func (s *ListRoleResponseBodyRoleListRoleItem) GetRole() *ListRoleResponseBodyRoleListRoleItemRole {
	return s.Role
}

func (s *ListRoleResponseBodyRoleListRoleItem) SetActionList(v *ListRoleResponseBodyRoleListRoleItemActionList) *ListRoleResponseBodyRoleListRoleItem {
	s.ActionList = v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItem) SetRole(v *ListRoleResponseBodyRoleListRoleItemRole) *ListRoleResponseBodyRoleListRoleItem {
	s.Role = v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItem) Validate() error {
	if s.ActionList != nil {
		if err := s.ActionList.Validate(); err != nil {
			return err
		}
	}
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRoleResponseBodyRoleListRoleItemActionList struct {
	Action []*ListRoleResponseBodyRoleListRoleItemActionListAction `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
}

func (s ListRoleResponseBodyRoleListRoleItemActionList) String() string {
	return dara.Prettify(s)
}

func (s ListRoleResponseBodyRoleListRoleItemActionList) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleListRoleItemActionList) GetAction() []*ListRoleResponseBodyRoleListRoleItemActionListAction {
	return s.Action
}

func (s *ListRoleResponseBodyRoleListRoleItemActionList) SetAction(v []*ListRoleResponseBodyRoleListRoleItemActionListAction) *ListRoleResponseBodyRoleListRoleItemActionList {
	s.Action = v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemActionList) Validate() error {
	if s.Action != nil {
		for _, item := range s.Action {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoleResponseBodyRoleListRoleItemActionListAction struct {
	// The serial number of the permission that is granted to the role.
	//
	// example:
	//
	// 1
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the permission to be granted to the role.
	//
	// example:
	//
	// Operations in operation records
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the permission group to which the permission that is granted to the role belongs.
	//
	// example:
	//
	// 31
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the permission to be granted to the role.
	//
	// example:
	//
	// Operation records
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListRoleResponseBodyRoleListRoleItemActionListAction) String() string {
	return dara.Prettify(s)
}

func (s ListRoleResponseBodyRoleListRoleItemActionListAction) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) GetCode() *string {
	return s.Code
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) GetDescription() *string {
	return s.Description
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) GetGroupId() *string {
	return s.GroupId
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) GetName() *string {
	return s.Name
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) SetCode(v string) *ListRoleResponseBodyRoleListRoleItemActionListAction {
	s.Code = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) SetDescription(v string) *ListRoleResponseBodyRoleListRoleItemActionListAction {
	s.Description = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) SetGroupId(v string) *ListRoleResponseBodyRoleListRoleItemActionListAction {
	s.GroupId = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) SetName(v string) *ListRoleResponseBodyRoleListRoleItemActionListAction {
	s.Name = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) Validate() error {
	return dara.Validate(s)
}

type ListRoleResponseBodyRoleListRoleItemRole struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// test**@aliyun.com
	AdminUserId *string `json:"AdminUserId,omitempty" xml:"AdminUserId,omitempty"`
	// The timestamp when the role was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1542717260156
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the role.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the role is a default role.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// Super Admin(All privileges)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The timestamp when the role was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1542717260156
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListRoleResponseBodyRoleListRoleItemRole) String() string {
	return dara.Prettify(s)
}

func (s ListRoleResponseBodyRoleListRoleItemRole) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) GetAdminUserId() *string {
	return s.AdminUserId
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) GetId() *int32 {
	return s.Id
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) GetName() *string {
	return s.Name
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetAdminUserId(v string) *ListRoleResponseBodyRoleListRoleItemRole {
	s.AdminUserId = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetCreateTime(v int64) *ListRoleResponseBodyRoleListRoleItemRole {
	s.CreateTime = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetId(v int32) *ListRoleResponseBodyRoleListRoleItemRole {
	s.Id = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetIsDefault(v bool) *ListRoleResponseBodyRoleListRoleItemRole {
	s.IsDefault = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetName(v string) *ListRoleResponseBodyRoleListRoleItemRole {
	s.Name = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetUpdateTime(v int64) *ListRoleResponseBodyRoleListRoleItemRole {
	s.UpdateTime = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) Validate() error {
	return dara.Validate(s)
}
