// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAppResourcePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*RegisterAppResourcePermissionsRequestItems) *RegisterAppResourcePermissionsRequest
	GetItems() []*RegisterAppResourcePermissionsRequestItems
	SetResourceTenantId(v string) *RegisterAppResourcePermissionsRequest
	GetResourceTenantId() *string
	SetTenantId(v string) *RegisterAppResourcePermissionsRequest
	GetTenantId() *string
}

type RegisterAppResourcePermissionsRequest struct {
	// 批量授权项
	//
	// This parameter is required.
	Items []*RegisterAppResourcePermissionsRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 资源所属租户 ID；省略为调用人当前租户
	//
	// example:
	//
	// 10000
	ResourceTenantId *string `json:"resourceTenantId,omitempty" xml:"resourceTenantId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RegisterAppResourcePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsRequest) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsRequest) GetItems() []*RegisterAppResourcePermissionsRequestItems {
	return s.Items
}

func (s *RegisterAppResourcePermissionsRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *RegisterAppResourcePermissionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RegisterAppResourcePermissionsRequest) SetItems(v []*RegisterAppResourcePermissionsRequestItems) *RegisterAppResourcePermissionsRequest {
	s.Items = v
	return s
}

func (s *RegisterAppResourcePermissionsRequest) SetResourceTenantId(v string) *RegisterAppResourcePermissionsRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsRequest) SetTenantId(v string) *RegisterAppResourcePermissionsRequest {
	s.TenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsRequest) Validate() error {
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

type RegisterAppResourcePermissionsRequestItems struct {
	// 带时区的到期时间；省略表示永久
	//
	// example:
	//
	// 2026-12-31T23:59:59+08:00
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// 被授权主体
	//
	// This parameter is required.
	Grantees []*RegisterAppResourcePermissionsRequestItemsGrantees `json:"grantees,omitempty" xml:"grantees,omitempty" type:"Repeated"`
	// 授权权限点
	//
	// This parameter is required.
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	// 首段为 APP 应用键的点分资源键
	//
	// This parameter is required.
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// 固定为 APP_RESOURCE
	//
	// This parameter is required.
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s RegisterAppResourcePermissionsRequestItems) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsRequestItems) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsRequestItems) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *RegisterAppResourcePermissionsRequestItems) GetGrantees() []*RegisterAppResourcePermissionsRequestItemsGrantees {
	return s.Grantees
}

func (s *RegisterAppResourcePermissionsRequestItems) GetPermissions() []*string {
	return s.Permissions
}

func (s *RegisterAppResourcePermissionsRequestItems) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *RegisterAppResourcePermissionsRequestItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *RegisterAppResourcePermissionsRequestItems) SetExpireAt(v string) *RegisterAppResourcePermissionsRequestItems {
	s.ExpireAt = &v
	return s
}

func (s *RegisterAppResourcePermissionsRequestItems) SetGrantees(v []*RegisterAppResourcePermissionsRequestItemsGrantees) *RegisterAppResourcePermissionsRequestItems {
	s.Grantees = v
	return s
}

func (s *RegisterAppResourcePermissionsRequestItems) SetPermissions(v []*string) *RegisterAppResourcePermissionsRequestItems {
	s.Permissions = v
	return s
}

func (s *RegisterAppResourcePermissionsRequestItems) SetResourceKey(v string) *RegisterAppResourcePermissionsRequestItems {
	s.ResourceKey = &v
	return s
}

func (s *RegisterAppResourcePermissionsRequestItems) SetResourceType(v string) *RegisterAppResourcePermissionsRequestItems {
	s.ResourceType = &v
	return s
}

func (s *RegisterAppResourcePermissionsRequestItems) Validate() error {
	if s.Grantees != nil {
		for _, item := range s.Grantees {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RegisterAppResourcePermissionsRequestItemsGrantees struct {
	// 主体 ID；ALL_USERS 省略
	//
	// example:
	//
	// 10086
	GranteeId *string `json:"granteeId,omitempty" xml:"granteeId,omitempty"`
	// 跨租户 USER 所属租户
	//
	// example:
	//
	// 10000
	GranteeTenantId *string `json:"granteeTenantId,omitempty" xml:"granteeTenantId,omitempty"`
	// 主体类型
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
}

func (s RegisterAppResourcePermissionsRequestItemsGrantees) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsRequestItemsGrantees) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsRequestItemsGrantees) GetGranteeId() *string {
	return s.GranteeId
}

func (s *RegisterAppResourcePermissionsRequestItemsGrantees) GetGranteeTenantId() *string {
	return s.GranteeTenantId
}

func (s *RegisterAppResourcePermissionsRequestItemsGrantees) GetGranteeType() *string {
	return s.GranteeType
}

func (s *RegisterAppResourcePermissionsRequestItemsGrantees) SetGranteeId(v string) *RegisterAppResourcePermissionsRequestItemsGrantees {
	s.GranteeId = &v
	return s
}

func (s *RegisterAppResourcePermissionsRequestItemsGrantees) SetGranteeTenantId(v string) *RegisterAppResourcePermissionsRequestItemsGrantees {
	s.GranteeTenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsRequestItemsGrantees) SetGranteeType(v string) *RegisterAppResourcePermissionsRequestItemsGrantees {
	s.GranteeType = &v
	return s
}

func (s *RegisterAppResourcePermissionsRequestItemsGrantees) Validate() error {
	return dara.Validate(s)
}
