// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeBudgetsResponseBody
	GetCurrentPage() *int32
	SetData(v []*DescribeBudgetsResponseBodyData) *DescribeBudgetsResponseBody
	GetData() []*DescribeBudgetsResponseBodyData
	SetPageSize(v int32) *DescribeBudgetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBudgetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBudgetsResponseBody
	GetTotalCount() *int32
}

type DescribeBudgetsResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*DescribeBudgetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// DB2A9097-289C-11CE-AA74-235FCFD39204
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBudgetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBudgetsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBudgetsResponseBody) GetData() []*DescribeBudgetsResponseBodyData {
	return s.Data
}

func (s *DescribeBudgetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBudgetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBudgetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBudgetsResponseBody) SetCurrentPage(v int32) *DescribeBudgetsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBudgetsResponseBody) SetData(v []*DescribeBudgetsResponseBodyData) *DescribeBudgetsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBudgetsResponseBody) SetPageSize(v int32) *DescribeBudgetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBudgetsResponseBody) SetRequestId(v string) *DescribeBudgetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBudgetsResponseBody) SetTotalCount(v int32) *DescribeBudgetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBudgetsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBudgetsResponseBodyData struct {
	// example:
	//
	// department1-test
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// example:
	//
	// CONSUME
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	Comment    *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 2026-12
	CycleEndPeriod *string                                      `json:"CycleEndPeriod,omitempty" xml:"CycleEndPeriod,omitempty"`
	CycleQuota     []*DescribeBudgetsResponseBodyDataCycleQuota `json:"CycleQuota,omitempty" xml:"CycleQuota,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-10
	CycleStartPeriod *string `json:"CycleStartPeriod,omitempty" xml:"CycleStartPeriod,omitempty"`
	// example:
	//
	// MONTHLY
	CycleType            *string                                              `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	EcIdAccountIdsFilter *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter `json:"EcIdAccountIdsFilter,omitempty" xml:"EcIdAccountIdsFilter,omitempty" type:"Struct"`
	// example:
	//
	// NOT_EXPIRED
	ExpireStatus *string `json:"ExpireStatus,omitempty" xml:"ExpireStatus,omitempty"`
	// example:
	//
	// REQUIRE_AMOUNT
	Metric      *string                                       `json:"Metric,omitempty" xml:"Metric,omitempty"`
	QueryFilter []*DescribeBudgetsResponseBodyDataQueryFilter `json:"QueryFilter,omitempty" xml:"QueryFilter,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// FIXED
	QuotaType *string                                     `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	WarnConfs []*DescribeBudgetsResponseBodyDataWarnConfs `json:"WarnConfs,omitempty" xml:"WarnConfs,omitempty" type:"Repeated"`
}

func (s DescribeBudgetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBudgetsResponseBodyData) GetBudgetName() *string {
	return s.BudgetName
}

func (s *DescribeBudgetsResponseBodyData) GetBudgetType() *string {
	return s.BudgetType
}

func (s *DescribeBudgetsResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *DescribeBudgetsResponseBodyData) GetCycleEndPeriod() *string {
	return s.CycleEndPeriod
}

func (s *DescribeBudgetsResponseBodyData) GetCycleQuota() []*DescribeBudgetsResponseBodyDataCycleQuota {
	return s.CycleQuota
}

func (s *DescribeBudgetsResponseBodyData) GetCycleStartPeriod() *string {
	return s.CycleStartPeriod
}

func (s *DescribeBudgetsResponseBodyData) GetCycleType() *string {
	return s.CycleType
}

func (s *DescribeBudgetsResponseBodyData) GetEcIdAccountIdsFilter() *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter {
	return s.EcIdAccountIdsFilter
}

func (s *DescribeBudgetsResponseBodyData) GetExpireStatus() *string {
	return s.ExpireStatus
}

func (s *DescribeBudgetsResponseBodyData) GetMetric() *string {
	return s.Metric
}

func (s *DescribeBudgetsResponseBodyData) GetQueryFilter() []*DescribeBudgetsResponseBodyDataQueryFilter {
	return s.QueryFilter
}

func (s *DescribeBudgetsResponseBodyData) GetQuota() *string {
	return s.Quota
}

func (s *DescribeBudgetsResponseBodyData) GetQuotaType() *string {
	return s.QuotaType
}

func (s *DescribeBudgetsResponseBodyData) GetWarnConfs() []*DescribeBudgetsResponseBodyDataWarnConfs {
	return s.WarnConfs
}

