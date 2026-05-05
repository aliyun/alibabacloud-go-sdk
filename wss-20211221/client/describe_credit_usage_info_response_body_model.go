// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditUsageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCreditUsageInfoResponseBody
	GetRequestId() *string
	SetUsageInfoList(v []*DescribeCreditUsageInfoResponseBodyUsageInfoList) *DescribeCreditUsageInfoResponseBody
	GetUsageInfoList() []*DescribeCreditUsageInfoResponseBodyUsageInfoList
}

type DescribeCreditUsageInfoResponseBody struct {
	// example:
	//
	// 68BD3312-53D8-123E-BB32-1A9F25E07A03
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageInfoList []*DescribeCreditUsageInfoResponseBodyUsageInfoList `json:"UsageInfoList,omitempty" xml:"UsageInfoList,omitempty" type:"Repeated"`
}

func (s DescribeCreditUsageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCreditUsageInfoResponseBody) GetUsageInfoList() []*DescribeCreditUsageInfoResponseBodyUsageInfoList {
	return s.UsageInfoList
}

func (s *DescribeCreditUsageInfoResponseBody) SetRequestId(v string) *DescribeCreditUsageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBody) SetUsageInfoList(v []*DescribeCreditUsageInfoResponseBodyUsageInfoList) *DescribeCreditUsageInfoResponseBody {
	s.UsageInfoList = v
	return s
}

func (s *DescribeCreditUsageInfoResponseBody) Validate() error {
	if s.UsageInfoList != nil {
		for _, item := range s.UsageInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCreditUsageInfoResponseBodyUsageInfoList struct {
	UsageInfo *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo `json:"UsageInfo,omitempty" xml:"UsageInfo,omitempty" type:"Struct"`
	// example:
	//
	// agent-abc
	UsageInfoKey *string `json:"UsageInfoKey,omitempty" xml:"UsageInfoKey,omitempty"`
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoList) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoList) GetUsageInfo() *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	return s.UsageInfo
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoList) GetUsageInfoKey() *string {
	return s.UsageInfoKey
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoList) SetUsageInfo(v *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) *DescribeCreditUsageInfoResponseBodyUsageInfoList {
	s.UsageInfo = v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoList) SetUsageInfoKey(v string) *DescribeCreditUsageInfoResponseBodyUsageInfoList {
	s.UsageInfoKey = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoList) Validate() error {
	if s.UsageInfo != nil {
		if err := s.UsageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo struct {
	CreditTrendList []*DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList `json:"CreditTrendList,omitempty" xml:"CreditTrendList,omitempty" type:"Repeated"`
	// example:
	//
	// cp-inst-001
	CurrentInstanceId   *string `json:"CurrentInstanceId,omitempty" xml:"CurrentInstanceId,omitempty"`
	CurrentRemainCredit *int64  `json:"CurrentRemainCredit,omitempty" xml:"CurrentRemainCredit,omitempty"`
	CurrentTotalCredit  *int64  `json:"CurrentTotalCredit,omitempty" xml:"CurrentTotalCredit,omitempty"`
	CurrentUsedCredit   *int64  `json:"CurrentUsedCredit,omitempty" xml:"CurrentUsedCredit,omitempty"`
	DayUsedCredit       *int64  `json:"DayUsedCredit,omitempty" xml:"DayUsedCredit,omitempty"`
	// example:
	//
	// 300
	PeriodTotalCredit *int64 `json:"PeriodTotalCredit,omitempty" xml:"PeriodTotalCredit,omitempty"`
	// example:
	//
	// 120
	PeriodUsedCredit *int64 `json:"PeriodUsedCredit,omitempty" xml:"PeriodUsedCredit,omitempty"`
	RemainCredit     *int64 `json:"RemainCredit,omitempty" xml:"RemainCredit,omitempty"`
	TotalCredit      *int64 `json:"TotalCredit,omitempty" xml:"TotalCredit,omitempty"`
	TotalUsedCredit  *int64 `json:"TotalUsedCredit,omitempty" xml:"TotalUsedCredit,omitempty"`
	// example:
	//
	// 80
	WarnPercent    *int32 `json:"WarnPercent,omitempty" xml:"WarnPercent,omitempty"`
	WeekUsedCredit *int64 `json:"WeekUsedCredit,omitempty" xml:"WeekUsedCredit,omitempty"`
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetCreditTrendList() []*DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList {
	return s.CreditTrendList
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetCurrentInstanceId() *string {
	return s.CurrentInstanceId
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetCurrentRemainCredit() *int64 {
	return s.CurrentRemainCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetCurrentTotalCredit() *int64 {
	return s.CurrentTotalCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetCurrentUsedCredit() *int64 {
	return s.CurrentUsedCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetDayUsedCredit() *int64 {
	return s.DayUsedCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetPeriodTotalCredit() *int64 {
	return s.PeriodTotalCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetPeriodUsedCredit() *int64 {
	return s.PeriodUsedCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetRemainCredit() *int64 {
	return s.RemainCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetTotalCredit() *int64 {
	return s.TotalCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetTotalUsedCredit() *int64 {
	return s.TotalUsedCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetWarnPercent() *int32 {
	return s.WarnPercent
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetWeekUsedCredit() *int64 {
	return s.WeekUsedCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetCreditTrendList(v []*DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.CreditTrendList = v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetCurrentInstanceId(v string) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.CurrentInstanceId = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetCurrentRemainCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.CurrentRemainCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetCurrentTotalCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.CurrentTotalCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetCurrentUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.CurrentUsedCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetDayUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.DayUsedCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetPeriodTotalCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.PeriodTotalCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetPeriodUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.PeriodUsedCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetRemainCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.RemainCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetTotalCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.TotalCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetTotalUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.TotalUsedCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetWarnPercent(v int32) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.WarnPercent = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetWeekUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.WeekUsedCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) Validate() error {
	if s.CreditTrendList != nil {
		for _, item := range s.CreditTrendList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList struct {
	// example:
	//
	// 2026-05-02 10
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	// example:
	//
	// 12
	UsedCredit *int64 `json:"UsedCredit,omitempty" xml:"UsedCredit,omitempty"`
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList) GetTimePoint() *string {
	return s.TimePoint
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList) GetUsedCredit() *int64 {
	return s.UsedCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList) SetTimePoint(v string) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList {
	s.TimePoint = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList) SetUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList {
	s.UsedCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList) Validate() error {
	return dara.Validate(s)
}
