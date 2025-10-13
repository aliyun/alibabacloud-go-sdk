// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePipelineResponseBody
	GetCode() *string
	SetData(v *DescribePipelineResponseBodyData) *DescribePipelineResponseBody
	GetData() *DescribePipelineResponseBodyData
	SetErrorCode(v string) *DescribePipelineResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribePipelineResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePipelineResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribePipelineResponseBody
	GetTraceId() *string
}

type DescribePipelineResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The batch information.
	Data *DescribePipelineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned for the operation.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 559B4247-C41C-4D9E-B866-B55D360B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the batch information was obtained. Valid values:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: The image failed to be found.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0be3e0c316390414649128666e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePipelineResponseBody) GetData() *DescribePipelineResponseBodyData {
	return s.Data
}

func (s *DescribePipelineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribePipelineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePipelineResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribePipelineResponseBody) SetCode(v string) *DescribePipelineResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePipelineResponseBody) SetData(v *DescribePipelineResponseBodyData) *DescribePipelineResponseBody {
	s.Data = v
	return s
}

func (s *DescribePipelineResponseBody) SetErrorCode(v string) *DescribePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribePipelineResponseBody) SetMessage(v string) *DescribePipelineResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePipelineResponseBody) SetRequestId(v string) *DescribePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePipelineResponseBody) SetSuccess(v bool) *DescribePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePipelineResponseBody) SetTraceId(v string) *DescribePipelineResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribePipelineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePipelineResponseBodyData struct {
	// The status of the change order for the batch.
	//
	// example:
	//
	// Successful
	CoStatus *string `json:"CoStatus,omitempty" xml:"CoStatus,omitempty"`
	// The ID of the batch processing stage.
	//
	// example:
	//
	// c3a55592-4c30-4d84-ac2d-e98c18ec****
	CurrentStageId *string `json:"CurrentStageId,omitempty" xml:"CurrentStageId,omitempty"`
	// The ID of the next batch.
	//
	// example:
	//
	// b77b1c98-5772-4f05-95fc-c7bee5fa****
	NextPipelineId *string `json:"NextPipelineId,omitempty" xml:"NextPipelineId,omitempty"`
	// The ID of the batch.
	//
	// example:
	//
	// 917660ba-5092-44ca-b8e0-80012c96****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The name of the batch.
	//
	// example:
	//
	// First batch
	PipelineName *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	// The batch status. Valid values:
	//
	// 	- **0**: The batch is prepared for processing.
	//
	// 	- **1**: The task is being executed.
	//
	// 	- **2**: successful
	//
	// 	- **3**: The processing failed in this stage.
	//
	// 	- **6**: The batch processing was terminated.
	//
	// 	- **10**: The batch could not be processed due to a system exception.
	//
	// example:
	//
	// 2
	PipelineStatus *int32 `json:"PipelineStatus,omitempty" xml:"PipelineStatus,omitempty"`
	// Indicates whether to start processing the next batch. Valid values:
	//
	// 	- **false**: indicates that the next batch cannot be processed yet.
	//
	// 	- **true**: indicates that the next batch can be processed now.
	//
	// example:
	//
	// false
	ShowBatch *bool `json:"ShowBatch,omitempty" xml:"ShowBatch,omitempty"`
	// The list of batch processing stages.
	StageList []*DescribePipelineResponseBodyDataStageList `json:"StageList,omitempty" xml:"StageList,omitempty" type:"Repeated"`
}

func (s DescribePipelineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBodyData) GetCoStatus() *string {
	return s.CoStatus
}

func (s *DescribePipelineResponseBodyData) GetCurrentStageId() *string {
	return s.CurrentStageId
}

func (s *DescribePipelineResponseBodyData) GetNextPipelineId() *string {
	return s.NextPipelineId
}

func (s *DescribePipelineResponseBodyData) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DescribePipelineResponseBodyData) GetPipelineName() *string {
	return s.PipelineName
}

func (s *DescribePipelineResponseBodyData) GetPipelineStatus() *int32 {
	return s.PipelineStatus
}

func (s *DescribePipelineResponseBodyData) GetShowBatch() *bool {
	return s.ShowBatch
}

func (s *DescribePipelineResponseBodyData) GetStageList() []*DescribePipelineResponseBodyDataStageList {
	return s.StageList
}

