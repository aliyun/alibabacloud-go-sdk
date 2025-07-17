// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSyntheticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSyntheticTaskResponseBody
	GetCode() *string
	SetData(v *CreateSyntheticTaskResponseBodyData) *CreateSyntheticTaskResponseBody
	GetData() *CreateSyntheticTaskResponseBodyData
	SetMsg(v string) *CreateSyntheticTaskResponseBody
	GetMsg() *string
	SetRequestId(v string) *CreateSyntheticTaskResponseBody
	GetRequestId() *string
}

type CreateSyntheticTaskResponseBody struct {
	// The status code returned.
	//
	// 	- 1001: The request was successful.
	//
	// 	- 1002: The request failed.
	//
	// 	- 1003: Parameter errors occurred.
	//
	// 	- 1004: Authentication failed.
	//
	// 	- 1006: The task does not exist.
	//
	// 	- 1099: Internal errors occurred.
	//
	// example:
	//
	// 1001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the synthetic monitoring task.
	Data *CreateSyntheticTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned when the task failed to be created.
	//
	// example:
	//
	// null
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSyntheticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSyntheticTaskResponseBody) GetData() *CreateSyntheticTaskResponseBodyData {
	return s.Data
}

func (s *CreateSyntheticTaskResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CreateSyntheticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSyntheticTaskResponseBody) SetCode(v string) *CreateSyntheticTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSyntheticTaskResponseBody) SetData(v *CreateSyntheticTaskResponseBodyData) *CreateSyntheticTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateSyntheticTaskResponseBody) SetMsg(v string) *CreateSyntheticTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateSyntheticTaskResponseBody) SetRequestId(v string) *CreateSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSyntheticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskResponseBodyData struct {
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSyntheticTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateSyntheticTaskResponseBodyData) SetTaskId(v int64) *CreateSyntheticTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateSyntheticTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
