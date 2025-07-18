// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPAL7ConfigRewriteOp interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *PAL7ConfigRewriteOp
	GetKey() *string
	SetOldValue(v string) *PAL7ConfigRewriteOp
	GetOldValue() *string
	SetOp(v string) *PAL7ConfigRewriteOp
	GetOp() *string
	SetValue(v string) *PAL7ConfigRewriteOp
	GetValue() *string
	SetValueVariable(v string) *PAL7ConfigRewriteOp
	GetValueVariable() *string
}

type PAL7ConfigRewriteOp struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	// example:
	//
	// add,set,delete,replace
	Op            *string `json:"Op,omitempty" xml:"Op,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueVariable *string `json:"ValueVariable,omitempty" xml:"ValueVariable,omitempty"`
}

func (s PAL7ConfigRewriteOp) String() string {
	return dara.Prettify(s)
}

func (s PAL7ConfigRewriteOp) GoString() string {
	return s.String()
}

func (s *PAL7ConfigRewriteOp) GetKey() *string {
	return s.Key
}

func (s *PAL7ConfigRewriteOp) GetOldValue() *string {
	return s.OldValue
}

func (s *PAL7ConfigRewriteOp) GetOp() *string {
	return s.Op
}

func (s *PAL7ConfigRewriteOp) GetValue() *string {
	return s.Value
}

func (s *PAL7ConfigRewriteOp) GetValueVariable() *string {
	return s.ValueVariable
}

func (s *PAL7ConfigRewriteOp) SetKey(v string) *PAL7ConfigRewriteOp {
	s.Key = &v
	return s
}

func (s *PAL7ConfigRewriteOp) SetOldValue(v string) *PAL7ConfigRewriteOp {
	s.OldValue = &v
	return s
}

func (s *PAL7ConfigRewriteOp) SetOp(v string) *PAL7ConfigRewriteOp {
	s.Op = &v
	return s
}

func (s *PAL7ConfigRewriteOp) SetValue(v string) *PAL7ConfigRewriteOp {
	s.Value = &v
	return s
}

func (s *PAL7ConfigRewriteOp) SetValueVariable(v string) *PAL7ConfigRewriteOp {
	s.ValueVariable = &v
	return s
}

func (s *PAL7ConfigRewriteOp) Validate() error {
	return dara.Validate(s)
}
