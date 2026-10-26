// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromExecTime(v int32) *DescribeHistoryTasksRequest
	GetFromExecTime() *int32
	SetFromStartTime(v string) *DescribeHistoryTasksRequest
	GetFromStartTime() *string
	SetInstanceId(v string) *DescribeHistoryTasksRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeHistoryTasksRequest
	GetInstanceType() *string
	SetOwnerId(v int64) *DescribeHistoryTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeHistoryTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHistoryTasksRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeHistoryTasksRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeHistoryTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHistoryTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeHistoryTasksRequest
	GetSecurityToken() *string
	SetStatus(v string) *DescribeHistoryTasksRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeHistoryTasksRequest
	GetTaskId() *string
	SetTaskType(v string) *DescribeHistoryTasksRequest
	GetTaskType() *string
	SetToExecTime(v int32) *DescribeHistoryTasksRequest
	GetToExecTime() *int32
	SetToStartTime(v string) *DescribeHistoryTasksRequest
	GetToStartTime() *string
}

type DescribeHistoryTasksRequest struct {
	// The minimum execution duration of the task. This parameter is used to filter tasks whose execution duration is greater than the specified value. Unit: seconds. Default value: 0, which indicates no limit.
	//
	// example:
	//
	// 0
	FromExecTime *int32 `json:"FromExecTime,omitempty" xml:"FromExecTime,omitempty"`
	// The start time of the time range to query the task start time. This parameter specifies that the start time of the task is later than the specified time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The earliest supported time is 30 days ago. If the specified time is more than 30 days ago, it is automatically converted to 30 days ago.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-02T11:31:03Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The instance ID, which is used to query the tasks of the corresponding instance. Separate multiple instance IDs with commas (,). You can specify up to 30 instance IDs. If this parameter is left empty, no limit is applied.
	//
	// example:
	//
	// rm-uf62br2491p5l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type. Only Instance is supported.
	//
	// example:
	//
	// Instance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the query result. Valid values: positive integers.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **10 to 100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the pending event. You can call the DescribeRegions operation to obtain the region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-ae****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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
	// Separate multiple values with commas (,). If this parameter is left empty, all statuses are queried.
	//
	// example:
	//
	// Scheduled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID. You can call the DescribeTasks operation to query the task ID. Separate multiple task IDs with commas (,). You can specify up to 30 task IDs. If this parameter is left empty, no limit is applied.
	//
	// example:
	//
	// t-83br18hloy3faf****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type, which is used to query tasks of a specific type. Separate multiple task types with commas (,). You can specify up to 30 task types. If this parameter is left empty, no limit is applied.
	//
	// example:
	//
	// autotest_dispatch_cases
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The maximum execution duration of the task. This parameter is used to filter tasks whose execution duration is greater than or equal to the specified value. Unit: seconds. Default value: 0, which indicates no limit.
	//
	// example:
	//
	// 0
	ToExecTime *int32 `json:"ToExecTime,omitempty" xml:"ToExecTime,omitempty"`
	// The end time of the time range to query the task start time. This parameter specifies that the start time of the task is earlier than the specified time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-03-02T11:31:03Z
	ToStartTime *string `json:"ToStartTime,omitempty" xml:"ToStartTime,omitempty"`
}

func (s DescribeHistoryTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksRequest) GetFromExecTime() *int32 {
	return s.FromExecTime
}

func (s *DescribeHistoryTasksRequest) GetFromStartTime() *string {
	return s.FromStartTime
}

func (s *DescribeHistoryTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHistoryTasksRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeHistoryTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHistoryTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHistoryTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHistoryTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHistoryTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeHistoryTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeHistoryTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHistoryTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeHistoryTasksRequest) GetToExecTime() *int32 {
	return s.ToExecTime
}

func (s *DescribeHistoryTasksRequest) GetToStartTime() *string {
	return s.ToStartTime
}

func (s *DescribeHistoryTasksRequest) SetFromExecTime(v int32) *DescribeHistoryTasksRequest {
	s.FromExecTime = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetFromStartTime(v string) *DescribeHistoryTasksRequest {
	s.FromStartTime = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetInstanceId(v string) *DescribeHistoryTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetInstanceType(v string) *DescribeHistoryTasksRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetOwnerId(v int64) *DescribeHistoryTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetPageNumber(v int32) *DescribeHistoryTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetPageSize(v int32) *DescribeHistoryTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetRegionId(v string) *DescribeHistoryTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetResourceGroupId(v string) *DescribeHistoryTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetResourceOwnerAccount(v string) *DescribeHistoryTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetResourceOwnerId(v int64) *DescribeHistoryTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetSecurityToken(v string) *DescribeHistoryTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetStatus(v string) *DescribeHistoryTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetTaskId(v string) *DescribeHistoryTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetTaskType(v string) *DescribeHistoryTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetToExecTime(v int32) *DescribeHistoryTasksRequest {
	s.ToExecTime = &v
	return s
}

func (s *DescribeHistoryTasksRequest) SetToStartTime(v string) *DescribeHistoryTasksRequest {
	s.ToStartTime = &v
	return s
}

func (s *DescribeHistoryTasksRequest) Validate() error {
	return dara.Validate(s)
}
