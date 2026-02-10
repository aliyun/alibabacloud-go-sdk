// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaTableIndex interface {
	dara.Model
	String() string
	GoString() string
	SetColumnNames(v []*string) *OneMetaTableIndex
	GetColumnNames() []*string
	SetDescription(v string) *OneMetaTableIndex
	GetDescription() *string
	SetIndexName(v string) *OneMetaTableIndex
	GetIndexName() *string
	SetIndexType(v string) *OneMetaTableIndex
	GetIndexType() *string
	SetPrimary(v bool) *OneMetaTableIndex
	GetPrimary() *bool
	SetRealColumnNames(v []*string) *OneMetaTableIndex
	GetRealColumnNames() []*string
	SetUnique(v bool) *OneMetaTableIndex
	GetUnique() *bool
}

type OneMetaTableIndex struct {
	ColumnNames     []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	IndexName       *string   `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	IndexType       *string   `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	Primary         *bool     `json:"Primary,omitempty" xml:"Primary,omitempty"`
	RealColumnNames []*string `json:"RealColumnNames,omitempty" xml:"RealColumnNames,omitempty" type:"Repeated"`
	Unique          *bool     `json:"Unique,omitempty" xml:"Unique,omitempty"`
}

func (s OneMetaTableIndex) String() string {
	return dara.Prettify(s)
}

func (s OneMetaTableIndex) GoString() string {
	return s.String()
}

func (s *OneMetaTableIndex) GetColumnNames() []*string {
	return s.ColumnNames
}

func (s *OneMetaTableIndex) GetDescription() *string {
	return s.Description
}

func (s *OneMetaTableIndex) GetIndexName() *string {
	return s.IndexName
}

func (s *OneMetaTableIndex) GetIndexType() *string {
	return s.IndexType
}

func (s *OneMetaTableIndex) GetPrimary() *bool {
	return s.Primary
}

func (s *OneMetaTableIndex) GetRealColumnNames() []*string {
	return s.RealColumnNames
}

func (s *OneMetaTableIndex) GetUnique() *bool {
	return s.Unique
}

func (s *OneMetaTableIndex) SetColumnNames(v []*string) *OneMetaTableIndex {
	s.ColumnNames = v
	return s
}

func (s *OneMetaTableIndex) SetDescription(v string) *OneMetaTableIndex {
	s.Description = &v
	return s
}

func (s *OneMetaTableIndex) SetIndexName(v string) *OneMetaTableIndex {
	s.IndexName = &v
	return s
}

func (s *OneMetaTableIndex) SetIndexType(v string) *OneMetaTableIndex {
	s.IndexType = &v
	return s
}

func (s *OneMetaTableIndex) SetPrimary(v bool) *OneMetaTableIndex {
	s.Primary = &v
	return s
}

func (s *OneMetaTableIndex) SetRealColumnNames(v []*string) *OneMetaTableIndex {
	s.RealColumnNames = v
	return s
}

func (s *OneMetaTableIndex) SetUnique(v bool) *OneMetaTableIndex {
	s.Unique = &v
	return s
}

func (s *OneMetaTableIndex) Validate() error {
	return dara.Validate(s)
}
