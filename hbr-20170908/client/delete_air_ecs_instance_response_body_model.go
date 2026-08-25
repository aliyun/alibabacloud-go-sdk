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
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. The value "successful" is returned for a successful request. An error message is returned for a failed request.
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
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The asynchronous task ID. You can call DescribeTask to obtain the task result.
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
