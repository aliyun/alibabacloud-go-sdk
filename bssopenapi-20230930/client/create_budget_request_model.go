// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBudgetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *CreateBudgetRequest
	GetBudgetName() *string
	SetBudgetType(v string) *CreateBudgetRequest
	GetBudgetType() *string
	SetComment(v string) *CreateBudgetRequest
	GetComment() *string
	SetCycleEndPeriod(v string) *CreateBudgetRequest
	GetCycleEndPeriod() *string
	SetCycleQuota(v []*CreateBudgetRequestCycleQuota) *CreateBudgetRequest
	GetCycleQuota() []*CreateBudgetRequestCycleQuota
	SetCycleStartPeriod(v string) *CreateBudgetRequest
	GetCycleStartPeriod() *string
	SetCycleType(v string) *CreateBudgetRequest
	GetCycleType() *string
	SetEcIdAccountIds(v []*CreateBudgetRequestEcIdAccountIds) *CreateBudgetRequest
	GetEcIdAccountIds() []*CreateBudgetRequestEcIdAccountIds
	SetMetric(v string) *CreateBudgetRequest
	GetMetric() *string
	SetNbid(v string) *CreateBudgetRequest
	GetNbid() *string
	SetQueryFilter(v []*CreateBudgetRequestQueryFilter) *CreateBudgetRequest
	GetQueryFilter() []*CreateBudgetRequestQueryFilter
	SetQuota(v string) *CreateBudgetRequest
	GetQuota() *string
	SetQuotaType(v string) *CreateBudgetRequest
	GetQuotaType() *string
	SetWarnConfs(v []*CreateBudgetRequestWarnConfs) *CreateBudgetRequest
	GetWarnConfs() []*CreateBudgetRequestWarnConfs
}

type CreateBudgetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Department_dev_budget
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
	CycleQuota     []*CreateBudgetRequestCycleQuota `json:"CycleQuota,omitempty" xml:"CycleQuota,omitempty" type:"Repeated"`
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
	EcIdAccountIds []*CreateBudgetRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// REQUIRE_AMOUNT
	Metric      *string                           `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Nbid        *string                           `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	QueryFilter []*CreateBudgetRequestQueryFilter `json:"QueryFilter,omitempty" xml:"QueryFilter,omitempty" type:"Repeated"`
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
	WarnConfs []*CreateBudgetRequestWarnConfs `json:"WarnConfs,omitempty" xml:"WarnConfs,omitempty" type:"Repeated"`
}

func (s CreateBudgetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetRequest) GoString() string {
	return s.String()
}

func (s *CreateBudgetRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *CreateBudgetRequest) GetBudgetType() *string {
	return s.BudgetType
}

func (s *CreateBudgetRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateBudgetRequest) GetCycleEndPeriod() *string {
	return s.CycleEndPeriod
}

func (s *CreateBudgetRequest) GetCycleQuota() []*CreateBudgetRequestCycleQuota {
	return s.CycleQuota
}

func (s *CreateBudgetRequest) GetCycleStartPeriod() *string {
	return s.CycleStartPeriod
}

func (s *CreateBudgetRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *CreateBudgetRequest) GetEcIdAccountIds() []*CreateBudgetRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *CreateBudgetRequest) GetMetric() *string {
	return s.Metric
}

func (s *CreateBudgetRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateBudgetRequest) GetQueryFilter() []*CreateBudgetRequestQueryFilter {
	return s.QueryFilter
}

func (s *CreateBudgetRequest) GetQuota() *string {
	return s.Quota
}

func (s *CreateBudgetRequest) GetQuotaType() *string {
	return s.QuotaType
}

func (s *CreateBudgetRequest) GetWarnConfs() []*CreateBudgetRequestWarnConfs {
	return s.WarnConfs
}

func (s *CreateBudgetRequest) SetBudgetName(v string) *CreateBudgetRequest {
	s.BudgetName = &v
	return s
}

