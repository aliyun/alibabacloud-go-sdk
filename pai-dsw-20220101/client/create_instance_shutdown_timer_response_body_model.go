// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceShutdownTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceShutdownTimerResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateInstanceShutdownTimerResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *CreateInstanceShutdownTimerResponseBody
	GetInstanceId() *string
	SetMessage(v string) *CreateInstanceShutdownTimerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceShutdownTimerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceShutdownTimerResponseBody
	GetSuccess() *bool
}

type CreateInstanceShutdownTimerResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: an internal error. All errors, except for parameter validation errors, are classified as internal errors.
	//
	// 	- ValidationError: a parameter validation error.
	//
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. Valid values:
	//
	// 	- 400
	//
	// 	- 404
	//
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response message.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
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

func (s CreateInstanceShutdownTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceShutdownTimerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceShutdownTimerResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceShutdownTimerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceShutdownTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceShutdownTimerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceShutdownTimerResponseBody) SetCode(v string) *CreateInstanceShutdownTimerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetHttpStatusCode(v int32) *CreateInstanceShutdownTimerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetInstanceId(v string) *CreateInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetMessage(v string) *CreateInstanceShutdownTimerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetRequestId(v string) *CreateInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) SetSuccess(v bool) *CreateInstanceShutdownTimerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
