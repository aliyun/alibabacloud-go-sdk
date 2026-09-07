// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentMJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentMJobInfoResponseBody
	GetCode() *string
	SetData(v []*ListAgentMJobInfoResponseBodyData) *ListAgentMJobInfoResponseBody
	GetData() []*ListAgentMJobInfoResponseBodyData
	SetMessage(v string) *ListAgentMJobInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAgentMJobInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentMJobInfoResponseBody
	GetSuccess() *bool
}

type ListAgentMJobInfoResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListAgentMJobInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message, if an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D10B9203-1A6A-49DA-AE56-4D160DD37DBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false/null: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgentMJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentMJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentMJobInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentMJobInfoResponseBody) GetData() []*ListAgentMJobInfoResponseBodyData {
	return s.Data
}

func (s *ListAgentMJobInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentMJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentMJobInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentMJobInfoResponseBody) SetCode(v string) *ListAgentMJobInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentMJobInfoResponseBody) SetData(v []*ListAgentMJobInfoResponseBodyData) *ListAgentMJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentMJobInfoResponseBody) SetMessage(v string) *ListAgentMJobInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentMJobInfoResponseBody) SetRequestId(v string) *ListAgentMJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentMJobInfoResponseBody) SetSuccess(v bool) *ListAgentMJobInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentMJobInfoResponseBody) Validate() error {
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

type ListAgentMJobInfoResponseBodyData struct {
	// The end time of the scan range.
	//
	// example:
	//
	// 2026-08-26 20:00:00
	DataEndTime *string `json:"DataEndTime,omitempty" xml:"DataEndTime,omitempty"`
	// The start time of the scan range.
	//
	// example:
	//
	// 2026-08-26 19:00:00
	DataStartTime *string `json:"DataStartTime,omitempty" xml:"DataStartTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message, if an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task status. Valid values:
	//
	// - queing: The task is queued.
	//
	// - readyAnalysis: The task is pending analysis.
	//
	// - running: The task is running.
	//
	// - error: The task failed.
	//
	// - finish: The task is complete.
	//
	// - fileUploadUser: The user-specified file is uploaded.
	//
	// - fileUploadSystem: The system-generated file is uploaded.
	//
	// - expired: The task has expired.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The actual end time of the task.
	//
	// example:
	//
	// 2026-08-26 20:00:00
	TaskEndTime *string `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	// The scheduled task ID.
	//
	// example:
	//
	// 20250728-8B43DF47-24DB-1CED-8D74-2AB204187D45
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The actual start time of the task.
	//
	// example:
	//
	// 2026-08-26 19:00:00
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s ListAgentMJobInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentMJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentMJobInfoResponseBodyData) GetDataEndTime() *string {
	return s.DataEndTime
}

func (s *ListAgentMJobInfoResponseBodyData) GetDataStartTime() *string {
	return s.DataStartTime
}

func (s *ListAgentMJobInfoResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAgentMJobInfoResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListAgentMJobInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAgentMJobInfoResponseBodyData) GetTaskEndTime() *string {
	return s.TaskEndTime
}

func (s *ListAgentMJobInfoResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAgentMJobInfoResponseBodyData) GetTaskStartTime() *string {
	return s.TaskStartTime
}

func (s *ListAgentMJobInfoResponseBodyData) SetDataEndTime(v string) *ListAgentMJobInfoResponseBodyData {
	s.DataEndTime = &v
	return s
}

func (s *ListAgentMJobInfoResponseBodyData) SetDataStartTime(v string) *ListAgentMJobInfoResponseBodyData {
	s.DataStartTime = &v
	return s
}

func (s *ListAgentMJobInfoResponseBodyData) SetId(v int64) *ListAgentMJobInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAgentMJobInfoResponseBodyData) SetMessage(v string) *ListAgentMJobInfoResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListAgentMJobInfoResponseBodyData) SetStatus(v string) *ListAgentMJobInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAgentMJobInfoResponseBodyData) SetTaskEndTime(v string) *ListAgentMJobInfoResponseBodyData {
	s.TaskEndTime = &v
	return s
}

func (s *ListAgentMJobInfoResponseBodyData) SetTaskId(v string) *ListAgentMJobInfoResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListAgentMJobInfoResponseBodyData) SetTaskStartTime(v string) *ListAgentMJobInfoResponseBodyData {
	s.TaskStartTime = &v
	return s
}

func (s *ListAgentMJobInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
