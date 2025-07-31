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
	// The minimum execution duration of the task. This parameter is used to filter tasks whose execution duration is longer than the minimum execution duration. Unit: seconds. The default value is 0, which indicates that no limit is imposed for the query.
	//
	// example:
	//
	// 0
	FromExecTime *int32 `json:"FromExecTime,omitempty" xml:"FromExecTime,omitempty"`
	// The start time of the O\\&M task to perform. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. You can query data within the last 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-15T03:41:26Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The instance ID. Separate multiple instance IDs with commas (,). You can specify up to 30 instance IDs. This parameter is empty by default, which indicates that tasks of all instances are queried.
	//
	// example:
	//
	// dds-8vb38f0e7933xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type of the instance. Set the value to Instance.
	//
	// example:
	//
	// Instance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of the page to return. The value must be a positive integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the pending event. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2inrfrnw3xby
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task status. Valid values:
	//
	// 	- Scheduled: The task is waiting to be executed.
	//
	// 	- Running: The task is running.
	//
	// 	- Succeed: The task is successful.
	//
	// 	- Failed: The task failed.
	//
	// 	- Cancelling: The task is being terminated.
	//
	// 	- Canceled: The task has been terminated.
	//
	// 	- Waiting: The task is waiting for scheduled time.
	//
	// Separate multiple states with commas (,). This parameter is empty by default, which indicates that tasks in all states are queried.
	//
	// example:
	//
	// Succeed,Running,Waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID. Separate multiple task IDs with commas (,). You can specify up to 30 task IDs. This parameter is empty by default, which indicates that all tasks are queried.
	//
	// example:
	//
	// t-0mq1yyhm3ffl2bxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. This parameter is left empty by default, which indicates that all types of tasks are queried. Valid values:
	//
	// 	- CreateIns: Create an instance.
	//
	// 	- DeleteIns: Delete an instance.
	//
	// 	- ChangeVariable: Modify parameter settings for an instance.
	//
	// 	- ModifyInsConfig: Change the configurations of an instance.
	//
	// 	- RestartIns: Restart an instance.
	//
	// 	- HaSwitch: Perform a primary/secondary switchover on an instance.
	//
	// 	- CloneIns: Clone an instance.
	//
	// 	- KernelVersionUpgrade: Update the minor version of an instance.
	//
	// 	- ProxyVersionUpgrade: Upgrade the agent version of an instance.
	//
	// 	- ModifyAccount: Change the account of an instance.
	//
	// 	- ModifyInsSpec: Change the specifications of an instance or perform a data migration on the instance.
	//
	// 	- CreateReadIns: Create a read-only instance.
	//
	// 	- StartIns: Start an instance.
	//
	// 	- StopIns: Stop an instance.
	//
	// 	- ModifyNetwork: Modify the network type for an instance.
	//
	// 	- LockIns: Lock an instance.
	//
	// 	- UnlockIns: Unlock an instance.
	//
	// 	- DiskOnlineExpansion: Scale out the disks of an instance online.
	//
	// 	- StorageOnlineExpansion: Expend the storage capacity of an instance online.
	//
	// 	- AddInsNode: Add a node to an instance.
	//
	// 	- DeleteInsNode: Delete a node from an instance.
	//
	// 	- ManualBackupIns: Manually back up an instance.
	//
	// 	- ModifyInsStorageType: Modify the storage type for an instance.
	//
	// example:
	//
	// DeleteInsNode
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The maximum execution duration of the task. This parameter is used to filter tasks whose execution duration is shorter than or equal to the maximum execution duration. Unit: seconds. The default value is 0, which indicates that no limit is imposed for the query.
	//
	// example:
	//
	// 0
	ToExecTime *int32 `json:"ToExecTime,omitempty" xml:"ToExecTime,omitempty"`
	// The end time of the O\\&M task to perform. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. You can query data within the last 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-16T07:21:31Z
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
