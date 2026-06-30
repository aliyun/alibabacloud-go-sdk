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
	// The request ID.
	//
	// example:
	//
	// 68BD3312-53D8-123E-BB32-1A9F25E07A03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of usage data.
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
	// The usage data details.
	UsageInfo *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo `json:"UsageInfo,omitempty" xml:"UsageInfo,omitempty" type:"Struct"`
	// The usage primary key. When `UsageType=User`, this is the `aliUid`. When `UsageType=CreditPackage`, this is the credit package instance ID. When `UsageType=Agent`, this is the `AgentId`.
	//
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
	// The hourly consumption samples of the current credit package.
	CreditTrendList []*DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList `json:"CreditTrendList,omitempty" xml:"CreditTrendList,omitempty" type:"Repeated"`
	// The instance ID of the current active credit package.
	//
	// example:
	//
	// cp-inst-001
	CurrentInstanceId *string `json:"CurrentInstanceId,omitempty" xml:"CurrentInstanceId,omitempty"`
	// The remaining credits of the current active credit package.
	//
	// example:
	//
	// 当前周期积分余量
	CurrentRemainCredit *int64 `json:"CurrentRemainCredit,omitempty" xml:"CurrentRemainCredit,omitempty"`
	// The total credits of the current active credit package.
	//
	// example:
	//
	// 当前周期积分配额
	CurrentTotalCredit *int64 `json:"CurrentTotalCredit,omitempty" xml:"CurrentTotalCredit,omitempty"`
	// The used credits of the current active credit package.
	//
	// example:
	//
	// 当前周期积分消耗
	CurrentUsedCredit *int64 `json:"CurrentUsedCredit,omitempty" xml:"CurrentUsedCredit,omitempty"`
	// The credit usage in the last 1 day.
	//
	// example:
	//
	// 最近一天消耗积分
	DayUsedCredit *int64 `json:"DayUsedCredit,omitempty" xml:"DayUsedCredit,omitempty"`
	// The shared credit quota in the current active period.
	//
	// example:
	//
	// 300
	PeriodTotalCredit *int64 `json:"PeriodTotalCredit,omitempty" xml:"PeriodTotalCredit,omitempty"`
	// The shared credit usage in the current active period.
	//
	// example:
	//
	// 120
	PeriodUsedCredit *int64 `json:"PeriodUsedCredit,omitempty" xml:"PeriodUsedCredit,omitempty"`
	// The cumulative remaining credits.
	//
	// example:
	//
	// 积分余量
	RemainCredit *int64  `json:"RemainCredit,omitempty" xml:"RemainCredit,omitempty"`
	TodayUsed    *string `json:"TodayUsed,omitempty" xml:"TodayUsed,omitempty"`
	// The cumulative total credits.
	//
	// example:
	//
	// 积分配额
	TotalCredit *int64  `json:"TotalCredit,omitempty" xml:"TotalCredit,omitempty"`
	TotalUsed   *string `json:"TotalUsed,omitempty" xml:"TotalUsed,omitempty"`
	// The cumulative credit usage.
	//
	// example:
	//
	// 共计消耗积分
	TotalUsedCredit *int64 `json:"TotalUsedCredit,omitempty" xml:"TotalUsedCredit,omitempty"`
	// The alert threshold percentage (0–100).
	//
	// example:
	//
	// 80
	WarnPercent *int32 `json:"WarnPercent,omitempty" xml:"WarnPercent,omitempty"`
	// The credit usage in the last 1 week.
	//
	// example:
	//
	// 最近一周消耗积分
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetTodayUsed() *string {
	return s.TodayUsed
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetTotalCredit() *int64 {
	return s.TotalCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetTotalUsed() *string {
	return s.TotalUsed
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetTodayUsed(v string) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.TodayUsed = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetTotalCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.TotalCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetTotalUsed(v string) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.TotalUsed = &v
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
	// The time point in the format of `yyyy-MM-dd HH` (accurate to the hour).
	//
	// example:
	//
	// 2026-05-02 10
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	// The number of credits consumed during the hour.
	//
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
