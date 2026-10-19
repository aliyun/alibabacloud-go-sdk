// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAppResourcePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterAppResourcePermissionsResponseBody
	GetCode() *string
	SetMessage(v string) *RegisterAppResourcePermissionsResponseBody
	GetMessage() *string
	SetRecords(v []*RegisterAppResourcePermissionsResponseBodyRecords) *RegisterAppResourcePermissionsResponseBody
	GetRecords() []*RegisterAppResourcePermissionsResponseBodyRecords
	SetRequestId(v string) *RegisterAppResourcePermissionsResponseBody
	GetRequestId() *string
}

type RegisterAppResourcePermissionsResponseBody struct {
	// 业务状态码；成功为 200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述；成功时为空
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 新建或幂等命中的权限记录
	Records []*RegisterAppResourcePermissionsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RegisterAppResourcePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterAppResourcePermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterAppResourcePermissionsResponseBody) GetRecords() []*RegisterAppResourcePermissionsResponseBodyRecords {
	return s.Records
}

func (s *RegisterAppResourcePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterAppResourcePermissionsResponseBody) SetCode(v string) *RegisterAppResourcePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBody) SetMessage(v string) *RegisterAppResourcePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBody) SetRecords(v []*RegisterAppResourcePermissionsResponseBodyRecords) *RegisterAppResourcePermissionsResponseBody {
	s.Records = v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBody) SetRequestId(v string) *RegisterAppResourcePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RegisterAppResourcePermissionsResponseBodyRecords struct {
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

func (s RegisterAppResourcePermissionsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGrantedBy() *string {
	return s.GrantedBy
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGranteeId() *string {
	return s.GranteeId
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGranteeName() *string {
	return s.GranteeName
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGranteeTenantId() *string {
	return s.GranteeTenantId
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetGranteeType() *string {
	return s.GranteeType
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetPermission() *string {
	return s.Permission
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetPermissionId() *int64 {
	return s.PermissionId
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetResourceType() *string {
	return s.ResourceType
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetCreatedAt(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.CreatedAt = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetExpireAt(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.ExpireAt = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGrantedBy(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GrantedBy = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGranteeId(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GranteeId = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGranteeName(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GranteeName = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGranteeTenantId(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GranteeTenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetGranteeType(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.GranteeType = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetPermission(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.Permission = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetPermissionId(v int64) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.PermissionId = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetResourceKey(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.ResourceKey = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetResourceType(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.ResourceType = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) SetUpdatedAt(v string) *RegisterAppResourcePermissionsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
