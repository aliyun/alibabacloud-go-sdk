// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisposeStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetIncidentUuid(v string) *ListDisposeStrategyRequest
	GetIncidentUuid() *string
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
	SetRegionId(v string) *ListDisposeStrategyRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDisposeStrategyRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListDisposeStrategyRequest
	GetRoleType() *int32
	SetSophonTaskId(v string) *ListDisposeStrategyRequest
	GetSophonTaskId() *string
	SetStartTime(v int64) *ListDisposeStrategyRequest
	GetStartTime() *int64
}

type ListDisposeStrategyRequest struct {
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- 0: invalid
	//
	// 	- 1: valid
	//
	// example:
	//
	// 0
	EffectiveStatus *int32 `json:"EffectiveStatus,omitempty" xml:"EffectiveStatus,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The feature value of the entity. Fuzzy match is supported.
	//
	// example:
	//
	// test22.php
	EntityIdentity *string `json:"EntityIdentity,omitempty" xml:"EntityIdentity,omitempty"`
	// The entity type of the playbook. Valid values:
	//
	// 	- ip
	//
	// 	- process
	//
	// 	- file
	//
	// example:
	//
	// ip
	EntityType   *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The sort order. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sort field. Valid values:
	//
	// 	- GmtModified: sorts the policies by update time.
	//
	// 	- GmtCreate: sorts the policies by creation time.
	//
	// 	- FinishTime: sorts the policies by end time.
	//
	// example:
	//
	// GmtModified
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the playbook, which is the unique identifier of the playbook.
	//
	// example:
	//
	// WafBlockIP
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- system: user-triggered playbook
	//
	// 	- custom: event-triggered playbook
	//
	// 	- custom_alert: alert-triggered playbook
	//
	// 	- soar-manual: user-run playbook
	//
	// 	- soar-mdr: MDR-run playbook
	//
	// example:
	//
	// system
	PlaybookTypes *string `json:"PlaybookTypes,omitempty" xml:"PlaybookTypes,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the SOAR handling policy.
	//
	// example:
	//
	// a50a49b7-6044-4593-ab15-2b46567caadd
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDisposeStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyRequest) GoString() string {
	return s.String()
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

func (s *ListDisposeStrategyRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
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

func (s *ListDisposeStrategyRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *ListDisposeStrategyRequest) SetIncidentUuid(v string) *ListDisposeStrategyRequest {
	s.IncidentUuid = &v
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

func (s *ListDisposeStrategyRequest) SetRegionId(v string) *ListDisposeStrategyRequest {
	s.RegionId = &v
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

func (s *ListDisposeStrategyRequest) Validate() error {
	return dara.Validate(s)
}
