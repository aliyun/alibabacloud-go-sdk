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
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// required in UpdateComment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// required in AddDialect/UpdateDialect/DropDialect
	Dialect *string `json:"dialect,omitempty" xml:"dialect,omitempty"`
	// required in SetOption/RemoveOption
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// required in AddDialect/UpdateDialect
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// required in SetOption
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
