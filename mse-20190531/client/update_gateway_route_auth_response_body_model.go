// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayRouteAuthResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayRouteAuthResponseBody
	GetData() *int64
	SetDynamicCode(v string) *UpdateGatewayRouteAuthResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateGatewayRouteAuthResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *UpdateGatewayRouteAuthResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdateGatewayRouteAuthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayRouteAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayRouteAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayRouteAuthResponseBody
	GetSuccess() *bool
}

type UpdateGatewayRouteAuthResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The route ID.
	//
	// example:
	//
	// 123
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 54973C90-F379-4372-9AA5-053A3F7****
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

func (s UpdateGatewayRouteAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteAuthResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteAuthResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteAuthResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayRouteAuthResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateGatewayRouteAuthResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateGatewayRouteAuthResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGatewayRouteAuthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayRouteAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayRouteAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayRouteAuthResponseBody) SetCode(v int32) *UpdateGatewayRouteAuthResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) SetData(v int64) *UpdateGatewayRouteAuthResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) SetDynamicCode(v string) *UpdateGatewayRouteAuthResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) SetDynamicMessage(v string) *UpdateGatewayRouteAuthResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) SetErrorCode(v string) *UpdateGatewayRouteAuthResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteAuthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) SetMessage(v string) *UpdateGatewayRouteAuthResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) SetRequestId(v string) *UpdateGatewayRouteAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) SetSuccess(v bool) *UpdateGatewayRouteAuthResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayRouteAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
