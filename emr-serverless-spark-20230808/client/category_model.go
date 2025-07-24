// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategory interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *Category
	GetBizId() *string
	SetCreator(v int64) *Category
	GetCreator() *int64
	SetGmtCreated(v string) *Category
	GetGmtCreated() *string
	SetGmtModified(v string) *Category
	GetGmtModified() *string
	SetModifier(v int64) *Category
	GetModifier() *int64
	SetName(v string) *Category
	GetName() *string
	SetParentBizId(v string) *Category
	GetParentBizId() *string
	SetType(v string) *Category
	GetType() *string
}

type Category struct {
	// This parameter is required.
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	Creator *int64 `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Modifier *int64 `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// This parameter is required.
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentBizId *string `json:"parentBizId,omitempty" xml:"parentBizId,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Category) String() string {
	return dara.Prettify(s)
}

func (s Category) GoString() string {
	return s.String()
}

func (s *Category) GetBizId() *string {
	return s.BizId
}

func (s *Category) GetCreator() *int64 {
	return s.Creator
}

func (s *Category) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *Category) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Category) GetModifier() *int64 {
	return s.Modifier
}

func (s *Category) GetName() *string {
	return s.Name
}

func (s *Category) GetParentBizId() *string {
	return s.ParentBizId
}

func (s *Category) GetType() *string {
	return s.Type
}

func (s *Category) SetBizId(v string) *Category {
	s.BizId = &v
	return s
}

func (s *Category) SetCreator(v int64) *Category {
	s.Creator = &v
	return s
}

func (s *Category) SetGmtCreated(v string) *Category {
	s.GmtCreated = &v
	return s
}

func (s *Category) SetGmtModified(v string) *Category {
	s.GmtModified = &v
	return s
}

func (s *Category) SetModifier(v int64) *Category {
	s.Modifier = &v
	return s
}

func (s *Category) SetName(v string) *Category {
	s.Name = &v
	return s
}

func (s *Category) SetParentBizId(v string) *Category {
	s.ParentBizId = &v
	return s
}

func (s *Category) SetType(v string) *Category {
	s.Type = &v
	return s
}

func (s *Category) Validate() error {
	return dara.Validate(s)
}
