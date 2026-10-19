// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResourcePermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemsShrink(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetItemsShrink() *string
	SetResourceTenantId(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetResourceTenantId() *string
	SetTenantId(v string) *DeleteAppResourcePermissionsShrinkRequest
	GetTenantId() *string
}

type DeleteAppResourcePermissionsShrinkRequest struct {
	// 撤权范围
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

func (s DeleteAppResourcePermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResourcePermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetItemsShrink() *string {
	return s.ItemsShrink
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *DeleteAppResourcePermissionsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetItemsShrink(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetResourceTenantId(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.ResourceTenantId = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) SetTenantId(v string) *DeleteAppResourcePermissionsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteAppResourcePermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
