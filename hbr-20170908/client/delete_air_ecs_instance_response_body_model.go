// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAirEcsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAirEcsInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAirEcsInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAirEcsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAirEcsInstanceResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *DeleteAirEcsInstanceResponseBody
	GetTaskId() *string
}

type DeleteAirEcsInstanceResponseBody struct {
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
	// Id of the request
	//
	// example:
	//
	// 33AA3AAE-89E1-5D3A-A51D-0C0A80850F68
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
	// t-*********************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteAirEcsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirEcsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAirEcsInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAirEcsInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAirEcsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAirEcsInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAirEcsInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteAirEcsInstanceResponseBody) SetCode(v string) *DeleteAirEcsInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAirEcsInstanceResponseBody) SetMessage(v string) *DeleteAirEcsInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAirEcsInstanceResponseBody) SetRequestId(v string) *DeleteAirEcsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAirEcsInstanceResponseBody) SetSuccess(v bool) *DeleteAirEcsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAirEcsInstanceResponseBody) SetTaskId(v string) *DeleteAirEcsInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteAirEcsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
