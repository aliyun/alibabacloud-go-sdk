// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *DescribeBudgetResponseBody
	GetBudgetName() *string
	SetBudgetType(v string) *DescribeBudgetResponseBody
	GetBudgetType() *string
	SetComment(v string) *DescribeBudgetResponseBody
	GetComment() *string
	SetCycleEndPeriod(v string) *DescribeBudgetResponseBody
	GetCycleEndPeriod() *string
	SetCycleQuota(v []*DescribeBudgetResponseBodyCycleQuota) *DescribeBudgetResponseBody
	GetCycleQuota() []*DescribeBudgetResponseBodyCycleQuota
	SetCycleStartPeriod(v string) *DescribeBudgetResponseBody
	GetCycleStartPeriod() *string
	SetCycleType(v string) *DescribeBudgetResponseBody
	GetCycleType() *string
	SetEcIdAccountIdsFilter(v *DescribeBudgetResponseBodyEcIdAccountIdsFilter) *DescribeBudgetResponseBody
	GetEcIdAccountIdsFilter() *DescribeBudgetResponseBodyEcIdAccountIdsFilter
	SetMetadata(v interface{}) *DescribeBudgetResponseBody
	GetMetadata() interface{}
	SetMetric(v string) *DescribeBudgetResponseBody
	GetMetric() *string
	SetQueryFilter(v []*DescribeBudgetResponseBodyQueryFilter) *DescribeBudgetResponseBody
	GetQueryFilter() []*DescribeBudgetResponseBodyQueryFilter
	SetQuota(v string) *DescribeBudgetResponseBody
	GetQuota() *string
	SetQuotaType(v string) *DescribeBudgetResponseBody
	GetQuotaType() *string
	SetRequestId(v string) *DescribeBudgetResponseBody
	GetRequestId() *string
	SetWarnConfs(v []*DescribeBudgetResponseBodyWarnConfs) *DescribeBudgetResponseBody
	GetWarnConfs() []*DescribeBudgetResponseBodyWarnConfs
}

type DescribeBudgetResponseBody struct {
	// Budget name.
	//
	// example:
	//
	// department1
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// Budget type.
	//
	// example:
	//
	// CONSUME
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// Remarks.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// End cycle period.
	//
	// example:
	//
	// 2026-12
	CycleEndPeriod *string `json:"CycleEndPeriod,omitempty" xml:"CycleEndPeriod,omitempty"`
	// Quota specified per cycle.
	CycleQuota []*DescribeBudgetResponseBodyCycleQuota `json:"CycleQuota,omitempty" xml:"CycleQuota,omitempty" type:"Repeated"`
	// Start cycle period.
	//
	// example:
	//
	// 2025-10
	CycleStartPeriod *string `json:"CycleStartPeriod,omitempty" xml:"CycleStartPeriod,omitempty"`
	// Cycle type.
	//
	// example:
	//
	// MONTHLY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// Enterprise multi-account filter conditions.
	EcIdAccountIdsFilter *DescribeBudgetResponseBodyEcIdAccountIdsFilter `json:"EcIdAccountIdsFilter,omitempty" xml:"EcIdAccountIdsFilter,omitempty" type:"Struct"`
	// Response structure metadata.
	//
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Budget metric.
	//
	// example:
	//
	// REQUIRE_AMOUNT
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// Filter conditions.
	QueryFilter []*DescribeBudgetResponseBodyQueryFilter `json:"QueryFilter,omitempty" xml:"QueryFilter,omitempty" type:"Repeated"`
	// Fixed quota value.
	//
	// example:
	//
	// 1000
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// Quota type.
	//
	// example:
	//
	// FIXED
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Alert configurations.
	WarnConfs []*DescribeBudgetResponseBodyWarnConfs `json:"WarnConfs,omitempty" xml:"WarnConfs,omitempty" type:"Repeated"`
}

func (s DescribeBudgetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBudgetResponseBody) GetBudgetName() *string {
	return s.BudgetName
}

func (s *DescribeBudgetResponseBody) GetBudgetType() *string {
	return s.BudgetType
}

func (s *DescribeBudgetResponseBody) GetComment() *string {
	return s.Comment
}

func (s *DescribeBudgetResponseBody) GetCycleEndPeriod() *string {
	return s.CycleEndPeriod
}

func (s *DescribeBudgetResponseBody) GetCycleQuota() []*DescribeBudgetResponseBodyCycleQuota {
	return s.CycleQuota
}

func (s *DescribeBudgetResponseBody) GetCycleStartPeriod() *string {
	return s.CycleStartPeriod
}

func (s *DescribeBudgetResponseBody) GetCycleType() *string {
	return s.CycleType
}

