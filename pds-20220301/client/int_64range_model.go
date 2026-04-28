// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInt64Range interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *Int64Range
	GetFrom() *int64
	SetTo(v int64) *Int64Range
	GetTo() *int64
}

type Int64Range struct {
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	To   *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s Int64Range) String() string {
	return dara.Prettify(s)
}

func (s Int64Range) GoString() string {
	return s.String()
}

func (s *Int64Range) GetFrom() *int64 {
	return s.From
}

func (s *Int64Range) GetTo() *int64 {
	return s.To
}

func (s *Int64Range) SetFrom(v int64) *Int64Range {
	s.From = &v
	return s
}

func (s *Int64Range) SetTo(v int64) *Int64Range {
	s.To = &v
	return s
}

func (s *Int64Range) Validate() error {
	return dara.Validate(s)
}
