// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFieldSchemaModel interface {
	dara.Model
	String() string
	GoString() string
	SetAutoIncrement(v bool) *FieldSchemaModel
	GetAutoIncrement() *bool
	SetColumnRawName(v string) *FieldSchemaModel
	GetColumnRawName() *string
	SetComment(v string) *FieldSchemaModel
	GetComment() *string
	SetCompressFloatUseShort(v bool) *FieldSchemaModel
	GetCompressFloatUseShort() *bool
	SetCompression(v string) *FieldSchemaModel
	GetCompression() *string
	SetCreateTime(v string) *FieldSchemaModel
	GetCreateTime() *string
	SetDataType(v string) *FieldSchemaModel
	GetDataType() *string
	SetDatabaseName(v string) *FieldSchemaModel
	GetDatabaseName() *string
	SetDefaultValue(v string) *FieldSchemaModel
	GetDefaultValue() *string
	SetDelimiter(v string) *FieldSchemaModel
	GetDelimiter() *string
	SetEncode(v string) *FieldSchemaModel
	GetEncode() *string
	SetIsPartitionKey(v bool) *FieldSchemaModel
	GetIsPartitionKey() *bool
	SetMappedName(v string) *FieldSchemaModel
	GetMappedName() *string
	SetName(v string) *FieldSchemaModel
	GetName() *string
	SetNullable(v bool) *FieldSchemaModel
	GetNullable() *bool
	SetOnUpdate(v string) *FieldSchemaModel
	GetOnUpdate() *string
	SetOrdinalPosition(v int64) *FieldSchemaModel
	GetOrdinalPosition() *int64
	SetPhysicalColumnName(v string) *FieldSchemaModel
	GetPhysicalColumnName() *string
	SetPkPosition(v int64) *FieldSchemaModel
	GetPkPosition() *int64
	SetPrecision(v int64) *FieldSchemaModel
	GetPrecision() *int64
	SetPrimarykey(v bool) *FieldSchemaModel
	GetPrimarykey() *bool
	SetScale(v int64) *FieldSchemaModel
	GetScale() *int64
	SetTableName(v string) *FieldSchemaModel
	GetTableName() *string
	SetTokenizer(v string) *FieldSchemaModel
	GetTokenizer() *string
	SetType(v string) *FieldSchemaModel
	GetType() *string
	SetUpdateTime(v string) *FieldSchemaModel
	GetUpdateTime() *string
	SetValueType(v string) *FieldSchemaModel
	GetValueType() *string
}

