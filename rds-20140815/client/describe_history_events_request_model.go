// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveStatus(v string) *DescribeHistoryEventsRequest
	GetArchiveStatus() *string
	SetEventCategory(v string) *DescribeHistoryEventsRequest
	GetEventCategory() *string
	SetEventId(v string) *DescribeHistoryEventsRequest
	GetEventId() *string
	SetEventLevel(v string) *DescribeHistoryEventsRequest
	GetEventLevel() *string
	SetEventStatus(v string) *DescribeHistoryEventsRequest
	GetEventStatus() *string
	SetEventType(v string) *DescribeHistoryEventsRequest
	GetEventType() *string
	SetFromStartTime(v string) *DescribeHistoryEventsRequest
	GetFromStartTime() *string
	SetInstanceId(v string) *DescribeHistoryEventsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeHistoryEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryEventsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHistoryEventsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeHistoryEventsRequest
	GetResourceGroupId() *string
	SetResourceType(v string) *DescribeHistoryEventsRequest
	GetResourceType() *string
	SetSecurityToken(v string) *DescribeHistoryEventsRequest
	GetSecurityToken() *string
	SetTaskId(v string) *DescribeHistoryEventsRequest
	GetTaskId() *string
	SetToStartTime(v string) *DescribeHistoryEventsRequest
	GetToStartTime() *string
}

type DescribeHistoryEventsRequest struct {
	// The resource status. Valid values: **importing**, failed, checksuccess, and deleted.
	//
	// example:
	//
	// deleted
	ArchiveStatus *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	// The system event category. For more information, see [View the event history of an ApsaraDB RDS instance](https://help.aliyun.com/document_detail/129759.html).
	//
	// example:
	//
	// Exception
	EventCategory *string `json:"EventCategory,omitempty" xml:"EventCategory,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 5345398
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The event level. Valid values: ***high***, **medium**, and **low**.
	//
	// example:
	//
	// high
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The status of the exception. Valid values:
	//
	// 	- 1: pending
	//
	// 	- 2: ignored
	//
	// 	- 4: confirmed
	//
	// 	- 8: marked as false positive
	//
	// 	- 16: handling
	//
	// 	- 32: handled
	//
	// 	- 64: expired
	//
	// example:
	//
	// 1
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The system event type. This parameter takes effect only when InstanceEventType.N is not specified. Valid values:
	//
	// 	- SystemMaintenance.Reboot: The instance is restarted due to system maintenance.
	//
	// 	- SystemMaintenance.Redeploy: The instance is redeployed due to system maintenance.
	//
	// 	- SystemFailure.Reboot: The instance is restarted due to a system error.
	//
	// 	- SystemFailure.Redeploy: The instance is redeployed due to a system error.
	//
	// 	- SystemFailure.Delete: The instance is released due to an instance creation failure.
	//
	// 	- InstanceFailure.Reboot: The instance is restarted due to an instance error.
	//
	// 	- InstanceExpiration.Stop: The subscription instance is stopped due to expiration.
	//
	// 	- InstanceExpiration.Delete: The subscription instance is released due to expiration.
	//
	// 	- AccountUnbalanced.Stop: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// 	- AccountUnbalanced.Delete: The pay-as-you-go instance is released due to an overdue payment.
	//
	// >  For more information, see Overview. The values of this parameter are applicable only to instance system events, but not to disk system events.
	//
	// example:
	//
	// SystemFailure.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The beginning of the time range to query. Only tasks that have a start time later than or equal to the time specified by this parameter are queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The start time can be up to 30 days earlier than the current time. If you set this parameter to a time more than 30 days earlier than the current time, this time is automatically converted to a time that is exactly 30 days earlier than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-02T11:31:03Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf62br2491p5l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 30.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource type. Set the value to **INSTANCE**.
	//
	// example:
	//
	// INSTANCE
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task ID. This value is used to query the data of a specific task.
	//
	// example:
	//
	// 241535739
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The end of the time range to query. Only tasks that have a start time earlier than or equal to the time specified by this parameter are queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-12T07:06:19Z
	ToStartTime *string `json:"ToStartTime,omitempty" xml:"ToStartTime,omitempty"`
}

func (s DescribeHistoryEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsRequest) GetArchiveStatus() *string {
	return s.ArchiveStatus
}

func (s *DescribeHistoryEventsRequest) GetEventCategory() *string {
	return s.EventCategory
}

func (s *DescribeHistoryEventsRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeHistoryEventsRequest) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeHistoryEventsRequest) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeHistoryEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeHistoryEventsRequest) GetFromStartTime() *string {
	return s.FromStartTime
}

func (s *DescribeHistoryEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHistoryEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHistoryEventsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeHistoryEventsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeHistoryEventsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHistoryEventsRequest) GetToStartTime() *string {
	return s.ToStartTime
}

func (s *DescribeHistoryEventsRequest) SetArchiveStatus(v string) *DescribeHistoryEventsRequest {
	s.ArchiveStatus = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetEventCategory(v string) *DescribeHistoryEventsRequest {
	s.EventCategory = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetEventId(v string) *DescribeHistoryEventsRequest {
	s.EventId = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetEventLevel(v string) *DescribeHistoryEventsRequest {
	s.EventLevel = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetEventStatus(v string) *DescribeHistoryEventsRequest {
	s.EventStatus = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetEventType(v string) *DescribeHistoryEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetFromStartTime(v string) *DescribeHistoryEventsRequest {
	s.FromStartTime = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetInstanceId(v string) *DescribeHistoryEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetPageNumber(v int32) *DescribeHistoryEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetPageSize(v int32) *DescribeHistoryEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetRegionId(v string) *DescribeHistoryEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetResourceGroupId(v string) *DescribeHistoryEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetResourceType(v string) *DescribeHistoryEventsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetSecurityToken(v string) *DescribeHistoryEventsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetTaskId(v string) *DescribeHistoryEventsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetToStartTime(v string) *DescribeHistoryEventsRequest {
	s.ToStartTime = &v
	return s
}

func (s *DescribeHistoryEventsRequest) Validate() error {
	return dara.Validate(s)
}
