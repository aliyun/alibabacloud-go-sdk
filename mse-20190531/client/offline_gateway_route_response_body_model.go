// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineGatewayRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OfflineGatewayRouteResponseBody
	GetCode() *int32
	SetData(v bool) *OfflineGatewayRouteResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *OfflineGatewayRouteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OfflineGatewayRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *OfflineGatewayRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OfflineGatewayRouteResponseBody
	GetSuccess() *bool
}

type OfflineGatewayRouteResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the route is unpublished.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
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
	// 8EB24742-AB15-5953-BEFB-1FA18E770EEE
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

func (s OfflineGatewayRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineGatewayRouteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OfflineGatewayRouteResponseBody) GetData() *bool {
	return s.Data
}

func (s *OfflineGatewayRouteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OfflineGatewayRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OfflineGatewayRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineGatewayRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflineGatewayRouteResponseBody) SetCode(v int32) *OfflineGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineGatewayRouteResponseBody) SetData(v bool) *OfflineGatewayRouteResponseBody {
	s.Data = &v
	return s
}

func (s *OfflineGatewayRouteResponseBody) SetHttpStatusCode(v int32) *OfflineGatewayRouteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OfflineGatewayRouteResponseBody) SetMessage(v string) *OfflineGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineGatewayRouteResponseBody) SetRequestId(v string) *OfflineGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineGatewayRouteResponseBody) SetSuccess(v bool) *OfflineGatewayRouteResponseBody {
	s.Success = &v
	return s
}

func (s *OfflineGatewayRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
