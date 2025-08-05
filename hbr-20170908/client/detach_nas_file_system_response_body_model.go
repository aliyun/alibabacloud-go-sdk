// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachNasFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetachNasFileSystemResponseBody
	GetCode() *string
	SetMessage(v string) *DetachNasFileSystemResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachNasFileSystemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetachNasFileSystemResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *DetachNasFileSystemResponseBody
	GetTaskId() *string
}

type DetachNasFileSystemResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the request is successful, a value of successful is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous job. You can call the DescribeTask operation to query the execution result of the asynchronous job.
	//
	// example:
	//
	// t-*********************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetachNasFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachNasFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *DetachNasFileSystemResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetachNasFileSystemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachNasFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachNasFileSystemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetachNasFileSystemResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DetachNasFileSystemResponseBody) SetCode(v string) *DetachNasFileSystemResponseBody {
	s.Code = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) SetMessage(v string) *DetachNasFileSystemResponseBody {
	s.Message = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) SetRequestId(v string) *DetachNasFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) SetSuccess(v bool) *DetachNasFileSystemResponseBody {
	s.Success = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) SetTaskId(v string) *DetachNasFileSystemResponseBody {
	s.TaskId = &v
	return s
}

func (s *DetachNasFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}
