// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSchema interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ViewSchema
	GetComment() *string
	SetDialects(v map[string]*string) *ViewSchema
	GetDialects() map[string]*string
	SetFields(v []*DataField) *ViewSchema
	GetFields() []*DataField
	SetOptions(v map[string]*string) *ViewSchema
	GetOptions() map[string]*string
	SetQuery(v string) *ViewSchema
	GetQuery() *string
}

type ViewSchema struct {
	// The comment.
	//
	// example:
	//
	// comment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// Queries for different engines.
	Dialects map[string]*string `json:"dialects,omitempty" xml:"dialects,omitempty"`
	// The list of field definitions for the view.
	Fields []*DataField `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The configuration parameters for the view.
	Options map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	// The default query for the view.
	//
	// example:
	//
	// select id from t1;
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ViewSchema) String() string {
	return dara.Prettify(s)
}

func (s ViewSchema) GoString() string {
	return s.String()
}

func (s *ViewSchema) GetComment() *string {
	return s.Comment
}

func (s *ViewSchema) GetDialects() map[string]*string {
	return s.Dialects
}

func (s *ViewSchema) GetFields() []*DataField {
	return s.Fields
}

func (s *ViewSchema) GetOptions() map[string]*string {
	return s.Options
}

func (s *ViewSchema) GetQuery() *string {
	return s.Query
}

func (s *ViewSchema) SetComment(v string) *ViewSchema {
	s.Comment = &v
	return s
}

func (s *ViewSchema) SetDialects(v map[string]*string) *ViewSchema {
	s.Dialects = v
	return s
}

func (s *ViewSchema) SetFields(v []*DataField) *ViewSchema {
	s.Fields = v
	return s
}

func (s *ViewSchema) SetOptions(v map[string]*string) *ViewSchema {
	s.Options = v
	return s
}

func (s *ViewSchema) SetQuery(v string) *ViewSchema {
	s.Query = &v
	return s
}

func (s *ViewSchema) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
