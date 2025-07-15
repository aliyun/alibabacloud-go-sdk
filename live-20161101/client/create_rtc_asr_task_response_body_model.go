// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtcAsrTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRtcAsrTaskResponseBody
	GetDescription() *string
	SetRequestId(v string) *CreateRtcAsrTaskResponseBody
	GetRequestId() *string
	SetRetCode(v int64) *CreateRtcAsrTaskResponseBody
	GetRetCode() *int64
	SetTaskId(v string) *CreateRtcAsrTaskResponseBody
	GetTaskId() *string
}

type CreateRtcAsrTaskResponseBody struct {
	// The result of the request. If success is returned, the request is successful. If an error message is returned, the request failed.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7FF5417D-06E9-5A2C-9A70-581F6149E6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned HTTP status code. HTTP status code 2000 indicates that the request is successful. If another HTTP status code is returned, the request failed.
	//
	// example:
	//
	// 2000
	RetCode *int64 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// asr-a6ac15e0-9118-4b4c-9e64-306163a0****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateRtcAsrTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRtcAsrTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRtcAsrTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateRtcAsrTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRtcAsrTaskResponseBody) GetRetCode() *int64 {
	return s.RetCode
}

func (s *CreateRtcAsrTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateRtcAsrTaskResponseBody) SetDescription(v string) *CreateRtcAsrTaskResponseBody {
	s.Description = &v
	return s
}

func (s *CreateRtcAsrTaskResponseBody) SetRequestId(v string) *CreateRtcAsrTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRtcAsrTaskResponseBody) SetRetCode(v int64) *CreateRtcAsrTaskResponseBody {
	s.RetCode = &v
	return s
}

func (s *CreateRtcAsrTaskResponseBody) SetTaskId(v string) *CreateRtcAsrTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateRtcAsrTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
