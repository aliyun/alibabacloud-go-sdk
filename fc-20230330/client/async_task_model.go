// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncTask interface {
	dara.Model
	String() string
	GoString() string
	SetAlreadyRetriedTimes(v int64) *AsyncTask
	GetAlreadyRetriedTimes() *int64
	SetDestinationStatus(v string) *AsyncTask
	GetDestinationStatus() *string
	SetDurationMs(v int64) *AsyncTask
	GetDurationMs() *int64
	SetEndTime(v int64) *AsyncTask
	GetEndTime() *int64
	SetEvents(v []*AsyncTaskEvent) *AsyncTask
	GetEvents() []*AsyncTaskEvent
	SetFunctionArn(v string) *AsyncTask
	GetFunctionArn() *string
	SetInstanceId(v string) *AsyncTask
	GetInstanceId() *string
	SetQualifier(v string) *AsyncTask
	GetQualifier() *string
	SetRequestId(v string) *AsyncTask
	GetRequestId() *string
	SetReturnPayload(v string) *AsyncTask
	GetReturnPayload() *string
	SetStartedTime(v int64) *AsyncTask
	GetStartedTime() *int64
	SetStatus(v string) *AsyncTask
	GetStatus() *string
	SetTaskErrorMessage(v string) *AsyncTask
	GetTaskErrorMessage() *string
	SetTaskId(v string) *AsyncTask
	GetTaskId() *string
	SetTaskPayload(v string) *AsyncTask
	GetTaskPayload() *string
}

type AsyncTask struct {
	// The number of retries after the asynchronous task fails.
	//
	// example:
	//
	// 3
	AlreadyRetriedTimes *int64 `json:"alreadyRetriedTimes,omitempty" xml:"alreadyRetriedTimes,omitempty"`
	// The final state of the asynchronous task.
	//
	// example:
	//
	// Succeeded
	DestinationStatus *string `json:"destinationStatus,omitempty" xml:"destinationStatus,omitempty"`
	// The execution duration of the asynchronous task.
	//
	// example:
	//
	// 1000
	DurationMs *int64 `json:"durationMs,omitempty" xml:"durationMs,omitempty"`
	// The end time of the asynchronous task. Unit: milliseconds.
	//
	// example:
	//
	// 1633449590000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The events of the asynchronous task.
	Events []*AsyncTaskEvent `json:"events" xml:"events" type:"Repeated"`
	// The Alibaba Cloud Resource Name (ARN) of the function.
	//
	// example:
	//
	// acs:fc:cn-shanghai:1234/functions/my-func
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// The ID of the instance that corresponds to the asynchronous task.
	//
	// example:
	//
	// D4-*******9FD1-882707E
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// prod
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The ID of the request corresponding to this asynchronous task.
	//
	// example:
	//
	// e026ae92-61e5-472f-b32d-1c9e3c4e****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The content of the response after the asynchronous task is executed. The maximum size is 1 MB. This parameter is in public preview. If you want to use this parameter, [contact us](https://help.aliyun.com/document_detail/2513733.html).
	//
	// example:
	//
	// result
	ReturnPayload *string `json:"returnPayload,omitempty" xml:"returnPayload,omitempty"`
	// The start time of the asynchronous task. Unit: milliseconds.
	//
	// example:
	//
	// 1633449590000
	StartedTime *int64 `json:"startedTime,omitempty" xml:"startedTime,omitempty"`
	// The state of the asynchronous task.
	//
	// 	- Enqueued: The asynchronous invocation is enqueued and waiting to be executed.
	//
	// 	- Succeeded: The invocation is successful.
	//
	// 	- Failed: The invocation fails.
	//
	// 	- Running: The invocation is being executed.
	//
	// 	- Stopped: The invocation is terminated.
	//
	// 	- Stopping: The invocation is being terminated.
	//
	// 	- Invalid: The invocation is invalid and not executed due to specific reasons. For example, the function is deleted.
	//
	// 	- Expired: The maximum validity period of messages is specified for asynchronous invocation. The invocation is discarded and not executed because the specified maximum validity period of messages expires.
	//
	// 	- Retrying: The asynchronous invocation is being retried due to an execution error.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The error message for an asynchronous task failure.
	//
	// example:
	//
	// UnhandledInvocationError
	TaskErrorMessage *string `json:"taskErrorMessage,omitempty" xml:"taskErrorMessage,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// e026ae92-61e5-472f-b32d-1c9e3c4e****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The content of the input parameter during asynchronous task execution.
	//
	// example:
	//
	// body
	TaskPayload *string `json:"taskPayload,omitempty" xml:"taskPayload,omitempty"`
}

func (s AsyncTask) String() string {
	return dara.Prettify(s)
}

func (s AsyncTask) GoString() string {
	return s.String()
}

func (s *AsyncTask) GetAlreadyRetriedTimes() *int64 {
	return s.AlreadyRetriedTimes
}

func (s *AsyncTask) GetDestinationStatus() *string {
	return s.DestinationStatus
}

func (s *AsyncTask) GetDurationMs() *int64 {
	return s.DurationMs
}

func (s *AsyncTask) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AsyncTask) GetEvents() []*AsyncTaskEvent {
	return s.Events
}

func (s *AsyncTask) GetFunctionArn() *string {
	return s.FunctionArn
}

func (s *AsyncTask) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AsyncTask) GetQualifier() *string {
	return s.Qualifier
}

func (s *AsyncTask) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncTask) GetReturnPayload() *string {
	return s.ReturnPayload
}

func (s *AsyncTask) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *AsyncTask) GetStatus() *string {
	return s.Status
}

func (s *AsyncTask) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *AsyncTask) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncTask) GetTaskPayload() *string {
	return s.TaskPayload
}

func (s *AsyncTask) SetAlreadyRetriedTimes(v int64) *AsyncTask {
	s.AlreadyRetriedTimes = &v
	return s
}

func (s *AsyncTask) SetDestinationStatus(v string) *AsyncTask {
	s.DestinationStatus = &v
	return s
}

func (s *AsyncTask) SetDurationMs(v int64) *AsyncTask {
	s.DurationMs = &v
	return s
}

func (s *AsyncTask) SetEndTime(v int64) *AsyncTask {
	s.EndTime = &v
	return s
}

func (s *AsyncTask) SetEvents(v []*AsyncTaskEvent) *AsyncTask {
	s.Events = v
	return s
}

func (s *AsyncTask) SetFunctionArn(v string) *AsyncTask {
	s.FunctionArn = &v
	return s
}

func (s *AsyncTask) SetInstanceId(v string) *AsyncTask {
	s.InstanceId = &v
	return s
}

func (s *AsyncTask) SetQualifier(v string) *AsyncTask {
	s.Qualifier = &v
	return s
}

func (s *AsyncTask) SetRequestId(v string) *AsyncTask {
	s.RequestId = &v
	return s
}

func (s *AsyncTask) SetReturnPayload(v string) *AsyncTask {
	s.ReturnPayload = &v
	return s
}

func (s *AsyncTask) SetStartedTime(v int64) *AsyncTask {
	s.StartedTime = &v
	return s
}

func (s *AsyncTask) SetStatus(v string) *AsyncTask {
	s.Status = &v
	return s
}

func (s *AsyncTask) SetTaskErrorMessage(v string) *AsyncTask {
	s.TaskErrorMessage = &v
	return s
}

func (s *AsyncTask) SetTaskId(v string) *AsyncTask {
	s.TaskId = &v
	return s
}

func (s *AsyncTask) SetTaskPayload(v string) *AsyncTask {
	s.TaskPayload = &v
	return s
}

func (s *AsyncTask) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
