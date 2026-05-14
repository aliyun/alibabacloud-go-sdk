// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBudgetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *UpdateBudgetRequest
	GetBudgetName() *string
	SetBudgetType(v string) *UpdateBudgetRequest
	GetBudgetType() *string
	SetComment(v string) *UpdateBudgetRequest
	GetComment() *string
	SetCycleEndPeriod(v string) *UpdateBudgetRequest
	GetCycleEndPeriod() *string
	SetCycleQuota(v []*UpdateBudgetRequestCycleQuota) *UpdateBudgetRequest
	GetCycleQuota() []*UpdateBudgetRequestCycleQuota
	SetCycleStartPeriod(v string) *UpdateBudgetRequest
	GetCycleStartPeriod() *string
	SetCycleType(v string) *UpdateBudgetRequest
	GetCycleType() *string
	SetEcIdAccountIds(v []*UpdateBudgetRequestEcIdAccountIds) *UpdateBudgetRequest
	GetEcIdAccountIds() []*UpdateBudgetRequestEcIdAccountIds
	SetMetric(v string) *UpdateBudgetRequest
	GetMetric() *string
	SetNbid(v string) *UpdateBudgetRequest
	GetNbid() *string
	SetOriginalBudgetName(v string) *UpdateBudgetRequest
	GetOriginalBudgetName() *string
	SetQueryFilter(v []*UpdateBudgetRequestQueryFilter) *UpdateBudgetRequest
	GetQueryFilter() []*UpdateBudgetRequestQueryFilter
	SetQuota(v string) *UpdateBudgetRequest
	GetQuota() *string
	SetQuotaType(v string) *UpdateBudgetRequest
	GetQuotaType() *string
	SetWarnConfs(v []*UpdateBudgetRequestWarnConfs) *UpdateBudgetRequest
	GetWarnConfs() []*UpdateBudgetRequestWarnConfs
}

type UpdateBudgetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// NewBudgetName
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CONSUME
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	Comment    *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-12
	CycleEndPeriod *string                          `json:"CycleEndPeriod,omitempty" xml:"CycleEndPeriod,omitempty"`
	CycleQuota     []*UpdateBudgetRequestCycleQuota `json:"CycleQuota,omitempty" xml:"CycleQuota,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-01
	CycleStartPeriod *string `json:"CycleStartPeriod,omitempty" xml:"CycleStartPeriod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MONTHLY
	CycleType      *string                              `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	EcIdAccountIds []*UpdateBudgetRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// REQUIRE_AMOUNT
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Nbid   *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OriginalBudgetName
	OriginalBudgetName *string                           `json:"OriginalBudgetName,omitempty" xml:"OriginalBudgetName,omitempty"`
	QueryFilter        []*UpdateBudgetRequestQueryFilter `json:"QueryFilter,omitempty" xml:"QueryFilter,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FIXED
	QuotaType *string                         `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	WarnConfs []*UpdateBudgetRequestWarnConfs `json:"WarnConfs,omitempty" xml:"WarnConfs,omitempty" type:"Repeated"`
}

func (s UpdateBudgetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBudgetRequest) GoString() string {
	return s.String()
}

func (s *UpdateBudgetRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *UpdateBudgetRequest) GetBudgetType() *string {
	return s.BudgetType
}

func (s *UpdateBudgetRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateBudgetRequest) GetCycleEndPeriod() *string {
	return s.CycleEndPeriod
}

func (s *UpdateBudgetRequest) GetCycleQuota() []*UpdateBudgetRequestCycleQuota {
	return s.CycleQuota
}

func (s *UpdateBudgetRequest) GetCycleStartPeriod() *string {
	return s.CycleStartPeriod
}

func (s *UpdateBudgetRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *UpdateBudgetRequest) GetEcIdAccountIds() []*UpdateBudgetRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *UpdateBudgetRequest) GetMetric() *string {
	return s.Metric
}

func (s *UpdateBudgetRequest) GetNbid() *string {
	return s.Nbid
}

func (s *UpdateBudgetRequest) GetOriginalBudgetName() *string {
	return s.OriginalBudgetName
}

func (s *UpdateBudgetRequest) GetQueryFilter() []*UpdateBudgetRequestQueryFilter {
	return s.QueryFilter
}

func (s *UpdateBudgetRequest) GetQuota() *string {
	return s.Quota
}

func (s *UpdateBudgetRequest) GetQuotaType() *string {
	return s.QuotaType
}

func (s *UpdateBudgetRequest) GetWarnConfs() []*UpdateBudgetRequestWarnConfs {
	return s.WarnConfs
}

func (s *UpdateBudgetRequest) SetBudgetName(v string) *UpdateBudgetRequest {
	s.BudgetName = &v
	return s
}

func (s *UpdateBudgetRequest) SetBudgetType(v string) *UpdateBudgetRequest {
	s.BudgetType = &v
	return s
}

func (s *UpdateBudgetRequest) SetComment(v string) *UpdateBudgetRequest {
	s.Comment = &v
	return s
}

func (s *UpdateBudgetRequest) SetCycleEndPeriod(v string) *UpdateBudgetRequest {
	s.CycleEndPeriod = &v
	return s
}

func (s *UpdateBudgetRequest) SetCycleQuota(v []*UpdateBudgetRequestCycleQuota) *UpdateBudgetRequest {
	s.CycleQuota = v
	return s
}

func (s *UpdateBudgetRequest) SetCycleStartPeriod(v string) *UpdateBudgetRequest {
	s.CycleStartPeriod = &v
	return s
}

func (s *UpdateBudgetRequest) SetCycleType(v string) *UpdateBudgetRequest {
	s.CycleType = &v
	return s
}

func (s *UpdateBudgetRequest) SetEcIdAccountIds(v []*UpdateBudgetRequestEcIdAccountIds) *UpdateBudgetRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *UpdateBudgetRequest) SetMetric(v string) *UpdateBudgetRequest {
	s.Metric = &v
	return s
}

func (s *UpdateBudgetRequest) SetNbid(v string) *UpdateBudgetRequest {
	s.Nbid = &v
	return s
}

func (s *UpdateBudgetRequest) SetOriginalBudgetName(v string) *UpdateBudgetRequest {
	s.OriginalBudgetName = &v
	return s
}

func (s *UpdateBudgetRequest) SetQueryFilter(v []*UpdateBudgetRequestQueryFilter) *UpdateBudgetRequest {
	s.QueryFilter = v
	return s
}

func (s *UpdateBudgetRequest) SetQuota(v string) *UpdateBudgetRequest {
	s.Quota = &v
	return s
}

func (s *UpdateBudgetRequest) SetQuotaType(v string) *UpdateBudgetRequest {
	s.QuotaType = &v
	return s
}

func (s *UpdateBudgetRequest) SetWarnConfs(v []*UpdateBudgetRequestWarnConfs) *UpdateBudgetRequest {
	s.WarnConfs = v
	return s
}

func (s *UpdateBudgetRequest) Validate() error {
	if s.CycleQuota != nil {
		for _, item := range s.CycleQuota {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type UpdateBudgetRequestCycleQuota struct {
	CyclePeriod *string `json:"CyclePeriod,omitempty" xml:"CyclePeriod,omitempty"`
	Quota       *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s UpdateBudgetRequestCycleQuota) String() string {
	return dara.Prettify(s)
}

func (s UpdateBudgetRequestCycleQuota) GoString() string {
	return s.String()
}

func (s *UpdateBudgetRequestCycleQuota) GetCyclePeriod() *string {
	return s.CyclePeriod
}

func (s *UpdateBudgetRequestCycleQuota) GetQuota() *string {
	return s.Quota
}

func (s *UpdateBudgetRequestCycleQuota) SetCyclePeriod(v string) *UpdateBudgetRequestCycleQuota {
	s.CyclePeriod = &v
	return s
}

func (s *UpdateBudgetRequestCycleQuota) SetQuota(v string) *UpdateBudgetRequestCycleQuota {
	s.Quota = &v
	return s
}

func (s *UpdateBudgetRequestCycleQuota) Validate() error {
	return dara.Validate(s)
}

type UpdateBudgetRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	EcId       *string  `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s UpdateBudgetRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateBudgetRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *UpdateBudgetRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *UpdateBudgetRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *UpdateBudgetRequestEcIdAccountIds) SetAccountIds(v []*int64) *UpdateBudgetRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *UpdateBudgetRequestEcIdAccountIds) SetEcId(v string) *UpdateBudgetRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *UpdateBudgetRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}

type UpdateBudgetRequestQueryFilter struct {
	Code       *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateBudgetRequestQueryFilter) String() string {
	return dara.Prettify(s)
}

