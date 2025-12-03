// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateMetaCategoryRequest
	GetDescription() *string
	SetName(v string) *CreateMetaCategoryRequest
	GetName() *string
	SetOwnerIds(v []*int64) *CreateMetaCategoryRequest
	GetOwnerIds() []*int64
	SetParentCategoryId(v int64) *CreateMetaCategoryRequest
	GetParentCategoryId() *int64
	SetRemark(v string) *CreateMetaCategoryRequest
	GetRemark() *string
	SetTid(v int64) *CreateMetaCategoryRequest
	GetTid() *int64
}

type CreateMetaCategoryRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the category.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Name     *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerIds []*int64 `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
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

func (s CreateMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaCategoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMetaCategoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetaCategoryRequest) GetOwnerIds() []*int64 {
	return s.OwnerIds
}

func (s *CreateMetaCategoryRequest) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *CreateMetaCategoryRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateMetaCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateMetaCategoryRequest) SetDescription(v string) *CreateMetaCategoryRequest {
	s.Description = &v
	return s
}

func (s *CreateMetaCategoryRequest) SetName(v string) *CreateMetaCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateMetaCategoryRequest) SetOwnerIds(v []*int64) *CreateMetaCategoryRequest {
	s.OwnerIds = v
	return s
}

func (s *CreateMetaCategoryRequest) SetParentCategoryId(v int64) *CreateMetaCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *CreateMetaCategoryRequest) SetRemark(v string) *CreateMetaCategoryRequest {
	s.Remark = &v
	return s
}

func (s *CreateMetaCategoryRequest) SetTid(v int64) *CreateMetaCategoryRequest {
	s.Tid = &v
	return s
}

func (s *CreateMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
