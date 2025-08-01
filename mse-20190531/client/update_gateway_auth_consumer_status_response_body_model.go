// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayAuthConsumerStatusResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateGatewayAuthConsumerStatusResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateGatewayAuthConsumerStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateGatewayAuthConsumerStatusResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *UpdateGatewayAuthConsumerStatusResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdateGatewayAuthConsumerStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayAuthConsumerStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayAuthConsumerStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayAuthConsumerStatusResponseBody
	GetSuccess() *bool
}

type UpdateGatewayAuthConsumerStatusResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// > If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 56D9E600-6348-4260-B35F-583413F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayAuthConsumerStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetCode(v int32) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetData(v bool) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetDynamicCode(v string) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetDynamicMessage(v string) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetErrorCode(v string) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetMessage(v string) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetRequestId(v string) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) SetSuccess(v bool) *UpdateGatewayAuthConsumerStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayAuthConsumerStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