func (s UpdateBudgetRequestQueryFilter) GoString() string {
	return s.String()
}

func (s *UpdateBudgetRequestQueryFilter) GetCode() *string {
	return s.Code
}

func (s *UpdateBudgetRequestQueryFilter) GetSelectType() *string {
	return s.SelectType
}

func (s *UpdateBudgetRequestQueryFilter) GetValues() []*string {
	return s.Values
}

func (s *UpdateBudgetRequestQueryFilter) SetCode(v string) *UpdateBudgetRequestQueryFilter {
	s.Code = &v
	return s
}

func (s *UpdateBudgetRequestQueryFilter) SetSelectType(v string) *UpdateBudgetRequestQueryFilter {
	s.SelectType = &v
	return s
}

func (s *UpdateBudgetRequestQueryFilter) SetValues(v []*string) *UpdateBudgetRequestQueryFilter {
	s.Values = v
	return s
}

func (s *UpdateBudgetRequestQueryFilter) Validate() error {
	return dara.Validate(s)
}

type UpdateBudgetRequestWarnConfs struct {
	Comment        *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	EventBridge    *bool     `json:"EventBridge,omitempty" xml:"EventBridge,omitempty"`
	MscChannels    []*string `json:"MscChannels,omitempty" xml:"MscChannels,omitempty" type:"Repeated"`
	MscContacts    []*string `json:"MscContacts,omitempty" xml:"MscContacts,omitempty" type:"Repeated"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	ThresholdType  *string   `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	ThresholdValue *string   `json:"ThresholdValue,omitempty" xml:"ThresholdValue,omitempty"`
	WarnTarget     *string   `json:"WarnTarget,omitempty" xml:"WarnTarget,omitempty"`
}

func (s UpdateBudgetRequestWarnConfs) String() string {
	return dara.Prettify(s)
}

func (s UpdateBudgetRequestWarnConfs) GoString() string {
	return s.String()
}

func (s *UpdateBudgetRequestWarnConfs) GetComment() *string {
	return s.Comment
}

func (s *UpdateBudgetRequestWarnConfs) GetEventBridge() *bool {
	return s.EventBridge
}

func (s *UpdateBudgetRequestWarnConfs) GetMscChannels() []*string {
	return s.MscChannels
}

func (s *UpdateBudgetRequestWarnConfs) GetMscContacts() []*string {
	return s.MscContacts
}

func (s *UpdateBudgetRequestWarnConfs) GetName() *string {
	return s.Name
}

func (s *UpdateBudgetRequestWarnConfs) GetThresholdType() *string {
	return s.ThresholdType
}

func (s *UpdateBudgetRequestWarnConfs) GetThresholdValue() *string {
	return s.ThresholdValue
}

func (s *UpdateBudgetRequestWarnConfs) GetWarnTarget() *string {
	return s.WarnTarget
}

func (s *UpdateBudgetRequestWarnConfs) SetComment(v string) *UpdateBudgetRequestWarnConfs {
	s.Comment = &v
	return s
}

func (s *UpdateBudgetRequestWarnConfs) SetEventBridge(v bool) *UpdateBudgetRequestWarnConfs {
	s.EventBridge = &v
	return s
}

func (s *UpdateBudgetRequestWarnConfs) SetMscChannels(v []*string) *UpdateBudgetRequestWarnConfs {
	s.MscChannels = v
	return s
}

func (s *UpdateBudgetRequestWarnConfs) SetMscContacts(v []*string) *UpdateBudgetRequestWarnConfs {
	s.MscContacts = v
	return s
}

func (s *UpdateBudgetRequestWarnConfs) SetName(v string) *UpdateBudgetRequestWarnConfs {
	s.Name = &v
	return s
}

func (s *UpdateBudgetRequestWarnConfs) SetThresholdType(v string) *UpdateBudgetRequestWarnConfs {
	s.ThresholdType = &v
	return s
}

func (s *UpdateBudgetRequestWarnConfs) SetThresholdValue(v string) *UpdateBudgetRequestWarnConfs {
	s.ThresholdValue = &v
	return s
}

func (s *UpdateBudgetRequestWarnConfs) SetWarnTarget(v string) *UpdateBudgetRequestWarnConfs {
	s.WarnTarget = &v
	return s
}

func (s *UpdateBudgetRequestWarnConfs) Validate() error {
	return dara.Validate(s)
}