func (s *DescribePipelineResponseBodyData) SetCoStatus(v string) *DescribePipelineResponseBodyData {
	s.CoStatus = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetCurrentStageId(v string) *DescribePipelineResponseBodyData {
	s.CurrentStageId = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetNextPipelineId(v string) *DescribePipelineResponseBodyData {
	s.NextPipelineId = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetPipelineId(v string) *DescribePipelineResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetPipelineName(v string) *DescribePipelineResponseBodyData {
	s.PipelineName = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetPipelineStatus(v int32) *DescribePipelineResponseBodyData {
	s.PipelineStatus = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetShowBatch(v bool) *DescribePipelineResponseBodyData {
	s.ShowBatch = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetStageList(v []*DescribePipelineResponseBodyDataStageList) *DescribePipelineResponseBodyData {
	s.StageList = v
	return s
}

func (s *DescribePipelineResponseBodyData) Validate() error {
	if s.StageList != nil {
		for _, item := range s.StageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePipelineResponseBodyDataStageList struct {
	// The execution type of the stage. Valid values:
	//
	// 	- **0**: in sequence.
	//
	// 	- **1**: in parallel.
	//
	// example:
	//
	// 0
	ExecutorType *int32 `json:"ExecutorType,omitempty" xml:"ExecutorType,omitempty"`
	// The ID of the stage.
	//
	// example:
	//
	// c3a55592-4c30-4d84-ac2d-e98c18ec****
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The name of the stage.
	//
	// example:
	//
	// Deploy an application
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The status of the batch processing stage. Valid values:
	//
	// 	- **0**: The batch is prepared for this processing stage.
	//
	// 	- **1**: The task is being executed.
	//
	// 	- **2**: successful
	//
	// 	- **3**: The processing failed in this stage.
	//
	// 	- **6**: The processing stage was terminated.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of task statuses.
	TaskList []*DescribePipelineResponseBodyDataStageListTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s DescribePipelineResponseBodyDataStageList) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineResponseBodyDataStageList) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBodyDataStageList) GetExecutorType() *int32 {
	return s.ExecutorType
}

func (s *DescribePipelineResponseBodyDataStageList) GetStageId() *string {
	return s.StageId
}

func (s *DescribePipelineResponseBodyDataStageList) GetStageName() *string {
	return s.StageName
}

func (s *DescribePipelineResponseBodyDataStageList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribePipelineResponseBodyDataStageList) GetTaskList() []*DescribePipelineResponseBodyDataStageListTaskList {
	return s.TaskList
}

func (s *DescribePipelineResponseBodyDataStageList) SetExecutorType(v int32) *DescribePipelineResponseBodyDataStageList {
	s.ExecutorType = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) SetStageId(v string) *DescribePipelineResponseBodyDataStageList {
	s.StageId = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) SetStageName(v string) *DescribePipelineResponseBodyDataStageList {
	s.StageName = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) SetStatus(v int32) *DescribePipelineResponseBodyDataStageList {
	s.Status = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) SetTaskList(v []*DescribePipelineResponseBodyDataStageListTaskList) *DescribePipelineResponseBodyDataStageList {
	s.TaskList = v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePipelineResponseBodyDataStageListTaskList struct {
	// The error code returned when the task could not be executed. If the task is successfully executed, this parameter is not returned.
	//
	// example:
	//
	// EDAS-10022
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether to execute the subsequent tasks when the task failed. Valid values:
	//
	// 	- **0**: The subsequent tasks cannot be executed.
	//
	// 	- **1**: The subsequent tasks can be executed.
	//
	// example:
	//
	// 0
	ErrorIgnore *int32 `json:"ErrorIgnore,omitempty" xml:"ErrorIgnore,omitempty"`
	// The error message returned when the task could not be executed. If the task is successfully executed, this parameter is not returned.
	//
	// example:
	//
	// EDAS-10022
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The returned message indicating the task execution result.
	//
	// example:
	//
	// init Namespace success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether a running task can be manually skipped. Valid values:
	//
	// 	- **true**: The running task can be skipped.
	//
	// 	- **false**: The zone does not allow you to change the network type of an ApsaraDB for Redis instance from classic network to VPC.
	//
	// example:
	//
	// false
	ShowManualIgnore *bool `json:"ShowManualIgnore,omitempty" xml:"ShowManualIgnore,omitempty"`
	// The ID of the stage.
	//
	// example:
	//
	// c3a55592-4c30-4d84-ac2d-e98c18ec****
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The task status. Valid values:
	//
	// 	- **0**: The task is prepared for execution.
	//
	// 	- **1**: The task is being executed.
	//
	// 	- **2**: successful
	//
	// 	- **3**: The task could not be executed.
	//
	// 	- **5**: The task is pending retry.
	//
	// 	- **6**: The task was terminated.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// bef0122f-de9a-4ab0-8223-b88bf8ad****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Environment initialization
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribePipelineResponseBodyDataStageListTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineResponseBodyDataStageListTaskList) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetErrorIgnore() *int32 {
	return s.ErrorIgnore
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetMessage() *string {
	return s.Message
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetShowManualIgnore() *bool {
	return s.ShowManualIgnore
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetStageId() *string {
	return s.StageId
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetErrorCode(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.ErrorCode = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetErrorIgnore(v int32) *DescribePipelineResponseBodyDataStageListTaskList {
	s.ErrorIgnore = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetErrorMessage(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.ErrorMessage = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetMessage(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.Message = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetShowManualIgnore(v bool) *DescribePipelineResponseBodyDataStageListTaskList {
	s.ShowManualIgnore = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetStageId(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.StageId = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetStatus(v int32) *DescribePipelineResponseBodyDataStageListTaskList {
	s.Status = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetTaskId(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetTaskName(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.TaskName = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) Validate() error {
	return dara.Validate(s)
}
