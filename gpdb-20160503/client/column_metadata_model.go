// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColumnMetadata interface {
	dara.Model
	String() string
	GoString() string
	SetColumnDefault(v string) *ColumnMetadata
	GetColumnDefault() *string
	SetComment(v string) *ColumnMetadata
	GetComment() *string
	SetDataType(v string) *ColumnMetadata
	GetDataType() *string
	SetIsCaseSensitive(v bool) *ColumnMetadata
	GetIsCaseSensitive() *bool
	SetIsCurrency(v bool) *ColumnMetadata
	GetIsCurrency() *bool
	SetIsPrimaryKey(v bool) *ColumnMetadata
	GetIsPrimaryKey() *bool
	SetIsSigned(v bool) *ColumnMetadata
	GetIsSigned() *bool
	SetMaxLength(v int32) *ColumnMetadata
	GetMaxLength() *int32
	SetName(v string) *ColumnMetadata
	GetName() *string
	SetNullable(v bool) *ColumnMetadata
	GetNullable() *bool
	SetPrecision(v int32) *ColumnMetadata
	GetPrecision() *int32
	SetScale(v int32) *ColumnMetadata
	GetScale() *int32
	SetSchemaName(v string) *ColumnMetadata
	GetSchemaName() *string
	SetTableName(v string) *ColumnMetadata
	GetTableName() *string
	SetUdtName(v string) *ColumnMetadata
	GetUdtName() *string
}

type ColumnMetadata struct {
	// The default value.
	ColumnDefault *string `json:"ColumnDefault,omitempty" xml:"ColumnDefault,omitempty"`
	// The column description.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The data type.
	//
	// example:
	//
	// text
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The case sensitivity.
	IsCaseSensitive *bool `json:"IsCaseSensitive,omitempty" xml:"IsCaseSensitive,omitempty"`
	// The currency value indicator.
	IsCurrency *bool `json:"IsCurrency,omitempty" xml:"IsCurrency,omitempty"`
	// The primary key indicator.
	IsPrimaryKey *bool `json:"IsPrimaryKey,omitempty" xml:"IsPrimaryKey,omitempty"`
	// The signed indicator.
	IsSigned *bool `json:"IsSigned,omitempty" xml:"IsSigned,omitempty"`
	// The maximum character length.
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The column name.
	//
	// example:
	//
	// content
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The nullability.
	Nullable *bool `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// The precision of the column, which indicates the total number of digits in the integer and fractional parts of a decimal number.
	Precision *int32 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	// The scale of the column, which indicates the number of digits in the fractional part of a decimal number.
	Scale *int32 `json:"Scale,omitempty" xml:"Scale,omitempty"`
	// The schema name.
	//
	// example:
	//
	// ns1
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// dc1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The name of the base data type used by the column.
	UdtName *string `json:"UdtName,omitempty" xml:"UdtName,omitempty"`
}

func (s ColumnMetadata) String() string {
	return dara.Prettify(s)
}

func (s ColumnMetadata) GoString() string {
	return s.String()
}

func (s *ColumnMetadata) GetColumnDefault() *string {
	return s.ColumnDefault
}

func (s *ColumnMetadata) GetComment() *string {
	return s.Comment
}

func (s *ColumnMetadata) GetDataType() *string {
	return s.DataType
}

func (s *ColumnMetadata) GetIsCaseSensitive() *bool {
	return s.IsCaseSensitive
}

func (s *ColumnMetadata) GetIsCurrency() *bool {
	return s.IsCurrency
}

func (s *ColumnMetadata) GetIsPrimaryKey() *bool {
	return s.IsPrimaryKey
}

func (s *ColumnMetadata) GetIsSigned() *bool {
	return s.IsSigned
}

func (s *ColumnMetadata) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *ColumnMetadata) GetName() *string {
	return s.Name
}

func (s *ColumnMetadata) GetNullable() *bool {
	return s.Nullable
}

func (s *ColumnMetadata) GetPrecision() *int32 {
	return s.Precision
}

func (s *ColumnMetadata) GetScale() *int32 {
	return s.Scale
}

func (s *ColumnMetadata) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ColumnMetadata) GetTableName() *string {
	return s.TableName
}

func (s *ColumnMetadata) GetUdtName() *string {
	return s.UdtName
}

func (s *ColumnMetadata) SetColumnDefault(v string) *ColumnMetadata {
	s.ColumnDefault = &v
	return s
}

func (s *ColumnMetadata) SetComment(v string) *ColumnMetadata {
	s.Comment = &v
	return s
}

func (s *ColumnMetadata) SetDataType(v string) *ColumnMetadata {
	s.DataType = &v
	return s
}

func (s *ColumnMetadata) SetIsCaseSensitive(v bool) *ColumnMetadata {
	s.IsCaseSensitive = &v
	return s
}

func (s *ColumnMetadata) SetIsCurrency(v bool) *ColumnMetadata {
	s.IsCurrency = &v
	return s
}

func (s *ColumnMetadata) SetIsPrimaryKey(v bool) *ColumnMetadata {
	s.IsPrimaryKey = &v
	return s
}

func (s *ColumnMetadata) SetIsSigned(v bool) *ColumnMetadata {
	s.IsSigned = &v
	return s
}

func (s *ColumnMetadata) SetMaxLength(v int32) *ColumnMetadata {
	s.MaxLength = &v
	return s
}

func (s *ColumnMetadata) SetName(v string) *ColumnMetadata {
	s.Name = &v
	return s
}

func (s *ColumnMetadata) SetNullable(v bool) *ColumnMetadata {
	s.Nullable = &v
	return s
}

func (s *ColumnMetadata) SetPrecision(v int32) *ColumnMetadata {
	s.Precision = &v
	return s
}

func (s *ColumnMetadata) SetScale(v int32) *ColumnMetadata {
	s.Scale = &v
	return s
}

func (s *ColumnMetadata) SetSchemaName(v string) *ColumnMetadata {
	s.SchemaName = &v
	return s
}

func (s *ColumnMetadata) SetTableName(v string) *ColumnMetadata {
	s.TableName = &v
	return s
}

func (s *ColumnMetadata) SetUdtName(v string) *ColumnMetadata {
	s.UdtName = &v
	return s
}

func (s *ColumnMetadata) Validate() error {
	return dara.Validate(s)
}
