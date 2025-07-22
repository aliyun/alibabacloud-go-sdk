// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceShutdownTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstanceShutdownTimerResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteInstanceShutdownTimerResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *DeleteInstanceShutdownTimerResponseBody
	GetInstanceId() *string
	SetMessage(v string) *DeleteInstanceShutdownTimerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstanceShutdownTimerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstanceShutdownTimerResponseBody
	GetSuccess() *bool
}

type DeleteInstanceShutdownTimerResponseBody struct {
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

func (s DeleteInstanceShutdownTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceShutdownTimerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceShutdownTimerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInstanceShutdownTimerResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceShutdownTimerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstanceShutdownTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceShutdownTimerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetCode(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceShutdownTimerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetInstanceId(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetMessage(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetRequestId(v string) *DeleteInstanceShutdownTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) SetSuccess(v bool) *DeleteInstanceShutdownTimerResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
