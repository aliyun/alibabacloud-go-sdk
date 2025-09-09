// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmoothExpandPreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsg(v string) *SubmitSmoothExpandPreCheckResponseBody
	GetMsg() *string
	SetRequestId(v string) *SubmitSmoothExpandPreCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSmoothExpandPreCheckResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *SubmitSmoothExpandPreCheckResponseBody
	GetTaskId() *int64
}

type SubmitSmoothExpandPreCheckResponseBody struct {
	// The result of the precheck task.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FE104D26-AC19-49B5-AC67-947F69******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the precheck task.
	//
	// example:
	//
	// 11111
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SubmitSmoothExpandPreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSmoothExpandPreCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSmoothExpandPreCheckResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetMsg(v string) *SubmitSmoothExpandPreCheckResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetRequestId(v string) *SubmitSmoothExpandPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetSuccess(v bool) *SubmitSmoothExpandPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) SetTaskId(v int64) *SubmitSmoothExpandPreCheckResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
