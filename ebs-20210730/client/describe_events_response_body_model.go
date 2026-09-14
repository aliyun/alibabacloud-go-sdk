// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeEventsResponseBody
	GetRequestId() *string
	SetResourceEvents(v []*DescribeEventsResponseBodyResourceEvents) *DescribeEventsResponseBody
	GetResourceEvents() []*DescribeEventsResponseBodyResourceEvents
	SetTotalCount(v int32) *DescribeEventsResponseBody
	GetTotalCount() *int32
}

type DescribeEventsResponseBody struct {
	// The token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of events.
	ResourceEvents []*DescribeEventsResponseBodyResourceEvents `json:"ResourceEvents,omitempty" xml:"ResourceEvents,omitempty" type:"Repeated"`
	// The total number of entries returned for the paged query.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventsResponseBody) GetResourceEvents() []*DescribeEventsResponseBodyResourceEvents {
	return s.ResourceEvents
}

func (s *DescribeEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEventsResponseBody) SetNextToken(v string) *DescribeEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetResourceEvents(v []*DescribeEventsResponseBodyResourceEvents) *DescribeEventsResponseBody {
	s.ResourceEvents = v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalCount(v int32) *DescribeEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsResponseBody) Validate() error {
	if s.ResourceEvents != nil {
		for _, item := range s.ResourceEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventsResponseBodyResourceEvents struct {
	// The event description.
	//
	// example:
	//
	// You can purchase 4296 provisioned IOPS for cost optimization. Based on your usage over the past 7 days, costs are estimated to decrease by 16%
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time of the event. This value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1679538083000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event level. Valid values:
	//
	// 1. INFO
	//
	// 2. WARN
	//
	// 3. CRITICAL
	//
	// example:
	//
	// INFO
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The event name. Valid values:
	//
	// - NoSnapshot: data protection
	//
	// - BurstIOTriggered: burst I/O
	//
	// - CostOptimizationNeeded: cost optimization
	//
	// - DiskSpecNotMatchedWithInstance: instance and disk specification mismatch
	//
	// - DiskIONo4kAligned: non-4K aligned read/write
	//
	// - DiskIOHang: disk IOHang occurred
	//
	// - InstanceIOPSExceedInstanceMaxLimit: instance IOPS reached the upper limit
	//
	// - InstanceBPSExceedInstanceMaxLimit: instance BPS reached the upper limit
	//
	// - DiskIOPSExceedInstanceMaxLimit: disk IOPS reached the instance upper limit
	//
	// - DiskBPSExceedInstanceMaxLimit: disk BPS reached the instance upper limit
	//
	// - DiskIOPSExceedDiskMaxLimit: disk IOPS reached the disk upper limit
	//
	// - DiskBPSExceedDiskMaxLimit: disk BPS reached the disk upper limit
	//
	// example:
	//
	// DiskIOHang
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The event type. Valid values:
	//
	// 1. Notification
	//
	// 2. SystemException
	//
	// 3. Alert
	//
	// example:
	//
	// Alert
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The additional properties. Possible fields:
	//
	// - EcsInstanceId: the ID of the ECS instance to which the cloud disk is attached.
	//
	// - Adapter: the mount point of the cloud disk.
	//
	// example:
	//
	// {\\"EcsInstanceId\\":\\"i-uf6dkn9qpcw6y94g7ag7\\",\\"Adapter\\":\\"hda\\"}
	ExtraAttributes *string `json:"ExtraAttributes,omitempty" xml:"ExtraAttributes,omitempty"`
	// The recommended action after the event occurs. Valid values:
	//
	// - ModifyDiskSpec: change disk specifications
	//
	// - CreateSnapshot: create a snapshot
	//
	// - ResizeDisk: expand disk capacity
	//
	// - AdjustProvision: adjust provisioned performance
	//
	// - ModifyInstanceSpec: change instance specifications
	//
	// example:
	//
	// AdjustProvision
	RecommendAction *string `json:"RecommendAction,omitempty" xml:"RecommendAction,omitempty"`
	// The parameters for the recommended action after the event occurs.
	//
	// example:
	//
	// 4296
	RecommendParams *string `json:"RecommendParams,omitempty" xml:"RecommendParams,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// disk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time of the event. This value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1684204822000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event status. Valid values:
	//
	// 1. WillExecute: pending
	//
	// 2. Executing: processing
	//
	// 3. Executed: processed
	//
	// 4. Ignore: ignored
	//
	// 5. Expired: expired
	//
	// 6. Deleted: deleted
	//
	// example:
	//
	// WillExecute
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventsResponseBodyResourceEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyResourceEvents) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyResourceEvents) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventsResponseBodyResourceEvents) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEventsResponseBodyResourceEvents) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeEventsResponseBodyResourceEvents) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventsResponseBodyResourceEvents) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventsResponseBodyResourceEvents) GetExtraAttributes() *string {
	return s.ExtraAttributes
}

func (s *DescribeEventsResponseBodyResourceEvents) GetRecommendAction() *string {
	return s.RecommendAction
}

func (s *DescribeEventsResponseBodyResourceEvents) GetRecommendParams() *string {
	return s.RecommendParams
}

func (s *DescribeEventsResponseBodyResourceEvents) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeEventsResponseBodyResourceEvents) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeEventsResponseBodyResourceEvents) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEventsResponseBodyResourceEvents) GetStatus() *string {
	return s.Status
}

func (s *DescribeEventsResponseBodyResourceEvents) SetDescription(v string) *DescribeEventsResponseBodyResourceEvents {
	s.Description = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetEndTime(v string) *DescribeEventsResponseBodyResourceEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetEventLevel(v string) *DescribeEventsResponseBodyResourceEvents {
	s.EventLevel = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetEventName(v string) *DescribeEventsResponseBodyResourceEvents {
	s.EventName = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetEventType(v string) *DescribeEventsResponseBodyResourceEvents {
	s.EventType = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetExtraAttributes(v string) *DescribeEventsResponseBodyResourceEvents {
	s.ExtraAttributes = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetRecommendAction(v string) *DescribeEventsResponseBodyResourceEvents {
	s.RecommendAction = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetRecommendParams(v string) *DescribeEventsResponseBodyResourceEvents {
	s.RecommendParams = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetResourceId(v string) *DescribeEventsResponseBodyResourceEvents {
	s.ResourceId = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetResourceType(v string) *DescribeEventsResponseBodyResourceEvents {
	s.ResourceType = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetStartTime(v string) *DescribeEventsResponseBodyResourceEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) SetStatus(v string) *DescribeEventsResponseBodyResourceEvents {
	s.Status = &v
	return s
}

func (s *DescribeEventsResponseBodyResourceEvents) Validate() error {
	return dara.Validate(s)
}
