// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCollaborationGroupsResponseBody
	GetCode() *string
	SetItems(v []*ListCollaborationGroupsResponseBodyItems) *ListCollaborationGroupsResponseBody
	GetItems() []*ListCollaborationGroupsResponseBodyItems
	SetMessage(v string) *ListCollaborationGroupsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListCollaborationGroupsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCollaborationGroupsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListCollaborationGroupsResponseBody
	GetTotal() *int64
}

type ListCollaborationGroupsResponseBody struct {
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
	// [{"collaborationGroupId":"cg_9401","name":"产品协作空间","status":"ACTIVE","currentUserRole":"OWNER"}]
	Items []*ListCollaborationGroupsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 满足查询条件的有效空间总数，不是当前页记录数。
	//
	// example:
	//
	// 39
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCollaborationGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCollaborationGroupsResponseBody) GetItems() []*ListCollaborationGroupsResponseBodyItems {
	return s.Items
}

func (s *ListCollaborationGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCollaborationGroupsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListCollaborationGroupsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCollaborationGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCollaborationGroupsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListCollaborationGroupsResponseBody) SetCode(v string) *ListCollaborationGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetItems(v []*ListCollaborationGroupsResponseBodyItems) *ListCollaborationGroupsResponseBody {
	s.Items = v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetMessage(v string) *ListCollaborationGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetPage(v int64) *ListCollaborationGroupsResponseBody {
	s.Page = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetPageSize(v int64) *ListCollaborationGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetRequestId(v string) *ListCollaborationGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) SetTotal(v int64) *ListCollaborationGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *ListCollaborationGroupsResponseBody) Validate() error {
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

type ListCollaborationGroupsResponseBodyItems struct {
	// 协作空间ID，可用于查询空间成员、查询节点详情及预览对话文件。
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 当前调用用户在空间内的角色：ADMIN为管理员，MEMBER为普通成员。
	//
	// example:
	//
	// ADMIN
	CurrentUserRole *string `json:"currentUserRole,omitempty" xml:"currentUserRole,omitempty"`
	// 协作空间显示名称。
	//
	// example:
	//
	// 工作流待办测试
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 协作空间状态，当前接口只返回ACTIVE。
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListCollaborationGroupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupsResponseBodyItems) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListCollaborationGroupsResponseBodyItems) GetCurrentUserRole() *string {
	return s.CurrentUserRole
}

func (s *ListCollaborationGroupsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListCollaborationGroupsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListCollaborationGroupsResponseBodyItems) SetCollaborationGroupId(v string) *ListCollaborationGroupsResponseBodyItems {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListCollaborationGroupsResponseBodyItems) SetCurrentUserRole(v string) *ListCollaborationGroupsResponseBodyItems {
	s.CurrentUserRole = &v
	return s
}

func (s *ListCollaborationGroupsResponseBodyItems) SetName(v string) *ListCollaborationGroupsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListCollaborationGroupsResponseBodyItems) SetStatus(v string) *ListCollaborationGroupsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListCollaborationGroupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
