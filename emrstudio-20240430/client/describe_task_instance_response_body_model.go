// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *DescribeTaskInstanceResponseBody
	GetDryRun() *string
	SetEmrClusterId(v string) *DescribeTaskInstanceResponseBody
	GetEmrClusterId() *string
	SetEndTime(v string) *DescribeTaskInstanceResponseBody
	GetEndTime() *string
	SetExternalAppId(v string) *DescribeTaskInstanceResponseBody
	GetExternalAppId() *string
	SetResourceGroupId(v string) *DescribeTaskInstanceResponseBody
	GetResourceGroupId() *string
	SetRetryTimes(v int32) *DescribeTaskInstanceResponseBody
	GetRetryTimes() *int32
	SetStartTime(v string) *DescribeTaskInstanceResponseBody
	GetStartTime() *string
	SetStatus(v string) *DescribeTaskInstanceResponseBody
	GetStatus() *string
	SetSubmitTime(v string) *DescribeTaskInstanceResponseBody
	GetSubmitTime() *string
	SetTaskId(v string) *DescribeTaskInstanceResponseBody
	GetTaskId() *string
	SetTaskInstanceId(v string) *DescribeTaskInstanceResponseBody
	GetTaskInstanceId() *string
	SetTaskInstanceName(v string) *DescribeTaskInstanceResponseBody
	GetTaskInstanceName() *string
	SetTaskParams(v string) *DescribeTaskInstanceResponseBody
	GetTaskParams() *string
	SetTaskType(v string) *DescribeTaskInstanceResponseBody
	GetTaskType() *string
	SetTaskVersion(v string) *DescribeTaskInstanceResponseBody
	GetTaskVersion() *string
	SetWorkflowInstanceId(v string) *DescribeTaskInstanceResponseBody
	GetWorkflowInstanceId() *string
	SetRequestId(v string) *DescribeTaskInstanceResponseBody
	GetRequestId() *string
}

type DescribeTaskInstanceResponseBody struct {
	// example:
	//
	// 0
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// c-b933c5aac7f7***
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// application_123_***
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// example:
	//
	// wg-123abc***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 0
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// t-3q9jo749ne5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ti-3q9jo749ne5****
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// example:
	//
	// test
	TaskInstanceName *string `json:"TaskInstanceName,omitempty" xml:"TaskInstanceName,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 1
	TaskVersion *string `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	// example:
	//
	// wi-3q9jo749ne5****
	WorkflowInstanceId *string `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskInstanceResponseBody) GetDryRun() *string {
	return s.DryRun
}

func (s *DescribeTaskInstanceResponseBody) GetEmrClusterId() *string {
	return s.EmrClusterId
}

func (s *DescribeTaskInstanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTaskInstanceResponseBody) GetExternalAppId() *string {
	return s.ExternalAppId
}

func (s *DescribeTaskInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTaskInstanceResponseBody) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *DescribeTaskInstanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTaskInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeTaskInstanceResponseBody) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *DescribeTaskInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskInstanceResponseBody) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *DescribeTaskInstanceResponseBody) GetTaskInstanceName() *string {
	return s.TaskInstanceName
}

func (s *DescribeTaskInstanceResponseBody) GetTaskParams() *string {
	return s.TaskParams
}

func (s *DescribeTaskInstanceResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskInstanceResponseBody) GetTaskVersion() *string {
	return s.TaskVersion
}

func (s *DescribeTaskInstanceResponseBody) GetWorkflowInstanceId() *string {
	return s.WorkflowInstanceId
}

func (s *DescribeTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskInstanceResponseBody) SetDryRun(v string) *DescribeTaskInstanceResponseBody {
	s.DryRun = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetEmrClusterId(v string) *DescribeTaskInstanceResponseBody {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetEndTime(v string) *DescribeTaskInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetExternalAppId(v string) *DescribeTaskInstanceResponseBody {
	s.ExternalAppId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetResourceGroupId(v string) *DescribeTaskInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetRetryTimes(v int32) *DescribeTaskInstanceResponseBody {
	s.RetryTimes = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetStartTime(v string) *DescribeTaskInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetStatus(v string) *DescribeTaskInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetSubmitTime(v string) *DescribeTaskInstanceResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskId(v string) *DescribeTaskInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskInstanceId(v string) *DescribeTaskInstanceResponseBody {
	s.TaskInstanceId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskInstanceName(v string) *DescribeTaskInstanceResponseBody {
	s.TaskInstanceName = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskParams(v string) *DescribeTaskInstanceResponseBody {
	s.TaskParams = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskType(v string) *DescribeTaskInstanceResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetTaskVersion(v string) *DescribeTaskInstanceResponseBody {
	s.TaskVersion = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetWorkflowInstanceId(v string) *DescribeTaskInstanceResponseBody {
	s.WorkflowInstanceId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) SetRequestId(v string) *DescribeTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
