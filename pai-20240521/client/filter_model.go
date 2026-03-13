// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilter interface {
	dara.Model
	String() string
	GoString() string
	SetFilterBy(v string) *Filter
	GetFilterBy() *string
	SetFilterCondition(v string) *Filter
	GetFilterCondition() *string
	SetFilterOperation(v string) *Filter
	GetFilterOperation() *string
}

type Filter struct {
	FilterBy        *string `json:"FilterBy,omitempty" xml:"FilterBy,omitempty"`
	FilterCondition *string `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	FilterOperation *string `json:"FilterOperation,omitempty" xml:"FilterOperation,omitempty"`
}

func (s Filter) String() string {
	return dara.Prettify(s)
}

func (s Filter) GoString() string {
	return s.String()
}

func (s *Filter) GetFilterBy() *string {
	return s.FilterBy
}

func (s *Filter) GetFilterCondition() *string {
	return s.FilterCondition
}

func (s *Filter) GetFilterOperation() *string {
	return s.FilterOperation
}

func (s *Filter) SetFilterBy(v string) *Filter {
	s.FilterBy = &v
	return s
}

func (s *Filter) SetFilterCondition(v string) *Filter {
	s.FilterCondition = &v
	return s
}

func (s *Filter) SetFilterOperation(v string) *Filter {
	s.FilterOperation = &v
	return s
}

func (s *Filter) Validate() error {
	return dara.Validate(s)
}
