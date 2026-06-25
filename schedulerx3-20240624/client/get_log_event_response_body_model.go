// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetLogEventResponseBody
	GetCode() *int32
	SetData(v *GetLogEventResponseBodyData) *GetLogEventResponseBody
	GetData() *GetLogEventResponseBodyData
	SetMessage(v string) *GetLogEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLogEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLogEventResponseBody
	GetSuccess() *bool
}

type GetLogEventResponseBody struct {
	// The HTTP status code. A value of `200` indicates a successful request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetLogEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request fails.
	//
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique, server-generated ID for the request. This ID is used for troubleshooting purposes.
	//
	// example:
	//
	// BAC1ADB5-EEB5-5834-93D8-522E067AF8D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLogEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogEventResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetLogEventResponseBody) GetData() *GetLogEventResponseBodyData {
	return s.Data
}

func (s *GetLogEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLogEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLogEventResponseBody) SetCode(v int32) *GetLogEventResponseBody {
	s.Code = &v
	return s
}

func (s *GetLogEventResponseBody) SetData(v *GetLogEventResponseBodyData) *GetLogEventResponseBody {
	s.Data = v
	return s
}

func (s *GetLogEventResponseBody) SetMessage(v string) *GetLogEventResponseBody {
	s.Message = &v
	return s
}

func (s *GetLogEventResponseBody) SetRequestId(v string) *GetLogEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogEventResponseBody) SetSuccess(v bool) *GetLogEventResponseBody {
	s.Success = &v
	return s
}

func (s *GetLogEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLogEventResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*GetLogEventResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 33
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLogEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLogEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLogEventResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetLogEventResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetLogEventResponseBodyData) GetRecords() []*GetLogEventResponseBodyDataRecords {
	return s.Records
}

func (s *GetLogEventResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetLogEventResponseBodyData) SetPageNumber(v int32) *GetLogEventResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetLogEventResponseBodyData) SetPageSize(v int32) *GetLogEventResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetLogEventResponseBodyData) SetRecords(v []*GetLogEventResponseBodyDataRecords) *GetLogEventResponseBodyData {
	s.Records = v
	return s
}

func (s *GetLogEventResponseBodyData) SetTotal(v int64) *GetLogEventResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetLogEventResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLogEventResponseBodyDataRecords struct {
	// The name of the application.
	//
	// example:
	//
	// portal-dev
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The log content.
	//
	// example:
	//
	// hello word
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The severity level of the event.
	//
	// example:
	//
	// info
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The type of the event.
	//
	// example:
	//
	// JOB
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The job execution ID.
	//
	// example:
	//
	// 101
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// test
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The time when the log was recorded. The time is in the `yyyy-MM-dd HH:mm:ss` format.
	//
	// example:
	//
	// 2024-10-31 16:43:51
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The address of the worker that runs the job.
	//
	// example:
	//
	// 030225016025_9357_60125@127.0.0.1:51363
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
	// The workflow execution ID.
	//
	// example:
	//
	// 1450568762586578000
	WorkflowExecutionId *string `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// 工作流0001
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s GetLogEventResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetLogEventResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetLogEventResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *GetLogEventResponseBodyDataRecords) GetContent() *string {
	return s.Content
}

func (s *GetLogEventResponseBodyDataRecords) GetEvent() *string {
	return s.Event
}

func (s *GetLogEventResponseBodyDataRecords) GetEventType() *string {
	return s.EventType
}

func (s *GetLogEventResponseBodyDataRecords) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetLogEventResponseBodyDataRecords) GetJobName() *string {
	return s.JobName
}

func (s *GetLogEventResponseBodyDataRecords) GetTime() *string {
	return s.Time
}

func (s *GetLogEventResponseBodyDataRecords) GetWorkerAddr() *string {
	return s.WorkerAddr
}

func (s *GetLogEventResponseBodyDataRecords) GetWorkflowExecutionId() *string {
	return s.WorkflowExecutionId
}

func (s *GetLogEventResponseBodyDataRecords) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *GetLogEventResponseBodyDataRecords) SetAppName(v string) *GetLogEventResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetContent(v string) *GetLogEventResponseBodyDataRecords {
	s.Content = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetEvent(v string) *GetLogEventResponseBodyDataRecords {
	s.Event = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetEventType(v string) *GetLogEventResponseBodyDataRecords {
	s.EventType = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetJobExecutionId(v string) *GetLogEventResponseBodyDataRecords {
	s.JobExecutionId = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetJobName(v string) *GetLogEventResponseBodyDataRecords {
	s.JobName = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetTime(v string) *GetLogEventResponseBodyDataRecords {
	s.Time = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetWorkerAddr(v string) *GetLogEventResponseBodyDataRecords {
	s.WorkerAddr = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetWorkflowExecutionId(v string) *GetLogEventResponseBodyDataRecords {
	s.WorkflowExecutionId = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) SetWorkflowName(v string) *GetLogEventResponseBodyDataRecords {
	s.WorkflowName = &v
	return s
}

func (s *GetLogEventResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
