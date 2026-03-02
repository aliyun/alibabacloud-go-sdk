// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageColumn interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *LineageColumn
	GetColumnName() *string
	SetColumnNativeType(v string) *LineageColumn
	GetColumnNativeType() *string
	SetColumnType(v string) *LineageColumn
	GetColumnType() *string
	SetCreatedAt(v int64) *LineageColumn
	GetCreatedAt() *int64
	SetCreator(v string) *LineageColumn
	GetCreator() *string
	SetDescription(v string) *LineageColumn
	GetDescription() *string
	SetId(v string) *LineageColumn
	GetId() *string
	SetModifiedAt(v int64) *LineageColumn
	GetModifiedAt() *int64
	SetModifier(v string) *LineageColumn
	GetModifier() *string
	SetNullable(v bool) *LineageColumn
	GetNullable() *bool
}

type LineageColumn struct {
	// The name of the column.
	//
	// example:
	//
	// name
	ColumnName *string `json:"columnName,omitempty" xml:"columnName,omitempty"`
	// The original type of the column.
	//
	// example:
	//
	// string
	ColumnNativeType *string `json:"columnNativeType,omitempty" xml:"columnNativeType,omitempty"`
	// The column type.
	//
	// example:
	//
	// string
	ColumnType *string `json:"columnType,omitempty" xml:"columnType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1723537978
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The user that creates the column.
	//
	// example:
	//
	// userA
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The description.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The field ID.
	//
	// example:
	//
	// 123ss3
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1723537978
	ModifiedAt *int64 `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// The ID of the account that is used to modify the column.
	//
	// example:
	//
	// userA
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Indicates whether the value is nullable.
	//
	// example:
	//
	// false
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
}

func (s LineageColumn) String() string {
	return dara.Prettify(s)
}

func (s LineageColumn) GoString() string {
	return s.String()
}

func (s *LineageColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *LineageColumn) GetColumnNativeType() *string {
	return s.ColumnNativeType
}

func (s *LineageColumn) GetColumnType() *string {
	return s.ColumnType
}

func (s *LineageColumn) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *LineageColumn) GetCreator() *string {
	return s.Creator
}

func (s *LineageColumn) GetDescription() *string {
	return s.Description
}

func (s *LineageColumn) GetId() *string {
	return s.Id
}

func (s *LineageColumn) GetModifiedAt() *int64 {
	return s.ModifiedAt
}

func (s *LineageColumn) GetModifier() *string {
	return s.Modifier
}

func (s *LineageColumn) GetNullable() *bool {
	return s.Nullable
}

func (s *LineageColumn) SetColumnName(v string) *LineageColumn {
	s.ColumnName = &v
	return s
}

func (s *LineageColumn) SetColumnNativeType(v string) *LineageColumn {
	s.ColumnNativeType = &v
	return s
}

func (s *LineageColumn) SetColumnType(v string) *LineageColumn {
	s.ColumnType = &v
	return s
}

func (s *LineageColumn) SetCreatedAt(v int64) *LineageColumn {
	s.CreatedAt = &v
	return s
}

func (s *LineageColumn) SetCreator(v string) *LineageColumn {
	s.Creator = &v
	return s
}

func (s *LineageColumn) SetDescription(v string) *LineageColumn {
	s.Description = &v
	return s
}

func (s *LineageColumn) SetId(v string) *LineageColumn {
	s.Id = &v
	return s
}

func (s *LineageColumn) SetModifiedAt(v int64) *LineageColumn {
	s.ModifiedAt = &v
	return s
}

func (s *LineageColumn) SetModifier(v string) *LineageColumn {
	s.Modifier = &v
	return s
}

func (s *LineageColumn) SetNullable(v bool) *LineageColumn {
	s.Nullable = &v
	return s
}

func (s *LineageColumn) Validate() error {
	return dara.Validate(s)
}
