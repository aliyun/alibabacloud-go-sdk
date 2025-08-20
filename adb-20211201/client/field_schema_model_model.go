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
	AutoIncrement         *bool   `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	ColumnRawName         *string `json:"ColumnRawName,omitempty" xml:"ColumnRawName,omitempty"`
	Comment               *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CompressFloatUseShort *bool   `json:"CompressFloatUseShort,omitempty" xml:"CompressFloatUseShort,omitempty"`
	Compression           *string `json:"Compression,omitempty" xml:"Compression,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataType              *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DatabaseName          *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DefaultValue          *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Delimiter             *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	Encode                *string `json:"Encode,omitempty" xml:"Encode,omitempty"`
	IsPartitionKey        *bool   `json:"IsPartitionKey,omitempty" xml:"IsPartitionKey,omitempty"`
	MappedName            *string `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nullable              *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	OnUpdate              *string `json:"OnUpdate,omitempty" xml:"OnUpdate,omitempty"`
	OrdinalPosition       *int64  `json:"OrdinalPosition,omitempty" xml:"OrdinalPosition,omitempty"`
	PhysicalColumnName    *string `json:"PhysicalColumnName,omitempty" xml:"PhysicalColumnName,omitempty"`
	PkPosition            *int64  `json:"PkPosition,omitempty" xml:"PkPosition,omitempty"`
	Precision             *int64  `json:"Precision,omitempty" xml:"Precision,omitempty"`
	Primarykey            *bool   `json:"Primarykey,omitempty" xml:"Primarykey,omitempty"`
	Scale                 *int64  `json:"Scale,omitempty" xml:"Scale,omitempty"`
	TableName             *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Tokenizer             *string `json:"Tokenizer,omitempty" xml:"Tokenizer,omitempty"`
	Type                  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime            *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ValueType             *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
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
