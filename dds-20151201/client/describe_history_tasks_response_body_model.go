// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeHistoryTasksResponseBodyItems) *DescribeHistoryTasksResponseBody
	GetItems() []*DescribeHistoryTasksResponseBodyItems
	SetPageNumber(v int32) *DescribeHistoryTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHistoryTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHistoryTasksResponseBody
	GetTotalCount() *int32
}

type DescribeHistoryTasksResponseBody struct {
	// The task objects.
	Items []*DescribeHistoryTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C4A2494-XXXX-XXXX-93CF-548DB3375193
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of tasks that meet these conditions without taking pagination into account.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksResponseBody) GetItems() []*DescribeHistoryTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeHistoryTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHistoryTasksResponseBody) SetItems(v []*DescribeHistoryTasksResponseBodyItems) *DescribeHistoryTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetPageNumber(v int32) *DescribeHistoryTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetPageSize(v int32) *DescribeHistoryTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetRequestId(v string) *DescribeHistoryTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetTotalCount(v int32) *DescribeHistoryTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHistoryTasksResponseBodyItems struct {
	// A set of allowed actions that can be taken on the task. The system matches the current step name and status of the task to the available actions specified by ActionInfo. If no matching action is found, the current status of the task does not support any action. Example:
	//
	//        "steps": [
	//
	//         {
	//
	//           "step_name": "exec_task", // The name of the step, which matches the value of CurrentStepName.
	//
	//           "action_info": {    // The actions supported for this step.
	//
	//             "Waiting": [      // The status, which matches the value of Status.
	//
	//               "modifySwitchTime" // The action. Multiple actions are supported.
	//
	//             ]
	//
	//           }
	//
	//         },
	//
	//         {
	//
	//           "step_name": "init_task", // The name of the step.
	//
	//           "action_info": {    // The actions supported for this step.
	//
	//             "Running": [      // The status.
	//
	//               "cancel",       // The action.
	//
	//               "pause"
	//
	//             ]
	//
	//           }
	//
	//         }
	//
	//       ]
	//
	//     }
	//
	// The system may support the following actions:
	//
	// 	- retry: makes another attempt.
	//
	// 	- cancel: makes a cancellation.
	//
	// 	- modifySwitchTime: changes the switching or restoration time.
	//
	// example:
	//
	// `{\\"steps\\":[{\\"action_info\\":{\\"Waiting\\":[\\"modifySwitchTime\\"]},\\"step_name\\":\\"exec_task\\"}]}`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The request source. Valid values: System and User.
	//
	// example:
	//
	// System
	CallerSource *string `json:"CallerSource,omitempty" xml:"CallerSource,omitempty"`
	// The ID of the user who made the request. If CallerSource is set to User, CallerUid indicates the unique ID (UID) of the user.
	//
	// example:
	//
	// 141345906006****
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The name of the current step. If this parameter is left empty, the task is not started.
	//
	// example:
	//
	// exec_task
	CurrentStepName *string `json:"CurrentStepName,omitempty" xml:"CurrentStepName,omitempty"`
	// The database type. The value is fixed to mongodb.
	//
	// example:
	//
	// mongodb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The end time of the performed O\\&M task. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-16T02:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID
	//
	// example:
	//
	// dds-t4n18194768fxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test-dds
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type of the instance. The value is fixed to Instance.
	//
	// example:
	//
	// Instance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The product. The value is fixed to dds.
	//
	// example:
	//
	// dds
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The current progress of the task. The valid values range from 0 to 100.
	//
	// example:
	//
	// 100.0
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The reason why the current task was initiated.
	//
	// example:
	//
	// test
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The estimated remaining execution time. Unit: seconds. The value 0 indicates that the task is completed.
	//
	// example:
	//
	// 1000
	RemainTime *int32 `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	// The start time of the performed O\\&M task. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-11T02:33Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// example:
	//
	// Succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the task. The task details vary based on the value of the taskType parameter.
	//
	// example:
	//
	// {\\"callerUid\\":\\"test\\"}
	TaskDetail *string `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-0mq1yyhm3ffl2bxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
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
	// CreateIns
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The ID of the user to which the resource belongs.
	//
	// example:
	//
	// 141345906006****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeHistoryTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksResponseBodyItems) GetActionInfo() *string {
	return s.ActionInfo
}

func (s *DescribeHistoryTasksResponseBodyItems) GetCallerSource() *string {
	return s.CallerSource
}

func (s *DescribeHistoryTasksResponseBodyItems) GetCallerUid() *string {
	return s.CallerUid
}

func (s *DescribeHistoryTasksResponseBodyItems) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *DescribeHistoryTasksResponseBodyItems) GetDbType() *string {
	return s.DbType
}

func (s *DescribeHistoryTasksResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeHistoryTasksResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHistoryTasksResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeHistoryTasksResponseBodyItems) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeHistoryTasksResponseBodyItems) GetProduct() *string {
	return s.Product
}

func (s *DescribeHistoryTasksResponseBodyItems) GetProgress() *float32 {
	return s.Progress
}

func (s *DescribeHistoryTasksResponseBodyItems) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *DescribeHistoryTasksResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryTasksResponseBodyItems) GetRemainTime() *int32 {
	return s.RemainTime
}

func (s *DescribeHistoryTasksResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeHistoryTasksResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeHistoryTasksResponseBodyItems) GetTaskDetail() *string {
	return s.TaskDetail
}

func (s *DescribeHistoryTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHistoryTasksResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeHistoryTasksResponseBodyItems) GetUid() *string {
	return s.Uid
}

func (s *DescribeHistoryTasksResponseBodyItems) SetActionInfo(v string) *DescribeHistoryTasksResponseBodyItems {
	s.ActionInfo = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetCallerSource(v string) *DescribeHistoryTasksResponseBodyItems {
	s.CallerSource = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetCallerUid(v string) *DescribeHistoryTasksResponseBodyItems {
	s.CallerUid = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetCurrentStepName(v string) *DescribeHistoryTasksResponseBodyItems {
	s.CurrentStepName = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetDbType(v string) *DescribeHistoryTasksResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetEndTime(v string) *DescribeHistoryTasksResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetInstanceId(v string) *DescribeHistoryTasksResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetInstanceName(v string) *DescribeHistoryTasksResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetInstanceType(v string) *DescribeHistoryTasksResponseBodyItems {
	s.InstanceType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetProduct(v string) *DescribeHistoryTasksResponseBodyItems {
	s.Product = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetProgress(v float32) *DescribeHistoryTasksResponseBodyItems {
	s.Progress = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetReasonCode(v string) *DescribeHistoryTasksResponseBodyItems {
	s.ReasonCode = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetRegionId(v string) *DescribeHistoryTasksResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetRemainTime(v int32) *DescribeHistoryTasksResponseBodyItems {
	s.RemainTime = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetStartTime(v string) *DescribeHistoryTasksResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetStatus(v string) *DescribeHistoryTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetTaskDetail(v string) *DescribeHistoryTasksResponseBodyItems {
	s.TaskDetail = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetTaskId(v string) *DescribeHistoryTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetTaskType(v string) *DescribeHistoryTasksResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetUid(v string) *DescribeHistoryTasksResponseBodyItems {
	s.Uid = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
