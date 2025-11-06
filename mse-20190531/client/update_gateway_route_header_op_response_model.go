// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteHeaderOpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayRouteHeaderOpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayRouteHeaderOpResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayRouteHeaderOpResponseBody) *UpdateGatewayRouteHeaderOpResponse
	GetBody() *UpdateGatewayRouteHeaderOpResponseBody
}

type UpdateGatewayRouteHeaderOpResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayRouteHeaderOpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayRouteHeaderOpResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteHeaderOpResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHeaderOpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayRouteHeaderOpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayRouteHeaderOpResponse) GetBody() *UpdateGatewayRouteHeaderOpResponseBody {
	return s.Body
}

func (s *UpdateGatewayRouteHeaderOpResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteHeaderOpResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponse) SetStatusCode(v int32) *UpdateGatewayRouteHeaderOpResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponse) SetBody(v *UpdateGatewayRouteHeaderOpResponseBody) *UpdateGatewayRouteHeaderOpResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayRouteHeaderOpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
