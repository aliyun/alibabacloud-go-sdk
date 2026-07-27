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
	End      *int64  `json:"end,omitempty" xml:"end,omitempty"`
	Expr     *string `json:"expr,omitempty" xml:"expr,omitempty"`
	Start    *int64  `json:"start,omitempty" xml:"start,omitempty"`
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	Window   *int64  `json:"window,omitempty" xml:"window,omitempty"`
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
