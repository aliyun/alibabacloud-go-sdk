// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJudgeNodeMetaDesc interface {
	dara.Model
	String() string
	GoString() string
	SetActualValue(v string) *JudgeNodeMetaDesc
	GetActualValue() *string
	SetDataType(v int32) *JudgeNodeMetaDesc
	GetDataType() *int32
	SetField(v string) *JudgeNodeMetaDesc
	GetField() *string
	SetFieldType(v int32) *JudgeNodeMetaDesc
	GetFieldType() *int32
	SetSymbol(v int32) *JudgeNodeMetaDesc
	GetSymbol() *int32
	SetValue(v string) *JudgeNodeMetaDesc
	GetValue() *string
}

type JudgeNodeMetaDesc struct {
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	DataType    *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Field       *string `json:"Field,omitempty" xml:"Field,omitempty"`
	FieldType   *int32  `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	Symbol      *int32  `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s JudgeNodeMetaDesc) String() string {
	return dara.Prettify(s)
}

func (s JudgeNodeMetaDesc) GoString() string {
	return s.String()
}

func (s *JudgeNodeMetaDesc) GetActualValue() *string {
	return s.ActualValue
}

func (s *JudgeNodeMetaDesc) GetDataType() *int32 {
	return s.DataType
}

func (s *JudgeNodeMetaDesc) GetField() *string {
	return s.Field
}

func (s *JudgeNodeMetaDesc) GetFieldType() *int32 {
	return s.FieldType
}

func (s *JudgeNodeMetaDesc) GetSymbol() *int32 {
	return s.Symbol
}

func (s *JudgeNodeMetaDesc) GetValue() *string {
	return s.Value
}

func (s *JudgeNodeMetaDesc) SetActualValue(v string) *JudgeNodeMetaDesc {
	s.ActualValue = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetDataType(v int32) *JudgeNodeMetaDesc {
	s.DataType = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetField(v string) *JudgeNodeMetaDesc {
	s.Field = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetFieldType(v int32) *JudgeNodeMetaDesc {
	s.FieldType = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetSymbol(v int32) *JudgeNodeMetaDesc {
	s.Symbol = &v
	return s
}

func (s *JudgeNodeMetaDesc) SetValue(v string) *JudgeNodeMetaDesc {
	s.Value = &v
	return s
}

func (s *JudgeNodeMetaDesc) Validate() error {
	return dara.Validate(s)
}
