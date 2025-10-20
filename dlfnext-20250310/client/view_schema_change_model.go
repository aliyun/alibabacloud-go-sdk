// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSchemaChange interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ViewSchemaChange
	GetAction() *string
	SetComment(v string) *ViewSchemaChange
	GetComment() *string
	SetDialect(v string) *ViewSchemaChange
	GetDialect() *string
	SetKey(v string) *ViewSchemaChange
	GetKey() *string
	SetQuery(v string) *ViewSchemaChange
	GetQuery() *string
	SetValue(v string) *ViewSchemaChange
	GetValue() *string
}

type ViewSchemaChange struct {
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

func (s ViewSchemaChange) String() string {
	return dara.Prettify(s)
}

func (s ViewSchemaChange) GoString() string {
	return s.String()
}

func (s *ViewSchemaChange) GetAction() *string {
	return s.Action
}

func (s *ViewSchemaChange) GetComment() *string {
	return s.Comment
}

func (s *ViewSchemaChange) GetDialect() *string {
	return s.Dialect
}

func (s *ViewSchemaChange) GetKey() *string {
	return s.Key
}

func (s *ViewSchemaChange) GetQuery() *string {
	return s.Query
}

func (s *ViewSchemaChange) GetValue() *string {
	return s.Value
}

func (s *ViewSchemaChange) SetAction(v string) *ViewSchemaChange {
	s.Action = &v
	return s
}

func (s *ViewSchemaChange) SetComment(v string) *ViewSchemaChange {
	s.Comment = &v
	return s
}

func (s *ViewSchemaChange) SetDialect(v string) *ViewSchemaChange {
	s.Dialect = &v
	return s
}

func (s *ViewSchemaChange) SetKey(v string) *ViewSchemaChange {
	s.Key = &v
	return s
}

func (s *ViewSchemaChange) SetQuery(v string) *ViewSchemaChange {
	s.Query = &v
	return s
}

func (s *ViewSchemaChange) SetValue(v string) *ViewSchemaChange {
	s.Value = &v
	return s
}

func (s *ViewSchemaChange) Validate() error {
	return dara.Validate(s)
}
