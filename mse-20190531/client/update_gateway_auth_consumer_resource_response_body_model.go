// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayAuthConsumerResourceResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateGatewayAuthConsumerResourceResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateGatewayAuthConsumerResourceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateGatewayAuthConsumerResourceResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *UpdateGatewayAuthConsumerResourceResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdateGatewayAuthConsumerResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayAuthConsumerResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayAuthConsumerResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayAuthConsumerResourceResponseBody
	GetSuccess() *bool
}

type UpdateGatewayAuthConsumerResourceResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the list of the resources was updated. Valid values:
	//
	// 	- true: The list of the resources was updated.
	//
	// 	- false: The list of the resources was not updated.
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
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned.
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
	// The request ID.
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

func (s UpdateGatewayAuthConsumerResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetCode(v int32) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetData(v bool) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetDynamicCode(v string) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetDynamicMessage(v string) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetErrorCode(v string) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetMessage(v string) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetRequestId(v string) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) SetSuccess(v bool) *UpdateGatewayAuthConsumerResourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
