// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMonitorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyMonitorGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *ModifyMonitorGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMonitorGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMonitorGroupResponseBody
	GetSuccess() *bool
}

type ModifyMonitorGroupResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C85A2870-5DF4-4269-BC50-ECB5E4591A80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMonitorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMonitorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyMonitorGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMonitorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMonitorGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMonitorGroupResponseBody) SetCode(v int32) *ModifyMonitorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMonitorGroupResponseBody) SetMessage(v string) *ModifyMonitorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMonitorGroupResponseBody) SetRequestId(v string) *ModifyMonitorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMonitorGroupResponseBody) SetSuccess(v bool) *ModifyMonitorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMonitorGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
