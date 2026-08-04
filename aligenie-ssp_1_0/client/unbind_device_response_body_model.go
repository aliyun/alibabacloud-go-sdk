// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UnbindDeviceResponseBody
	GetCode() *int32
	SetMessage(v string) *UnbindDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindDeviceResponseBody
	GetRequestId() *string
	SetResult(v bool) *UnbindDeviceResponseBody
	GetResult() *bool
}

type UnbindDeviceResponseBody struct {
	// The error code returned. A value of 200 indicates that the call succeeded.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return result of invoking this API.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the device was detached successfully.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UnbindDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UnbindDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindDeviceResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UnbindDeviceResponseBody) SetCode(v int32) *UnbindDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetMessage(v string) *UnbindDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetRequestId(v string) *UnbindDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetResult(v bool) *UnbindDeviceResponseBody {
	s.Result = &v
	return s
}

func (s *UnbindDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
