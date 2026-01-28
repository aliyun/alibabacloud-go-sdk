// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaTableColumnEngineMeta interface {
	dara.Model
	String() string
	GoString() string
	SetAutoIncrement(v bool) *OneMetaTableColumnEngineMeta
	GetAutoIncrement() *bool
	SetDataLength(v int64) *OneMetaTableColumnEngineMeta
	GetDataLength() *int64
	SetDataPrecision(v int32) *OneMetaTableColumnEngineMeta
	GetDataPrecision() *int32
	SetDataScale(v int32) *OneMetaTableColumnEngineMeta
	GetDataScale() *int32
	SetDefaultValue(v string) *OneMetaTableColumnEngineMeta
	GetDefaultValue() *string
	SetEncoding(v string) *OneMetaTableColumnEngineMeta
	GetEncoding() *string
	SetExtra(v string) *OneMetaTableColumnEngineMeta
	GetExtra() *string
	SetGenerationColumn(v bool) *OneMetaTableColumnEngineMeta
	GetGenerationColumn() *bool
	SetGenerationExpression(v string) *OneMetaTableColumnEngineMeta
	GetGenerationExpression() *string
	SetNullable(v bool) *OneMetaTableColumnEngineMeta
	GetNullable() *bool
}

type OneMetaTableColumnEngineMeta struct {
	AutoIncrement        *bool   `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	DataLength           *int64  `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
	DataPrecision        *int32  `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	DataScale            *int32  `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	DefaultValue         *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Encoding             *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	Extra                *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	GenerationColumn     *bool   `json:"GenerationColumn,omitempty" xml:"GenerationColumn,omitempty"`
	GenerationExpression *string `json:"GenerationExpression,omitempty" xml:"GenerationExpression,omitempty"`
	Nullable             *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
}

func (s OneMetaTableColumnEngineMeta) String() string {
	return dara.Prettify(s)
}

func (s OneMetaTableColumnEngineMeta) GoString() string {
	return s.String()
}

func (s *OneMetaTableColumnEngineMeta) GetAutoIncrement() *bool {
	return s.AutoIncrement
}

func (s *OneMetaTableColumnEngineMeta) GetDataLength() *int64 {
	return s.DataLength
}

func (s *OneMetaTableColumnEngineMeta) GetDataPrecision() *int32 {
	return s.DataPrecision
}

func (s *OneMetaTableColumnEngineMeta) GetDataScale() *int32 {
	return s.DataScale
}

func (s *OneMetaTableColumnEngineMeta) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *OneMetaTableColumnEngineMeta) GetEncoding() *string {
	return s.Encoding
}

func (s *OneMetaTableColumnEngineMeta) GetExtra() *string {
	return s.Extra
}

func (s *OneMetaTableColumnEngineMeta) GetGenerationColumn() *bool {
	return s.GenerationColumn
}

func (s *OneMetaTableColumnEngineMeta) GetGenerationExpression() *string {
	return s.GenerationExpression
}

func (s *OneMetaTableColumnEngineMeta) GetNullable() *bool {
	return s.Nullable
}

func (s *OneMetaTableColumnEngineMeta) SetAutoIncrement(v bool) *OneMetaTableColumnEngineMeta {
	s.AutoIncrement = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetDataLength(v int64) *OneMetaTableColumnEngineMeta {
	s.DataLength = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetDataPrecision(v int32) *OneMetaTableColumnEngineMeta {
	s.DataPrecision = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetDataScale(v int32) *OneMetaTableColumnEngineMeta {
	s.DataScale = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetDefaultValue(v string) *OneMetaTableColumnEngineMeta {
	s.DefaultValue = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetEncoding(v string) *OneMetaTableColumnEngineMeta {
	s.Encoding = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetExtra(v string) *OneMetaTableColumnEngineMeta {
	s.Extra = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetGenerationColumn(v bool) *OneMetaTableColumnEngineMeta {
	s.GenerationColumn = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetGenerationExpression(v string) *OneMetaTableColumnEngineMeta {
	s.GenerationExpression = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) SetNullable(v bool) *OneMetaTableColumnEngineMeta {
	s.Nullable = &v
	return s
}

func (s *OneMetaTableColumnEngineMeta) Validate() error {
	return dara.Validate(s)
}
