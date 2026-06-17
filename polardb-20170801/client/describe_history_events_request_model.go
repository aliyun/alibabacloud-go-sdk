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
	SetTaskId(v string) *DescribeHistoryEventsRequest
	GetTaskId() *string
	SetToStartTime(v string) *DescribeHistoryEventsRequest
	GetToStartTime() *string
}

type DescribeHistoryEventsRequest struct {
	// The event status. Valid values:
	//
	// - Archived: The event is archived.
	//
	// - UnArchived: The event is not archived.
	//
	// - All: All events.
	//
	// example:
	//
	// All
	ArchiveStatus *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	// The category of the system event. Valid values:
	//
	// - **Exception**: anomalous activity
	//
	// - **Optimize**: optimization events
	//
	// - **Notification**: notification events
	//
	// - **Maintenance**: scheduled O\\&M events
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
	// The event level. Valid values are:
	//
	// - **INFO**: Notification
	//
	// - **WARN**: Warning
	//
	// - **CRITICAL**: Critical
	//
	// example:
	//
	// high
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The event status. Valid values:
	//
	// - **Inquiring**: The event is being inquired.
	//
	// - **Scheduled**: The event is scheduled.
	//
	// - **Running**: The event is in progress.
	//
	// - **Succeed**: The event is successful.
	//
	// - **Failed**: The event failed.
	//
	// - **Canceled**: The event is canceled.
	//
	// > To query multiple statuses, separate them with commas (,).
	//
	// example:
	//
	// Succeed
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The type of the system event. This parameter is valid only when InstanceEventType.N is not specified. Valid values:
	//
	// - **SystemMaintenance.Reboot**: The instance is restarted due to system maintenance.
	//
	// - **SystemMaintenance.Redeploy**: The instance is redeployed due to system maintenance.
	//
	// - **SystemFailure.Reboot**: The instance is restarted due to a system fault.
	//
	// - **SystemFailure.Redeploy**: The instance is redeployed due to a system fault.
	//
	// - **SystemFailure.Delete**: The instance is released because it failed to be created.
	//
	// - **InstanceFailure.Reboot**: The instance is restarted due to an instance fault.
	//
	// - **InstanceExpiration.Stop**: The subscription instance is stopped because its subscription expires.
	//
	// - **InstanceExpiration.Delete**: The subscription instance is released because its subscription expires.
	//
	// - **AccountUnbalanced.Stop**: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// - **AccountUnbalanced.Delete**: The pay-as-you-go instance is released due to an overdue payment.
	//
	// > The value of this parameter can only be an instance system event, not a disk system event.
	//
	// example:
	//
	// SystemFailure.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The beginning of the time range to query tasks based on their start time. The tasks that started after this time are queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. You can query data from the last 30 days. If the specified time is more than 30 days ago, the system automatically sets it to 30 days ago.
	//
	// example:
	//
	// 2025-01-02T11:31:03Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The resource ID. Use this parameter to query tasks for a specific resource. To query tasks for multiple resources, separate the resource IDs with commas (,). You can specify up to 30 resource IDs. If you leave this parameter empty, all resources are queried.
	//
	// > Currently, only cluster IDs are supported.
	//
	// example:
	//
	// pc-2zed3m89cw***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to query. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-**********
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource type. Valid values:
	//
	// - **Instance**: instance resource
	//
	// - **Host**: host resource
	//
	// - **User**: user resource
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the scheduled task that you want to cancel.
	//
	// > - You can call the [DescribeScheduleTasks](https://help.aliyun.com/document_detail/199648.html) operation to view information about all scheduled tasks under the current account, including task IDs.
	//
	// - Only tasks that are pending execution (the `Status` parameter returns `pending`) can be canceled.
	//
	// example:
	//
	// 32077515
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The end of the time range to query tasks based on their start time. The tasks that started before this time are queried. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in `UTC`.
	//
	// example:
	//
	// 2025-01-03T12:31:03Z
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
