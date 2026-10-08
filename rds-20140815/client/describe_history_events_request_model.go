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
	// The event status. Valid values:
	//
	// - **Archived**: archived.
	//
	// - **UnArchived**: not archived.
	//
	// - **All**: all.
	//
	// example:
	//
	// All
	ArchiveStatus *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	// The system event categorization. Valid values:
	//
	// - **Exception**: abnormal event.
	//
	// - **Optimize**: optimization events.
	//
	// - **Notification**: notification event.
	//
	// - **Maintenance**: scheduled maintenance event.
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
	// The event level. Valid values:
	//
	// - **INFO**: notification.
	//
	// - **WARN**: warning.
	//
	// - **CRITICAL**: critical.
	//
	// example:
	//
	// INFO
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The event status. Valid values:
	//
	// - **Inquiring**: inquiring.
	//
	// - **Scheduled**: scheduled.
	//
	// - **Running**: running.
	//
	// - **Succeed**: completed.
	//
	// - **Failed**: failed.
	//
	// - **Canceled**: canceled.
	//
	// > To query multiple statuses, separate them with commas (,).
	//
	// example:
	//
	// Scheduled
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The system event type. This parameter takes effect only when InstanceEventType.N is not specified. Valid values:
	//
	// - **SystemMaintenance.Reboot**: The instance is restarted due to system maintenance.
	//
	// - **SystemMaintenance.Redeploy**: The instance is redeployed due to system maintenance.
	//
	// - **SystemFailure.Reboot**: The instance is restarted due to a system error.
	//
	// - **SystemFailure.Redeploy**: The instance is redeployed due to a system error.
	//
	// - **SystemFailure.Delete**: The instance is released due to an instance creation failure.
	//
	// - **InstanceFailure.Reboot**: The instance is restarted due to an instance error.
	//
	// - **InstanceExpiration.Stop**: The instance is stopped due to subscription expiration.
	//
	// - **InstanceExpiration.Delete**: The instance is released due to subscription expiration.
	//
	// - **AccountUnbalanced.Stop**: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// - **AccountUnbalanced.Delete**: The pay-as-you-go instance is released due to an overdue payment.
	//
	// > The value of this parameter can only be an instance system event, not a cloud disk system event.
	//
	// example:
	//
	// SystemFailure.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The beginning of the time range for the task start time. Tasks whose start time is later than this time are queried. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in `UTC +0`. The earliest supported time is 30 days before the current time. If the specified time is more than 30 days before the current time, it is automatically converted to 30 days before the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-02T11:31:03Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The ApsaraDB RDS instance ID.
	//
	// example:
	//
	// rm-uf62br2491p5l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value must be greater than 0 and cannot exceed the maximum value of the integer type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **30**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) to query the most recent region list.
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
	// The resource type. Valid values:
	//
	// - **Instance**: instance resource.
	//
	// - **Host**: host resource.
	//
	// - **User**: user resource.
	//
	// > If this parameter is not specified, all resource types are queried.
	//
	// example:
	//
	// Instance
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task ID. Specify this parameter to retrieve data for a specific task.
	//
	// example:
	//
	// 241535739
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The end of the time range for the task start time. Tasks whose start time is earlier than this time are queried. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in `UTC +0`.
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
