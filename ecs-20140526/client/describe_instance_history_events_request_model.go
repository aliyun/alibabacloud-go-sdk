// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceHistoryEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventPublishTime(v *DescribeInstanceHistoryEventsRequestEventPublishTime) *DescribeInstanceHistoryEventsRequest
	GetEventPublishTime() *DescribeInstanceHistoryEventsRequestEventPublishTime
	SetNotBefore(v *DescribeInstanceHistoryEventsRequestNotBefore) *DescribeInstanceHistoryEventsRequest
	GetNotBefore() *DescribeInstanceHistoryEventsRequestNotBefore
	SetEventCycleStatus(v string) *DescribeInstanceHistoryEventsRequest
	GetEventCycleStatus() *string
	SetEventId(v []*string) *DescribeInstanceHistoryEventsRequest
	GetEventId() []*string
	SetEventType(v string) *DescribeInstanceHistoryEventsRequest
	GetEventType() *string
	SetImpactLevel(v string) *DescribeInstanceHistoryEventsRequest
	GetImpactLevel() *string
	SetInstanceEventCycleStatus(v []*string) *DescribeInstanceHistoryEventsRequest
	GetInstanceEventCycleStatus() []*string
	SetInstanceEventType(v []*string) *DescribeInstanceHistoryEventsRequest
	GetInstanceEventType() []*string
	SetInstanceId(v string) *DescribeInstanceHistoryEventsRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *DescribeInstanceHistoryEventsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeInstanceHistoryEventsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeInstanceHistoryEventsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceHistoryEventsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstanceHistoryEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceHistoryEventsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstanceHistoryEventsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstanceHistoryEventsRequest
	GetResourceGroupId() *string
	SetResourceId(v []*string) *DescribeInstanceHistoryEventsRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *DescribeInstanceHistoryEventsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceHistoryEventsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeInstanceHistoryEventsRequest
	GetResourceType() *string
	SetTag(v []*DescribeInstanceHistoryEventsRequestTag) *DescribeInstanceHistoryEventsRequest
	GetTag() []*DescribeInstanceHistoryEventsRequestTag
}

