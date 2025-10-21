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
	ColumnName       *string `json:"columnName,omitempty" xml:"columnName,omitempty"`
	ColumnNativeType *string `json:"columnNativeType,omitempty" xml:"columnNativeType,omitempty"`
	ColumnType       *string `json:"columnType,omitempty" xml:"columnType,omitempty"`
	CreatedAt        *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Creator          *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	Id               *string `json:"id,omitempty" xml:"id,omitempty"`
	ModifiedAt       *int64  `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Modifier         *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Nullable         *bool   `json:"nullable,omitempty" xml:"nullable,omitempty"`
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