func (s *DescribeBudgetsResponseBodyData) SetBudgetName(v string) *DescribeBudgetsResponseBodyData {
	s.BudgetName = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetBudgetType(v string) *DescribeBudgetsResponseBodyData {
	s.BudgetType = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetComment(v string) *DescribeBudgetsResponseBodyData {
	s.Comment = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetCycleEndPeriod(v string) *DescribeBudgetsResponseBodyData {
	s.CycleEndPeriod = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetCycleQuota(v []*DescribeBudgetsResponseBodyDataCycleQuota) *DescribeBudgetsResponseBodyData {
	s.CycleQuota = v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetCycleStartPeriod(v string) *DescribeBudgetsResponseBodyData {
	s.CycleStartPeriod = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetCycleType(v string) *DescribeBudgetsResponseBodyData {
	s.CycleType = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetEcIdAccountIdsFilter(v *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter) *DescribeBudgetsResponseBodyData {
	s.EcIdAccountIdsFilter = v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetExpireStatus(v string) *DescribeBudgetsResponseBodyData {
	s.ExpireStatus = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetMetric(v string) *DescribeBudgetsResponseBodyData {
	s.Metric = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetQueryFilter(v []*DescribeBudgetsResponseBodyDataQueryFilter) *DescribeBudgetsResponseBodyData {
	s.QueryFilter = v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetQuota(v string) *DescribeBudgetsResponseBodyData {
	s.Quota = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetQuotaType(v string) *DescribeBudgetsResponseBodyData {
	s.QuotaType = &v
	return s
}

func (s *DescribeBudgetsResponseBodyData) SetWarnConfs(v []*DescribeBudgetsResponseBodyDataWarnConfs) *DescribeBudgetsResponseBodyData {
	s.WarnConfs = v
	return s
}

func (s *DescribeBudgetsResponseBodyData) Validate() error {
	if s.CycleQuota != nil {
		for _, item := range s.CycleQuota {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EcIdAccountIdsFilter != nil {
		if err := s.EcIdAccountIdsFilter.Validate(); err != nil {
			return err
		}
	}
	if s.QueryFilter != nil {
		for _, item := range s.QueryFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WarnConfs != nil {
		for _, item := range s.WarnConfs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBudgetsResponseBodyDataCycleQuota struct {
	CyclePeriod *string `json:"CyclePeriod,omitempty" xml:"CyclePeriod,omitempty"`
	Quota       *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s DescribeBudgetsResponseBodyDataCycleQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetsResponseBodyDataCycleQuota) GoString() string {
	return s.String()
}

func (s *DescribeBudgetsResponseBodyDataCycleQuota) GetCyclePeriod() *string {
	return s.CyclePeriod
}

func (s *DescribeBudgetsResponseBodyDataCycleQuota) GetQuota() *string {
	return s.Quota
}

func (s *DescribeBudgetsResponseBodyDataCycleQuota) SetCyclePeriod(v string) *DescribeBudgetsResponseBodyDataCycleQuota {
	s.CyclePeriod = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataCycleQuota) SetQuota(v string) *DescribeBudgetsResponseBodyDataCycleQuota {
	s.Quota = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataCycleQuota) Validate() error {
	return dara.Validate(s)
}

type DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	EcId       *string  `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter) GoString() string {
	return s.String()
}

func (s *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter) GetEcId() *string {
	return s.EcId
}

func (s *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter) SetAccountIds(v []*int64) *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter {
	s.AccountIds = v
	return s
}

func (s *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter) SetEcId(v string) *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter {
	s.EcId = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataEcIdAccountIdsFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeBudgetsResponseBodyDataQueryFilter struct {
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeBudgetsResponseBodyDataQueryFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetsResponseBodyDataQueryFilter) GoString() string {
	return s.String()
}

func (s *DescribeBudgetsResponseBodyDataQueryFilter) GetCode() *string {
	return s.Code
}

func (s *DescribeBudgetsResponseBodyDataQueryFilter) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeBudgetsResponseBodyDataQueryFilter) GetValues() []*string {
	return s.Values
}

func (s *DescribeBudgetsResponseBodyDataQueryFilter) SetCode(v string) *DescribeBudgetsResponseBodyDataQueryFilter {
	s.Code = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataQueryFilter) SetSelectType(v string) *DescribeBudgetsResponseBodyDataQueryFilter {
	s.SelectType = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataQueryFilter) SetValues(v []*string) *DescribeBudgetsResponseBodyDataQueryFilter {
	s.Values = v
	return s
}

func (s *DescribeBudgetsResponseBodyDataQueryFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeBudgetsResponseBodyDataWarnConfs struct {
	Comment        *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	EventBridge    *bool     `json:"EventBridge,omitempty" xml:"EventBridge,omitempty"`
	MscChannels    []*string `json:"MscChannels,omitempty" xml:"MscChannels,omitempty" type:"Repeated"`
	MscContacts    []*string `json:"MscContacts,omitempty" xml:"MscContacts,omitempty" type:"Repeated"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Sequence       *int32    `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	ThresholdType  *string   `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	ThresholdValue *string   `json:"ThresholdValue,omitempty" xml:"ThresholdValue,omitempty"`
	WarnTarget     *string   `json:"WarnTarget,omitempty" xml:"WarnTarget,omitempty"`
}

func (s DescribeBudgetsResponseBodyDataWarnConfs) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetsResponseBodyDataWarnConfs) GoString() string {
	return s.String()
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetComment() *string {
	return s.Comment
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetEventBridge() *bool {
	return s.EventBridge
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetMscChannels() []*string {
	return s.MscChannels
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetMscContacts() []*string {
	return s.MscContacts
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetName() *string {
	return s.Name
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetSequence() *int32 {
	return s.Sequence
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetThresholdType() *string {
	return s.ThresholdType
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetThresholdValue() *string {
	return s.ThresholdValue
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) GetWarnTarget() *string {
	return s.WarnTarget
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetComment(v string) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.Comment = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetEventBridge(v bool) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.EventBridge = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetMscChannels(v []*string) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.MscChannels = v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetMscContacts(v []*string) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.MscContacts = v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetName(v string) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.Name = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetSequence(v int32) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.Sequence = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetThresholdType(v string) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.ThresholdType = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetThresholdValue(v string) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.ThresholdValue = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) SetWarnTarget(v string) *DescribeBudgetsResponseBodyDataWarnConfs {
	s.WarnTarget = &v
	return s
}

func (s *DescribeBudgetsResponseBodyDataWarnConfs) Validate() error {
	return dara.Validate(s)
}
