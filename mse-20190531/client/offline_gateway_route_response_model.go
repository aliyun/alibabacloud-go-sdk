// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineGatewayRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineGatewayRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineGatewayRouteResponse
	GetStatusCode() *int32
	SetBody(v *OfflineGatewayRouteResponseBody) *OfflineGatewayRouteResponse
	GetBody() *OfflineGatewayRouteResponseBody
}

type OfflineGatewayRouteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineGatewayRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineGatewayRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineGatewayRouteResponse) GoString() string {
	return s.String()
}

func (s *OfflineGatewayRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineGatewayRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineGatewayRouteResponse) GetBody() *OfflineGatewayRouteResponseBody {
	return s.Body
}

func (s *OfflineGatewayRouteResponse) SetHeaders(v map[string]*string) *OfflineGatewayRouteResponse {
	s.Headers = v
	return s
}

func (s *OfflineGatewayRouteResponse) SetStatusCode(v int32) *OfflineGatewayRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineGatewayRouteResponse) SetBody(v *OfflineGatewayRouteResponseBody) *OfflineGatewayRouteResponse {
	s.Body = v
	return s
}

func (s *OfflineGatewayRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
