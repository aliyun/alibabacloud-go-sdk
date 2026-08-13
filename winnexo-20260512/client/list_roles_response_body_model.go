// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRolesResponseBody
	GetCode() *string
	SetItems(v []*ListRolesResponseBodyItems) *ListRolesResponseBody
	GetItems() []*ListRolesResponseBodyItems
	SetMessage(v string) *ListRolesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRolesResponseBody
	GetRequestId() *string
}

type ListRolesResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 系统内置角色列表（固定 7 个）
	Items []*ListRolesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRolesResponseBody) GetItems() []*ListRolesResponseBodyItems {
	return s.Items
}

func (s *ListRolesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRolesResponseBody) SetCode(v string) *ListRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRolesResponseBody) SetItems(v []*ListRolesResponseBodyItems) *ListRolesResponseBody {
	s.Items = v
	return s
}

func (s *ListRolesResponseBody) SetMessage(v string) *ListRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRolesResponseBodyItems struct {
	// 角色说明（按请求 Accept-Language 国际化）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 角色标识码，用于 createUser/updateUser 的 roleCodes 参数
	//
	// example:
	//
	// string_value
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 角色显示名称（按请求 Accept-Language 国际化）
	//
	// example:
	//
	// string_value
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 是否允许启用/停用操作（超级管理员和应用用户不可切换）
	//
	// example:
	//
	// true
	Toggleable *bool `json:"toggleable,omitempty" xml:"toggleable,omitempty"`
}

func (s ListRolesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListRolesResponseBodyItems) GetRoleCode() *string {
	return s.RoleCode
}

func (s *ListRolesResponseBodyItems) GetRoleName() *string {
	return s.RoleName
}

func (s *ListRolesResponseBodyItems) GetToggleable() *bool {
	return s.Toggleable
}

func (s *ListRolesResponseBodyItems) SetDescription(v string) *ListRolesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListRolesResponseBodyItems) SetRoleCode(v string) *ListRolesResponseBodyItems {
	s.RoleCode = &v
	return s
}

func (s *ListRolesResponseBodyItems) SetRoleName(v string) *ListRolesResponseBodyItems {
	s.RoleName = &v
	return s
}

func (s *ListRolesResponseBodyItems) SetToggleable(v bool) *ListRolesResponseBodyItems {
	s.Toggleable = &v
	return s
}

func (s *ListRolesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