func (s *CreateBudgetRequest) SetBudgetType(v string) *CreateBudgetRequest {
	s.BudgetType = &v
	return s
}

func (s *CreateBudgetRequest) SetComment(v string) *CreateBudgetRequest {
	s.Comment = &v
	return s
}

func (s *CreateBudgetRequest) SetCycleEndPeriod(v string) *CreateBudgetRequest {
	s.CycleEndPeriod = &v
	return s
}

func (s *CreateBudgetRequest) SetCycleQuota(v []*CreateBudgetRequestCycleQuota) *CreateBudgetRequest {
	s.CycleQuota = v
	return s
}

func (s *CreateBudgetRequest) SetCycleStartPeriod(v string) *CreateBudgetRequest {
	s.CycleStartPeriod = &v
	return s
}

func (s *CreateBudgetRequest) SetCycleType(v string) *CreateBudgetRequest {
	s.CycleType = &v
	return s
}

func (s *CreateBudgetRequest) SetEcIdAccountIds(v []*CreateBudgetRequestEcIdAccountIds) *CreateBudgetRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *CreateBudgetRequest) SetMetric(v string) *CreateBudgetRequest {
	s.Metric = &v
	return s
}

func (s *CreateBudgetRequest) SetNbid(v string) *CreateBudgetRequest {
	s.Nbid = &v
	return s
}

func (s *CreateBudgetRequest) SetQueryFilter(v []*CreateBudgetRequestQueryFilter) *CreateBudgetRequest {
	s.QueryFilter = v
	return s
}

func (s *CreateBudgetRequest) SetQuota(v string) *CreateBudgetRequest {
	s.Quota = &v
	return s
}

func (s *CreateBudgetRequest) SetQuotaType(v string) *CreateBudgetRequest {
	s.QuotaType = &v
	return s
}

func (s *CreateBudgetRequest) SetWarnConfs(v []*CreateBudgetRequestWarnConfs) *CreateBudgetRequest {
	s.WarnConfs = v
	return s
}

func (s *CreateBudgetRequest) Validate() error {
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

type CreateBudgetRequestCycleQuota struct {
	CyclePeriod *string `json:"CyclePeriod,omitempty" xml:"CyclePeriod,omitempty"`
	Quota       *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s CreateBudgetRequestCycleQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetRequestCycleQuota) GoString() string {
	return s.String()
}

func (s *CreateBudgetRequestCycleQuota) GetCyclePeriod() *string {
	return s.CyclePeriod
}

func (s *CreateBudgetRequestCycleQuota) GetQuota() *string {
	return s.Quota
}

func (s *CreateBudgetRequestCycleQuota) SetCyclePeriod(v string) *CreateBudgetRequestCycleQuota {
	s.CyclePeriod = &v
	return s
}

func (s *CreateBudgetRequestCycleQuota) SetQuota(v string) *CreateBudgetRequestCycleQuota {
	s.Quota = &v
	return s
}

func (s *CreateBudgetRequestCycleQuota) Validate() error {
	return dara.Validate(s)
}

type CreateBudgetRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	EcId       *string  `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s CreateBudgetRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *CreateBudgetRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *CreateBudgetRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *CreateBudgetRequestEcIdAccountIds) SetAccountIds(v []*int64) *CreateBudgetRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *CreateBudgetRequestEcIdAccountIds) SetEcId(v string) *CreateBudgetRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *CreateBudgetRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}

type CreateBudgetRequestQueryFilter struct {
	// example:
	//
	// RESOURCE_OWNER_ACCOUNT
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// IN
	SelectType *string   `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	Values     []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateBudgetRequestQueryFilter) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetRequestQueryFilter) GoString() string {
	return s.String()
}

func (s *CreateBudgetRequestQueryFilter) GetCode() *string {
	return s.Code
}

func (s *CreateBudgetRequestQueryFilter) GetSelectType() *string {
	return s.SelectType
}

func (s *CreateBudgetRequestQueryFilter) GetValues() []*string {
	return s.Values
}

