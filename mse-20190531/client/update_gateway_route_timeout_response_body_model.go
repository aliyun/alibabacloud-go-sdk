// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteTimeoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayRouteTimeoutResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayRouteTimeoutResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayRouteTimeoutResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayRouteTimeoutResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayRouteTimeoutResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayRouteTimeoutResponseBody
	GetSuccess() *bool
}

type UpdateGatewayRouteTimeoutResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the node.
	//
	// example:
	//
	// 400
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5359729E-A059-5942-98AF-E981E89C99FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayRouteTimeoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteTimeoutResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteTimeoutResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteTimeoutResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayRouteTimeoutResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayRouteTimeoutResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayRouteTimeoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteTimeoutResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayRouteTimeoutResponseBody) SetCode(v int32) *UpdateGatewayRouteTimeoutResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponseBody) SetData(v int64) *UpdateGatewayRouteTimeoutResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteTimeoutResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponseBody) SetMessage(v string) *UpdateGatewayRouteTimeoutResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponseBody) SetRequestId(v string) *UpdateGatewayRouteTimeoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponseBody) SetSuccess(v bool) *UpdateGatewayRouteTimeoutResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutResponseBody) Validate() error {
	return dara.Validate(s)
}
