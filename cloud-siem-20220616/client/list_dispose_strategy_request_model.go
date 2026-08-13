// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisposeStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertUuid(v string) *ListDisposeStrategyRequest
	GetAlertUuid() *string
	SetCurrentPage(v int32) *ListDisposeStrategyRequest
	GetCurrentPage() *int32
	SetEffectiveStatus(v int32) *ListDisposeStrategyRequest
	GetEffectiveStatus() *int32
	SetEndTime(v int64) *ListDisposeStrategyRequest
	GetEndTime() *int64
	SetEntityIdentity(v string) *ListDisposeStrategyRequest
	GetEntityIdentity() *string
	SetEntityType(v string) *ListDisposeStrategyRequest
	GetEntityType() *string
	SetEntityUuidList(v []*string) *ListDisposeStrategyRequest
	GetEntityUuidList() []*string
	SetGroupBy(v string) *ListDisposeStrategyRequest
	GetGroupBy() *string
	SetGroupKey(v string) *ListDisposeStrategyRequest
	GetGroupKey() *string
	SetIncidentUuid(v string) *ListDisposeStrategyRequest
	GetIncidentUuid() *string
	SetMaxResults(v int32) *ListDisposeStrategyRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDisposeStrategyRequest
	GetNextToken() *string
	SetOrder(v string) *ListDisposeStrategyRequest
	GetOrder() *string
	SetOrderField(v string) *ListDisposeStrategyRequest
	GetOrderField() *string
	SetPageSize(v int32) *ListDisposeStrategyRequest
	GetPageSize() *int32
	SetPlaybookName(v string) *ListDisposeStrategyRequest
	GetPlaybookName() *string
	SetPlaybookTypes(v string) *ListDisposeStrategyRequest
	GetPlaybookTypes() *string
	SetPlaybookUuid(v string) *ListDisposeStrategyRequest
	GetPlaybookUuid() *string
	SetQueryMode(v string) *ListDisposeStrategyRequest
	GetQueryMode() *string
	SetRegionId(v string) *ListDisposeStrategyRequest
	GetRegionId() *string
	SetResponseRuleId(v string) *ListDisposeStrategyRequest
	GetResponseRuleId() *string
	SetRoleFor(v int64) *ListDisposeStrategyRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListDisposeStrategyRequest
	GetRoleType() *int32
	SetSophonTaskId(v string) *ListDisposeStrategyRequest
	GetSophonTaskId() *string
	SetStartTime(v int64) *ListDisposeStrategyRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListDisposeStrategyRequest
	GetStatus() *int32
	SetStrategyId(v string) *ListDisposeStrategyRequest
	GetStrategyId() *string
}

