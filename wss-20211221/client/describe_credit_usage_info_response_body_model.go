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
	AvailableAmount   *int32    `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	ContactGroupNames []*string `json:"ContactGroupNames,omitempty" xml:"ContactGroupNames,omitempty" type:"Repeated"`
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
	// Current period remaining credits
	CurrentRemainCredit *int64 `json:"CurrentRemainCredit,omitempty" xml:"CurrentRemainCredit,omitempty"`
	// The total credits of the current active credit package.
	//
	// example:
	//
	// Current period credit quota
	CurrentTotalCredit *int64 `json:"CurrentTotalCredit,omitempty" xml:"CurrentTotalCredit,omitempty"`
	// The used credits of the current active credit package.
	//
	// example:
	//
	// Current period credits consumed
	CurrentUsedCredit *int64 `json:"CurrentUsedCredit,omitempty" xml:"CurrentUsedCredit,omitempty"`
	// The credit usage in the last 1 day.
	//
	// example:
	//
	// Credits consumed in the last day
	DayUsedCredit   *int64  `json:"DayUsedCredit,omitempty" xml:"DayUsedCredit,omitempty"`
	LastTriggeredAt *string `json:"LastTriggeredAt,omitempty" xml:"LastTriggeredAt,omitempty"`
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
	// Remaining credits
	RemainCredit     *int64                                                                     `json:"RemainCredit,omitempty" xml:"RemainCredit,omitempty"`
	RemainCreditInfo *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo `json:"RemainCreditInfo,omitempty" xml:"RemainCreditInfo,omitempty" type:"Struct"`
	// The quota used today.
	TodayUsed *string `json:"TodayUsed,omitempty" xml:"TodayUsed,omitempty"`
	// The total cumulative credits.
	//
	// example:
	//
	// Credit quota
	TotalCredit *int64 `json:"TotalCredit,omitempty" xml:"TotalCredit,omitempty"`
	// The cumulative used quota.
	TotalUsed *string `json:"TotalUsed,omitempty" xml:"TotalUsed,omitempty"`
	// The cumulative credit usage.
	//
	// example:
	//
	// Total credits consumed
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
	// Credits consumed in the last week
	WeekUsedCredit *int64 `json:"WeekUsedCredit,omitempty" xml:"WeekUsedCredit,omitempty"`
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetAvailableAmount() *int32 {
	return s.AvailableAmount
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetContactGroupNames() []*string {
	return s.ContactGroupNames
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetLastTriggeredAt() *string {
	return s.LastTriggeredAt
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetRemainCreditInfo() *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo {
	return s.RemainCreditInfo
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetAvailableAmount(v int32) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.AvailableAmount = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetContactGroupNames(v []*string) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.ContactGroupNames = v
	return s
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetLastTriggeredAt(v string) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.LastTriggeredAt = &v
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetRemainCreditInfo(v *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.RemainCreditInfo = v
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
	if s.RemainCreditInfo != nil {
		if err := s.RemainCreditInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList struct {
	// The time point in the format `yyyy-MM-dd HH` (accurate to the hour).
	//
	// example:
	//
	// 2026-05-02 10
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	// The number of credits consumed in this hour.
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

type DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo struct {
	DeductingAmount *int32 `json:"DeductingAmount,omitempty" xml:"DeductingAmount,omitempty"`
	PendingAmount   *int32 `json:"PendingAmount,omitempty" xml:"PendingAmount,omitempty"`
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo) GetDeductingAmount() *int32 {
	return s.DeductingAmount
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo) GetPendingAmount() *int32 {
	return s.PendingAmount
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo) SetDeductingAmount(v int32) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo {
	s.DeductingAmount = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo) SetPendingAmount(v int32) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo {
	s.PendingAmount = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo) Validate() error {
	return dara.Validate(s)
}