func (s *DescribeBudgetResponseBody) GetEcIdAccountIdsFilter() *DescribeBudgetResponseBodyEcIdAccountIdsFilter {
	return s.EcIdAccountIdsFilter
}

func (s *DescribeBudgetResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *DescribeBudgetResponseBody) GetMetric() *string {
	return s.Metric
}

func (s *DescribeBudgetResponseBody) GetQueryFilter() []*DescribeBudgetResponseBodyQueryFilter {
	return s.QueryFilter
}

func (s *DescribeBudgetResponseBody) GetQuota() *string {
	return s.Quota
}

func (s *DescribeBudgetResponseBody) GetQuotaType() *string {
	return s.QuotaType
}

func (s *DescribeBudgetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBudgetResponseBody) GetWarnConfs() []*DescribeBudgetResponseBodyWarnConfs {
	return s.WarnConfs
}

func (s *DescribeBudgetResponseBody) SetBudgetName(v string) *DescribeBudgetResponseBody {
	s.BudgetName = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetBudgetType(v string) *DescribeBudgetResponseBody {
	s.BudgetType = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetComment(v string) *DescribeBudgetResponseBody {
	s.Comment = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetCycleEndPeriod(v string) *DescribeBudgetResponseBody {
	s.CycleEndPeriod = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetCycleQuota(v []*DescribeBudgetResponseBodyCycleQuota) *DescribeBudgetResponseBody {
	s.CycleQuota = v
	return s
}

func (s *DescribeBudgetResponseBody) SetCycleStartPeriod(v string) *DescribeBudgetResponseBody {
	s.CycleStartPeriod = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetCycleType(v string) *DescribeBudgetResponseBody {
	s.CycleType = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetEcIdAccountIdsFilter(v *DescribeBudgetResponseBodyEcIdAccountIdsFilter) *DescribeBudgetResponseBody {
	s.EcIdAccountIdsFilter = v
	return s
}

func (s *DescribeBudgetResponseBody) SetMetadata(v interface{}) *DescribeBudgetResponseBody {
	s.Metadata = v
	return s
}

func (s *DescribeBudgetResponseBody) SetMetric(v string) *DescribeBudgetResponseBody {
	s.Metric = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetQueryFilter(v []*DescribeBudgetResponseBodyQueryFilter) *DescribeBudgetResponseBody {
	s.QueryFilter = v
	return s
}

func (s *DescribeBudgetResponseBody) SetQuota(v string) *DescribeBudgetResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetQuotaType(v string) *DescribeBudgetResponseBody {
	s.QuotaType = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetRequestId(v string) *DescribeBudgetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBudgetResponseBody) SetWarnConfs(v []*DescribeBudgetResponseBodyWarnConfs) *DescribeBudgetResponseBody {
	s.WarnConfs = v
	return s
}

func (s *DescribeBudgetResponseBody) Validate() error {
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

type DescribeBudgetResponseBodyCycleQuota struct {
	// Cycle period.
	//
	// example:
	//
	// 202601
	CyclePeriod *string `json:"CyclePeriod,omitempty" xml:"CyclePeriod,omitempty"`
	// Quota.
	//
	// example:
	//
	// 100
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s DescribeBudgetResponseBodyCycleQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetResponseBodyCycleQuota) GoString() string {
	return s.String()
}

func (s *DescribeBudgetResponseBodyCycleQuota) GetCyclePeriod() *string {
	return s.CyclePeriod
}

func (s *DescribeBudgetResponseBodyCycleQuota) GetQuota() *string {
	return s.Quota
}

func (s *DescribeBudgetResponseBodyCycleQuota) SetCyclePeriod(v string) *DescribeBudgetResponseBodyCycleQuota {
	s.CyclePeriod = &v
	return s
}

func (s *DescribeBudgetResponseBodyCycleQuota) SetQuota(v string) *DescribeBudgetResponseBodyCycleQuota {
	s.Quota = &v
	return s
}

func (s *DescribeBudgetResponseBodyCycleQuota) Validate() error {
	return dara.Validate(s)
}

type DescribeBudgetResponseBodyEcIdAccountIdsFilter struct {
	// Member account IDs.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// Enterprise entity ID.
	//
	// example:
	//
	// E2024112210463400001
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeBudgetResponseBodyEcIdAccountIdsFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetResponseBodyEcIdAccountIdsFilter) GoString() string {
	return s.String()
}

func (s *DescribeBudgetResponseBodyEcIdAccountIdsFilter) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *DescribeBudgetResponseBodyEcIdAccountIdsFilter) GetEcId() *string {
	return s.EcId
}

func (s *DescribeBudgetResponseBodyEcIdAccountIdsFilter) SetAccountIds(v []*int64) *DescribeBudgetResponseBodyEcIdAccountIdsFilter {
	s.AccountIds = v
	return s
}

func (s *DescribeBudgetResponseBodyEcIdAccountIdsFilter) SetEcId(v string) *DescribeBudgetResponseBodyEcIdAccountIdsFilter {
	s.EcId = &v
	return s
}

func (s *DescribeBudgetResponseBodyEcIdAccountIdsFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeBudgetResponseBodyQueryFilter struct {
	// Parameter code.
	//
	// example:
	//
	// RESOURCE_OWNER_ACCOUNT
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Selection mode.
	//
	// example:
	//
	// IN
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	// Filter value list.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeBudgetResponseBodyQueryFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetResponseBodyQueryFilter) GoString() string {
	return s.String()
}

func (s *DescribeBudgetResponseBodyQueryFilter) GetCode() *string {
	return s.Code
}

func (s *DescribeBudgetResponseBodyQueryFilter) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeBudgetResponseBodyQueryFilter) GetValues() []*string {
	return s.Values
}

func (s *DescribeBudgetResponseBodyQueryFilter) SetCode(v string) *DescribeBudgetResponseBodyQueryFilter {
	s.Code = &v
	return s
}

func (s *DescribeBudgetResponseBodyQueryFilter) SetSelectType(v string) *DescribeBudgetResponseBodyQueryFilter {
	s.SelectType = &v
	return s
}

func (s *DescribeBudgetResponseBodyQueryFilter) SetValues(v []*string) *DescribeBudgetResponseBodyQueryFilter {
	s.Values = v
	return s
}

func (s *DescribeBudgetResponseBodyQueryFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeBudgetResponseBodyWarnConfs struct {
	// Remarks.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Specifies whether to enable EventBridge.
	EventBridge *bool `json:"EventBridge,omitempty" xml:"EventBridge,omitempty"`
	// Message center notification channel list.
	MscChannels []*string `json:"MscChannels,omitempty" xml:"MscChannels,omitempty" type:"Repeated"`
	// Message center contact list.
	MscContacts []*string `json:"MscContacts,omitempty" xml:"MscContacts,omitempty" type:"Repeated"`
	// Alert name. User-defined and optional. If not specified, the backend automatically generates a name.
	//
	// example:
	//
	// Alter-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Sequence number. Only present in responses. Alerts are numbered in ascending order of alert ID, starting from 1.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Threshold type.
	//
	// example:
	//
	// FIXED
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// Threshold value.
	//
	// example:
	//
	// 1000
	ThresholdValue *string `json:"ThresholdValue,omitempty" xml:"ThresholdValue,omitempty"`
	// Alert target.
	//
	// example:
	//
	// ACTUAL
	WarnTarget *string `json:"WarnTarget,omitempty" xml:"WarnTarget,omitempty"`
}

func (s DescribeBudgetResponseBodyWarnConfs) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetResponseBodyWarnConfs) GoString() string {
	return s.String()
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetComment() *string {
	return s.Comment
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetEventBridge() *bool {
	return s.EventBridge
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetMscChannels() []*string {
	return s.MscChannels
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetMscContacts() []*string {
	return s.MscContacts
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetName() *string {
	return s.Name
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetSequence() *int32 {
	return s.Sequence
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetThresholdType() *string {
	return s.ThresholdType
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetThresholdValue() *string {
	return s.ThresholdValue
}

func (s *DescribeBudgetResponseBodyWarnConfs) GetWarnTarget() *string {
	return s.WarnTarget
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetComment(v string) *DescribeBudgetResponseBodyWarnConfs {
	s.Comment = &v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetEventBridge(v bool) *DescribeBudgetResponseBodyWarnConfs {
	s.EventBridge = &v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetMscChannels(v []*string) *DescribeBudgetResponseBodyWarnConfs {
	s.MscChannels = v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetMscContacts(v []*string) *DescribeBudgetResponseBodyWarnConfs {
	s.MscContacts = v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetName(v string) *DescribeBudgetResponseBodyWarnConfs {
	s.Name = &v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetSequence(v int32) *DescribeBudgetResponseBodyWarnConfs {
	s.Sequence = &v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetThresholdType(v string) *DescribeBudgetResponseBodyWarnConfs {
	s.ThresholdType = &v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetThresholdValue(v string) *DescribeBudgetResponseBodyWarnConfs {
	s.ThresholdValue = &v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) SetWarnTarget(v string) *DescribeBudgetResponseBodyWarnConfs {
	s.WarnTarget = &v
	return s
}

func (s *DescribeBudgetResponseBodyWarnConfs) Validate() error {
	return dara.Validate(s)
}
