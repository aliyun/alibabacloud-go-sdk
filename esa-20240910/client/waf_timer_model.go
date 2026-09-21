// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWafTimer interface {
	dara.Model
	String() string
	GoString() string
	SetPeriods(v []*WafTimerPeriods) *WafTimer
	GetPeriods() []*WafTimerPeriods
	SetScopes(v string) *WafTimer
	GetScopes() *string
	SetWeeklyPeriods(v []*WafTimerWeeklyPeriods) *WafTimer
	GetWeeklyPeriods() []*WafTimerWeeklyPeriods
	SetZone(v int32) *WafTimer
	GetZone() *int32
}

type WafTimer struct {
	// The effective periods.
	Periods []*WafTimerPeriods `json:"Periods,omitempty" xml:"Periods,omitempty" type:"Repeated"`
	// The timer type: permanently effective (permanent/empty), effective during time periods (periods), or periodically effective (weekly).
	//
	// example:
	//
	// permanent
	Scopes *string `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
	// The weekly effective periods.
	WeeklyPeriods []*WafTimerWeeklyPeriods `json:"WeeklyPeriods,omitempty" xml:"WeeklyPeriods,omitempty" type:"Repeated"`
	// The time zone. Default value: UTC+00:00.<br>Example: 8 indicates UTC+8, -8 indicates UTC-8.<br>Valid range: -12 to +14.
	//
	// example:
	//
	// 8
	Zone *int32 `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s WafTimer) String() string {
	return dara.Prettify(s)
}

func (s WafTimer) GoString() string {
	return s.String()
}

func (s *WafTimer) GetPeriods() []*WafTimerPeriods {
	return s.Periods
}

func (s *WafTimer) GetScopes() *string {
	return s.Scopes
}

func (s *WafTimer) GetWeeklyPeriods() []*WafTimerWeeklyPeriods {
	return s.WeeklyPeriods
}

func (s *WafTimer) GetZone() *int32 {
	return s.Zone
}

func (s *WafTimer) SetPeriods(v []*WafTimerPeriods) *WafTimer {
	s.Periods = v
	return s
}

func (s *WafTimer) SetScopes(v string) *WafTimer {
	s.Scopes = &v
	return s
}

func (s *WafTimer) SetWeeklyPeriods(v []*WafTimerWeeklyPeriods) *WafTimer {
	s.WeeklyPeriods = v
	return s
}

func (s *WafTimer) SetZone(v int32) *WafTimer {
	s.Zone = &v
	return s
}

func (s *WafTimer) Validate() error {
	if s.Periods != nil {
		for _, item := range s.Periods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WeeklyPeriods != nil {
		for _, item := range s.WeeklyPeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type WafTimerPeriods struct {
	// The end time in RFC 3339 UTC format.
	//
	// example:
	//
	// 2025-01-01T01:00:00Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start time in RFC 3339 UTC format.
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s WafTimerPeriods) String() string {
	return dara.Prettify(s)
}

func (s WafTimerPeriods) GoString() string {
	return s.String()
}

func (s *WafTimerPeriods) GetEnd() *string {
	return s.End
}

func (s *WafTimerPeriods) GetStart() *string {
	return s.Start
}

func (s *WafTimerPeriods) SetEnd(v string) *WafTimerPeriods {
	s.End = &v
	return s
}

func (s *WafTimerPeriods) SetStart(v string) *WafTimerPeriods {
	s.Start = &v
	return s
}

func (s *WafTimerPeriods) Validate() error {
	return dara.Validate(s)
}

type WafTimerWeeklyPeriods struct {
	// The effective period within the epoch.
	DailyPeriods []*WafTimerWeeklyPeriodsDailyPeriods `json:"DailyPeriods,omitempty" xml:"DailyPeriods,omitempty" type:"Repeated"`
	// The days of the week, separated by commas. Values 1-7 represent Monday through Sunday.<br>Example: Monday and Wednesday is "1,3".
	//
	// example:
	//
	// 1
	Days *string `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s WafTimerWeeklyPeriods) String() string {
	return dara.Prettify(s)
}

func (s WafTimerWeeklyPeriods) GoString() string {
	return s.String()
}

func (s *WafTimerWeeklyPeriods) GetDailyPeriods() []*WafTimerWeeklyPeriodsDailyPeriods {
	return s.DailyPeriods
}

func (s *WafTimerWeeklyPeriods) GetDays() *string {
	return s.Days
}

func (s *WafTimerWeeklyPeriods) SetDailyPeriods(v []*WafTimerWeeklyPeriodsDailyPeriods) *WafTimerWeeklyPeriods {
	s.DailyPeriods = v
	return s
}

func (s *WafTimerWeeklyPeriods) SetDays(v string) *WafTimerWeeklyPeriods {
	s.Days = &v
	return s
}

func (s *WafTimerWeeklyPeriods) Validate() error {
	if s.DailyPeriods != nil {
		for _, item := range s.DailyPeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type WafTimerWeeklyPeriodsDailyPeriods struct {
	// The end time in HH:mm:ss format.
	//
	// example:
	//
	// 01:00:00
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start time in HH:mm:ss format.
	//
	// example:
	//
	// 00:00:00
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s WafTimerWeeklyPeriodsDailyPeriods) String() string {
	return dara.Prettify(s)
}

func (s WafTimerWeeklyPeriodsDailyPeriods) GoString() string {
	return s.String()
}

func (s *WafTimerWeeklyPeriodsDailyPeriods) GetEnd() *string {
	return s.End
}

func (s *WafTimerWeeklyPeriodsDailyPeriods) GetStart() *string {
	return s.Start
}

func (s *WafTimerWeeklyPeriodsDailyPeriods) SetEnd(v string) *WafTimerWeeklyPeriodsDailyPeriods {
	s.End = &v
	return s
}

func (s *WafTimerWeeklyPeriodsDailyPeriods) SetStart(v string) *WafTimerWeeklyPeriodsDailyPeriods {
	s.Start = &v
	return s
}

func (s *WafTimerWeeklyPeriodsDailyPeriods) Validate() error {
	return dara.Validate(s)
}
