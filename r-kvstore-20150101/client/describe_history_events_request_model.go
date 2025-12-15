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
	SetOwnerAccount(v string) *DescribeHistoryEventsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHistoryEventsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeHistoryEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryEventsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHistoryEventsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeHistoryEventsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeHistoryEventsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHistoryEventsRequest
	GetResourceOwnerId() *int64
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
	// The status of the event. Valid values:
	//
	// 	- **Archived**
	//
	// 	- **UnArchived**
	//
	// 	- **All**
	//
	// example:
	//
	// Archived
	ArchiveStatus *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	// The system event category. Valid values:
	//
	// 	- **Exception**: abnormal events.
	//
	// 	- **Optimize**: optimization events.
	//
	// 	- **Notification**: Notification events.
	//
	// 	- **Maintenance**: scheduled maintenance events.
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
	// The level of the event. Valid values:
	//
	// 	- **INFO**
	//
	// 	- **WARN**
	//
	// 	- **CRITICAL**
	//
	// example:
	//
	// INFO
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- **Inquiring**
	//
	// 	- **Scheduled**: Planned
	//
	// 	- **Running**
	//
	// 	- **Succeed**
	//
	// 	- **Failed**
	//
	// 	- **Canceled**
	//
	// > Separate multiple states with commas (,).
	//
	// example:
	//
	// Scheduled
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
	// > For more information about event types, see [Overview of system events](https://help.aliyun.com/document_detail/66574.html). The values of this parameter are applicable only to instance system events, but not to disk system events.
	//
	// example:
	//
	// SystemFailure.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-02T11:31:03Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfnuslkubs****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **Instance**: instance resources.
	//
	// 	- **Host**: host resources.
	//
	// 	- **User**: user resources
	//
	// > If you do not specify this parameter, events of all resource types are queried.
	//
	// example:
	//
	// Instance
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task IDs.
	//
	// example:
	//
	// 578678678
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Only events that have a start time earlier than or equal to the time specified by this parameter are queried.
	//
	// example:
	//
	// 2022-02-02T11:31:03Z
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

func (s *DescribeHistoryEventsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHistoryEventsRequest) GetOwnerId() *int64 {
	return s.OwnerId
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

func (s *DescribeHistoryEventsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHistoryEventsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *DescribeHistoryEventsRequest) SetOwnerAccount(v string) *DescribeHistoryEventsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetOwnerId(v int64) *DescribeHistoryEventsRequest {
	s.OwnerId = &v
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

func (s *DescribeHistoryEventsRequest) SetResourceOwnerAccount(v string) *DescribeHistoryEventsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHistoryEventsRequest) SetResourceOwnerId(v int64) *DescribeHistoryEventsRequest {
	s.ResourceOwnerId = &v
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
