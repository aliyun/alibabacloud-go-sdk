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
	// The list of task objects.
	Items []*DescribeHistoryTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of tasks that meet the query conditions, regardless of pagination.
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
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHistoryTasksResponseBodyItems struct {
	// The information about allowed operations. When you use this parameter, the system matches the operation action based on the CurrentStepName and Status in this information. If no action is matched, the current task status does not support the operation. Example:
	//
	// ```
	//
	// {"steps": [
	//
	//     {
	//
	//       "step_name": "exec_task", // The step name, which matches CurrentStepName.
	//
	//       "action_info": {    // The operations supported by the step.
	//
	//         "Waiting": [      // The status, which matches Status.
	//
	//           "modifySwitchTime" // The operation action. Multiple actions indicate that multiple operations are supported.
	//
	//         ]
	//
	//       }
	//
	//     },
	//
	//     {
	//
	//       "step_name": "init_task", // The step name.
	//
	//       "action_info": {    // The operations supported by the step.
	//
	//         "Running": [      // The status.
	//
	//           "cancel",       // The operation.
	//
	//           "pause"
	//
	//         ]
	//
	//       }
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// ```
	//
	// Operations that may be supported by the system:
	//
	// 	- **retry**: Retry.
	//
	// 	- **cancel**: Cancel.
	//
	// 	- **modifySwitchTime**: Modify the switchover time or recovery time.
	//
	// example:
	//
	// {"steps":[{"action_info":{"Waiting":["modifySwitchTime"]},"step_name":"exec_task"}]}
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The ID of the requesting user. If CallerSource is set to User, this parameter represents the user ID.
	//
	// example:
	//
	// 141345906006****
	CallerSource *string `json:"CallerSource,omitempty" xml:"CallerSource,omitempty"`
	// The request source. Valid values:
	//
	// 	- **System**: Automatic system operation.
	//
	// 	- **User**: User.
	//
	// example:
	//
	// User
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The name of the current step being executed. If this parameter is empty, the task has not started.
	//
	// example:
	//
	// exec_task
	CurrentStepName *string `json:"CurrentStepName,omitempty" xml:"CurrentStepName,omitempty"`
	// The database type. The value is fixed as redis.
	//
	// example:
	//
	// redis
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The end time of the task. The time is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2022-02-03T12:06:17Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// dba-tair-test-qcloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. The value is fixed as Instance.
	//
	// example:
	//
	// Instance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The product. The value is fixed as kvstore.
	//
	// example:
	//
	// kvstore
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The current progress. Valid values: 0 to 100.
	//
	// example:
	//
	// 79.0
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The reason why the current task was initiated.
	//
	// example:
	//
	// Minor version upgrade
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The estimated remaining execution time. Unit: seconds. A value of 0 indicates that the execution is complete.
	//
	// example:
	//
	// 1000
	RemainTime *int32 `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	// The start time of the task. The time is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2022-02-03T11:31:03Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task details. The output varies based on the task type.
	//
	// example:
	//
	// {\\"callerUid\\":\\"test\\"}
	TaskDetail *string `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-83br18hloy3faf****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// - **ModifyInsSpec**: Instance specification change or migration.
	//
	// - **DeleteInsNode**: Shard deletion.
	//
	// - **AddInsNode**: Shard addition.
	//
	// - **HaSwitch**: Primary/secondary switchover.
	//
	// - **RestartIns**: Instance restart.
	//
	// - **CreateIns**: Instance creation.
	//
	// - **ModifyInsConfig**: Configuration change.
	//
	// example:
	//
	// ModifyInsSpec
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The ID of the user to whom the resource belongs.
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