type FieldSchemaModel struct {
	// Indicates whether the column is auto-incremented.
	//
	// example:
	//
	// true
	AutoIncrement *bool `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	// The original name of the column.
	//
	// example:
	//
	// ColumnRawName
	ColumnRawName *string `json:"ColumnRawName,omitempty" xml:"ColumnRawName,omitempty"`
	// The description of the column.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Indicates whether FLOAT data is compressed to SHORT data.
	//
	// example:
	//
	// false
	CompressFloatUseShort *bool `json:"CompressFloatUseShort,omitempty" xml:"CompressFloatUseShort,omitempty"`
	// The compression method of the column.
	//
	// example:
	//
	// compression
	Compression *string `json:"Compression,omitempty" xml:"Compression,omitempty"`
	// The time when the column was created.
	//
	// example:
	//
	// 2023-01-05 13:17:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// long
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The logical name of the database.
	//
	// example:
	//
	// databaseName
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The default value of the column.
	//
	// example:
	//
	// default
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The delimiter of the column.
	//
	// example:
	//
	// delimiter
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// The encryption method of the column.
	//
	// example:
	//
	// encode
	Encode *string `json:"Encode,omitempty" xml:"Encode,omitempty"`
	// Indicates whether the column is the partition key.
	//
	// example:
	//
	// false
	IsPartitionKey *bool `json:"IsPartitionKey,omitempty" xml:"IsPartitionKey,omitempty"`
	// The mapping name.
	//
	// example:
	//
	// mappedName
	MappedName *string `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the column can be empty.
	//
	// example:
	//
	// true
	Nullable *bool `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// The update condition of the column.
	//
	// example:
	//
	// onUpdate
	OnUpdate *string `json:"OnUpdate,omitempty" xml:"OnUpdate,omitempty"`
	// The location of the column.
	//
	// example:
	//
	// -1
	OrdinalPosition *int64 `json:"OrdinalPosition,omitempty" xml:"OrdinalPosition,omitempty"`
	// The physical name of the column.
	//
	// example:
	//
	// PhysicalColumnName
	PhysicalColumnName *string `json:"PhysicalColumnName,omitempty" xml:"PhysicalColumnName,omitempty"`
	// The location of the primary key.
	//
	// example:
	//
	// -1
	PkPosition *int64 `json:"PkPosition,omitempty" xml:"PkPosition,omitempty"`
	// The precision of the column.
	//
	// example:
	//
	// 1
	Precision *int64 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	// Indicates whether the column is the primary key.
	//
	// example:
	//
	// true
	Primarykey *bool `json:"Primarykey,omitempty" xml:"Primarykey,omitempty"`
	// The scale of the column.
	//
	// example:
	//
	// 1
	Scale *int64 `json:"Scale,omitempty" xml:"Scale,omitempty"`
	// The logical name of the table.
	//
	// example:
	//
	// tableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The token of the column.
	//
	// example:
	//
	// tokenizer
	Tokenizer *string `json:"Tokenizer,omitempty" xml:"Tokenizer,omitempty"`
	// The type of the column.
	//
	// example:
	//
	// long
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the index was updated.
	//
	// example:
	//
	// 2023-01-05 13:17:55
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The value type of the column.
	//
	// example:
	//
	// valueType
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s FieldSchemaModel) String() string {
	return dara.Prettify(s)
}

func (s FieldSchemaModel) GoString() string {
	return s.String()
}

func (s *FieldSchemaModel) GetAutoIncrement() *bool {
	return s.AutoIncrement
}

func (s *FieldSchemaModel) GetColumnRawName() *string {
	return s.ColumnRawName
}

func (s *FieldSchemaModel) GetComment() *string {
	return s.Comment
}

func (s *FieldSchemaModel) GetCompressFloatUseShort() *bool {
	return s.CompressFloatUseShort
}

func (s *FieldSchemaModel) GetCompression() *string {
	return s.Compression
}

func (s *FieldSchemaModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *FieldSchemaModel) GetDataType() *string {
	return s.DataType
}

func (s *FieldSchemaModel) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *FieldSchemaModel) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *FieldSchemaModel) GetDelimiter() *string {
	return s.Delimiter
}

func (s *FieldSchemaModel) GetEncode() *string {
	return s.Encode
}

func (s *FieldSchemaModel) GetIsPartitionKey() *bool {
	return s.IsPartitionKey
}

func (s *FieldSchemaModel) GetMappedName() *string {
	return s.MappedName
}

func (s *FieldSchemaModel) GetName() *string {
	return s.Name
}

func (s *FieldSchemaModel) GetNullable() *bool {
	return s.Nullable
}

func (s *FieldSchemaModel) GetOnUpdate() *string {
	return s.OnUpdate
}

func (s *FieldSchemaModel) GetOrdinalPosition() *int64 {
	return s.OrdinalPosition
}

func (s *FieldSchemaModel) GetPhysicalColumnName() *string {
	return s.PhysicalColumnName
}

func (s *FieldSchemaModel) GetPkPosition() *int64 {
	return s.PkPosition
}

func (s *FieldSchemaModel) GetPrecision() *int64 {
	return s.Precision
}

func (s *FieldSchemaModel) GetPrimarykey() *bool {
	return s.Primarykey
}

func (s *FieldSchemaModel) GetScale() *int64 {
	return s.Scale
}

func (s *FieldSchemaModel) GetTableName() *string {
	return s.TableName
}

func (s *FieldSchemaModel) GetTokenizer() *string {
	return s.Tokenizer
}

func (s *FieldSchemaModel) GetType() *string {
	return s.Type
}

func (s *FieldSchemaModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *FieldSchemaModel) GetValueType() *string {
	return s.ValueType
}

func (s *FieldSchemaModel) SetAutoIncrement(v bool) *FieldSchemaModel {
	s.AutoIncrement = &v
	return s
}

func (s *FieldSchemaModel) SetColumnRawName(v string) *FieldSchemaModel {
	s.ColumnRawName = &v
	return s
}

func (s *FieldSchemaModel) SetComment(v string) *FieldSchemaModel {
	s.Comment = &v
	return s
}

func (s *FieldSchemaModel) SetCompressFloatUseShort(v bool) *FieldSchemaModel {
	s.CompressFloatUseShort = &v
	return s
}

func (s *FieldSchemaModel) SetCompression(v string) *FieldSchemaModel {
	s.Compression = &v
	return s
}

func (s *FieldSchemaModel) SetCreateTime(v string) *FieldSchemaModel {
	s.CreateTime = &v
	return s
}

func (s *FieldSchemaModel) SetDataType(v string) *FieldSchemaModel {
	s.DataType = &v
	return s
}

func (s *FieldSchemaModel) SetDatabaseName(v string) *FieldSchemaModel {
	s.DatabaseName = &v
	return s
}

func (s *FieldSchemaModel) SetDefaultValue(v string) *FieldSchemaModel {
	s.DefaultValue = &v
	return s
}

func (s *FieldSchemaModel) SetDelimiter(v string) *FieldSchemaModel {
	s.Delimiter = &v
	return s
}

func (s *FieldSchemaModel) SetEncode(v string) *FieldSchemaModel {
	s.Encode = &v
	return s
}

func (s *FieldSchemaModel) SetIsPartitionKey(v bool) *FieldSchemaModel {
	s.IsPartitionKey = &v
	return s
}

func (s *FieldSchemaModel) SetMappedName(v string) *FieldSchemaModel {
	s.MappedName = &v
	return s
}

func (s *FieldSchemaModel) SetName(v string) *FieldSchemaModel {
	s.Name = &v
	return s
}

func (s *FieldSchemaModel) SetNullable(v bool) *FieldSchemaModel {
	s.Nullable = &v
	return s
}

func (s *FieldSchemaModel) SetOnUpdate(v string) *FieldSchemaModel {
	s.OnUpdate = &v
	return s
}

func (s *FieldSchemaModel) SetOrdinalPosition(v int64) *FieldSchemaModel {
	s.OrdinalPosition = &v
	return s
}

func (s *FieldSchemaModel) SetPhysicalColumnName(v string) *FieldSchemaModel {
	s.PhysicalColumnName = &v
	return s
}

func (s *FieldSchemaModel) SetPkPosition(v int64) *FieldSchemaModel {
	s.PkPosition = &v
	return s
}

func (s *FieldSchemaModel) SetPrecision(v int64) *FieldSchemaModel {
	s.Precision = &v
	return s
}

func (s *FieldSchemaModel) SetPrimarykey(v bool) *FieldSchemaModel {
	s.Primarykey = &v
	return s
}

func (s *FieldSchemaModel) SetScale(v int64) *FieldSchemaModel {
	s.Scale = &v
	return s
}

func (s *FieldSchemaModel) SetTableName(v string) *FieldSchemaModel {
	s.TableName = &v
	return s
}

func (s *FieldSchemaModel) SetTokenizer(v string) *FieldSchemaModel {
	s.Tokenizer = &v
	return s
}

func (s *FieldSchemaModel) SetType(v string) *FieldSchemaModel {
	s.Type = &v
	return s
}

func (s *FieldSchemaModel) SetUpdateTime(v string) *FieldSchemaModel {
	s.UpdateTime = &v
	return s
}

func (s *FieldSchemaModel) SetValueType(v string) *FieldSchemaModel {
	s.ValueType = &v
	return s
}

func (s *FieldSchemaModel) Validate() error {
	return dara.Validate(s)
}
