// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOtaTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApproveOtaTaskResponseBody
	GetCode() *string
	SetMessage(v string) *ApproveOtaTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApproveOtaTaskResponseBody
	GetRequestId() *string
}

type ApproveOtaTaskResponseBody struct {
	// The execution result. If the request was successful, `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// OtaTask.Running
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. If the value of `Code` is `success`, this parameter is not returned.
	//
	// example:
	//
	// The task is running and cannot be sumitted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveOtaTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApproveOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApproveOtaTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApproveOtaTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApproveOtaTaskResponseBody) SetCode(v string) *ApproveOtaTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) SetMessage(v string) *ApproveOtaTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) SetRequestId(v string) *ApproveOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
