// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteWafStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteWafStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteWafStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteWafStatusResponseBody) *UpdateGatewayRouteWafStatusResponse
	GetBody() *UpdateGatewayRouteWafStatusResponseBody
}

type UpdateGatewayRouteWafStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteWafStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteWafStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteWafStatusResponse) GetBody() *UpdateGatewayRouteWafStatusResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteWafStatusResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteWafStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponse) SetStatusCode(v int32) *UpdateGatewayRouteWafStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponse) SetBody(v *UpdateGatewayRouteWafStatusResponseBody) *UpdateGatewayRouteWafStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
