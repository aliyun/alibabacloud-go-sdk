// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsWithEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeAlertsWithEntityRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeAlertsWithEntityRequest
	GetEndTime() *int64
	SetEntityId(v int64) *DescribeAlertsWithEntityRequest
	GetEntityId() *int64
	SetEntityUuid(v string) *DescribeAlertsWithEntityRequest
	GetEntityUuid() *string
	SetIncidentUuid(v string) *DescribeAlertsWithEntityRequest
	GetIncidentUuid() *string
	SetPageSize(v int32) *DescribeAlertsWithEntityRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAlertsWithEntityRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertsWithEntityRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertsWithEntityRequest
	GetRoleType() *int32
	SetSophonTaskId(v string) *DescribeAlertsWithEntityRequest
	GetSophonTaskId() *string
	SetStartTime(v int64) *DescribeAlertsWithEntityRequest
	GetStartTime() *int64
}

type DescribeAlertsWithEntityRequest struct {
	// The number of the page to return. The value must be greater than or equal to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query. The unit is milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the entity.
	//
	// example:
	//
	// 123456789
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The ID of the entity.
	//
	// example:
	//
	// 123456789
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The globally unique ID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The number of entries to return on each page. The maximum value is 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region of the data management center for threat analysis. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: your assets are in the Chinese mainland and China (Hong Kong).
	//
	// - ap-southeast-1: your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that are managed by your enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The ID of the Security Orchestration, Automation, and Response (SOAR) handling policy.
	//
	// example:
	//
	// 577bbf90-a770-44a7-8154-586aa2d318fa
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The start time of the query. The unit is milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertsWithEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEntityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAlertsWithEntityRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlertsWithEntityRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DescribeAlertsWithEntityRequest) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *DescribeAlertsWithEntityRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeAlertsWithEntityRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertsWithEntityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertsWithEntityRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertsWithEntityRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertsWithEntityRequest) GetSophonTaskId() *string {
	return s.SophonTaskId
}

func (s *DescribeAlertsWithEntityRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlertsWithEntityRequest) SetCurrentPage(v int32) *DescribeAlertsWithEntityRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetEndTime(v int64) *DescribeAlertsWithEntityRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetEntityId(v int64) *DescribeAlertsWithEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetEntityUuid(v string) *DescribeAlertsWithEntityRequest {
	s.EntityUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetIncidentUuid(v string) *DescribeAlertsWithEntityRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetPageSize(v int32) *DescribeAlertsWithEntityRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetRegionId(v string) *DescribeAlertsWithEntityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetRoleFor(v int64) *DescribeAlertsWithEntityRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetRoleType(v int32) *DescribeAlertsWithEntityRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetSophonTaskId(v string) *DescribeAlertsWithEntityRequest {
	s.SophonTaskId = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) SetStartTime(v int64) *DescribeAlertsWithEntityRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsWithEntityRequest) Validate() error {
	return dara.Validate(s)
}
