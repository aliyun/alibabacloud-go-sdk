// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutePrivateZoneInCenToVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RoutePrivateZoneInCenToVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RoutePrivateZoneInCenToVpcResponse
	GetStatusCode() *int32
	SetBody(v *RoutePrivateZoneInCenToVpcResponseBody) *RoutePrivateZoneInCenToVpcResponse
	GetBody() *RoutePrivateZoneInCenToVpcResponseBody
}

type RoutePrivateZoneInCenToVpcResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RoutePrivateZoneInCenToVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RoutePrivateZoneInCenToVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s RoutePrivateZoneInCenToVpcResponse) GoString() string {
	return s.String()
}

func (s *RoutePrivateZoneInCenToVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RoutePrivateZoneInCenToVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RoutePrivateZoneInCenToVpcResponse) GetBody() *RoutePrivateZoneInCenToVpcResponseBody {
	return s.Body
}

func (s *RoutePrivateZoneInCenToVpcResponse) SetHeaders(v map[string]*string) *RoutePrivateZoneInCenToVpcResponse {
	s.Headers = v
	return s
}

func (s *RoutePrivateZoneInCenToVpcResponse) SetStatusCode(v int32) *RoutePrivateZoneInCenToVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcResponse) SetBody(v *RoutePrivateZoneInCenToVpcResponseBody) *RoutePrivateZoneInCenToVpcResponse {
	s.Body = v
	return s
}

func (s *RoutePrivateZoneInCenToVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
