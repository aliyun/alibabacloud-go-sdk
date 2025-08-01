// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteRetryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayRouteRetryResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayRouteRetryResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayRouteRetryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayRouteRetryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayRouteRetryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayRouteRetryResponseBody
	GetSuccess() *bool
}

type UpdateGatewayRouteRetryResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 640
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
	// 9297B27D-D932-5E9F-93B9-99D6C5F3A879
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

func (s UpdateGatewayRouteRetryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRetryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRetryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteRetryResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayRouteRetryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayRouteRetryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayRouteRetryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteRetryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayRouteRetryResponseBody) SetCode(v int32) *UpdateGatewayRouteRetryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteRetryResponseBody) SetData(v int64) *UpdateGatewayRouteRetryResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayRouteRetryResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteRetryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteRetryResponseBody) SetMessage(v string) *UpdateGatewayRouteRetryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteRetryResponseBody) SetRequestId(v string) *UpdateGatewayRouteRetryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteRetryResponseBody) SetSuccess(v bool) *UpdateGatewayRouteRetryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayRouteRetryResponseBody) Validate() error {
	return dara.Validate(s)
}
