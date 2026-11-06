// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCollaborationGroupMembersResponseBody
	GetCode() *string
	SetItems(v []*ListCollaborationGroupMembersResponseBodyItems) *ListCollaborationGroupMembersResponseBody
	GetItems() []*ListCollaborationGroupMembersResponseBodyItems
	SetMessage(v string) *ListCollaborationGroupMembersResponseBody
	GetMessage() *string
	SetPage(v int64) *ListCollaborationGroupMembersResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupMembersResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCollaborationGroupMembersResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListCollaborationGroupMembersResponseBody
	GetTotal() *int64
}

type ListCollaborationGroupMembersResponseBody struct {
	// 公共业务状态码，由后端Code字段映射；错误时用于识别失败原因。
	//
	// example:
	//
	// ERR.Robject.Auth.TenantAccessDenied
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 当前页记录列表；没有匹配记录时返回空数组。
	//
	// example:
	//
	// [{"userId":"5","displayName":"张三","role":"OWNER","isActive":true}]
	Items []*ListCollaborationGroupMembersResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 公共错误说明，由后端Message字段映射；错误时描述失败原因。
	//
	// example:
	//
	// 无权访问租户 2147483647，用户不属于该租户
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 当前返回页的页码，从1开始。
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 当前分页设置的每页记录数。
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪ID，由后端RequestId字段映射，用于排查调用失败。
	//
	// example:
	//
	// c9cf6cb3-3f5c-4994-a8a1-c93379a6f5cf
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 满足查询条件的记录总数，不是当前页记录数。
	//
	// example:
	//
	// 10
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCollaborationGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCollaborationGroupMembersResponseBody) GetItems() []*ListCollaborationGroupMembersResponseBodyItems {
	return s.Items
}

func (s *ListCollaborationGroupMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCollaborationGroupMembersResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListCollaborationGroupMembersResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCollaborationGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCollaborationGroupMembersResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCollaborationGroupMembersResponseBody) SetCode(v string) *ListCollaborationGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetItems(v []*ListCollaborationGroupMembersResponseBodyItems) *ListCollaborationGroupMembersResponseBody {
	s.Items = v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetMessage(v string) *ListCollaborationGroupMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetPage(v int64) *ListCollaborationGroupMembersResponseBody {
	s.Page = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetPageSize(v int64) *ListCollaborationGroupMembersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetRequestId(v string) *ListCollaborationGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) SetTotal(v int64) *ListCollaborationGroupMembersResponseBody {
	s.Total = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBody) Validate() error {
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

type ListCollaborationGroupMembersResponseBodyItems struct {
	// 成员显示名称；未设置时为空。
	//
	// example:
	//
	// TEST001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 租户账号是否启用。本接口只返回有效且账号启用的成员，值为true；不代表HR在职状态。
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// 成员在空间内的角色：ADMIN为管理员，MEMBER为普通成员。
	//
	// example:
	//
	// ADMIN
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 平台用户ID，字符串形式，可用作转交待办的目标用户ID。
	//
	// example:
	//
	// 5
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListCollaborationGroupMembersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupMembersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupMembersResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListCollaborationGroupMembersResponseBodyItems) GetIsActive() *bool {
	return s.IsActive
}

func (s *ListCollaborationGroupMembersResponseBodyItems) GetRole() *string {
	return s.Role
}

func (s *ListCollaborationGroupMembersResponseBodyItems) GetUserId() *string {
	return s.UserId
}

func (s *ListCollaborationGroupMembersResponseBodyItems) SetDisplayName(v string) *ListCollaborationGroupMembersResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBodyItems) SetIsActive(v bool) *ListCollaborationGroupMembersResponseBodyItems {
	s.IsActive = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBodyItems) SetRole(v string) *ListCollaborationGroupMembersResponseBodyItems {
	s.Role = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBodyItems) SetUserId(v string) *ListCollaborationGroupMembersResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *ListCollaborationGroupMembersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
