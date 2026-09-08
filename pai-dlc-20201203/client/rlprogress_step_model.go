// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressStep interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *RLProgressStep
	GetCurrent() *int32
	SetEtaSec(v int64) *RLProgressStep
	GetEtaSec() *int64
	SetPaceSec(v float64) *RLProgressStep
	GetPaceSec() *float64
	SetPct(v float64) *RLProgressStep
	GetPct() *float64
	SetTime(v int64) *RLProgressStep
	GetTime() *int64
	SetTotal(v int32) *RLProgressStep
	GetTotal() *int32
}

type RLProgressStep struct {
	// The current step.
	//
	// example:
	//
	// 3
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// The estimated remaining seconds, calculated as (Total - Current) × PaceSec.
	//
	// example:
	//
	// 0
	EtaSec *int64 `json:"EtaSec,omitempty" xml:"EtaSec,omitempty"`
	// The per-step duration, calculated as the differential between contiguous step marks, in seconds.
	//
	// example:
	//
	// 14
	PaceSec *float64 `json:"PaceSec,omitempty" xml:"PaceSec,omitempty"`
	// The progress percentage, which is the ratio of Current to Total.
	//
	// example:
	//
	// 100
	Pct *float64 `json:"Pct,omitempty" xml:"Pct,omitempty"`
	// The latest step mark time, in UNIX seconds.
	//
	// example:
	//
	// 1787474487
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The total number of steps, obtained from the configuration dump.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RLProgressStep) String() string {
	return dara.Prettify(s)
}

func (s RLProgressStep) GoString() string {
	return s.String()
}

func (s *RLProgressStep) GetCurrent() *int32 {
	return s.Current
}

func (s *RLProgressStep) GetEtaSec() *int64 {
	return s.EtaSec
}

func (s *RLProgressStep) GetPaceSec() *float64 {
	return s.PaceSec
}

func (s *RLProgressStep) GetPct() *float64 {
	return s.Pct
}

func (s *RLProgressStep) GetTime() *int64 {
	return s.Time
}

func (s *RLProgressStep) GetTotal() *int32 {
	return s.Total
}

func (s *RLProgressStep) SetCurrent(v int32) *RLProgressStep {
	s.Current = &v
	return s
}

func (s *RLProgressStep) SetEtaSec(v int64) *RLProgressStep {
	s.EtaSec = &v
	return s
}

func (s *RLProgressStep) SetPaceSec(v float64) *RLProgressStep {
	s.PaceSec = &v
	return s
}

func (s *RLProgressStep) SetPct(v float64) *RLProgressStep {
	s.Pct = &v
	return s
}

func (s *RLProgressStep) SetTime(v int64) *RLProgressStep {
	s.Time = &v
	return s
}

func (s *RLProgressStep) SetTotal(v int32) *RLProgressStep {
	s.Total = &v
	return s
}

func (s *RLProgressStep) Validate() error {
	return dara.Validate(s)
}