func (s *CreateBudgetRequestQueryFilter) SetCode(v string) *CreateBudgetRequestQueryFilter {
	s.Code = &v
	return s
}

func (s *CreateBudgetRequestQueryFilter) SetSelectType(v string) *CreateBudgetRequestQueryFilter {
	s.SelectType = &v
	return s
}

func (s *CreateBudgetRequestQueryFilter) SetValues(v []*string) *CreateBudgetRequestQueryFilter {
	s.Values = v
	return s
}

func (s *CreateBudgetRequestQueryFilter) Validate() error {
	return dara.Validate(s)
}

type CreateBudgetRequestWarnConfs struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// true
	EventBridge *bool     `json:"EventBridge,omitempty" xml:"EventBridge,omitempty"`
	MscChannels []*string `json:"MscChannels,omitempty" xml:"MscChannels,omitempty" type:"Repeated"`
	MscContacts []*string `json:"MscContacts,omitempty" xml:"MscContacts,omitempty" type:"Repeated"`
	// example:
	//
	// Alter-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FIXED
	ThresholdType *string `json:"ThresholdType,omitempty" xml:"ThresholdType,omitempty"`
	// example:
	//
	// 2000
	ThresholdValue *string `json:"ThresholdValue,omitempty" xml:"ThresholdValue,omitempty"`
	// example:
	//
	// FORECAST
	WarnTarget *string `json:"WarnTarget,omitempty" xml:"WarnTarget,omitempty"`
}

func (s CreateBudgetRequestWarnConfs) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetRequestWarnConfs) GoString() string {
	return s.String()
}

func (s *CreateBudgetRequestWarnConfs) GetComment() *string {
	return s.Comment
}

func (s *CreateBudgetRequestWarnConfs) GetEventBridge() *bool {
	return s.EventBridge
}

func (s *CreateBudgetRequestWarnConfs) GetMscChannels() []*string {
	return s.MscChannels
}

func (s *CreateBudgetRequestWarnConfs) GetMscContacts() []*string {
	return s.MscContacts
}

func (s *CreateBudgetRequestWarnConfs) GetName() *string {
	return s.Name
}

func (s *CreateBudgetRequestWarnConfs) GetThresholdType() *string {
	return s.ThresholdType
}

func (s *CreateBudgetRequestWarnConfs) GetThresholdValue() *string {
	return s.ThresholdValue
}

func (s *CreateBudgetRequestWarnConfs) GetWarnTarget() *string {
	return s.WarnTarget
}

func (s *CreateBudgetRequestWarnConfs) SetComment(v string) *CreateBudgetRequestWarnConfs {
	s.Comment = &v
	return s
}

func (s *CreateBudgetRequestWarnConfs) SetEventBridge(v bool) *CreateBudgetRequestWarnConfs {
	s.EventBridge = &v
	return s
}

func (s *CreateBudgetRequestWarnConfs) SetMscChannels(v []*string) *CreateBudgetRequestWarnConfs {
	s.MscChannels = v
	return s
}

func (s *CreateBudgetRequestWarnConfs) SetMscContacts(v []*string) *CreateBudgetRequestWarnConfs {
	s.MscContacts = v
	return s
}

func (s *CreateBudgetRequestWarnConfs) SetName(v string) *CreateBudgetRequestWarnConfs {
	s.Name = &v
	return s
}

func (s *CreateBudgetRequestWarnConfs) SetThresholdType(v string) *CreateBudgetRequestWarnConfs {
	s.ThresholdType = &v
	return s
}

func (s *CreateBudgetRequestWarnConfs) SetThresholdValue(v string) *CreateBudgetRequestWarnConfs {
	s.ThresholdValue = &v
	return s
}

func (s *CreateBudgetRequestWarnConfs) SetWarnTarget(v string) *CreateBudgetRequestWarnConfs {
	s.WarnTarget = &v
	return s
}

func (s *CreateBudgetRequestWarnConfs) Validate() error {
	return dara.Validate(s)
}
