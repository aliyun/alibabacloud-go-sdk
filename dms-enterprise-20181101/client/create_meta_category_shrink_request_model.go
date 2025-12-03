// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCategoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateMetaCategoryShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateMetaCategoryShrinkRequest
	GetName() *string
	SetOwnerIdsShrink(v string) *CreateMetaCategoryShrinkRequest
	GetOwnerIdsShrink() *string
	SetParentCategoryId(v int64) *CreateMetaCategoryShrinkRequest
	GetParentCategoryId() *int64
	SetRemark(v string) *CreateMetaCategoryShrinkRequest
	GetRemark() *string
	SetTid(v int64) *CreateMetaCategoryShrinkRequest
	GetTid() *int64
}

type CreateMetaCategoryShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the category.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerIdsShrink *string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty"`
	// The ID of the parent category. The new category is created under this parent category. If this value is left empty, the new category is of the first level.
	//
	// example:
	//
	// 30000322682
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	Remark           *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateMetaCategoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCategoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaCategoryShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMetaCategoryShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetaCategoryShrinkRequest) GetOwnerIdsShrink() *string {
	return s.OwnerIdsShrink
}

func (s *CreateMetaCategoryShrinkRequest) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *CreateMetaCategoryShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateMetaCategoryShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateMetaCategoryShrinkRequest) SetDescription(v string) *CreateMetaCategoryShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateMetaCategoryShrinkRequest) SetName(v string) *CreateMetaCategoryShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateMetaCategoryShrinkRequest) SetOwnerIdsShrink(v string) *CreateMetaCategoryShrinkRequest {
	s.OwnerIdsShrink = &v
	return s
}

func (s *CreateMetaCategoryShrinkRequest) SetParentCategoryId(v int64) *CreateMetaCategoryShrinkRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *CreateMetaCategoryShrinkRequest) SetRemark(v string) *CreateMetaCategoryShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateMetaCategoryShrinkRequest) SetTid(v int64) *CreateMetaCategoryShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateMetaCategoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
