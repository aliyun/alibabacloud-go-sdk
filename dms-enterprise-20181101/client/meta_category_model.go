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
	SetDescription(v string) *MetaCategory
	GetDescription() *string
	SetName(v string) *MetaCategory
	GetName() *string
	SetOwnerIds(v []*int64) *MetaCategory
	GetOwnerIds() []*int64
	SetOwnerNickNames(v []*string) *MetaCategory
	GetOwnerNickNames() []*string
	SetParentCategoryId(v int64) *MetaCategory
	GetParentCategoryId() *int64
	SetRemark(v string) *MetaCategory
	GetRemark() *string
}

type MetaCategory struct {
	CategoryId       *int64    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreateTime       *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Depth            *int32    `json:"Depth,omitempty" xml:"Depth,omitempty"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerIds         []*int64  `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
	OwnerNickNames   []*string `json:"OwnerNickNames,omitempty" xml:"OwnerNickNames,omitempty" type:"Repeated"`
	ParentCategoryId *int64    `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	Remark           *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
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

func (s *MetaCategory) GetDescription() *string {
	return s.Description
}

func (s *MetaCategory) GetName() *string {
	return s.Name
}

func (s *MetaCategory) GetOwnerIds() []*int64 {
	return s.OwnerIds
}

func (s *MetaCategory) GetOwnerNickNames() []*string {
	return s.OwnerNickNames
}

func (s *MetaCategory) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *MetaCategory) GetRemark() *string {
	return s.Remark
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

func (s *MetaCategory) SetDescription(v string) *MetaCategory {
	s.Description = &v
	return s
}

func (s *MetaCategory) SetName(v string) *MetaCategory {
	s.Name = &v
	return s
}

func (s *MetaCategory) SetOwnerIds(v []*int64) *MetaCategory {
	s.OwnerIds = v
	return s
}

func (s *MetaCategory) SetOwnerNickNames(v []*string) *MetaCategory {
	s.OwnerNickNames = v
	return s
}

func (s *MetaCategory) SetParentCategoryId(v int64) *MetaCategory {
	s.ParentCategoryId = &v
	return s
}

func (s *MetaCategory) SetRemark(v string) *MetaCategory {
	s.Remark = &v
	return s
}

func (s *MetaCategory) Validate() error {
	return dara.Validate(s)
}
