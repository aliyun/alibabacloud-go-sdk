// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotExpandPreCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsg(v string) *SubmitHotExpandPreCheckTaskResponseBody
	GetMsg() *string
	SetRequestId(v string) *SubmitHotExpandPreCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitHotExpandPreCheckTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *SubmitHotExpandPreCheckTaskResponseBody
	GetTaskId() *int64
}

type SubmitHotExpandPreCheckTaskResponseBody struct {
	// The result of the task.
	//
	// example:
	//
	// scucess
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FE104D26-AC19-49B5-AC67-947F69*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 11111
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitHotExpandPreCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetMsg(v string) *SubmitHotExpandPreCheckTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetRequestId(v string) *SubmitHotExpandPreCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetSuccess(v bool) *SubmitHotExpandPreCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) SetTaskId(v int64) *SubmitHotExpandPreCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
