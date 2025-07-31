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
	SetResourceOwnerAccount(v string) *DescribeHistoryTasksStatRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHistoryTasksStatRequest
	GetResourceOwnerId() *int64
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
	// 2023-04-16T02:46:21Z
	FromStartTime *string `json:"FromStartTime,omitempty" xml:"FromStartTime,omitempty"`
	// The instance ID. Separate multiple instance IDs with commas (,). You can specify up to 30 instance IDs. This parameter is empty by default, which indicates that the tasks of all instances are queried.
	//
	// example:
	//
	// dds-8vb38f0e7933xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the pending event. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
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
	// Succeed
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
	// DeleteIns
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
	// 2023-05-16T02:15:52Z
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

func (s *DescribeHistoryTasksStatRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHistoryTasksStatRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *DescribeHistoryTasksStatRequest) SetResourceOwnerAccount(v string) *DescribeHistoryTasksStatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHistoryTasksStatRequest) SetResourceOwnerId(v int64) *DescribeHistoryTasksStatRequest {
	s.ResourceOwnerId = &v
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
