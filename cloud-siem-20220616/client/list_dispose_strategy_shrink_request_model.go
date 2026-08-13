// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisposeStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertUuid(v string) *ListDisposeStrategyShrinkRequest
	GetAlertUuid() *string
	SetCurrentPage(v int32) *ListDisposeStrategyShrinkRequest
	GetCurrentPage() *int32
	SetEffectiveStatus(v int32) *ListDisposeStrategyShrinkRequest
	GetEffectiveStatus() *int32
	SetEndTime(v int64) *ListDisposeStrategyShrinkRequest
	GetEndTime() *int64
	SetEntityIdentity(v string) *ListDisposeStrategyShrinkRequest
	GetEntityIdentity() *string
	SetEntityType(v string) *ListDisposeStrategyShrinkRequest
	GetEntityType() *string
	SetEntityUuidListShrink(v string) *ListDisposeStrategyShrinkRequest
	GetEntityUuidListShrink() *string
	SetGroupBy(v string) *ListDisposeStrategyShrinkRequest
	GetGroupBy() *string
	SetGroupKey(v string) *ListDisposeStrategyShrinkRequest
	GetGroupKey() *string
	SetIncidentUuid(v string) *ListDisposeStrategyShrinkRequest
	GetIncidentUuid() *string
	SetMaxResults(v int32) *ListDisposeStrategyShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDisposeStrategyShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *ListDisposeStrategyShrinkRequest
	GetOrder() *string
	SetOrderField(v string) *ListDisposeStrategyShrinkRequest
	GetOrderField() *string
	SetPageSize(v int32) *ListDisposeStrategyShrinkRequest
	GetPageSize() *int32
	SetPlaybookName(v string) *ListDisposeStrategyShrinkRequest
	GetPlaybookName() *string
	SetPlaybookTypes(v string) *ListDisposeStrategyShrinkRequest
	GetPlaybookTypes() *string
	SetPlaybookUuid(v string) *ListDisposeStrategyShrinkRequest
	GetPlaybookUuid() *string
	SetQueryMode(v string) *ListDisposeStrategyShrinkRequest
	GetQueryMode() *string
	SetRegionId(v string) *ListDisposeStrategyShrinkRequest
	GetRegionId() *string
	SetResponseRuleId(v string) *ListDisposeStrategyShrinkRequest
	GetResponseRuleId() *string
	SetRoleFor(v int64) *ListDisposeStrategyShrinkRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListDisposeStrategyShrinkRequest
	GetRoleType() *int32
	SetSophonTaskId(v string) *ListDisposeStrategyShrinkRequest
	GetSophonTaskId() *string
	SetStartTime(v int64) *ListDisposeStrategyShrinkRequest
	GetStartTime() *int64
	SetStatus(v int32) *ListDisposeStrategyShrinkRequest
	GetStatus() *int32
	SetStrategyId(v string) *ListDisposeStrategyShrinkRequest
	GetStrategyId() *string
}

type ListDisposeStrategyShrinkRequest struct {
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
	EntityType           *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityUuidListShrink *string `json:"EntityUuidList,omitempty" xml:"EntityUuidList,omitempty"`
	GroupBy              *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	GroupKey             *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
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

func (s ListDisposeStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyShrinkRequest) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *ListDisposeStrategyShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDisposeStrategyShrinkRequest) GetEffectiveStatus() *int32 {
	return s.EffectiveStatus
}

func (s *ListDisposeStrategyShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDisposeStrategyShrinkRequest) GetEntityIdentity() *string {
	return s.EntityIdentity
}

func (s *ListDisposeStrategyShrinkRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListDisposeStrategyShrinkRequest) GetEntityUuidListShrink() *string {
	return s.EntityUuidListShrink
}

func (s *ListDisposeStrategyShrinkRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ListDisposeStrategyShrinkRequest) GetGroupKey() *string {
	return s.GroupKey
}

func (s *ListDisposeStrategyShrinkRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *ListDisposeStrategyShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDisposeStrategyShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDisposeStrategyShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDisposeStrategyShrinkRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListDisposeStrategyShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDisposeStrategyShrinkRequest) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *ListDisposeStrategyShrinkRequest) GetPlaybookTypes() *string {
	return s.PlaybookTypes
}

func (s *ListDisposeStrategyShrinkRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ListDisposeStrategyShrinkRequest) GetQueryMode() *string {
	return s.QueryMode
}

func (s *ListDisposeStrategyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDisposeStrategyShrinkRequest) GetResponseRuleId() *string {
	return s.ResponseRuleId
}

func (s *ListDisposeStrategyShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDisposeStrategyShrinkRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListDisposeStrategyShrinkRequest) GetSophonTaskId() *string {
	return s.SophonTaskId
}

func (s *ListDisposeStrategyShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDisposeStrategyShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListDisposeStrategyShrinkRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ListDisposeStrategyShrinkRequest) SetAlertUuid(v string) *ListDisposeStrategyShrinkRequest {
	s.AlertUuid = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetCurrentPage(v int32) *ListDisposeStrategyShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetEffectiveStatus(v int32) *ListDisposeStrategyShrinkRequest {
	s.EffectiveStatus = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetEndTime(v int64) *ListDisposeStrategyShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetEntityIdentity(v string) *ListDisposeStrategyShrinkRequest {
	s.EntityIdentity = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetEntityType(v string) *ListDisposeStrategyShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetEntityUuidListShrink(v string) *ListDisposeStrategyShrinkRequest {
	s.EntityUuidListShrink = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetGroupBy(v string) *ListDisposeStrategyShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetGroupKey(v string) *ListDisposeStrategyShrinkRequest {
	s.GroupKey = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetIncidentUuid(v string) *ListDisposeStrategyShrinkRequest {
	s.IncidentUuid = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetMaxResults(v int32) *ListDisposeStrategyShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetNextToken(v string) *ListDisposeStrategyShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetOrder(v string) *ListDisposeStrategyShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetOrderField(v string) *ListDisposeStrategyShrinkRequest {
	s.OrderField = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetPageSize(v int32) *ListDisposeStrategyShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetPlaybookName(v string) *ListDisposeStrategyShrinkRequest {
	s.PlaybookName = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetPlaybookTypes(v string) *ListDisposeStrategyShrinkRequest {
	s.PlaybookTypes = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetPlaybookUuid(v string) *ListDisposeStrategyShrinkRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetQueryMode(v string) *ListDisposeStrategyShrinkRequest {
	s.QueryMode = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetRegionId(v string) *ListDisposeStrategyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetResponseRuleId(v string) *ListDisposeStrategyShrinkRequest {
	s.ResponseRuleId = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetRoleFor(v int64) *ListDisposeStrategyShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetRoleType(v int32) *ListDisposeStrategyShrinkRequest {
	s.RoleType = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetSophonTaskId(v string) *ListDisposeStrategyShrinkRequest {
	s.SophonTaskId = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetStartTime(v int64) *ListDisposeStrategyShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetStatus(v int32) *ListDisposeStrategyShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) SetStrategyId(v string) *ListDisposeStrategyShrinkRequest {
	s.StrategyId = &v
	return s
}

func (s *ListDisposeStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
