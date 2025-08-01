// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteHTTPRewriteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayRouteHTTPRewriteResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateGatewayRouteHTTPRewriteResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateGatewayRouteHTTPRewriteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayRouteHTTPRewriteResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayRouteHTTPRewriteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayRouteHTTPRewriteResponseBody
	GetSuccess() *bool
}

type UpdateGatewayRouteHTTPRewriteResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 411
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
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AE60DA1E-B234-5105-897B-68372453A5A2
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

func (s UpdateGatewayRouteHTTPRewriteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteHTTPRewriteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetCode(v int32) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetData(v int64) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetMessage(v string) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetRequestId(v string) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetSuccess(v bool) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) Validate() error {
	return dara.Validate(s)
}
