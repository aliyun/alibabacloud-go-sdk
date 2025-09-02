// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddToMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *AddToMetaCategoryRequest
	GetCategoryId() *int64
	SetTableGuid(v string) *AddToMetaCategoryRequest
	GetTableGuid() *string
}

type AddToMetaCategoryRequest struct {
	// The ID of the category.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The GUID of the metatable.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s AddToMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddToMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddToMetaCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *AddToMetaCategoryRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *AddToMetaCategoryRequest) SetCategoryId(v int64) *AddToMetaCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *AddToMetaCategoryRequest) SetTableGuid(v string) *AddToMetaCategoryRequest {
	s.TableGuid = &v
	return s
}

func (s *AddToMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
