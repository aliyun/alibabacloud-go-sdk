// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmoothExpandPreCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsg(v string) *SubmitSmoothExpandPreCheckTaskResponseBody
	GetMsg() *string
	SetRequestId(v string) *SubmitSmoothExpandPreCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSmoothExpandPreCheckTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *SubmitSmoothExpandPreCheckTaskResponseBody
	GetTaskId() *int64
}

type SubmitSmoothExpandPreCheckTaskResponseBody struct {
	// Indicates whether the precheck task was submitted.
	//
	// example:
	//
	// scucess
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-23ERW
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
	// 2321
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetMsg(v string) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetRequestId(v string) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetSuccess(v bool) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) SetTaskId(v int64) *SubmitSmoothExpandPreCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
