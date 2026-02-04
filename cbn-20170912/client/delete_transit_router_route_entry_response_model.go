// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterRouteEntryResponseBody) *DeleteTransitRouterRouteEntryResponse
	GetBody() *DeleteTransitRouterRouteEntryResponseBody
}

type DeleteTransitRouterRouteEntryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterRouteEntryResponse) GetBody() *DeleteTransitRouterRouteEntryResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterRouteEntryResponse) SetStatusCode(v int32) *DeleteTransitRouterRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryResponse) SetBody(v *DeleteTransitRouterRouteEntryResponseBody) *DeleteTransitRouterRouteEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
