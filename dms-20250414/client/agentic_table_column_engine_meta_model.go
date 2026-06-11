// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticTableColumnEngineMeta interface {
	dara.Model
	String() string
	GoString() string
	SetAutoIncrement(v bool) *AgenticTableColumnEngineMeta
	GetAutoIncrement() *bool
	SetDataLength(v int64) *AgenticTableColumnEngineMeta
	GetDataLength() *int64
	SetDataPrecision(v int32) *AgenticTableColumnEngineMeta
	GetDataPrecision() *int32
	SetDataScale(v int32) *AgenticTableColumnEngineMeta
	GetDataScale() *int32
	SetDefaultValue(v string) *AgenticTableColumnEngineMeta
	GetDefaultValue() *string
	SetEncoding(v string) *AgenticTableColumnEngineMeta
	GetEncoding() *string
	SetExtra(v string) *AgenticTableColumnEngineMeta
	GetExtra() *string
	SetGenerationColumn(v bool) *AgenticTableColumnEngineMeta
	GetGenerationColumn() *bool
	SetGenerationExpression(v string) *AgenticTableColumnEngineMeta
	GetGenerationExpression() *string
	SetNullable(v bool) *AgenticTableColumnEngineMeta
	GetNullable() *bool
}

type AgenticTableColumnEngineMeta struct {
	// Indicates whether the column uses auto-increment.
	AutoIncrement *bool `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	// The data length of the column. This parameter typically applies to string and binary data types.
	DataLength *int64 `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
	// The data precision of the column, which is the total number of digits in a numeric type.
	DataPrecision *int32 `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	// The data scale of the column, which is the number of digits to the right of the decimal point in a numeric type.
	DataScale *int32 `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	// The default value of the column.
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The character encoding of the column.
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// Engine-specific attributes or flags for the column.
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// Indicates whether the column is a generated column.
	GenerationColumn *bool `json:"GenerationColumn,omitempty" xml:"GenerationColumn,omitempty"`
	// The expression used to generate the column\\"s value. Applies only if `GenerationColumn` is `true`.
	GenerationExpression *string `json:"GenerationExpression,omitempty" xml:"GenerationExpression,omitempty"`
	// Indicates whether the column is nullable.
	Nullable *bool `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
}

func (s AgenticTableColumnEngineMeta) String() string {
	return dara.Prettify(s)
}

func (s AgenticTableColumnEngineMeta) GoString() string {
	return s.String()
}

func (s *AgenticTableColumnEngineMeta) GetAutoIncrement() *bool {
	return s.AutoIncrement
}

func (s *AgenticTableColumnEngineMeta) GetDataLength() *int64 {
	return s.DataLength
}

func (s *AgenticTableColumnEngineMeta) GetDataPrecision() *int32 {
	return s.DataPrecision
}

func (s *AgenticTableColumnEngineMeta) GetDataScale() *int32 {
	return s.DataScale
}

func (s *AgenticTableColumnEngineMeta) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *AgenticTableColumnEngineMeta) GetEncoding() *string {
	return s.Encoding
}

func (s *AgenticTableColumnEngineMeta) GetExtra() *string {
	return s.Extra
}

func (s *AgenticTableColumnEngineMeta) GetGenerationColumn() *bool {
	return s.GenerationColumn
}

func (s *AgenticTableColumnEngineMeta) GetGenerationExpression() *string {
	return s.GenerationExpression
}

func (s *AgenticTableColumnEngineMeta) GetNullable() *bool {
	return s.Nullable
}

func (s *AgenticTableColumnEngineMeta) SetAutoIncrement(v bool) *AgenticTableColumnEngineMeta {
	s.AutoIncrement = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetDataLength(v int64) *AgenticTableColumnEngineMeta {
	s.DataLength = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetDataPrecision(v int32) *AgenticTableColumnEngineMeta {
	s.DataPrecision = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetDataScale(v int32) *AgenticTableColumnEngineMeta {
	s.DataScale = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetDefaultValue(v string) *AgenticTableColumnEngineMeta {
	s.DefaultValue = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetEncoding(v string) *AgenticTableColumnEngineMeta {
	s.Encoding = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetExtra(v string) *AgenticTableColumnEngineMeta {
	s.Extra = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetGenerationColumn(v bool) *AgenticTableColumnEngineMeta {
	s.GenerationColumn = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetGenerationExpression(v string) *AgenticTableColumnEngineMeta {
	s.GenerationExpression = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) SetNullable(v bool) *AgenticTableColumnEngineMeta {
	s.Nullable = &v
	return s
}

func (s *AgenticTableColumnEngineMeta) Validate() error {
	return dara.Validate(s)
}
