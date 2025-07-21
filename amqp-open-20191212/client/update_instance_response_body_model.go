// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateInstanceResponseBody
	GetCode() *int32
	SetData(v interface{}) *UpdateInstanceResponseBody
	GetData() interface{}
	SetMessage(v string) *UpdateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceResponseBody
	GetRequestId() *string
	SetStatusCode(v string) *UpdateInstanceResponseBody
	GetStatusCode() *string
	SetSuccess(v string) *UpdateInstanceResponseBody
	GetSuccess() *string
}

type UpdateInstanceResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data, which includes orderId and instanceId. Sample returned data:
	//
	// ```json
	//
	// "Data": {
	//
	//     "instanceId": "amqp-cn-xxxxx",
	//
	//     "orderId": 22222
	//
	//   }
	//
	// ```
	//
	// example:
	//
	// {“instanceId”: “amqp-cn-jtexxxxx”, “orderId”: 2222222}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// InstanceNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 628705FD-03EE-4ABE-BB21-E1672960***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateInstanceResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResponseBody) GetStatusCode() *string {
	return s.StatusCode
}

func (s *UpdateInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateInstanceResponseBody) SetCode(v int32) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetData(v interface{}) *UpdateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetStatusCode(v string) *UpdateInstanceResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v string) *UpdateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
