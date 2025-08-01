// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayAuthConsumerResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayAuthConsumerResponseBody
	GetData() *int64
	SetDynamicCode(v string) *UpdateGatewayAuthConsumerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateGatewayAuthConsumerResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *UpdateGatewayAuthConsumerResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdateGatewayAuthConsumerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayAuthConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayAuthConsumerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayAuthConsumerResponseBody
	GetSuccess() *bool
}

type UpdateGatewayAuthConsumerResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the consumer.
	//
	// example:
	//
	// 2
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
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
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
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

func (s UpdateGatewayAuthConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayAuthConsumerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetCode(v int32) *UpdateGatewayAuthConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetData(v int64) *UpdateGatewayAuthConsumerResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetDynamicCode(v string) *UpdateGatewayAuthConsumerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetDynamicMessage(v string) *UpdateGatewayAuthConsumerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetErrorCode(v string) *UpdateGatewayAuthConsumerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayAuthConsumerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetMessage(v string) *UpdateGatewayAuthConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetRequestId(v string) *UpdateGatewayAuthConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) SetSuccess(v bool) *UpdateGatewayAuthConsumerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayAuthConsumerResponseBody) Validate() error {
	return dara.Validate(s)
}
