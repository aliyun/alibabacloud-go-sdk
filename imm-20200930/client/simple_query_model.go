// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleQuery interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *SimpleQuery
	GetField() *string
	SetOperation(v string) *SimpleQuery
	GetOperation() *string
	SetSubQueries(v []*SimpleQuery) *SimpleQuery
	GetSubQueries() []*SimpleQuery
	SetValue(v string) *SimpleQuery
	GetValue() *string
}

type SimpleQuery struct {
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// eq / gt / gte / lt / lte / match / prefix / and / or / not
	Operation  *string        `json:"Operation,omitempty" xml:"Operation,omitempty"`
	SubQueries []*SimpleQuery `json:"SubQueries,omitempty" xml:"SubQueries,omitempty" type:"Repeated"`
	Value      *string        `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SimpleQuery) String() string {
	return dara.Prettify(s)
}

func (s SimpleQuery) GoString() string {
	return s.String()
}

func (s *SimpleQuery) GetField() *string {
	return s.Field
}

func (s *SimpleQuery) GetOperation() *string {
	return s.Operation
}

func (s *SimpleQuery) GetSubQueries() []*SimpleQuery {
	return s.SubQueries
}

func (s *SimpleQuery) GetValue() *string {
	return s.Value
}

func (s *SimpleQuery) SetField(v string) *SimpleQuery {
	s.Field = &v
	return s
}

func (s *SimpleQuery) SetOperation(v string) *SimpleQuery {
	s.Operation = &v
	return s
}

func (s *SimpleQuery) SetSubQueries(v []*SimpleQuery) *SimpleQuery {
	s.SubQueries = v
	return s
}

func (s *SimpleQuery) SetValue(v string) *SimpleQuery {
	s.Value = &v
	return s
}

func (s *SimpleQuery) Validate() error {
	if s.SubQueries != nil {
		for _, item := range s.SubQueries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
