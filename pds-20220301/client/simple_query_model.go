// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleQuery interface {
	dara.Model
	String() string
	GoString() string
	SetField(v []byte) *SimpleQuery
	GetField() []byte
	SetOperation(v []byte) *SimpleQuery
	GetOperation() []byte
	SetSubQueries(v []*SimpleQuery) *SimpleQuery
	GetSubQueries() []*SimpleQuery
	SetValue(v []byte) *SimpleQuery
	GetValue() []byte
}

type SimpleQuery struct {
	Field      []byte         `json:"field,omitempty" xml:"field,omitempty"`
	Operation  []byte         `json:"operation,omitempty" xml:"operation,omitempty"`
	SubQueries []*SimpleQuery `json:"sub_queries,omitempty" xml:"sub_queries,omitempty" type:"Repeated"`
	Value      []byte         `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SimpleQuery) String() string {
	return dara.Prettify(s)
}

func (s SimpleQuery) GoString() string {
	return s.String()
}

func (s *SimpleQuery) GetField() []byte {
	return s.Field
}

func (s *SimpleQuery) GetOperation() []byte {
	return s.Operation
}

func (s *SimpleQuery) GetSubQueries() []*SimpleQuery {
	return s.SubQueries
}

func (s *SimpleQuery) GetValue() []byte {
	return s.Value
}

func (s *SimpleQuery) SetField(v []byte) *SimpleQuery {
	s.Field = v
	return s
}

func (s *SimpleQuery) SetOperation(v []byte) *SimpleQuery {
	s.Operation = v
	return s
}

func (s *SimpleQuery) SetSubQueries(v []*SimpleQuery) *SimpleQuery {
	s.SubQueries = v
	return s
}

func (s *SimpleQuery) SetValue(v []byte) *SimpleQuery {
	s.Value = v
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
