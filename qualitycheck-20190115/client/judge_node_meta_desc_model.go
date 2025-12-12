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
	SetExpressionMetaDesc(v *JudgeNodeMetaDescExpressionMetaDesc) *JudgeNodeMetaDesc
	GetExpressionMetaDesc() *JudgeNodeMetaDescExpressionMetaDesc
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
	ActualValue        *string                              `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	DataType           *int32                               `json:"DataType,omitempty" xml:"DataType,omitempty"`
	ExpressionMetaDesc *JudgeNodeMetaDescExpressionMetaDesc `json:"ExpressionMetaDesc,omitempty" xml:"ExpressionMetaDesc,omitempty" type:"Struct"`
	Field              *string                              `json:"Field,omitempty" xml:"Field,omitempty"`
	FieldType          *int32                               `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	Symbol             *int32                               `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	Value              *string                              `json:"Value,omitempty" xml:"Value,omitempty"`
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

func (s *JudgeNodeMetaDesc) GetExpressionMetaDesc() *JudgeNodeMetaDescExpressionMetaDesc {
	return s.ExpressionMetaDesc
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

func (s *JudgeNodeMetaDesc) SetExpressionMetaDesc(v *JudgeNodeMetaDescExpressionMetaDesc) *JudgeNodeMetaDesc {
	s.ExpressionMetaDesc = v
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
	if s.ExpressionMetaDesc != nil {
		if err := s.ExpressionMetaDesc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type JudgeNodeMetaDescExpressionMetaDesc struct {
	LeftFieldType  *int32  `json:"LeftFieldType,omitempty" xml:"LeftFieldType,omitempty"`
	LeftOperand    *string `json:"LeftOperand,omitempty" xml:"LeftOperand,omitempty"`
	Operator       *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RightFieldType *int32  `json:"RightFieldType,omitempty" xml:"RightFieldType,omitempty"`
	RightOperand   *string `json:"RightOperand,omitempty" xml:"RightOperand,omitempty"`
	RoundingMode   *string `json:"RoundingMode,omitempty" xml:"RoundingMode,omitempty"`
}

func (s JudgeNodeMetaDescExpressionMetaDesc) String() string {
	return dara.Prettify(s)
}

func (s JudgeNodeMetaDescExpressionMetaDesc) GoString() string {
	return s.String()
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) GetLeftFieldType() *int32 {
	return s.LeftFieldType
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) GetLeftOperand() *string {
	return s.LeftOperand
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) GetOperator() *string {
	return s.Operator
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) GetRightFieldType() *int32 {
	return s.RightFieldType
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) GetRightOperand() *string {
	return s.RightOperand
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) GetRoundingMode() *string {
	return s.RoundingMode
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) SetLeftFieldType(v int32) *JudgeNodeMetaDescExpressionMetaDesc {
	s.LeftFieldType = &v
	return s
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) SetLeftOperand(v string) *JudgeNodeMetaDescExpressionMetaDesc {
	s.LeftOperand = &v
	return s
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) SetOperator(v string) *JudgeNodeMetaDescExpressionMetaDesc {
	s.Operator = &v
	return s
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) SetRightFieldType(v int32) *JudgeNodeMetaDescExpressionMetaDesc {
	s.RightFieldType = &v
	return s
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) SetRightOperand(v string) *JudgeNodeMetaDescExpressionMetaDesc {
	s.RightOperand = &v
	return s
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) SetRoundingMode(v string) *JudgeNodeMetaDescExpressionMetaDesc {
	s.RoundingMode = &v
	return s
}

func (s *JudgeNodeMetaDescExpressionMetaDesc) Validate() error {
	return dara.Validate(s)
}
