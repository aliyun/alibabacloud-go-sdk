// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupsAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHanaBackupsAsyncResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeHanaBackupsAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHanaBackupsAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHanaBackupsAsyncResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *DescribeHanaBackupsAsyncResponseBody
	GetTaskId() *string
}

type DescribeHanaBackupsAsyncResponseBody struct {
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
	// 31F97233-8563-563D-8880-914B00EEA928
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
	// t-0006xmbplrqebt9dhkth
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeHanaBackupsAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupsAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupsAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHanaBackupsAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaBackupsAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHanaBackupsAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHanaBackupsAsyncResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetCode(v string) *DescribeHanaBackupsAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetMessage(v string) *DescribeHanaBackupsAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetRequestId(v string) *DescribeHanaBackupsAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetSuccess(v bool) *DescribeHanaBackupsAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) SetTaskId(v string) *DescribeHanaBackupsAsyncResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
