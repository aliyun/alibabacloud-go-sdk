// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryListQuery interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIds(v []*int64) *CategoryListQuery
	GetCategoryIds() []*int64
	SetParentCategoryId(v int64) *CategoryListQuery
	GetParentCategoryId() *int64
}

type CategoryListQuery struct {
	CategoryIds []*int64 `json:"categoryIds,omitempty" xml:"categoryIds,omitempty" type:"Repeated"`
	// example:
	//
	// 5200001
	ParentCategoryId *int64 `json:"parentCategoryId,omitempty" xml:"parentCategoryId,omitempty"`
}

func (s CategoryListQuery) String() string {
	return dara.Prettify(s)
}

func (s CategoryListQuery) GoString() string {
	return s.String()
}

func (s *CategoryListQuery) GetCategoryIds() []*int64 {
	return s.CategoryIds
}

func (s *CategoryListQuery) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *CategoryListQuery) SetCategoryIds(v []*int64) *CategoryListQuery {
	s.CategoryIds = v
	return s
}

func (s *CategoryListQuery) SetParentCategoryId(v int64) *CategoryListQuery {
	s.ParentCategoryId = &v
	return s
}

func (s *CategoryListQuery) Validate() error {
	return dara.Validate(s)
}
