// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResourceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayAuthConsumerResourceStatusResponseBody
	GetSuccess() *bool
}

type UpdateGatewayAuthConsumerResourceStatusResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the execution was successful. Valid values:
	//
	// 	- true: The execution was successful.
	//
	// 	- false: The execution failed.
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
	// The dynamic part in the error message. This parameter is used to replace the `%s` variable in the `ErrMessage` parameter.
	//
	// >  If the return value of the `ErrMessage` parameter is `The Value of Input Parameter %s is not valid` and the return value of the `DynamicMessage` parameter is `DtsJobId`, the specified `DtsJobId` parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayAuthConsumerResourceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResourceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetCode(v int32) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetData(v bool) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetDynamicCode(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetDynamicMessage(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetErrorCode(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetMessage(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetRequestId(v string) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) SetSuccess(v bool) *UpdateGatewayAuthConsumerResourceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
