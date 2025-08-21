// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateScheduleTaskResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateScheduleTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateScheduleTaskResponseBody
	GetRequestId() *string
	SetResult(v int64) *CreateScheduleTaskResponseBody
	GetResult() *int64
}

type CreateScheduleTaskResponseBody struct {
	// example:
	//
	// 200
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F7E2****B7C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1234567
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateScheduleTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduleTaskResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *CreateScheduleTaskResponseBody) SetCode(v int32) *CreateScheduleTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScheduleTaskResponseBody) SetMessage(v string) *CreateScheduleTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScheduleTaskResponseBody) SetRequestId(v string) *CreateScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleTaskResponseBody) SetResult(v int64) *CreateScheduleTaskResponseBody {
	s.Result = &v
	return s
}

func (s *CreateScheduleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
