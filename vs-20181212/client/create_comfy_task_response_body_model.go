// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateComfyTaskResponseBody
	GetCode() *int64
	SetMessage(v string) *CreateComfyTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateComfyTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateComfyTaskResponseBody
	GetTaskId() *string
}

type CreateComfyTaskResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6c8234f4-d1e1-4cea-b08b-7926fbdea144
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateComfyTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComfyTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateComfyTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateComfyTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComfyTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateComfyTaskResponseBody) SetCode(v int64) *CreateComfyTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComfyTaskResponseBody) SetMessage(v string) *CreateComfyTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateComfyTaskResponseBody) SetRequestId(v string) *CreateComfyTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComfyTaskResponseBody) SetTaskId(v string) *CreateComfyTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateComfyTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
