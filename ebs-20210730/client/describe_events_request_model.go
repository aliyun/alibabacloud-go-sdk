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
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-01T04:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The severity level of the event. Valid values:
	//
	// 	- **INFO**
	//
	// 	- **WARN**
	//
	// 	- **CRITICAL**
	//
	// example:
	//
	// WARN
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The name of the event. Valid values:
	//
	// 	- NoSnapshot: indicates the event that is triggered because no snapshot is created for a disk to protect data on the disk.
	//
	// 	- BurstIOTriggered: indicates the event that is triggered when a burst I/O operation is performed on a disk.
	//
	// 	- CostOptimizationNeeded: indicates the event that is triggered when cost optimization is required.
	//
	// 	- DiskSpecNotMatchedWithInstance: indicates the event that is triggered because the specifications of a disk do not match the instance to which the disk is attached.
	//
	// 	- DiskIONo4kAligned: indicates the event that is triggered because the physical and logical sectors involved in a read or write operation are not 4K aligned.
	//
	// 	- DiskIOHang: indicates the event that is triggered when an I/O hang occurs on a disk.
	//
	// 	- InstanceIOPSExceedInstanceMaxLimit: indicates the event that is triggered when the number of IOPS on an instance reaches the upper limit.
	//
	// 	- InstanceBPSExceedInstanceMaxLimit: indicates the event that is triggered when the number of BPS on an instance reaches the upper limit.
	//
	// 	- DiskIOPSExceedInstanceMaxLimit: indicates the event that is triggered when the number of IOPS on a disk reaches the upper limit for the associated instance.
	//
	// 	- DiskBPSExceedInstanceMaxLimit: indicates the event that is triggered when the number of BPS on a disk reaches the upper limit for the associated instance.
	//
	// 	- DiskIOPSExceedDiskMaxLimit: indicates the event that is triggered when the number of IOPS on a disk reaches the upper limit for the disk.
	//
	// 	- DiskBPSExceedDiskMaxLimit: indicates the event that is triggered when the number of BPS on a disk reaches the upper limit for the disk.
	//
	// example:
	//
	// DiskIOHang
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The number of entries to return on each page. If you specify MaxResults, `MaxResults` and `NextToken` are used for a paged query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID . You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions supported.
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
	// The type of resource. Valid values:
	//
	// 	- disk.
	//
	// Default value: disk.
	//
	// example:
	//
	// disk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-01T03:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of event. Valid values:
	//
	// - WillExecute
	//
	// - Executing
	//
	// - Executed
	//
	// - Ignore
	//
	// - Expired
	//
	// - Deleted
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
