// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopOnlineEvalTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopOnlineEvalTaskResponseBody
	GetCode() *string
	SetMessage(v string) *StopOnlineEvalTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopOnlineEvalTaskResponseBody
	GetRequestId() *string
}

type StopOnlineEvalTaskResponseBody struct {
	// Internal error code. Set only when the response is in error.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response error message. Set only when the response is in error.
	//
	// example:
	//
	// task id is empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the POP request
	//
	// example:
	//
	// 31E5FBC2-792D-5B5C-A5EB-3019984ABFC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopOnlineEvalTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopOnlineEvalTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopOnlineEvalTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopOnlineEvalTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopOnlineEvalTaskResponseBody) SetCode(v string) *StopOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopOnlineEvalTaskResponseBody) SetMessage(v string) *StopOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopOnlineEvalTaskResponseBody) SetRequestId(v string) *StopOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopOnlineEvalTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
