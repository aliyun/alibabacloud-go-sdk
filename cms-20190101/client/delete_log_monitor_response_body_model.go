// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLogMonitorResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteLogMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLogMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLogMonitorResponseBody
	GetSuccess() *bool
}

type DeleteLogMonitorResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// metric not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 42BFFC2B-5E4D-4FDE-BCC6-E91EE33C5967
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
}

func (s DeleteLogMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLogMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLogMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLogMonitorResponseBody) SetCode(v string) *DeleteLogMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLogMonitorResponseBody) SetMessage(v string) *DeleteLogMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLogMonitorResponseBody) SetRequestId(v string) *DeleteLogMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogMonitorResponseBody) SetSuccess(v bool) *DeleteLogMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLogMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
