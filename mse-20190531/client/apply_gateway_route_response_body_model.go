// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyGatewayRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ApplyGatewayRouteResponseBody
	GetCode() *int32
	SetData(v bool) *ApplyGatewayRouteResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *ApplyGatewayRouteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ApplyGatewayRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyGatewayRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyGatewayRouteResponseBody
	GetSuccess() *bool
}

type ApplyGatewayRouteResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the node.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1DB1AB36-8395-540E-AACC-7571D6B2057A
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

func (s ApplyGatewayRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyGatewayRouteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ApplyGatewayRouteResponseBody) GetData() *bool {
	return s.Data
}

func (s *ApplyGatewayRouteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ApplyGatewayRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyGatewayRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyGatewayRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyGatewayRouteResponseBody) SetCode(v int32) *ApplyGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyGatewayRouteResponseBody) SetData(v bool) *ApplyGatewayRouteResponseBody {
	s.Data = &v
	return s
}

func (s *ApplyGatewayRouteResponseBody) SetHttpStatusCode(v int32) *ApplyGatewayRouteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ApplyGatewayRouteResponseBody) SetMessage(v string) *ApplyGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyGatewayRouteResponseBody) SetRequestId(v string) *ApplyGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyGatewayRouteResponseBody) SetSuccess(v bool) *ApplyGatewayRouteResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyGatewayRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
