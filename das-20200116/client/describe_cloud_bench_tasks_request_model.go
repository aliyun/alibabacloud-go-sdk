// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudBenchTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCloudBenchTasksRequest
	GetEndTime() *string
	SetPageNo(v string) *DescribeCloudBenchTasksRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeCloudBenchTasksRequest
	GetPageSize() *string
	SetStartTime(v string) *DescribeCloudBenchTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeCloudBenchTasksRequest
	GetStatus() *string
	SetTaskType(v string) *DescribeCloudBenchTasksRequest
	GetTaskType() *string
}

type DescribeCloudBenchTasksRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 1596177993001
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. The value must be a positive integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be a positive integer. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1596177993000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the stress testing task. Valid values:
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **IGNORED**: The task is ignored.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **EXCEPTION**: The task is abnormal.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the stress testing task. Valid values:
	//
	// 	- **pressure test*	- (default): A task of this type replays the traffic that is captured from the source instance on the destination instance at the maximum playback rate that is supported by the destination instance.
	//
	// 	- **smart pressure test**: A task of this type analyzes the traffic that is captured from the source instance over a short period of time and generates traffic on the destination instance for continuous stress testing. The business model based on which the traffic is generated on the destination instance and the traffic distribution are consistent with those on the source instance. Stress testing tasks of this type can help you reduce the amount of time that is consumed to collect data from the source instance and reduce storage costs and performance overheads.
	//
	// example:
	//
	// pressure test
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeCloudBenchTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCloudBenchTasksRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeCloudBenchTasksRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCloudBenchTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCloudBenchTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudBenchTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeCloudBenchTasksRequest) SetEndTime(v string) *DescribeCloudBenchTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetPageNo(v string) *DescribeCloudBenchTasksRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetPageSize(v string) *DescribeCloudBenchTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetStartTime(v string) *DescribeCloudBenchTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetStatus(v string) *DescribeCloudBenchTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) SetTaskType(v string) *DescribeCloudBenchTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeCloudBenchTasksRequest) Validate() error {
	return dara.Validate(s)
}
