// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTimingSyntheticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateTimingSyntheticTaskResponseBody
	GetCode() *int64
	SetData(v *CreateTimingSyntheticTaskResponseBodyData) *CreateTimingSyntheticTaskResponseBody
	GetData() *CreateTimingSyntheticTaskResponseBodyData
	SetMessage(v string) *CreateTimingSyntheticTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTimingSyntheticTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTimingSyntheticTaskResponseBody
	GetSuccess() *bool
}

type CreateTimingSyntheticTaskResponseBody struct {
	// The HTTP status code returned. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *CreateTimingSyntheticTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 53B5874D-EBC1-5567-B787-E4B7267F5CEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTimingSyntheticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateTimingSyntheticTaskResponseBody) GetData() *CreateTimingSyntheticTaskResponseBodyData {
	return s.Data
}

func (s *CreateTimingSyntheticTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTimingSyntheticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTimingSyntheticTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTimingSyntheticTaskResponseBody) SetCode(v int64) *CreateTimingSyntheticTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTimingSyntheticTaskResponseBody) SetData(v *CreateTimingSyntheticTaskResponseBodyData) *CreateTimingSyntheticTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateTimingSyntheticTaskResponseBody) SetMessage(v string) *CreateTimingSyntheticTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTimingSyntheticTaskResponseBody) SetRequestId(v string) *CreateTimingSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTimingSyntheticTaskResponseBody) SetSuccess(v bool) *CreateTimingSyntheticTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTimingSyntheticTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTimingSyntheticTaskResponseBodyData struct {
	// The task status. Valid values:
	//
	// - INIT: The task is in the initial state.
	//
	// - RELEASE: The task is being parsed.
	//
	// - RUNNING: The task is running.
	//
	// - STOP: The task is suspended.
	//
	// - SYSTEM_STOP: The task is suspended by the system.
	//
	// - CANCEL: The task is canceled.
	//
	// - SYSTEM_CANCEL: The task is canceled by the system.
	//
	// - DONE: The task is complete.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 5308a2691f59422c8c3b7aeccec9cd3b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateTimingSyntheticTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateTimingSyntheticTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTimingSyntheticTaskResponseBodyData) SetStatus(v string) *CreateTimingSyntheticTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTimingSyntheticTaskResponseBodyData) SetTaskId(v string) *CreateTimingSyntheticTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateTimingSyntheticTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
