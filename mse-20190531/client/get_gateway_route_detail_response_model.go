// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayRouteDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayRouteDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayRouteDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayRouteDetailResponseBody) *GetGatewayRouteDetailResponse
	GetBody() *GetGatewayRouteDetailResponseBody
}

type GetGatewayRouteDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayRouteDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayRouteDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayRouteDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayRouteDetailResponse) GetBody() *GetGatewayRouteDetailResponseBody {
	return s.Body
}

func (s *GetGatewayRouteDetailResponse) SetHeaders(v map[string]*string) *GetGatewayRouteDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayRouteDetailResponse) SetStatusCode(v int32) *GetGatewayRouteDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayRouteDetailResponse) SetBody(v *GetGatewayRouteDetailResponseBody) *GetGatewayRouteDetailResponse {
	s.Body = v
	return s
}

func (s *GetGatewayRouteDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
