// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateMonitorGroupInstancesResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateMonitorGroupInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMonitorGroupInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMonitorGroupInstancesResponseBody
	GetSuccess() *bool
}

type CreateMonitorGroupInstancesResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1BC69FEB-56CD-4555-A0E2-02536A24A946
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

func (s CreateMonitorGroupInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateMonitorGroupInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMonitorGroupInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMonitorGroupInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMonitorGroupInstancesResponseBody) SetCode(v int32) *CreateMonitorGroupInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorGroupInstancesResponseBody) SetMessage(v string) *CreateMonitorGroupInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorGroupInstancesResponseBody) SetRequestId(v string) *CreateMonitorGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorGroupInstancesResponseBody) SetSuccess(v bool) *CreateMonitorGroupInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMonitorGroupInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