type ListDisposeStrategyRequest struct {
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The current page number, which must be greater than or equal to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The policy status. Valid values:
	//
	// example:
	//
	// 0
	EffectiveStatus *int32 `json:"EffectiveStatus,omitempty" xml:"EffectiveStatus,omitempty"`
	// The query end time, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The entity characteristic value, which can be used for fuzzy match on response entities.
	//
	// example:
	//
	// test22.php
	EntityIdentity *string `json:"EntityIdentity,omitempty" xml:"EntityIdentity,omitempty"`
	// The entity type. Valid values:
	//
	// example:
	//
	// ip
	EntityType     *string   `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityUuidList []*string `json:"EntityUuidList,omitempty" xml:"EntityUuidList,omitempty" type:"Repeated"`
	GroupBy        *string   `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	GroupKey       *string   `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 49670d3bbf7aa9556a2fff3dbaa9****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort direction. Valid values:
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field used to sort results. Valid values:
	//
	// - GmtModified: sorts results by update time.
	//
	// - GmtCreate: sorts results by creation time.
	//
	// - FinishTime: sorts results by policy end time.
	//
	// example:
	//
	// GmtModified
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page, with a maximum of 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier name of the playbook.
	//
	// example:
	//
	// WafBlockIP
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The playbook type. Valid values:
	//
	// - system: manual handling
	//
	// - custom: event-triggered playbook
	//
	// - custom_alert: alert-triggered playbook
	//
	// - soar-manual: manually run playbook
	//
	// - soar-mdr: MDR-run playbook
	//
	// example:
	//
	// system
	PlaybookTypes *string `json:"PlaybookTypes,omitempty" xml:"PlaybookTypes,omitempty"`
	// The playbook UUID.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	QueryMode    *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// The region where the data management center of Cloud Threat Detection and Response (CTDR) is located. Specify the management center based on the region of your assets. Valid values:
	//
	// example:
	//
	// cn-hangzhou
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResponseRuleId *string `json:"ResponseRuleId,omitempty" xml:"ResponseRuleId,omitempty"`
	// The Alibaba Cloud account ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The SOAR response policy ID.
	//
	// example:
	//
	// a50a49b7-6044-4593-ab15-2b46567c****
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The query start time, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The response policy status.
	//
	// example:
	//
	// 200
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s ListDisposeStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyRequest) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *ListDisposeStrategyRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDisposeStrategyRequest) GetEffectiveStatus() *int32 {
	return s.EffectiveStatus
}

func (s *ListDisposeStrategyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDisposeStrategyRequest) GetEntityIdentity() *string {
	return s.EntityIdentity
}

func (s *ListDisposeStrategyRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListDisposeStrategyRequest) GetEntityUuidList() []*string {
	return s.EntityUuidList
}

func (s *ListDisposeStrategyRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ListDisposeStrategyRequest) GetGroupKey() *string {
	return s.GroupKey
}

func (s *ListDisposeStrategyRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *ListDisposeStrategyRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDisposeStrategyRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDisposeStrategyRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDisposeStrategyRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListDisposeStrategyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDisposeStrategyRequest) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *ListDisposeStrategyRequest) GetPlaybookTypes() *string {
	return s.PlaybookTypes
}

func (s *ListDisposeStrategyRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ListDisposeStrategyRequest) GetQueryMode() *string {
	return s.QueryMode
}

func (s *ListDisposeStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDisposeStrategyRequest) GetResponseRuleId() *string {
	return s.ResponseRuleId
}

func (s *ListDisposeStrategyRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDisposeStrategyRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListDisposeStrategyRequest) GetSophonTaskId() *string {
	return s.SophonTaskId
}

func (s *ListDisposeStrategyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDisposeStrategyRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListDisposeStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListDisposeStrategyRequest) SetAlertUuid(v string) *ListDisposeStrategyRequest {
	s.AlertUuid = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetCurrentPage(v int32) *ListDisposeStrategyRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEffectiveStatus(v int32) *ListDisposeStrategyRequest {
	s.EffectiveStatus = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEndTime(v int64) *ListDisposeStrategyRequest {
	s.EndTime = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEntityIdentity(v string) *ListDisposeStrategyRequest {
	s.EntityIdentity = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEntityType(v string) *ListDisposeStrategyRequest {
	s.EntityType = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEntityUuidList(v []*string) *ListDisposeStrategyRequest {
	s.EntityUuidList = v
	return s
}

func (s *ListDisposeStrategyRequest) SetGroupBy(v string) *ListDisposeStrategyRequest {
	s.GroupBy = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetGroupKey(v string) *ListDisposeStrategyRequest {
	s.GroupKey = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetIncidentUuid(v string) *ListDisposeStrategyRequest {
	s.IncidentUuid = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetMaxResults(v int32) *ListDisposeStrategyRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetNextToken(v string) *ListDisposeStrategyRequest {
	s.NextToken = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetOrder(v string) *ListDisposeStrategyRequest {
	s.Order = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetOrderField(v string) *ListDisposeStrategyRequest {
	s.OrderField = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPageSize(v int32) *ListDisposeStrategyRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookName(v string) *ListDisposeStrategyRequest {
	s.PlaybookName = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookTypes(v string) *ListDisposeStrategyRequest {
	s.PlaybookTypes = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookUuid(v string) *ListDisposeStrategyRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetQueryMode(v string) *ListDisposeStrategyRequest {
	s.QueryMode = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetRegionId(v string) *ListDisposeStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetResponseRuleId(v string) *ListDisposeStrategyRequest {
	s.ResponseRuleId = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetRoleFor(v int64) *ListDisposeStrategyRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetRoleType(v int32) *ListDisposeStrategyRequest {
	s.RoleType = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetSophonTaskId(v string) *ListDisposeStrategyRequest {
	s.SophonTaskId = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetStartTime(v int64) *ListDisposeStrategyRequest {
	s.StartTime = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetStatus(v int32) *ListDisposeStrategyRequest {
	s.Status = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetStrategyId(v string) *ListDisposeStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *ListDisposeStrategyRequest) Validate() error {
	return dara.Validate(s)
}
