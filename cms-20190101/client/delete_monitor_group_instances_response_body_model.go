// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteMonitorGroupInstancesResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteMonitorGroupInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMonitorGroupInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMonitorGroupInstancesResponseBody
	GetSuccess() *bool
}

type DeleteMonitorGroupInstancesResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5351B0F2-26A9-4BC9-87FF-1B74034D12C3
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

func (s DeleteMonitorGroupInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupInstancesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteMonitorGroupInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMonitorGroupInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMonitorGroupInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMonitorGroupInstancesResponseBody) SetCode(v int32) *DeleteMonitorGroupInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitorGroupInstancesResponseBody) SetMessage(v string) *DeleteMonitorGroupInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitorGroupInstancesResponseBody) SetRequestId(v string) *DeleteMonitorGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupInstancesResponseBody) SetSuccess(v bool) *DeleteMonitorGroupInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMonitorGroupInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
