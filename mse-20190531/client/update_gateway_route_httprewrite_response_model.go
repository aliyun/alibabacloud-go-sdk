// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteHTTPRewriteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteHTTPRewriteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteHTTPRewriteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteHTTPRewriteResponseBody) *UpdateGatewayRouteHTTPRewriteResponse
	GetBody() *UpdateGatewayRouteHTTPRewriteResponseBody
}

type UpdateGatewayRouteHTTPRewriteResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteHTTPRewriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteHTTPRewriteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteHTTPRewriteResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) GetBody() *UpdateGatewayRouteHTTPRewriteResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteHTTPRewriteResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) SetStatusCode(v int32) *UpdateGatewayRouteHTTPRewriteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) SetBody(v *UpdateGatewayRouteHTTPRewriteResponseBody) *UpdateGatewayRouteHTTPRewriteResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) Validate() error {
	return dara.Validate(s)
}
