// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayRouteResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayRouteResponseBody
	GetData() *int64
	SetErrorCode(v string) *UpdateGatewayRouteResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *UpdateGatewayRouteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayRouteResponseBody
	GetSuccess() *bool
}

type UpdateGatewayRouteResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 337
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// NoPermission
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
	// 73785414-2C5B-5E28-BD91-B2C742269C48
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

func (s UpdateGatewayRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayRouteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGatewayRouteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayRouteResponseBody) SetCode(v int32) *UpdateGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) SetData(v int64) *UpdateGatewayRouteResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) SetErrorCode(v string) *UpdateGatewayRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) SetMessage(v string) *UpdateGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) SetRequestId(v string) *UpdateGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) SetSuccess(v bool) *UpdateGatewayRouteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
