// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMonitorGroupInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyMonitorGroupInstancesResponseBody
	GetCode() *int32
	SetMessage(v string) *ModifyMonitorGroupInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMonitorGroupInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMonitorGroupInstancesResponseBody
	GetSuccess() *bool
}

type ModifyMonitorGroupInstancesResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
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
	// The ID of the request.
	//
	// example:
	//
	// FEC7EDB3-9B08-4AC0-A42A-329F5D14B95A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. The value true indicates a success. The value false indicates a failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMonitorGroupInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMonitorGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupInstancesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyMonitorGroupInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMonitorGroupInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMonitorGroupInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMonitorGroupInstancesResponseBody) SetCode(v int32) *ModifyMonitorGroupInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMonitorGroupInstancesResponseBody) SetMessage(v string) *ModifyMonitorGroupInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMonitorGroupInstancesResponseBody) SetRequestId(v string) *ModifyMonitorGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMonitorGroupInstancesResponseBody) SetSuccess(v bool) *ModifyMonitorGroupInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMonitorGroupInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
