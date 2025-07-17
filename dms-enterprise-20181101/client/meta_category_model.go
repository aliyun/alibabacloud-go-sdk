// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaCategory interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *MetaCategory
	GetCategoryId() *int64
	SetCreateTime(v string) *MetaCategory
	GetCreateTime() *string
	SetDepth(v int32) *MetaCategory
	GetDepth() *int32
	SetName(v string) *MetaCategory
	GetName() *string
	SetParentCategoryId(v int64) *MetaCategory
	GetParentCategoryId() *int64
}

type MetaCategory struct {
	CategoryId       *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Depth            *int32  `json:"Depth,omitempty" xml:"Depth,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCategoryId *int64  `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s MetaCategory) String() string {
	return dara.Prettify(s)
}

func (s MetaCategory) GoString() string {
	return s.String()
}

func (s *MetaCategory) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *MetaCategory) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MetaCategory) GetDepth() *int32 {
	return s.Depth
}

func (s *MetaCategory) GetName() *string {
	return s.Name
}

func (s *MetaCategory) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *MetaCategory) SetCategoryId(v int64) *MetaCategory {
	s.CategoryId = &v
	return s
}

func (s *MetaCategory) SetCreateTime(v string) *MetaCategory {
	s.CreateTime = &v
	return s
}

func (s *MetaCategory) SetDepth(v int32) *MetaCategory {
	s.Depth = &v
	return s
}

func (s *MetaCategory) SetName(v string) *MetaCategory {
	s.Name = &v
	return s
}

func (s *MetaCategory) SetParentCategoryId(v int64) *MetaCategory {
	s.ParentCategoryId = &v
	return s
}

func (s *MetaCategory) Validate() error {
	return dara.Validate(s)
}