type DescribeInstanceHistoryEventsRequest struct {
	EventPublishTime *DescribeInstanceHistoryEventsRequestEventPublishTime `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty" type:"Struct"`
	NotBefore        *DescribeInstanceHistoryEventsRequestNotBefore        `json:"NotBefore,omitempty" xml:"NotBefore,omitempty" type:"Struct"`
	// The lifecycle status of the system event. EventCycleStatus takes effect only when InstanceEventCycleStatus.N is not specified. Valid values:
	//
	// - Scheduled: The event is waiting to be executed.
	//
	// - Avoided: The event has been avoided.
	//
	// - Executing: The event is being executed.
	//
	// - Executed: The event has been executed.
	//
	// - Canceled: The event has been canceled.
	//
	// - Failed: The event execution failed.
	//
	// - Inquiring: The event is being inquired.
	//
	// example:
	//
	// Executed
	EventCycleStatus *string `json:"EventCycleStatus,omitempty" xml:"EventCycleStatus,omitempty"`
	// One or more system event IDs. Valid values of N: 1 to 100. Specify multiple values in a repeated list format.
	//
	// example:
	//
	// e-uf64yvznlao4jl2c****
	EventId []*string `json:"EventId,omitempty" xml:"EventId,omitempty" type:"Repeated"`
	// The type of the system event. EventType takes effect only when InstanceEventType.N is not specified. Valid values:
	//
	// - SystemMaintenance.Reboot: The instance is restarted due to system maintenance.
	//
	// - SystemMaintenance.Redeploy: The instance is redeployed due to system maintenance.
	//
	// - SystemFailure.Reboot: The instance is restarted due to a system error.
	//
	// - SystemFailure.Redeploy: The instance is redeployed due to a system error.
	//
	// - SystemFailure.Delete: The instance is released due to an instance creation failure.
	//
	// - InstanceFailure.Reboot: The instance is restarted due to an instance error.
	//
	// - InstanceExpiration.Stop: The instance is stopped due to subscription expiration.
	//
	// - InstanceExpiration.Delete: The instance is released due to subscription expiration.
	//
	// - AccountUnbalanced.Stop: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// - AccountUnbalanced.Delete: The pay-as-you-go instance is released due to an overdue payment.
	//
	// > For more information about event types, see [System event overview](https://help.aliyun.com/document_detail/66574.html). The value of this parameter must be an instance system event, not a disk system event.
	//
	// example:
	//
	// SystemMaintenance.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// > This parameter is not yet available.
	//
	// example:
	//
	// null
	ImpactLevel *string `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
	// One or more lifecycle statuses of the system event. Valid values of N: 1 to 7. Specify multiple values in a repeated list format. Valid values:
	//
	// - Scheduled: The event is waiting to be executed.
	//
	// - Avoided: The event has been avoided.
	//
	// - Executing: The event is being executed.
	//
	// - Executed: The event has been executed.
	//
	// - Canceled: The event has been canceled.
	//
	// - Failed: The event execution failed.
	//
	// - Inquiring: The event is being inquired.
	//
	// example:
	//
	// Executed
	InstanceEventCycleStatus []*string `json:"InstanceEventCycleStatus,omitempty" xml:"InstanceEventCycleStatus,omitempty" type:"Repeated"`
	// One or more types of the system event. Valid values of N: 1 to 30. Specify multiple values in a repeated list format. Valid values:
	//
	// - SystemMaintenance.Reboot: The instance is restarted due to system maintenance.
	//
	// - SystemMaintenance.Redeploy: The instance is redeployed due to system maintenance.
	//
	// - SystemFailure.Reboot: The instance is restarted due to a system error.
	//
	// - SystemFailure.Redeploy: The instance is redeployed due to a system error.
	//
	// - SystemFailure.Delete: The instance is released due to an instance creation failure.
	//
	// - InstanceFailure.Reboot: The instance is restarted due to an instance error.
	//
	// - InstanceExpiration.Stop: The instance is stopped due to subscription expiration.
	//
	// - InstanceExpiration.Delete: The instance is released due to subscription expiration.
	//
	// - AccountUnbalanced.Stop: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// - AccountUnbalanced.Delete: The pay-as-you-go instance is released due to an overdue payment.
	//
	// > For more information about event types, see [System event overview](https://help.aliyun.com/document_detail/66574.html). The value of this parameter must be an instance system event, not a disk system event.
	//
	// example:
	//
	// SystemMaintenance.Reboot
	InstanceEventType []*string `json:"InstanceEventType,omitempty" xml:"InstanceEventType,omitempty" type:"Repeated"`
	// The instance ID. If you do not specify an instance ID, the system events of all instances in the specified region are queried.
	//
	// example:
	//
	// i-uf678mass4zvr9n1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page for a paging query. Valid values: 10 to 100.
	//
	// Default value:
	//
	// 	- When the settings value is greater than 0 but less than 10, the default value is 10.
	//
	// 	- When the settings value is greater than 100, the default value is 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// f1c9fa9de5752***
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter will be deprecated. Use MaxResults or NextToken for paginated queries instead.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter will be deprecated. Use MaxResults or NextToken for paginated queries instead.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the resource. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// One or more resource IDs. Valid values of N: 1 to 100. Specify multiple values in a repeated list format. Valid values:
	//
	// - When `ResourceType=instance`, the resource ID is the ECS instance ID.
	//
	// - When `ResourceType=ddh`, the resource ID is the dedicated host ID.
	//
	// - When `ResourceType=managedhost`, the resource ID is the physical machine ID in an intelligent fully managed resource pool.
	//
	// If you do not specify this parameter, the system events of all resources of the specified resource type (`ResourceType`) in the specified region (`RegionId`) are queried.
	//
	// > Use `ResourceId.N` to specify one or more resource IDs. If you specify both `ResourceId.N` and `InstanceId`, `ResourceId.N` takes precedence by default.
	//
	// example:
	//
	// i-uf678mass4zvr9n1****
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// - instance: ECS instance.
	//
	// - ddh: dedicated host.
	//
	// - managedhost: physical machine in an intelligent fully managed resource pool.
	//
	// Default value: instance.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags supported by system events.
	Tag []*DescribeInstanceHistoryEventsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceHistoryEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsRequest) GetEventPublishTime() *DescribeInstanceHistoryEventsRequestEventPublishTime {
	return s.EventPublishTime
}

func (s *DescribeInstanceHistoryEventsRequest) GetNotBefore() *DescribeInstanceHistoryEventsRequestNotBefore {
	return s.NotBefore
}

func (s *DescribeInstanceHistoryEventsRequest) GetEventCycleStatus() *string {
	return s.EventCycleStatus
}

func (s *DescribeInstanceHistoryEventsRequest) GetEventId() []*string {
	return s.EventId
}

func (s *DescribeInstanceHistoryEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeInstanceHistoryEventsRequest) GetImpactLevel() *string {
	return s.ImpactLevel
}

func (s *DescribeInstanceHistoryEventsRequest) GetInstanceEventCycleStatus() []*string {
	return s.InstanceEventCycleStatus
}

func (s *DescribeInstanceHistoryEventsRequest) GetInstanceEventType() []*string {
	return s.InstanceEventType
}

func (s *DescribeInstanceHistoryEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceHistoryEventsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeInstanceHistoryEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceHistoryEventsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceHistoryEventsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceHistoryEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceHistoryEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceHistoryEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceHistoryEventsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeInstanceHistoryEventsRequest) GetTag() []*DescribeInstanceHistoryEventsRequestTag {
	return s.Tag
}

func (s *DescribeInstanceHistoryEventsRequest) SetEventPublishTime(v *DescribeInstanceHistoryEventsRequestEventPublishTime) *DescribeInstanceHistoryEventsRequest {
	s.EventPublishTime = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetNotBefore(v *DescribeInstanceHistoryEventsRequestNotBefore) *DescribeInstanceHistoryEventsRequest {
	s.NotBefore = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetEventCycleStatus(v string) *DescribeInstanceHistoryEventsRequest {
	s.EventCycleStatus = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetEventId(v []*string) *DescribeInstanceHistoryEventsRequest {
	s.EventId = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetEventType(v string) *DescribeInstanceHistoryEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetImpactLevel(v string) *DescribeInstanceHistoryEventsRequest {
	s.ImpactLevel = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetInstanceEventCycleStatus(v []*string) *DescribeInstanceHistoryEventsRequest {
	s.InstanceEventCycleStatus = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetInstanceEventType(v []*string) *DescribeInstanceHistoryEventsRequest {
	s.InstanceEventType = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetInstanceId(v string) *DescribeInstanceHistoryEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetMaxResults(v int64) *DescribeInstanceHistoryEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetNextToken(v string) *DescribeInstanceHistoryEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetOwnerAccount(v string) *DescribeInstanceHistoryEventsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetOwnerId(v int64) *DescribeInstanceHistoryEventsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetPageNumber(v int32) *DescribeInstanceHistoryEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetPageSize(v int32) *DescribeInstanceHistoryEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetRegionId(v string) *DescribeInstanceHistoryEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceGroupId(v string) *DescribeInstanceHistoryEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceId(v []*string) *DescribeInstanceHistoryEventsRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceOwnerAccount(v string) *DescribeInstanceHistoryEventsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceOwnerId(v int64) *DescribeInstanceHistoryEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetResourceType(v string) *DescribeInstanceHistoryEventsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) SetTag(v []*DescribeInstanceHistoryEventsRequestTag) *DescribeInstanceHistoryEventsRequest {
	s.Tag = v
	return s
}

func (s *DescribeInstanceHistoryEventsRequest) Validate() error {
	if s.EventPublishTime != nil {
		if err := s.EventPublishTime.Validate(); err != nil {
			return err
		}
	}
	if s.NotBefore != nil {
		if err := s.NotBefore.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceHistoryEventsRequestEventPublishTime struct {
	// The end of the time range during which the system event is published. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-01T06:32:31Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start of the time range during which the system event is published. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeInstanceHistoryEventsRequestEventPublishTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsRequestEventPublishTime) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) GetEnd() *string {
	return s.End
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) GetStart() *string {
	return s.Start
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) SetEnd(v string) *DescribeInstanceHistoryEventsRequestEventPublishTime {
	s.End = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) SetStart(v string) *DescribeInstanceHistoryEventsRequestEventPublishTime {
	s.Start = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestEventPublishTime) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHistoryEventsRequestNotBefore struct {
	// The end of the time range during which the system event is scheduled to execute. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-01T06:32:31Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start of the time range during which the system event is scheduled to execute. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeInstanceHistoryEventsRequestNotBefore) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsRequestNotBefore) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) GetEnd() *string {
	return s.End
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) GetStart() *string {
	return s.Start
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) SetEnd(v string) *DescribeInstanceHistoryEventsRequestNotBefore {
	s.End = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) SetStart(v string) *DescribeInstanceHistoryEventsRequestNotBefore {
	s.Start = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestNotBefore) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHistoryEventsRequestTag struct {
	// The tag key of the resource.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceHistoryEventsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHistoryEventsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHistoryEventsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceHistoryEventsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceHistoryEventsRequestTag) SetKey(v string) *DescribeInstanceHistoryEventsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestTag) SetValue(v string) *DescribeInstanceHistoryEventsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInstanceHistoryEventsRequestTag) Validate() error {
	return dara.Validate(s)
}
