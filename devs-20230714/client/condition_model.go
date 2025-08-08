// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCondition interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *Condition
	GetExpression() *string
}

type Condition struct {
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
}

func (s Condition) String() string {
	return dara.Prettify(s)
}

func (s Condition) GoString() string {
	return s.String()
}

func (s *Condition) GetExpression() *string {
	return s.Expression
}

func (s *Condition) SetExpression(v string) *Condition {
	s.Expression = &v
	return s
}

func (s *Condition) Validate() error {
	return dara.Validate(s)
}
