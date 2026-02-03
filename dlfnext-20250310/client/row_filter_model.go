// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRowFilter interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *RowFilter
	GetExpression() *string
	SetPredicate(v string) *RowFilter
	GetPredicate() *string
}

type RowFilter struct {
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	Predicate  *string `json:"predicate,omitempty" xml:"predicate,omitempty"`
}

func (s RowFilter) String() string {
	return dara.Prettify(s)
}

func (s RowFilter) GoString() string {
	return s.String()
}

func (s *RowFilter) GetExpression() *string {
	return s.Expression
}

func (s *RowFilter) GetPredicate() *string {
	return s.Predicate
}

func (s *RowFilter) SetExpression(v string) *RowFilter {
	s.Expression = &v
	return s
}

func (s *RowFilter) SetPredicate(v string) *RowFilter {
	s.Predicate = &v
	return s
}

func (s *RowFilter) Validate() error {
	return dara.Validate(s)
}
