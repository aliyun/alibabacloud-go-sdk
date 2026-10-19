// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResourcePermissionRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAppResourcePermissionRecordsResponseBody
	GetCode() *string
	SetItems(v []*ListAppResourcePermissionRecordsResponseBodyItems) *ListAppResourcePermissionRecordsResponseBody
	GetItems() []*ListAppResourcePermissionRecordsResponseBodyItems
	SetMessage(v string) *ListAppResourcePermissionRecordsResponseBody
	GetMessage() *string
	SetPage(v int64) *ListAppResourcePermissionRecordsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListAppResourcePermissionRecordsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAppResourcePermissionRecordsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAppResourcePermissionRecordsResponseBody
	GetTotal() *int64
}

type ListAppResourcePermissionRecordsResponseBody struct {
	// 业务状态码；成功为 200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 当前页授权记录
	Items []*ListAppResourcePermissionRecordsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述；成功时为空
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 记录总数
	//
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAppResourcePermissionRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourcePermissionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppResourcePermissionRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAppResourcePermissionRecordsResponseBody) GetItems() []*ListAppResourcePermissionRecordsResponseBodyItems {
	return s.Items
}

func (s *ListAppResourcePermissionRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppResourcePermissionRecordsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListAppResourcePermissionRecordsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAppResourcePermissionRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppResourcePermissionRecordsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAppResourcePermissionRecordsResponseBody) SetCode(v string) *ListAppResourcePermissionRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBody) SetItems(v []*ListAppResourcePermissionRecordsResponseBodyItems) *ListAppResourcePermissionRecordsResponseBody {
	s.Items = v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBody) SetMessage(v string) *ListAppResourcePermissionRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBody) SetPage(v int64) *ListAppResourcePermissionRecordsResponseBody {
	s.Page = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBody) SetPageSize(v int64) *ListAppResourcePermissionRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBody) SetRequestId(v string) *ListAppResourcePermissionRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBody) SetTotal(v int64) *ListAppResourcePermissionRecordsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBody) Validate() error {
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

type ListAppResourcePermissionRecordsResponseBodyItems struct {
	// 创建时间
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 到期时间
	//
	// example:
	//
	// 2026-12-31T23:59:59+08:00
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// 授权人平台用户 ID
	//
	// example:
	//
	// 10086
	GrantedBy *string `json:"grantedBy,omitempty" xml:"grantedBy,omitempty"`
	// 主体 ID
	//
	// example:
	//
	// 10086
	GranteeId *string `json:"granteeId,omitempty" xml:"granteeId,omitempty"`
	// 主体名称
	//
	// example:
	//
	// 张三
	GranteeName *string `json:"granteeName,omitempty" xml:"granteeName,omitempty"`
	// 跨租户 USER 所属租户
	//
	// example:
	//
	// 10000
	GranteeTenantId *string `json:"granteeTenantId,omitempty" xml:"granteeTenantId,omitempty"`
	// 主体类型
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
	// 权限点
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// 权限记录 ID
	//
	// example:
	//
	// 1001
	PermissionId *int64 `json:"permissionId,omitempty" xml:"permissionId,omitempty"`
	// 资源键
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// 资源类型
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListAppResourcePermissionRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourcePermissionRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetGrantedBy() *string {
	return s.GrantedBy
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetGranteeId() *string {
	return s.GranteeId
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetGranteeName() *string {
	return s.GranteeName
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetGranteeTenantId() *string {
	return s.GranteeTenantId
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetGranteeType() *string {
	return s.GranteeType
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetPermission() *string {
	return s.Permission
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetPermissionId() *int64 {
	return s.PermissionId
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetCreatedAt(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetExpireAt(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.ExpireAt = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetGrantedBy(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.GrantedBy = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetGranteeId(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.GranteeId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetGranteeName(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.GranteeName = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetGranteeTenantId(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.GranteeTenantId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetGranteeType(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.GranteeType = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetPermission(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.Permission = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetPermissionId(v int64) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.PermissionId = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetResourceKey(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.ResourceKey = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetResourceType(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) SetUpdatedAt(v string) *ListAppResourcePermissionRecordsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
