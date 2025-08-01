// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteHeaderOpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayRouteHeaderOpResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayRouteHeaderOpResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayRouteHeaderOpResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayRouteHeaderOpResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayRouteHeaderOpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayRouteHeaderOpResponseBody
	GetSuccess() *bool
}

type UpdateGatewayRouteHeaderOpResponseBody struct {
	// The status code returned. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 666
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7C8CB294-BED8-5F32-A416-7517725E8F27
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

func (s UpdateGatewayRouteHeaderOpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteHeaderOpResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) SetCode(v int32) *UpdateGatewayRouteHeaderOpResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) SetData(v int64) *UpdateGatewayRouteHeaderOpResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteHeaderOpResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) SetMessage(v string) *UpdateGatewayRouteHeaderOpResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) SetRequestId(v string) *UpdateGatewayRouteHeaderOpResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) SetSuccess(v bool) *UpdateGatewayRouteHeaderOpResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponseBody) Validate() error {
	return dara.Validate(s)
}
