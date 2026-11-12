// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetId(v string) *DescribeCloudSiemEventsRequest
	GetAssetId() *string
	SetCurrentPage(v int32) *DescribeCloudSiemEventsRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeCloudSiemEventsRequest
	GetEndTime() *int64
	SetEntityUuid(v string) *DescribeCloudSiemEventsRequest
	GetEntityUuid() *string
	SetEventName(v string) *DescribeCloudSiemEventsRequest
	GetEventName() *string
	SetIncidentUuid(v string) *DescribeCloudSiemEventsRequest
	GetIncidentUuid() *string
	SetOrder(v string) *DescribeCloudSiemEventsRequest
	GetOrder() *string
	SetOrderField(v string) *DescribeCloudSiemEventsRequest
	GetOrderField() *string
	SetPageSize(v int32) *DescribeCloudSiemEventsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCloudSiemEventsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeCloudSiemEventsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeCloudSiemEventsRequest
	GetRoleType() *int32
	SetStartTime(v int64) *DescribeCloudSiemEventsRequest
	GetStartTime() *int64
	SetStatus(v int32) *DescribeCloudSiemEventsRequest
	GetStatus() *int32
	SetThreadLevel(v []*string) *DescribeCloudSiemEventsRequest
	GetThreadLevel() []*string
}

type DescribeCloudSiemEventsRequest struct {
	// The ID of the asset associated with the event.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The current page number. The value must be greater than or equal to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time for querying events, in milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The UUID of the entity associated with the event.
	//
	// example:
	//
	// 6c740667-80b2-476d-8924-2e706feb****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// ECS unusual log in
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The sort order of the event list. Valid values:
	//
	// - desc: Descending order.
	//
	// - asc: Ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field used to sort the event list. Valid values:
	//
	// - GmtModified: Sorts by event generation time (default).
	//
	// - ThreatScore: Sorts by event threat score.
	//
	// example:
	//
	// ThreatScore
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center of Cloud Threat Detection and Response (CTDR) is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose view the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: Current Alibaba Cloud account view.
	//
	// - 1: View of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The start time for querying events, in milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the event. Valid values:
	//
	// - 0: Unhandled.
	//
	// - 1: Processing.
	//
	// - 5: Failed.
	//
	// - 10: Handled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The threat level of the event, in JSON array format. Valid values:
	//
	// - serious: High.
	//
	// - suspicious: Medium.
	//
	// - remind: Low.
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	ThreadLevel []*string `json:"ThreadLevel,omitempty" xml:"ThreadLevel,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsRequest) GetAssetId() *string {
	return s.AssetId
}

func (s *DescribeCloudSiemEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudSiemEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeCloudSiemEventsRequest) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *DescribeCloudSiemEventsRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeCloudSiemEventsRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeCloudSiemEventsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeCloudSiemEventsRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *DescribeCloudSiemEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudSiemEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudSiemEventsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeCloudSiemEventsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeCloudSiemEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeCloudSiemEventsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCloudSiemEventsRequest) GetThreadLevel() []*string {
	return s.ThreadLevel
}

func (s *DescribeCloudSiemEventsRequest) SetAssetId(v string) *DescribeCloudSiemEventsRequest {
	s.AssetId = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetCurrentPage(v int32) *DescribeCloudSiemEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetEndTime(v int64) *DescribeCloudSiemEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetEntityUuid(v string) *DescribeCloudSiemEventsRequest {
	s.EntityUuid = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetEventName(v string) *DescribeCloudSiemEventsRequest {
	s.EventName = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetIncidentUuid(v string) *DescribeCloudSiemEventsRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetOrder(v string) *DescribeCloudSiemEventsRequest {
	s.Order = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetOrderField(v string) *DescribeCloudSiemEventsRequest {
	s.OrderField = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetPageSize(v int32) *DescribeCloudSiemEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetRegionId(v string) *DescribeCloudSiemEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetRoleFor(v int64) *DescribeCloudSiemEventsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetRoleType(v int32) *DescribeCloudSiemEventsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetStartTime(v int64) *DescribeCloudSiemEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetStatus(v int32) *DescribeCloudSiemEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetThreadLevel(v []*string) *DescribeCloudSiemEventsRequest {
	s.ThreadLevel = v
	return s
}

func (s *DescribeCloudSiemEventsRequest) Validate() error {
	return dara.Validate(s)
}
