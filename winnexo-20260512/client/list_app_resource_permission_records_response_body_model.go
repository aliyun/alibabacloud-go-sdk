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
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The authorization records on the current page.
	Items []*ListAppResourcePermissionRecordsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The error message. This parameter is empty if the request is successful.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
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
	// The creation time.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-12-31T23:59:59+08:00
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// The platform user ID of the grantor.
	//
	// example:
	//
	// 10086
	GrantedBy *string `json:"grantedBy,omitempty" xml:"grantedBy,omitempty"`
	// The grantee ID.
	//
	// example:
	//
	// 10086
	GranteeId *string `json:"granteeId,omitempty" xml:"granteeId,omitempty"`
	// The grantee name.
	//
	// example:
	//
	// Zhang San
	GranteeName *string `json:"granteeName,omitempty" xml:"granteeName,omitempty"`
	// The ID of the tenant to which the cross-tenant user belongs.
	//
	// example:
	//
	// 10000
	GranteeTenantId *string `json:"granteeTenantId,omitempty" xml:"granteeTenantId,omitempty"`
	// The grantee type.
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
	// The permission.
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The ID of the permission record.
	//
	// example:
	//
	// 1001
	PermissionId *int64 `json:"permissionId,omitempty" xml:"permissionId,omitempty"`
	// The resource key.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// The resource type.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The update time.
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
