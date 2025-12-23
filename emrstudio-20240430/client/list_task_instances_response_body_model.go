// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListTaskInstancesResponseBodyData) *ListTaskInstancesResponseBody
	GetData() []*ListTaskInstancesResponseBodyData
	SetNextToken(v string) *ListTaskInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTaskInstancesResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *ListTaskInstancesResponseBody
	GetTotalSize() *int32
}

type ListTaskInstancesResponseBody struct {
	Data []*ListTaskInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA38***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBody) GetData() []*ListTaskInstancesResponseBodyData {
	return s.Data
}

func (s *ListTaskInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskInstancesResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListTaskInstancesResponseBody) SetData(v []*ListTaskInstancesResponseBodyData) *ListTaskInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskInstancesResponseBody) SetNextToken(v string) *ListTaskInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTaskInstancesResponseBody) SetRequestId(v string) *ListTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskInstancesResponseBody) SetTotalSize(v int32) *ListTaskInstancesResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListTaskInstancesResponseBody) Validate() error {
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

type ListTaskInstancesResponseBodyData struct {
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
	// wg-3q9jo749ne5****
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
}

func (s ListTaskInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponseBodyData) GetDryRun() *string {
	return s.DryRun
}

func (s *ListTaskInstancesResponseBodyData) GetEmrClusterId() *string {
	return s.EmrClusterId
}

func (s *ListTaskInstancesResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTaskInstancesResponseBodyData) GetExternalAppId() *string {
	return s.ExternalAppId
}

func (s *ListTaskInstancesResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTaskInstancesResponseBodyData) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *ListTaskInstancesResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTaskInstancesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListTaskInstancesResponseBodyData) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListTaskInstancesResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTaskInstancesResponseBodyData) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *ListTaskInstancesResponseBodyData) GetTaskInstanceName() *string {
	return s.TaskInstanceName
}

func (s *ListTaskInstancesResponseBodyData) GetTaskParams() *string {
	return s.TaskParams
}

func (s *ListTaskInstancesResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTaskInstancesResponseBodyData) GetTaskVersion() *string {
	return s.TaskVersion
}

func (s *ListTaskInstancesResponseBodyData) GetWorkflowInstanceId() *string {
	return s.WorkflowInstanceId
}

func (s *ListTaskInstancesResponseBodyData) SetDryRun(v string) *ListTaskInstancesResponseBodyData {
	s.DryRun = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetEmrClusterId(v string) *ListTaskInstancesResponseBodyData {
	s.EmrClusterId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetEndTime(v string) *ListTaskInstancesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetExternalAppId(v string) *ListTaskInstancesResponseBodyData {
	s.ExternalAppId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetResourceGroupId(v string) *ListTaskInstancesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetRetryTimes(v int32) *ListTaskInstancesResponseBodyData {
	s.RetryTimes = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetStartTime(v string) *ListTaskInstancesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetStatus(v string) *ListTaskInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetSubmitTime(v string) *ListTaskInstancesResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskId(v string) *ListTaskInstancesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskInstanceId(v string) *ListTaskInstancesResponseBodyData {
	s.TaskInstanceId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskInstanceName(v string) *ListTaskInstancesResponseBodyData {
	s.TaskInstanceName = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskParams(v string) *ListTaskInstancesResponseBodyData {
	s.TaskParams = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskType(v string) *ListTaskInstancesResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetTaskVersion(v string) *ListTaskInstancesResponseBodyData {
	s.TaskVersion = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) SetWorkflowInstanceId(v string) *ListTaskInstancesResponseBodyData {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
