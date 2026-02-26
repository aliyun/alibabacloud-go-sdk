// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategory interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *Category
	GetCategoryId() *int64
	SetIsLeaf(v bool) *Category
	GetIsLeaf() *bool
	SetLevel(v int32) *Category
	GetLevel() *int32
	SetName(v string) *Category
	GetName() *string
	SetParentId(v int64) *Category
	GetParentId() *int64
}

type Category struct {
	// example:
	//
	// 201792301
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// example:
	//
	// false
	IsLeaf *bool `json:"isLeaf,omitempty" xml:"isLeaf,omitempty"`
	// example:
	//
	// 0
	Level *int32  `json:"level,omitempty" xml:"level,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s Category) String() string {
	return dara.Prettify(s)
}

func (s Category) GoString() string {
	return s.String()
}

func (s *Category) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *Category) GetIsLeaf() *bool {
	return s.IsLeaf
}

func (s *Category) GetLevel() *int32 {
	return s.Level
}

func (s *Category) GetName() *string {
	return s.Name
}

func (s *Category) GetParentId() *int64 {
	return s.ParentId
}

func (s *Category) SetCategoryId(v int64) *Category {
	s.CategoryId = &v
	return s
}

func (s *Category) SetIsLeaf(v bool) *Category {
	s.IsLeaf = &v
	return s
}

func (s *Category) SetLevel(v int32) *Category {
	s.Level = &v
	return s
}

func (s *Category) SetName(v string) *Category {
	s.Name = &v
	return s
}

func (s *Category) SetParentId(v int64) *Category {
	s.ParentId = &v
	return s
}

func (s *Category) Validate() error {
	return dara.Validate(s)
}
