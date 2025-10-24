// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeTasksResponseBodyData) *DescribeTasksResponseBody
	GetData() []*DescribeTasksResponseBodyData
	SetNextToken(v string) *DescribeTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTasksResponseBody
	GetTotalCount() *int32
}

type DescribeTasksResponseBody struct {
	// The objects that are returned.
	Data []*DescribeTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) GetData() []*DescribeTasksResponseBodyData {
	return s.Data
}

func (s *DescribeTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTasksResponseBody) SetData(v []*DescribeTasksResponseBodyData) *DescribeTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTasksResponseBody) SetNextToken(v string) *DescribeTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalCount(v int32) *DescribeTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTasksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTasksResponseBodyData struct {
	// The error code.
	//
	// example:
	//
	// SendFileFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// connect error.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The total number of failed subtasks.
	//
	// example:
	//
	// 2
	FailedChildCount *int32 `json:"FailedChildCount,omitempty" xml:"FailedChildCount,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2022-10-11T08:53:32Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the cloud phone instance.
	//
	// example:
	//
	// defaultInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The state of the cloud phone instance.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The ID of the command execution.
	//
	// example:
	//
	// B8ED2BA9-0C6A-5643-818F-B5D60A64****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The level of the task.
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The operator.
	//
	// example:
	//
	// test
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parameters of the task.
	//
	// example:
	//
	// param
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The ID of the parent task.
	//
	// example:
	//
	// t-41oan3tza16vs****
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// acp-25nt4kk9whhok****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The execution result of the task.
	//
	// example:
	//
	// {\\"Success\\": True}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The total number of the subtasks that are running.
	//
	// example:
	//
	// 0
	RunningChildCount *int32 `json:"RunningChildCount,omitempty" xml:"RunningChildCount,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2022-10-11T08:53:32Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of successfully executed subtasks.
	//
	// example:
	//
	// 98
	SuccessChildCount *int32 `json:"SuccessChildCount,omitempty" xml:"SuccessChildCount,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The state of the task.
	//
	// example:
	//
	// Processing
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// StartInstance
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of subtasks of the current batch task.
	//
	// example:
	//
	// 100
	TotalChildCount *int32 `json:"TotalChildCount,omitempty" xml:"TotalChildCount,omitempty"`
}

func (s DescribeTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeTasksResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeTasksResponseBodyData) GetFailedChildCount() *int32 {
	return s.FailedChildCount
}

func (s *DescribeTasksResponseBodyData) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeTasksResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTasksResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTasksResponseBodyData) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTasksResponseBodyData) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeTasksResponseBodyData) GetLevel() *int32 {
	return s.Level
}

func (s *DescribeTasksResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *DescribeTasksResponseBodyData) GetParam() *string {
	return s.Param
}

func (s *DescribeTasksResponseBodyData) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *DescribeTasksResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTasksResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTasksResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *DescribeTasksResponseBodyData) GetRunningChildCount() *int32 {
	return s.RunningChildCount
}

func (s *DescribeTasksResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTasksResponseBodyData) GetSuccessChildCount() *int32 {
	return s.SuccessChildCount
}

func (s *DescribeTasksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTasksResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTasksResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTasksResponseBodyData) GetTotalChildCount() *int32 {
	return s.TotalChildCount
}

func (s *DescribeTasksResponseBodyData) SetErrorCode(v string) *DescribeTasksResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetErrorMsg(v string) *DescribeTasksResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetFailedChildCount(v int32) *DescribeTasksResponseBodyData {
	s.FailedChildCount = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetFinishTime(v string) *DescribeTasksResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetInstanceId(v string) *DescribeTasksResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetInstanceName(v string) *DescribeTasksResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetInstanceStatus(v string) *DescribeTasksResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetInvokeId(v string) *DescribeTasksResponseBodyData {
	s.InvokeId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetLevel(v int32) *DescribeTasksResponseBodyData {
	s.Level = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetOperator(v string) *DescribeTasksResponseBodyData {
	s.Operator = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetParam(v string) *DescribeTasksResponseBodyData {
	s.Param = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetParentTaskId(v string) *DescribeTasksResponseBodyData {
	s.ParentTaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetRegionId(v string) *DescribeTasksResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetResourceId(v string) *DescribeTasksResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetResult(v string) *DescribeTasksResponseBodyData {
	s.Result = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetRunningChildCount(v int32) *DescribeTasksResponseBodyData {
	s.RunningChildCount = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetStartTime(v string) *DescribeTasksResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetSuccessChildCount(v int32) *DescribeTasksResponseBodyData {
	s.SuccessChildCount = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTaskId(v string) *DescribeTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTaskStatus(v string) *DescribeTasksResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTaskType(v string) *DescribeTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTotalChildCount(v int32) *DescribeTasksResponseBodyData {
	s.TotalChildCount = &v
	return s
}

func (s *DescribeTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
