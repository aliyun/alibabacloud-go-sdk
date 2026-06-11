// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagCondition interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *TagCondition
	GetKey() *string
	SetOp(v string) *TagCondition
	GetOp() *string
	SetValue(v string) *TagCondition
	GetValue() *string
}

type TagCondition struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Op    *string `json:"op,omitempty" xml:"op,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TagCondition) String() string {
	return dara.Prettify(s)
}

func (s TagCondition) GoString() string {
	return s.String()
}

func (s *TagCondition) GetKey() *string {
	return s.Key
}

func (s *TagCondition) GetOp() *string {
	return s.Op
}

func (s *TagCondition) GetValue() *string {
	return s.Value
}

func (s *TagCondition) SetKey(v string) *TagCondition {
	s.Key = &v
	return s
}

func (s *TagCondition) SetOp(v string) *TagCondition {
	s.Op = &v
	return s
}

func (s *TagCondition) SetValue(v string) *TagCondition {
	s.Value = &v
	return s
}

func (s *TagCondition) Validate() error {
	return dara.Validate(s)
}
