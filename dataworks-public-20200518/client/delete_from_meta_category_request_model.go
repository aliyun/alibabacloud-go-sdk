// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFromMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *DeleteFromMetaCategoryRequest
	GetCategoryId() *int64
	SetTableGuid(v string) *DeleteFromMetaCategoryRequest
	GetTableGuid() *string
}

type DeleteFromMetaCategoryRequest struct {
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

func (s DeleteFromMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFromMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteFromMetaCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DeleteFromMetaCategoryRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *DeleteFromMetaCategoryRequest) SetCategoryId(v int64) *DeleteFromMetaCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *DeleteFromMetaCategoryRequest) SetTableGuid(v string) *DeleteFromMetaCategoryRequest {
	s.TableGuid = &v
	return s
}

func (s *DeleteFromMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
