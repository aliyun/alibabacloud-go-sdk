// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopHanaDatabaseAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopHanaDatabaseAsyncResponseBody
	GetCode() *string
	SetMessage(v string) *StopHanaDatabaseAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopHanaDatabaseAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopHanaDatabaseAsyncResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *StopHanaDatabaseAsyncResponseBody
	GetTaskId() *string
}

type StopHanaDatabaseAsyncResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CD8B903B-DE8F-5969-9414-B2C634D504D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous job. You can call the DescribeTask operation to query the execution result of an asynchronous job.
	//
	// example:
	//
	// t-0007o3vqfukgd3y5bxxr
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopHanaDatabaseAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopHanaDatabaseAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *StopHanaDatabaseAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopHanaDatabaseAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopHanaDatabaseAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopHanaDatabaseAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopHanaDatabaseAsyncResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopHanaDatabaseAsyncResponseBody) SetCode(v string) *StopHanaDatabaseAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) SetMessage(v string) *StopHanaDatabaseAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) SetRequestId(v string) *StopHanaDatabaseAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) SetSuccess(v bool) *StopHanaDatabaseAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) SetTaskId(v string) *StopHanaDatabaseAsyncResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopHanaDatabaseAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
