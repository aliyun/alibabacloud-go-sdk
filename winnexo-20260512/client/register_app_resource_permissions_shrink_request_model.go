// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAppResourcePermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemsShrink(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetItemsShrink() *string
	SetResourceTenantId(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetResourceTenantId() *string
	SetTenantId(v string) *RegisterAppResourcePermissionsShrinkRequest
	GetTenantId() *string
}

type RegisterAppResourcePermissionsShrinkRequest struct {
	// 批量授权项
	//
	// This parameter is required.
	ItemsShrink *string `json:"items,omitempty" xml:"items,omitempty"`
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

func (s RegisterAppResourcePermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetItemsShrink() *string {
	return s.ItemsShrink
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *RegisterAppResourcePermissionsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetItemsShrink(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetResourceTenantId(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) SetTenantId(v string) *RegisterAppResourcePermissionsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *RegisterAppResourcePermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
