// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceHistoryEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventPublishTime(v *DescribeRCInstanceHistoryEventsRequestEventPublishTime) *DescribeRCInstanceHistoryEventsRequest
	GetEventPublishTime() *DescribeRCInstanceHistoryEventsRequestEventPublishTime
	SetNotBefore(v *DescribeRCInstanceHistoryEventsRequestNotBefore) *DescribeRCInstanceHistoryEventsRequest
	GetNotBefore() *DescribeRCInstanceHistoryEventsRequestNotBefore
	SetEventCycleStatus(v string) *DescribeRCInstanceHistoryEventsRequest
	GetEventCycleStatus() *string
	SetEventId(v []*string) *DescribeRCInstanceHistoryEventsRequest
	GetEventId() []*string
	SetEventType(v string) *DescribeRCInstanceHistoryEventsRequest
	GetEventType() *string
	SetImpactLevel(v string) *DescribeRCInstanceHistoryEventsRequest
	GetImpactLevel() *string
	SetInstanceEventCycleStatus(v []*string) *DescribeRCInstanceHistoryEventsRequest
	GetInstanceEventCycleStatus() []*string
	SetInstanceEventType(v []*string) *DescribeRCInstanceHistoryEventsRequest
	GetInstanceEventType() []*string
	SetInstanceId(v string) *DescribeRCInstanceHistoryEventsRequest
	GetInstanceId() *string
	SetMaxResults(v string) *DescribeRCInstanceHistoryEventsRequest
	GetMaxResults() *string
	SetPageNumber(v string) *DescribeRCInstanceHistoryEventsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeRCInstanceHistoryEventsRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeRCInstanceHistoryEventsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRCInstanceHistoryEventsRequest
	GetResourceGroupId() *string
	SetResourceId(v []*string) *DescribeRCInstanceHistoryEventsRequest
	GetResourceId() []*string
	SetTag(v []*DescribeRCInstanceHistoryEventsRequestTag) *DescribeRCInstanceHistoryEventsRequest
	GetTag() []*DescribeRCInstanceHistoryEventsRequestTag
}

