// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribeHistoryEventsResponseBodyEvents) *DescribeHistoryEventsResponseBody
	GetEvents() []*DescribeHistoryEventsResponseBodyEvents
	SetPageNumber(v int32) *DescribeHistoryEventsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHistoryEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHistoryEventsResponseBody
	GetTotalCount() *int32
}

type DescribeHistoryEventsResponseBody struct {
	// The list of events.
	Events []*DescribeHistoryEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5359599C-F656-57BD-8A0D-329A2FD511A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsResponseBody) GetEvents() []*DescribeHistoryEventsResponseBodyEvents {
	return s.Events
}

func (s *DescribeHistoryEventsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHistoryEventsResponseBody) SetEvents(v []*DescribeHistoryEventsResponseBodyEvents) *DescribeHistoryEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeHistoryEventsResponseBody) SetPageNumber(v int32) *DescribeHistoryEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryEventsResponseBody) SetPageSize(v int32) *DescribeHistoryEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryEventsResponseBody) SetRequestId(v string) *DescribeHistoryEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryEventsResponseBody) SetTotalCount(v int32) *DescribeHistoryEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHistoryEventsResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHistoryEventsResponseBodyEvents struct {
	// The event ID.
	//
	// example:
	//
	// e-d71ff150945b9c02eb6ebc0016328468
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The event level. Valid values:
	//
	// - CRITICAL: Critical.
	//
	// - WARN: Warning.
	//
	// - INFO: Information.
	//
	// example:
	//
	// WARN
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The event status. Valid values:
	//
	// - Inquiring: Inquiring.
	//
	// - Scheduled: Scheduled.
	//
	// - Executing: Executing.
	//
	// - Executed: Executed.
	//
	// - Failed: Failed.
	//
	// - Canceled: Canceled.
	//
	// - Avoided: Avoided.
	//
	// example:
	//
	// Inquiring
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The event type. Valid values:
	//
	// - Instance:SystemFailure.Redeploy: The instance deployment due to a system fault.
	//
	// - Instance:SystemFailure.Reboot: The instance is restarted due to a system fault.
	//
	// - Instance:RegionNetworkDown: The edge zone network break.
	//
	// - Disk:Stalled: The disk performance is degraded.
	//
	// - EnsRegion:NetworkMigration: Instance migration due to a system fault.
	//
	// - IP:SafeRisk: IP alerting.
	//
	// - IP:SafeBan: The IP address is blocked.
	//
	// - Instance:SystemUpgrade.Migrate: The instance needs to be migrated due to an underlying upgrade.
	//
	// - Instance:SystemMaintenance.Redeploy: The instance deployment due to system maintenance.
	//
	// example:
	//
	// Instance:SystemFailure.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The extended attribute.
	//
	// example:
	//
	// {}
	ExtendedAttribute *string `json:"ExtendedAttribute,omitempty" xml:"ExtendedAttribute,omitempty"`
	// The scheduled execution time of the event, in milliseconds.
	//
	// example:
	//
	// 1715578245000
	NotBefore        *int64  `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	NotBeforeUtcTime *string `json:"NotBeforeUtcTime,omitempty" xml:"NotBeforeUtcTime,omitempty"`
	// The publish time, in milliseconds.
	//
	// example:
	//
	// 1715578245000
	PublishTime    *int64  `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	PublishUtcTime *string `json:"PublishUtcTime,omitempty" xml:"PublishUtcTime,omitempty"`
	// The reason for the event.
	//
	// example:
	//
	// The current system policy quota is insufficient
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the associated resource.
	//
	// example:
	//
	// i-55qi8m11rr53c4i964md8a00l
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeHistoryEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetEventId() *string {
	return s.EventId
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetEventType() *string {
	return s.EventType
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetExtendedAttribute() *string {
	return s.ExtendedAttribute
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetNotBeforeUtcTime() *string {
	return s.NotBeforeUtcTime
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetPublishTime() *int64 {
	return s.PublishTime
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetPublishUtcTime() *string {
	return s.PublishUtcTime
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetReason() *string {
	return s.Reason
}

func (s *DescribeHistoryEventsResponseBodyEvents) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetEventId(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetEventLevel(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.EventLevel = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetEventStatus(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.EventStatus = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetEventType(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.EventType = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetExtendedAttribute(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.ExtendedAttribute = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetNotBefore(v int64) *DescribeHistoryEventsResponseBodyEvents {
	s.NotBefore = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetNotBeforeUtcTime(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.NotBeforeUtcTime = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetPublishTime(v int64) *DescribeHistoryEventsResponseBodyEvents {
	s.PublishTime = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetPublishUtcTime(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.PublishUtcTime = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetReason(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.Reason = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) SetResourceId(v string) *DescribeHistoryEventsResponseBodyEvents {
	s.ResourceId = &v
	return s
}

func (s *DescribeHistoryEventsResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
