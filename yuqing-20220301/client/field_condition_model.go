// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFieldCondition interface {
	dara.Model
	String() string
	GoString() string
	SetFieldName(v string) *FieldCondition
	GetFieldName() *string
	SetNestFieldPath(v string) *FieldCondition
	GetNestFieldPath() *string
	SetNestFieldValue(v []*int64) *FieldCondition
	GetNestFieldValue() []*int64
	SetOperateType(v string) *FieldCondition
	GetOperateType() *string
	SetValue(v string) *FieldCondition
	GetValue() *string
}

type FieldCondition struct {
	// example:
	//
	// xxx
	FieldName      *string  `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	NestFieldPath  *string  `json:"nestFieldPath,omitempty" xml:"nestFieldPath,omitempty"`
	NestFieldValue []*int64 `json:"nestFieldValue,omitempty" xml:"nestFieldValue,omitempty" type:"Repeated"`
	// example:
	//
	// =
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// example:
	//
	// yyy
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FieldCondition) String() string {
	return dara.Prettify(s)
}

func (s FieldCondition) GoString() string {
	return s.String()
}

func (s *FieldCondition) GetFieldName() *string {
	return s.FieldName
}

func (s *FieldCondition) GetNestFieldPath() *string {
	return s.NestFieldPath
}

func (s *FieldCondition) GetNestFieldValue() []*int64 {
	return s.NestFieldValue
}

func (s *FieldCondition) GetOperateType() *string {
	return s.OperateType
}

func (s *FieldCondition) GetValue() *string {
	return s.Value
}

func (s *FieldCondition) SetFieldName(v string) *FieldCondition {
	s.FieldName = &v
	return s
}

func (s *FieldCondition) SetNestFieldPath(v string) *FieldCondition {
	s.NestFieldPath = &v
	return s
}

func (s *FieldCondition) SetNestFieldValue(v []*int64) *FieldCondition {
	s.NestFieldValue = v
	return s
}

func (s *FieldCondition) SetOperateType(v string) *FieldCondition {
	s.OperateType = &v
	return s
}

func (s *FieldCondition) SetValue(v string) *FieldCondition {
	s.Value = &v
	return s
}

func (s *FieldCondition) Validate() error {
	return dara.Validate(s)
}
