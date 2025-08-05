// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UninstallClientResponseBody
	GetCode() *string
	SetMessage(v string) *UninstallClientResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallClientResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UninstallClientResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *UninstallClientResponseBody
	GetTaskId() *string
}

type UninstallClientResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 048A2164-3732-5DF5-88B5-F97FA56DAEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous job. You can call the DescribeTask operation to query the execution result of an asynchronous job.
	//
	// example:
	//
	// t-0009qs5qcnvuvqrl2mxl
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UninstallClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallClientResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallClientResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallClientResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallClientResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UninstallClientResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallClientResponseBody) SetCode(v string) *UninstallClientResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallClientResponseBody) SetMessage(v string) *UninstallClientResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallClientResponseBody) SetRequestId(v string) *UninstallClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallClientResponseBody) SetSuccess(v bool) *UninstallClientResponseBody {
	s.Success = &v
	return s
}

func (s *UninstallClientResponseBody) SetTaskId(v string) *UninstallClientResponseBody {
	s.TaskId = &v
	return s
}

func (s *UninstallClientResponseBody) Validate() error {
	return dara.Validate(s)
}
