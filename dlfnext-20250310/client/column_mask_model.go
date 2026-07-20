// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColumnMask interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *ColumnMask
	GetExpression() *string
	SetTransform(v string) *ColumnMask
	GetTransform() *string
}

type ColumnMask struct {
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	Transform  *string `json:"transform,omitempty" xml:"transform,omitempty"`
}

func (s ColumnMask) String() string {
	return dara.Prettify(s)
}

func (s ColumnMask) GoString() string {
	return s.String()
}

func (s *ColumnMask) GetExpression() *string {
	return s.Expression
}

func (s *ColumnMask) GetTransform() *string {
	return s.Transform
}

func (s *ColumnMask) SetExpression(v string) *ColumnMask {
	s.Expression = &v
	return s
}

func (s *ColumnMask) SetTransform(v string) *ColumnMask {
	s.Transform = &v
	return s
}

func (s *ColumnMask) Validate() error {
	return dara.Validate(s)
}
