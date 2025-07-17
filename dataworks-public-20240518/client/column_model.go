// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColumn interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessMetadata(v *ColumnBusinessMetadata) *Column
	GetBusinessMetadata() *ColumnBusinessMetadata
	SetComment(v string) *Column
	GetComment() *string
	SetForeignKey(v bool) *Column
	GetForeignKey() *bool
	SetId(v string) *Column
	GetId() *string
	SetName(v string) *Column
	GetName() *string
	SetPartitionKey(v bool) *Column
	GetPartitionKey() *bool
	SetPosition(v int32) *Column
	GetPosition() *int32
	SetPrimaryKey(v bool) *Column
	GetPrimaryKey() *bool
	SetTableId(v string) *Column
	GetTableId() *string
	SetType(v string) *Column
	GetType() *string
}

type Column struct {
	BusinessMetadata *ColumnBusinessMetadata `json:"BusinessMetadata,omitempty" xml:"BusinessMetadata,omitempty" type:"Struct"`
	// example:
	//
	// 字段1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// false
	ForeignKey *bool `json:"ForeignKey,omitempty" xml:"ForeignKey,omitempty"`
	// example:
	//
	// maxcompute-column:123456::test_project:default:test_tbl:col1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// col1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	PartitionKey *bool `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// false
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// example:
	//
	// maxcompute-table:123456::test_project:default:test_tbl
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// example:
	//
	// bigint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Column) String() string {
	return dara.Prettify(s)
}

func (s Column) GoString() string {
	return s.String()
}

func (s *Column) GetBusinessMetadata() *ColumnBusinessMetadata {
	return s.BusinessMetadata
}

func (s *Column) GetComment() *string {
	return s.Comment
}

func (s *Column) GetForeignKey() *bool {
	return s.ForeignKey
}

func (s *Column) GetId() *string {
	return s.Id
}

func (s *Column) GetName() *string {
	return s.Name
}

func (s *Column) GetPartitionKey() *bool {
	return s.PartitionKey
}

func (s *Column) GetPosition() *int32 {
	return s.Position
}

func (s *Column) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *Column) GetTableId() *string {
	return s.TableId
}

func (s *Column) GetType() *string {
	return s.Type
}

func (s *Column) SetBusinessMetadata(v *ColumnBusinessMetadata) *Column {
	s.BusinessMetadata = v
	return s
}

func (s *Column) SetComment(v string) *Column {
	s.Comment = &v
	return s
}

func (s *Column) SetForeignKey(v bool) *Column {
	s.ForeignKey = &v
	return s
}

func (s *Column) SetId(v string) *Column {
	s.Id = &v
	return s
}

func (s *Column) SetName(v string) *Column {
	s.Name = &v
	return s
}

func (s *Column) SetPartitionKey(v bool) *Column {
	s.PartitionKey = &v
	return s
}

func (s *Column) SetPosition(v int32) *Column {
	s.Position = &v
	return s
}

func (s *Column) SetPrimaryKey(v bool) *Column {
	s.PrimaryKey = &v
	return s
}

func (s *Column) SetTableId(v string) *Column {
	s.TableId = &v
	return s
}

func (s *Column) SetType(v string) *Column {
	s.Type = &v
	return s
}

func (s *Column) Validate() error {
	return dara.Validate(s)
}

type ColumnBusinessMetadata struct {
	// example:
	//
	// 字段1的业务描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ColumnBusinessMetadata) String() string {
	return dara.Prettify(s)
}

func (s ColumnBusinessMetadata) GoString() string {
	return s.String()
}

func (s *ColumnBusinessMetadata) GetDescription() *string {
	return s.Description
}

func (s *ColumnBusinessMetadata) SetDescription(v string) *ColumnBusinessMetadata {
	s.Description = &v
	return s
}

func (s *ColumnBusinessMetadata) Validate() error {
	return dara.Validate(s)
}
