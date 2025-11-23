// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *UpdateMetaCategoryRequest
	GetCategoryId() *int64
	SetName(v string) *UpdateMetaCategoryRequest
	GetName() *string
	SetTid(v int64) *UpdateMetaCategoryRequest
	GetTid() *int64
}

type UpdateMetaCategoryRequest struct {
	// The category ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30000181325
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The updated name of the category.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 23****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateMetaCategoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMetaCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateMetaCategoryRequest) SetCategoryId(v int64) *UpdateMetaCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateMetaCategoryRequest) SetName(v string) *UpdateMetaCategoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateMetaCategoryRequest) SetTid(v int64) *UpdateMetaCategoryRequest {
	s.Tid = &v
	return s
}

func (s *UpdateMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
