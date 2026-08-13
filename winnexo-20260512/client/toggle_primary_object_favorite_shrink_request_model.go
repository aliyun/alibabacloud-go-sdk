// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTogglePrimaryObjectFavoriteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetAction() *string
	SetObjectIdsShrink(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetObjectIdsShrink() *string
	SetObjectType(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetTenantId() *string
}

type TogglePrimaryObjectFavoriteShrinkRequest struct {
	// 操作：add-关注，remove-取消关注
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 主对象业务ID列表
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ObjectIdsShrink *string `json:"objectIds,omitempty" xml:"objectIds,omitempty"`
	// 对象类型（如 customer、project）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 运营对象名称（如 customer_1）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s TogglePrimaryObjectFavoriteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TogglePrimaryObjectFavoriteShrinkRequest) GoString() string {
	return s.String()
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetAction() *string {
	return s.Action
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetObjectIdsShrink() *string {
	return s.ObjectIdsShrink
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetAction(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.Action = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetObjectIdsShrink(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetObjectType(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetOperatingObjectName(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetTenantId(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
