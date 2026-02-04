// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterRouteEntryResponseBody) *UpdateTransitRouterRouteEntryResponse
	GetBody() *UpdateTransitRouterRouteEntryResponseBody
}

type UpdateTransitRouterRouteEntryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterRouteEntryResponse) GetBody() *UpdateTransitRouterRouteEntryResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterRouteEntryResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterRouteEntryResponse) SetStatusCode(v int32) *UpdateTransitRouterRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryResponse) SetBody(v *UpdateTransitRouterRouteEntryResponseBody) *UpdateTransitRouterRouteEntryResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
