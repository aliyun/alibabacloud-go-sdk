// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteScheduleTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteScheduleTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteScheduleTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteScheduleTaskResponseBody
	GetResult() *bool
}

type DeleteScheduleTaskResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F7E2****B7C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteScheduleTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScheduleTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteScheduleTaskResponseBody) SetCode(v string) *DeleteScheduleTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScheduleTaskResponseBody) SetMessage(v string) *DeleteScheduleTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScheduleTaskResponseBody) SetRequestId(v string) *DeleteScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduleTaskResponseBody) SetResult(v bool) *DeleteScheduleTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteScheduleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
