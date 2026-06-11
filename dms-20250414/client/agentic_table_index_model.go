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
	// An array of column names included in the index.
	ColumnNames []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	// An optional, user-defined description for the index.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique name of the index within the table.
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The type of the index, such as PRIMARY, UNIQUE, or NORMAL.
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// Specifies if the index is the primary key. A table can have only one primary key.
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// An array of physical column names from the database. Use this parameter when the names in ColumnNames are aliases or logical names.
	RealColumnNames []*string `json:"RealColumnNames,omitempty" xml:"RealColumnNames,omitempty" type:"Repeated"`
	// Specifies if the index enforces a unique constraint, requiring all its values to be unique across rows.
	Unique *bool `json:"Unique,omitempty" xml:"Unique,omitempty"`
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
