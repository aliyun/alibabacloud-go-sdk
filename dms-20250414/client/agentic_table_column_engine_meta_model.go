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
	// Specifies whether the field is auto-incremented.
	//
	// example:
	//
	// false
	AutoIncrement *bool `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	// The defined length of the field.
	//
	// example:
	//
	// 20
	DataLength *int64 `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
	// The precision of the numeric field.
	//
	// example:
	//
	// 10
	DataPrecision *int32 `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	// The scale of the numeric field, which is the number of decimal places.
	//
	// example:
	//
	// 2
	DataScale *int32 `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	// The default value of the field.
	//
	// example:
	//
	// 0
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The character set of the field.
	//
	// example:
	//
	// utf8mb4
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The additional information of the field, such as auto_increment.
	//
	// example:
	//
	// auto_increment
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// Specifies whether the field is a generated column.
	//
	// example:
	//
	// false
	GenerationColumn *bool `json:"GenerationColumn,omitempty" xml:"GenerationColumn,omitempty"`
	// The calculation expression of the generated column.
	//
	// example:
	//
	// price 	- quantity
	GenerationExpression *string `json:"GenerationExpression,omitempty" xml:"GenerationExpression,omitempty"`
	// Specifies whether the field allows null values.
	//
	// example:
	//
	// false
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
