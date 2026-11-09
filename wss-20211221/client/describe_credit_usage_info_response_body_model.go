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
	// The business data array.
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
	// The primary key for usage. The value is `aliUid` when `UsageType` is `User`, the credit package instance ID when `UsageType` is `CreditPackage`, and the agent ID when `UsageType` is `Agent`.
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
	// The number of currently available instances for the reported resources.
	//
	// example:
	//
	// 9976.00
	AvailableAmount *int32 `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// The list of contact group names.
	ContactGroupNames []*string `json:"ContactGroupNames,omitempty" xml:"ContactGroupNames,omitempty" type:"Repeated"`
	// The hourly consumption samples of the current credit package.
	CreditTrendList []*DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoCreditTrendList `json:"CreditTrendList,omitempty" xml:"CreditTrendList,omitempty" type:"Repeated"`
	// The time when the current instance expires, in the `YYYY-MM-DD HH:mm:ss` format.
	//
	// example:
	//
	// 2026-12-02 00:00:00
	CurrentInstanceExpiredTime *string `json:"CurrentInstanceExpiredTime,omitempty" xml:"CurrentInstanceExpiredTime,omitempty"`
	// The instance ID of the currently active credit package.
	//
	// example:
	//
	// cp-inst-001
	CurrentInstanceId *string `json:"CurrentInstanceId,omitempty" xml:"CurrentInstanceId,omitempty"`
	// The remaining amount of the currently active credit package.
	//
	// example:
	//
	// Remaining credits for the current period
	CurrentRemainCredit *int64 `json:"CurrentRemainCredit,omitempty" xml:"CurrentRemainCredit,omitempty"`
	// The total amount of the currently active credit package.
	//
	// example:
	//
	// Credit quota for the current period
	CurrentTotalCredit *int64 `json:"CurrentTotalCredit,omitempty" xml:"CurrentTotalCredit,omitempty"`
	// The used amount of the currently active credit package.
	//
	// example:
	//
	// Credits consumed in the current period
	CurrentUsedCredit *int64 `json:"CurrentUsedCredit,omitempty" xml:"CurrentUsedCredit,omitempty"`
	// The credit usage in the last day.
	//
	// example:
	//
	// Credits consumed in the last day
	DayUsedCredit *int64 `json:"DayUsedCredit,omitempty" xml:"DayUsedCredit,omitempty"`
	// The last trigger time.
	//
	// example:
	//
	// 2026-05-02 10:10:10
	LastTriggeredAt *string `json:"LastTriggeredAt,omitempty" xml:"LastTriggeredAt,omitempty"`
	// The credit quota for the current period.
	//
	// example:
	//
	// 2000
	PeriodCreditQuota *int64 `json:"PeriodCreditQuota,omitempty" xml:"PeriodCreditQuota,omitempty"`
	// The pay-as-you-go credits consumed in the current period.
	//
	// example:
	//
	// 240
	PeriodPayGoUsedCredit *int64 `json:"PeriodPayGoUsedCredit,omitempty" xml:"PeriodPayGoUsedCredit,omitempty"`
	// The shared credits consumed in the current period.
	//
	// example:
	//
	// 120
	PeriodSharedUsedCredit *int64 `json:"PeriodSharedUsedCredit,omitempty" xml:"PeriodSharedUsedCredit,omitempty"`
	// The shared credit quota in the current active period.
	//
	// example:
	//
	// 300
	PeriodTotalCredit *int64 `json:"PeriodTotalCredit,omitempty" xml:"PeriodTotalCredit,omitempty"`
	// The total credits consumed in the current period.
	//
	// example:
	//
	// 360
	PeriodTotalUsedCredit *int64 `json:"PeriodTotalUsedCredit,omitempty" xml:"PeriodTotalUsedCredit,omitempty"`
	// The shared credit usage in the current active period.
	//
	// example:
	//
	// 120
	PeriodUsedCredit *int64 `json:"PeriodUsedCredit,omitempty" xml:"PeriodUsedCredit,omitempty"`
	// The cumulative remaining billable credits.
	//
	// example:
	//
	// Remaining credits
	RemainCredit *int64 `json:"RemainCredit,omitempty" xml:"RemainCredit,omitempty"`
	// The remaining credit information.
	RemainCreditInfo *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfoRemainCreditInfo `json:"RemainCreditInfo,omitempty" xml:"RemainCreditInfo,omitempty" type:"Struct"`
	// The used quota for today.
	//
	// example:
	//
	// 200
	TodayUsed *string `json:"TodayUsed,omitempty" xml:"TodayUsed,omitempty"`
	// The cumulative total credits.
	//
	// example:
	//
	// Credit quota
	TotalCredit *int64 `json:"TotalCredit,omitempty" xml:"TotalCredit,omitempty"`
	// The cumulative used quota.
	//
	// example:
	//
	// 100
	TotalUsed *string `json:"TotalUsed,omitempty" xml:"TotalUsed,omitempty"`
	// The cumulative credit usage.
	//
	// example:
	//
	// Total credits consumed
	TotalUsedCredit *int64 `json:"TotalUsedCredit,omitempty" xml:"TotalUsedCredit,omitempty"`
	// The alert threshold percentage, ranging from 0 to 100.
	//
	// example:
	//
	// 80
	WarnPercent *int32 `json:"WarnPercent,omitempty" xml:"WarnPercent,omitempty"`
	// The credit usage in the last week.
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetCurrentInstanceExpiredTime() *string {
	return s.CurrentInstanceExpiredTime
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetPeriodCreditQuota() *int64 {
	return s.PeriodCreditQuota
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetPeriodPayGoUsedCredit() *int64 {
	return s.PeriodPayGoUsedCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetPeriodSharedUsedCredit() *int64 {
	return s.PeriodSharedUsedCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetPeriodTotalCredit() *int64 {
	return s.PeriodTotalCredit
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) GetPeriodTotalUsedCredit() *int64 {
	return s.PeriodTotalUsedCredit
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetCurrentInstanceExpiredTime(v string) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.CurrentInstanceExpiredTime = &v
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

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetPeriodCreditQuota(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.PeriodCreditQuota = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetPeriodPayGoUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.PeriodPayGoUsedCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetPeriodSharedUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.PeriodSharedUsedCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetPeriodTotalCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.PeriodTotalCredit = &v
	return s
}

func (s *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo) SetPeriodTotalUsedCredit(v int64) *DescribeCreditUsageInfoResponseBodyUsageInfoListUsageInfo {
	s.PeriodTotalUsedCredit = &v
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
	// The time point in the `yyyy-MM-dd HH` format, accurate to the hour.
	//
	// example:
	//
	// 2026-05-02 10
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	// The number of credits consumed in the hour.
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
	// The number of instances being deducted.
	//
	// example:
	//
	// 23
	DeductingAmount *int32 `json:"DeductingAmount,omitempty" xml:"DeductingAmount,omitempty"`
	// The number of instances pending to take effect.
	//
	// example:
	//
	// 24
	PendingAmount *int32 `json:"PendingAmount,omitempty" xml:"PendingAmount,omitempty"`
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
