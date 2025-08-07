// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilter interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *Filter
	GetKey() *string
	SetOperator(v string) *Filter
	GetOperator() *string
	SetSubFilters(v []*Filter) *Filter
	GetSubFilters() []*Filter
	SetValues(v interface{}) *Filter
	GetValues() interface{}
}

type Filter struct {
	Key        *string     `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator   *string     `json:"Operator,omitempty" xml:"Operator,omitempty"`
	SubFilters []*Filter   `json:"SubFilters,omitempty" xml:"SubFilters,omitempty" type:"Repeated"`
	Values     interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s Filter) String() string {
	return dara.Prettify(s)
}

func (s Filter) GoString() string {
	return s.String()
}

func (s *Filter) GetKey() *string {
	return s.Key
}

func (s *Filter) GetOperator() *string {
	return s.Operator
}

func (s *Filter) GetSubFilters() []*Filter {
	return s.SubFilters
}

func (s *Filter) GetValues() interface{} {
	return s.Values
}

func (s *Filter) SetKey(v string) *Filter {
	s.Key = &v
	return s
}

func (s *Filter) SetOperator(v string) *Filter {
	s.Operator = &v
	return s
}

func (s *Filter) SetSubFilters(v []*Filter) *Filter {
	s.SubFilters = v
	return s
}

func (s *Filter) SetValues(v interface{}) *Filter {
	s.Values = v
	return s
}

func (s *Filter) Validate() error {
	return dara.Validate(s)
}
