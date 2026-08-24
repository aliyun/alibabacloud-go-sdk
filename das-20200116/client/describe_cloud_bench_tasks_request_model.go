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
	// The end time of the query task. Specify the value as a UNIX timestamp. Unit: milliseconds.
	//
	// >The end time of the query task must be later than the start time.
	//
	// example:
	//
	// 1596177993001
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. The value must be greater than 0 and cannot exceed the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of records per page. The value must be greater than 0 and cannot exceed the maximum value of the Integer data type. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query task. Specify the value as a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1596177993000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The running status of the task. Valid values:
	//
	// - **SUCCESS**: Successful.
	//
	// - **IGNORED**: Ignored.
	//
	// - **RUNNING**: Running.
	//
	// - **EXCEPTION**: Exception.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the stress testing task. Valid values:
	//
	// - **pressure test*	- (default): intelligent stress testing. Traffic captured from the target instance is replayed on the destination instance at the maximum speed supported by the destination instance specifications.
	//
	// - **smart pressure test**: generated stress testing. By analyzing and learning from traffic captured from the target instance within a short period, traffic that is consistent with the business model and traffic distribution of the original traffic is generated for continuous stress testing. This reduces the time required to collect data from the target instance and lowers storage costs and performance overhead.
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
