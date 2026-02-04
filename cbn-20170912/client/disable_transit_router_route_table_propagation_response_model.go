// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTransitRouterRouteTablePropagationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableTransitRouterRouteTablePropagationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableTransitRouterRouteTablePropagationResponse
	GetStatusCode() *int32
	SetBody(v *DisableTransitRouterRouteTablePropagationResponseBody) *DisableTransitRouterRouteTablePropagationResponse
	GetBody() *DisableTransitRouterRouteTablePropagationResponseBody
}

type DisableTransitRouterRouteTablePropagationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableTransitRouterRouteTablePropagationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableTransitRouterRouteTablePropagationResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableTransitRouterRouteTablePropagationResponse) GoString() string {
	return s.String()
}

func (s *DisableTransitRouterRouteTablePropagationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableTransitRouterRouteTablePropagationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableTransitRouterRouteTablePropagationResponse) GetBody() *DisableTransitRouterRouteTablePropagationResponseBody {
	return s.Body
}

func (s *DisableTransitRouterRouteTablePropagationResponse) SetHeaders(v map[string]*string) *DisableTransitRouterRouteTablePropagationResponse {
	s.Headers = v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationResponse) SetStatusCode(v int32) *DisableTransitRouterRouteTablePropagationResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationResponse) SetBody(v *DisableTransitRouterRouteTablePropagationResponseBody) *DisableTransitRouterRouteTablePropagationResponse {
	s.Body = v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