type DescribeRCInstanceHistoryEventsRequest struct {
	EventPublishTime *DescribeRCInstanceHistoryEventsRequestEventPublishTime `json:"EventPublishTime,omitempty" xml:"EventPublishTime,omitempty" type:"Struct"`
	NotBefore        *DescribeRCInstanceHistoryEventsRequestNotBefore        `json:"NotBefore,omitempty" xml:"NotBefore,omitempty" type:"Struct"`
	// The lifecycle state of the system event. This parameter is valid only when the **InstanceEventCycleStatus.N*	- parameter is not specified. Valid values:
	//
	// 	- **Scheduled**
	//
	// 	- **Avoided**
	//
	// 	- **Executing**
	//
	// 	- **Executed**
	//
	// 	- **Canceled**
	//
	// 	- **Failed**
	//
	// 	- **Inquiring**
	//
	// example:
	//
	// Executed
	EventCycleStatus *string `json:"EventCycleStatus,omitempty" xml:"EventCycleStatus,omitempty"`
	// The IDs of one or more system events.
	EventId []*string `json:"EventId,omitempty" xml:"EventId,omitempty" type:"Repeated"`
	// The system event type. This parameter is valid only when the **InstanceEventType.N*	- parameter is not specified. Valid values:
	//
	// 	- **SystemMaintenance.Reboot**: The instance was restarted due to system maintenance.
	//
	// 	- **SystemMaintenance.Redeploy**: The instance was redeployed due to system maintenance.
	//
	// 	- **SystemFailure.Reboot**: The instance was restarted due to system failures.
	//
	// 	- **SystemFailure.Redeploy**: The instance was redeployed due to system failures.
	//
	// 	- **SystemFailure.Delete**: The instance was released due to an instance creation failure.
	//
	// 	- **InstanceFailure.Reboot**: The instance was restarted due to an instance error.
	//
	// 	- **InstanceExpiration.Stop**: The subscription instance was stopped due to expiration.
	//
	// 	- **InstanceExpiration.Delete**: The subscription instance was released due to expiration.
	//
	// 	- **AccountUnbalanced.Stop**: The pay-as-you-go instance is stopped due to an overdue payment.
	//
	// 	- **AccountUnbalanced.Delete**: The pay-as-you-go instance was released due to an overdue payment.
	//
	// >  The values of this parameter are applicable only to instance system events, but not to disk system events.
	//
	// example:
	//
	// SystemMaintenance.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	ImpactLevel *string `json:"ImpactLevel,omitempty" xml:"ImpactLevel,omitempty"`
	// The lifecycle states of system events.
	InstanceEventCycleStatus []*string `json:"InstanceEventCycleStatus,omitempty" xml:"InstanceEventCycleStatus,omitempty" type:"Repeated"`
	// The type of system event N.
	InstanceEventType []*string `json:"InstanceEventType,omitempty" xml:"InstanceEventType,omitempty" type:"Repeated"`
	// The instance ID. If you do not specify an instance ID, system events of all instances in the specified region are queried.
	//
	// example:
	//
	// rc-yuf59nplc45t2tzn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group that you want to query.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of resource N.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// An array that consists of the tags that are supported by system events.
	Tag []*DescribeRCInstanceHistoryEventsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceHistoryEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetEventPublishTime() *DescribeRCInstanceHistoryEventsRequestEventPublishTime {
	return s.EventPublishTime
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetNotBefore() *DescribeRCInstanceHistoryEventsRequestNotBefore {
	return s.NotBefore
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetEventCycleStatus() *string {
	return s.EventCycleStatus
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetEventId() []*string {
	return s.EventId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetImpactLevel() *string {
	return s.ImpactLevel
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetInstanceEventCycleStatus() []*string {
	return s.InstanceEventCycleStatus
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetInstanceEventType() []*string {
	return s.InstanceEventType
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeRCInstanceHistoryEventsRequest) GetTag() []*DescribeRCInstanceHistoryEventsRequestTag {
	return s.Tag
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetEventPublishTime(v *DescribeRCInstanceHistoryEventsRequestEventPublishTime) *DescribeRCInstanceHistoryEventsRequest {
	s.EventPublishTime = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetNotBefore(v *DescribeRCInstanceHistoryEventsRequestNotBefore) *DescribeRCInstanceHistoryEventsRequest {
	s.NotBefore = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetEventCycleStatus(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.EventCycleStatus = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetEventId(v []*string) *DescribeRCInstanceHistoryEventsRequest {
	s.EventId = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetEventType(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetImpactLevel(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.ImpactLevel = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetInstanceEventCycleStatus(v []*string) *DescribeRCInstanceHistoryEventsRequest {
	s.InstanceEventCycleStatus = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetInstanceEventType(v []*string) *DescribeRCInstanceHistoryEventsRequest {
	s.InstanceEventType = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetInstanceId(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetMaxResults(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetPageNumber(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetPageSize(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetRegionId(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetResourceGroupId(v string) *DescribeRCInstanceHistoryEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetResourceId(v []*string) *DescribeRCInstanceHistoryEventsRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) SetTag(v []*DescribeRCInstanceHistoryEventsRequestTag) *DescribeRCInstanceHistoryEventsRequest {
	s.Tag = v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequest) Validate() error {
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

type DescribeRCInstanceHistoryEventsRequestEventPublishTime struct {
	// The end of the time range in which to query published system events. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-04-01T06:32:31Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The beginning of the time range in which to query published system events. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-03-30T06:32:31Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsRequestEventPublishTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsRequestEventPublishTime) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) GetEnd() *string {
	return s.End
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) GetStart() *string {
	return s.Start
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) SetEnd(v string) *DescribeRCInstanceHistoryEventsRequestEventPublishTime {
	s.End = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) SetStart(v string) *DescribeRCInstanceHistoryEventsRequestEventPublishTime {
	s.Start = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestEventPublishTime) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceHistoryEventsRequestNotBefore struct {
	// The end time of the scheduled execution period for the system event. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-04-01T06:32:31Z
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start time of the scheduled execution period for the system event. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-03-30T06:32:31Z
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsRequestNotBefore) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsRequestNotBefore) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) GetEnd() *string {
	return s.End
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) GetStart() *string {
	return s.Start
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) SetEnd(v string) *DescribeRCInstanceHistoryEventsRequestNotBefore {
	s.End = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) SetStart(v string) *DescribeRCInstanceHistoryEventsRequestNotBefore {
	s.Start = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestNotBefore) Validate() error {
	return dara.Validate(s)
}

type DescribeRCInstanceHistoryEventsRequestTag struct {
	// The key of the tag that is added to the resource.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the port list.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRCInstanceHistoryEventsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceHistoryEventsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) SetKey(v string) *DescribeRCInstanceHistoryEventsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) SetValue(v string) *DescribeRCInstanceHistoryEventsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeRCInstanceHistoryEventsRequestTag) Validate() error {
	return dara.Validate(s)
}
