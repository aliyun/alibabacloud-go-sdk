// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromExecTime(v int32) *DescribeHistoryTasksStatRequest
	GetFromExecTime() *int32
	SetFromStartTime(v string) *DescribeHistoryTasksStatRequest
	GetFromStartTime() *string
	SetInstanceId(v string) *DescribeHistoryTasksStatRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeHistoryTasksStatRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeHistoryTasksStatRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *DescribeHistoryTasksStatRequest
	GetSecurityToken() *string
	SetStatus(v string) *DescribeHistoryTasksStatRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeHistoryTasksStatRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeHistoryTasksStatRequest
	GetTaskType() *string
	SetToExecTime(v int32) *DescribeHistoryTasksStatRequest
	GetToExecTime() *int32
	SetToStartTime(v string) *DescribeHistoryTasksStatRequest
	GetToStartTime() *string
}

type DescribeHistoryTasksStatRequest struct {
	// The minimum execution duration of the task, in seconds. This parameter is used to filter tasks whose execution duration is greater than the specified value. Default value: 0, which indicates no limit.
	//
	// example:
	//
	// 1
	FromExecTime *int32 `json:"FromExecTime,omitempty" xml:"FromExecTime,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-08T07:04:17Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// abc
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task status. Valid values:
	//
	// - **Scheduled**: Pending execution.
	//
	// - **Running**: Running.
	//
	// - **Succeed**: Succeeded.
	//
	// - **Failed**: Failed.
	//
	// - **Cancelling**: Cancelling.
	//
	// - **Canceled**: Canceled.
	//
	// - **Waiting**: Waiting for the scheduled time.
	//
	// Separate multiple values with commas (,). By default, this parameter is left empty, which indicates that all statuses are queried.
	//
	// example:
	//
	// Scheduled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1564657730
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
	//
	// example:
	//
	// all
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The maximum execution duration of the task, in seconds. This parameter is used to filter tasks whose execution duration is less than or equal to the specified value. Default value: 0, which indicates no limit.
	//
	// example:
	//
	// 10
	ToExecTime *int32 `json:"ToExecTime,omitempty" xml:"ToExecTime,omitempty"`
	// The end time of the query. Only tasks whose start time is earlier than the specified end time are returned. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-24T10:01:37Z
	ToStartTime *string `json:"ToStartTime,omitempty" xml:"ToStartTime,omitempty"`
}

func (s DescribeHistoryTasksStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksStatRequest) GetFromExecTime() *int32 {
	return s.FromExecTime
}

func (s *DescribeHistoryTasksStatRequest) GetFromStartTime() *string {
	return s.FromStartTime
}

func (s *DescribeHistoryTasksStatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHistoryTasksStatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryTasksStatRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHistoryTasksStatRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeHistoryTasksStatRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeHistoryTasksStatRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHistoryTasksStatRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeHistoryTasksStatRequest) GetToExecTime() *int32 {
	return s.ToExecTime
}

func (s *DescribeHistoryTasksStatRequest) GetToStartTime() *string {
	return s.ToStartTime
}

func (s *DescribeHistoryTasksStatRequest) SetFromExecTime(v int32) *DescribeHistoryTasksStatRequest {
	s.FromExecTime = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetFromStartTime(v string) *DescribeHistoryTasksStatRequest {
	s.FromStartTime = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetInstanceId(v string) *DescribeHistoryTasksStatRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetRegionId(v string) *DescribeHistoryTasksStatRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetResourceGroupId(v string) *DescribeHistoryTasksStatRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetSecurityToken(v string) *DescribeHistoryTasksStatRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetStatus(v string) *DescribeHistoryTasksStatRequest {
	s.Status = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetTaskId(v string) *DescribeHistoryTasksStatRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetTaskType(v string) *DescribeHistoryTasksStatRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetToExecTime(v int32) *DescribeHistoryTasksStatRequest {
	s.ToExecTime = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetToStartTime(v string) *DescribeHistoryTasksStatRequest {
	s.ToStartTime = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) Validate() error {
	return dara.Validate(s)
}
