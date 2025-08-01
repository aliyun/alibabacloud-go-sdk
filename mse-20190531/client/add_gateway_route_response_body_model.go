// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddGatewayRouteResponseBody
	GetCode() *int32
	SetData(v int64) *AddGatewayRouteResponseBody
	GetData() *int64
	SetErrorCode(v string) *AddGatewayRouteResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *AddGatewayRouteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddGatewayRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddGatewayRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddGatewayRouteResponseBody
	GetSuccess() *bool
}

type AddGatewayRouteResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the created route.
	//
	// example:
	//
	// 555
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code that is returned.
	//
	// 	- InvalidParameter: The parameter is invalid. For example, the parameter is incomplete or the parameter format is invalid.
	//
	// 	- IllegalRequest: The request is invalid. Unauthorized access to data initiated by parameters is identified.
	//
	// 	- NoPermission: You are not authorized to perform this operation.
	//
	// 	- NotFound: The resource does not exist and error 404 is returned.
	//
	// 	- InternalError: An internal error occurred.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 03A3E2F4-6804-5663-9D5D-2EC47A1*****
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

func (s AddGatewayRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddGatewayRouteResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddGatewayRouteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddGatewayRouteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddGatewayRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddGatewayRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGatewayRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddGatewayRouteResponseBody) SetCode(v int32) *AddGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *AddGatewayRouteResponseBody) SetData(v int64) *AddGatewayRouteResponseBody {
	s.Data = &v
	return s
}

func (s *AddGatewayRouteResponseBody) SetErrorCode(v string) *AddGatewayRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddGatewayRouteResponseBody) SetHttpStatusCode(v int32) *AddGatewayRouteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddGatewayRouteResponseBody) SetMessage(v string) *AddGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *AddGatewayRouteResponseBody) SetRequestId(v string) *AddGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGatewayRouteResponseBody) SetSuccess(v bool) *AddGatewayRouteResponseBody {
	s.Success = &v
	return s
}

func (s *AddGatewayRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
