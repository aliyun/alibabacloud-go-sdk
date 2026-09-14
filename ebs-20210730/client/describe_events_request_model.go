// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeEventsRequest
	GetEndTime() *string
	SetEventLevel(v string) *DescribeEventsRequest
	GetEventLevel() *string
	SetEventName(v string) *DescribeEventsRequest
	GetEventName() *string
	SetMaxResults(v int32) *DescribeEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeEventsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeEventsRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeEventsRequest
	GetResourceId() *string
	SetResourceType(v string) *DescribeEventsRequest
	GetResourceType() *string
	SetStartTime(v string) *DescribeEventsRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeEventsRequest
	GetStatus() *string
}

type DescribeEventsRequest struct {
	// The end time of the event. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-01T04:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event level. Valid values:
	//
	// - **INFO**: Notification.
	//
	// - **WARN**: Warning.
	//
	// - **CRITICAL**: Critical.
	//
	// example:
	//
	// WARN
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
	// The maximum number of entries per page for a paged query. If you specify this parameter, the `MaxResults` and `NextToken` parameters are used together for the query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call DescribeRegions to query the list of regions supported by EBS Lens.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - disk: cloud disk
	//
	// example:
	//
	// disk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time of the event. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-01T03:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event status. Valid values:
	//
	// - WillExecute: pending
	//
	// - Executing: processing
	//
	// - Executed: processed
	//
	// - Ignore: ignored
	//
	// - Expired: expired
	//
	// - Deleted: deleted
	//
	// example:
	//
	// WillExecute
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEventsRequest) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeEventsRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeEventsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEventsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeEventsRequest) SetEndTime(v string) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetEventLevel(v string) *DescribeEventsRequest {
	s.EventLevel = &v
	return s
}

func (s *DescribeEventsRequest) SetEventName(v string) *DescribeEventsRequest {
	s.EventName = &v
	return s
}

func (s *DescribeEventsRequest) SetMaxResults(v int32) *DescribeEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeEventsRequest) SetNextToken(v string) *DescribeEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsRequest) SetRegionId(v string) *DescribeEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsRequest) SetResourceId(v string) *DescribeEventsRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeEventsRequest) SetResourceType(v string) *DescribeEventsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v string) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsRequest) SetStatus(v string) *DescribeEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeEventsRequest) Validate() error {
	return dara.Validate(s)
}
