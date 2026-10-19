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
	SetResourceOwnerId(v int64) *DescribeHistoryTasksStatRequest
	GetResourceOwnerId() *int64
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
	// The minimum execution duration used to filter tasks whose execution duration is greater than this value. Unit: seconds. Default value: 0, which indicates no restriction.
	//
	// example:
	//
	// 0
	FromExecTime *int32 `json:"FromExecTime,omitempty" xml:"FromExecTime,omitempty"`
	// The beginning of the query time range. Specify the time in the <i>yyyy-mm-dd</i>t<i>hh:mm</i>z format (UTC).
	//
	// This parameter is required.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the pending event. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task status. This parameter is used to filter tasks by status. Valid values:
	//
	// - **Scheduled**: Waiting to be executed.
	//
	// - **Running**: Running.
	//
	// - **Succeed**: Succeeded.
	//
	// - **Failed**: Failed.
	//
	// - **Cancelling**: Being stopped.
	//
	// - **Canceled**: Stopped.
	//
	// - **Waiting**: Waiting for the scheduled time.
	//
	// > By default, this parameter is empty, which indicates that tasks in all statuses are queried. To specify multiple statuses, separate them with commas (,).
	//
	// example:
	//
	// Scheduled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID. This parameter is used to query a task with a known ID. By default, this parameter is empty, which indicates no restriction. To specify multiple task IDs, separate them with commas (,). A maximum of 30 task IDs are supported.
	//
	// example:
	//
	// t-0mq1yyhm3ffl2bxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
	//
	// example:
	//
	// all
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The maximum execution duration used to filter tasks whose execution duration is not less than this value. Unit: seconds. Default value: 0, which indicates no restriction.
	//
	// example:
	//
	// 0
	ToExecTime *int32 `json:"ToExecTime,omitempty" xml:"ToExecTime,omitempty"`
	// The end of the query time range. Specify the time in the <i>yyyy-mm-dd</i>t<i>hh:mm</i>z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-02T11:31:03Z
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

func (s *DescribeHistoryTasksStatRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *DescribeHistoryTasksStatRequest) SetResourceOwnerId(v int64) *DescribeHistoryTasksStatRequest {
	s.ResourceOwnerId = &v
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
