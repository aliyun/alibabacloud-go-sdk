// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimeRange interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v string) *TimeRange
	GetEnd() *string
	SetStart(v string) *TimeRange
	GetStart() *string
}

type TimeRange struct {
	End   *string `json:"End,omitempty" xml:"End,omitempty"`
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s TimeRange) String() string {
	return dara.Prettify(s)
}

func (s TimeRange) GoString() string {
	return s.String()
}

func (s *TimeRange) GetEnd() *string {
	return s.End
}

func (s *TimeRange) GetStart() *string {
	return s.Start
}

func (s *TimeRange) SetEnd(v string) *TimeRange {
	s.End = &v
	return s
}

func (s *TimeRange) SetStart(v string) *TimeRange {
	s.Start = &v
	return s
}

func (s *TimeRange) Validate() error {
	return dara.Validate(s)
}
