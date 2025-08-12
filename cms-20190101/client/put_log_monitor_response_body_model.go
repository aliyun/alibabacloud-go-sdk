// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLogMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutLogMonitorResponseBody
	GetCode() *string
	SetLogId(v string) *PutLogMonitorResponseBody
	GetLogId() *string
	SetMessage(v string) *PutLogMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutLogMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutLogMonitorResponseBody
	GetSuccess() *bool
}

type PutLogMonitorResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the log monitoring metric.
	//
	// example:
	//
	// 16****
	LogId *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The returned message.
	//
	// 	- If the request was successful, `successful` is returned.
	//
	// 	- If the request failed, an error message is returned. Example: `alias of aggreate must be set value.`
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91561287-0802-5F9C-9BDE-404C50D41B06
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

func (s PutLogMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutLogMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *PutLogMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutLogMonitorResponseBody) GetLogId() *string {
	return s.LogId
}

func (s *PutLogMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutLogMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutLogMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutLogMonitorResponseBody) SetCode(v string) *PutLogMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *PutLogMonitorResponseBody) SetLogId(v string) *PutLogMonitorResponseBody {
	s.LogId = &v
	return s
}

func (s *PutLogMonitorResponseBody) SetMessage(v string) *PutLogMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *PutLogMonitorResponseBody) SetRequestId(v string) *PutLogMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutLogMonitorResponseBody) SetSuccess(v bool) *PutLogMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *PutLogMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
