// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSlsNamedQueryEntry interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v int64) *SlsNamedQueryEntry
	GetEnd() *int64
	SetExpr(v string) *SlsNamedQueryEntry
	GetExpr() *string
	SetStart(v int64) *SlsNamedQueryEntry
	GetStart() *int64
	SetTimeUnit(v string) *SlsNamedQueryEntry
	GetTimeUnit() *string
	SetWindow(v int64) *SlsNamedQueryEntry
	GetWindow() *int64
}

type SlsNamedQueryEntry struct {
	// The end offset of the time range. This parameter is mutually exclusive with window.
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// The SPL query expression.
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// The start offset of the time range. This parameter is mutually exclusive with window.
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// The time unit. Valid values: day, hour, minute, and second.
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// The snap window size. This parameter is mutually exclusive with start and end.
	Window *int64 `json:"window,omitempty" xml:"window,omitempty"`
}

func (s SlsNamedQueryEntry) String() string {
	return dara.Prettify(s)
}

func (s SlsNamedQueryEntry) GoString() string {
	return s.String()
}

func (s *SlsNamedQueryEntry) GetEnd() *int64 {
	return s.End
}

func (s *SlsNamedQueryEntry) GetExpr() *string {
	return s.Expr
}

func (s *SlsNamedQueryEntry) GetStart() *int64 {
	return s.Start
}

func (s *SlsNamedQueryEntry) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *SlsNamedQueryEntry) GetWindow() *int64 {
	return s.Window
}

func (s *SlsNamedQueryEntry) SetEnd(v int64) *SlsNamedQueryEntry {
	s.End = &v
	return s
}

func (s *SlsNamedQueryEntry) SetExpr(v string) *SlsNamedQueryEntry {
	s.Expr = &v
	return s
}

func (s *SlsNamedQueryEntry) SetStart(v int64) *SlsNamedQueryEntry {
	s.Start = &v
	return s
}

func (s *SlsNamedQueryEntry) SetTimeUnit(v string) *SlsNamedQueryEntry {
	s.TimeUnit = &v
	return s
}

func (s *SlsNamedQueryEntry) SetWindow(v int64) *SlsNamedQueryEntry {
	s.Window = &v
	return s
}

func (s *SlsNamedQueryEntry) Validate() error {
	return dara.Validate(s)
}
