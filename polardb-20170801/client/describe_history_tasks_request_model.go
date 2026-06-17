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
	SetResourceOwnerAccount(v int64) *DescribeHistoryTasksRequest
	GetResourceOwnerAccount() *int64
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
	// The minimum task execution time in seconds. Filters for tasks that took longer than this value. Default value: 0.
	//
	// example:
	//
	// 0
	FromExecTime *int32 `json:"FromExecTime,omitempty" xml:"FromExecTime,omitempty"`
	// The start of the time range to query, based on task start time. The time follows the ISO8601 standard and must be in `UTC+0` time. Format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// The earliest supported time is 30 days ago. If the specified time is more than 30 days ago, it will be automatically converted to 30 days ago.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-01-02T11:31:03Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The resource ID to filter by. You can provide a comma-separated list of up to 30 IDs. Default value: empty, indicating no restriction.
	//
	// > Currently, only PolarDB cluster IDs are supported.
	//
	// example:
	//
	// pc-2zed3m89cw***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Currently, only Instance is supported.
	//
	// example:
	//
	// Instance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: positive integers. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > For more information, see [DescribeRegions](https://help.aliyun.com/document_detail/98041.html).
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *int64  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task status. Valid values:
	//
	// - **Scheduled**: waiting for execution
	//
	// - **Running**: executing
	//
	// - **Succeed**: executed successfully
	//
	// - **Cancelling**: stopping
	//
	// - **Canceled**: stopped
	//
	// - **Waiting**: waiting for preset time
	//
	// You can provide a comma-separated list. Default value: empty, which indicates all statuses.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID. You can provide a comma-separated list of up to 30 IDs. Default value: empty, indicating no restriction.
	//
	// example:
	//
	// t-0mqi38ho0cgjv***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. You can provide a comma-separated list of up to 30 task types. Default value: empty, indicating no restriction.
	//
	// example:
	//
	// ChangeVariable
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The maximum task execution time in seconds. Filters for tasks that took less than this value. Default value: 0.
	//
	// example:
	//
	// 0
	ToExecTime *int32 `json:"ToExecTime,omitempty" xml:"ToExecTime,omitempty"`
	// The end of the time range to query, based on task start time. The time follows the ISO8601 standard and must be in `UTC+0` time. Format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-01-03T11:31:03Z
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

func (s *DescribeHistoryTasksRequest) GetResourceOwnerAccount() *int64 {
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

func (s *DescribeHistoryTasksRequest) SetResourceOwnerAccount(v int64) *DescribeHistoryTasksRequest {
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
