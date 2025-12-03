// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCategoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *UpdateMetaCategoryShrinkRequest
	GetCategoryId() *int64
	SetDescription(v string) *UpdateMetaCategoryShrinkRequest
	GetDescription() *string
	SetName(v string) *UpdateMetaCategoryShrinkRequest
	GetName() *string
	SetOwnerIdsShrink(v string) *UpdateMetaCategoryShrinkRequest
	GetOwnerIdsShrink() *string
	SetRemark(v string) *UpdateMetaCategoryShrinkRequest
	GetRemark() *string
	SetTid(v int64) *UpdateMetaCategoryShrinkRequest
	GetTid() *int64
}

type UpdateMetaCategoryShrinkRequest struct {
	// The category ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30000181325
	CategoryId  *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The updated name of the category.
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerIdsShrink *string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 23****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateMetaCategoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCategoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaCategoryShrinkRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateMetaCategoryShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMetaCategoryShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMetaCategoryShrinkRequest) GetOwnerIdsShrink() *string {
	return s.OwnerIdsShrink
}

func (s *UpdateMetaCategoryShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateMetaCategoryShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateMetaCategoryShrinkRequest) SetCategoryId(v int64) *UpdateMetaCategoryShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateMetaCategoryShrinkRequest) SetDescription(v string) *UpdateMetaCategoryShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateMetaCategoryShrinkRequest) SetName(v string) *UpdateMetaCategoryShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateMetaCategoryShrinkRequest) SetOwnerIdsShrink(v string) *UpdateMetaCategoryShrinkRequest {
	s.OwnerIdsShrink = &v
	return s
}

func (s *UpdateMetaCategoryShrinkRequest) SetRemark(v string) *UpdateMetaCategoryShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateMetaCategoryShrinkRequest) SetTid(v int64) *UpdateMetaCategoryShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateMetaCategoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
