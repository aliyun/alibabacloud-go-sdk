// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticTableIndex interface {
	dara.Model
	String() string
	GoString() string
	SetColumnNames(v []*string) *AgenticTableIndex
	GetColumnNames() []*string
	SetDescription(v string) *AgenticTableIndex
	GetDescription() *string
	SetIndexName(v string) *AgenticTableIndex
	GetIndexName() *string
	SetIndexType(v string) *AgenticTableIndex
	GetIndexType() *string
	SetPrimary(v bool) *AgenticTableIndex
	GetPrimary() *bool
	SetRealColumnNames(v []*string) *AgenticTableIndex
	GetRealColumnNames() []*string
	SetUnique(v bool) *AgenticTableIndex
	GetUnique() *bool
}

type AgenticTableIndex struct {
	ColumnNames     []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	IndexName       *string   `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	IndexType       *string   `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	Primary         *bool     `json:"Primary,omitempty" xml:"Primary,omitempty"`
	RealColumnNames []*string `json:"RealColumnNames,omitempty" xml:"RealColumnNames,omitempty" type:"Repeated"`
	Unique          *bool     `json:"Unique,omitempty" xml:"Unique,omitempty"`
}

func (s AgenticTableIndex) String() string {
	return dara.Prettify(s)
}

func (s AgenticTableIndex) GoString() string {
	return s.String()
}

func (s *AgenticTableIndex) GetColumnNames() []*string {
	return s.ColumnNames
}

func (s *AgenticTableIndex) GetDescription() *string {
	return s.Description
}

func (s *AgenticTableIndex) GetIndexName() *string {
	return s.IndexName
}

func (s *AgenticTableIndex) GetIndexType() *string {
	return s.IndexType
}

func (s *AgenticTableIndex) GetPrimary() *bool {
	return s.Primary
}

func (s *AgenticTableIndex) GetRealColumnNames() []*string {
	return s.RealColumnNames
}

func (s *AgenticTableIndex) GetUnique() *bool {
	return s.Unique
}

func (s *AgenticTableIndex) SetColumnNames(v []*string) *AgenticTableIndex {
	s.ColumnNames = v
	return s
}

func (s *AgenticTableIndex) SetDescription(v string) *AgenticTableIndex {
	s.Description = &v
	return s
}

func (s *AgenticTableIndex) SetIndexName(v string) *AgenticTableIndex {
	s.IndexName = &v
	return s
}

func (s *AgenticTableIndex) SetIndexType(v string) *AgenticTableIndex {
	s.IndexType = &v
	return s
}

func (s *AgenticTableIndex) SetPrimary(v bool) *AgenticTableIndex {
	s.Primary = &v
	return s
}

func (s *AgenticTableIndex) SetRealColumnNames(v []*string) *AgenticTableIndex {
	s.RealColumnNames = v
	return s
}

func (s *AgenticTableIndex) SetUnique(v bool) *AgenticTableIndex {
	s.Unique = &v
	return s
}

func (s *AgenticTableIndex) Validate() error {
	return dara.Validate(s)
}
