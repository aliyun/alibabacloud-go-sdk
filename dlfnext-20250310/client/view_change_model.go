// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewChange interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ViewChange
	GetAction() *string
	SetComment(v string) *ViewChange
	GetComment() *string
	SetDialect(v string) *ViewChange
	GetDialect() *string
	SetKey(v string) *ViewChange
	GetKey() *string
	SetQuery(v string) *ViewChange
	GetQuery() *string
	SetValue(v string) *ViewChange
	GetValue() *string
}

type ViewChange struct {
	// The type of the view change operation.
	//
	// example:
	//
	// updateComment
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The view comment. Required when you perform the updateComment operation.
	//
	// example:
	//
	// View update description
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The dialect used in the view query statement. Required when you perform the addDialect, updateDialect, or dropDialect operation.
	//
	// example:
	//
	// flink
	Dialect *string `json:"dialect,omitempty" xml:"dialect,omitempty"`
	// The key of the view property. Required when you perform the setOption or removeOption operation.
	//
	// example:
	//
	// view_option
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The view query statement. Required when you perform the addDialect or updateDialect operation.
	//
	// example:
	//
	// SELECT region, SUM(amount) AS total_amount FROM sales GROUP BY region
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The value of the view property. Required when you perform the setOption operation.
	//
	// example:
	//
	// enabled
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ViewChange) String() string {
	return dara.Prettify(s)
}

func (s ViewChange) GoString() string {
	return s.String()
}

func (s *ViewChange) GetAction() *string {
	return s.Action
}

func (s *ViewChange) GetComment() *string {
	return s.Comment
}

func (s *ViewChange) GetDialect() *string {
	return s.Dialect
}

func (s *ViewChange) GetKey() *string {
	return s.Key
}

func (s *ViewChange) GetQuery() *string {
	return s.Query
}

func (s *ViewChange) GetValue() *string {
	return s.Value
}

func (s *ViewChange) SetAction(v string) *ViewChange {
	s.Action = &v
	return s
}

func (s *ViewChange) SetComment(v string) *ViewChange {
	s.Comment = &v
	return s
}

func (s *ViewChange) SetDialect(v string) *ViewChange {
	s.Dialect = &v
	return s
}

func (s *ViewChange) SetKey(v string) *ViewChange {
	s.Key = &v
	return s
}

func (s *ViewChange) SetQuery(v string) *ViewChange {
	s.Query = &v
	return s
}

func (s *ViewChange) SetValue(v string) *ViewChange {
	s.Value = &v
	return s
}

func (s *ViewChange) Validate() error {
	return dara.Validate(s)
}
