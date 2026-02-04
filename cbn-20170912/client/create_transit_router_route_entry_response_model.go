// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterRouteEntryResponseBody) *CreateTransitRouterRouteEntryResponse
	GetBody() *CreateTransitRouterRouteEntryResponseBody
}

type CreateTransitRouterRouteEntryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterRouteEntryResponse) GetBody() *CreateTransitRouterRouteEntryResponseBody {
	return s.Body
}

func (s *CreateTransitRouterRouteEntryResponse) SetHeaders(v map[string]*string) *CreateTransitRouterRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterRouteEntryResponse) SetStatusCode(v int32) *CreateTransitRouterRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterRouteEntryResponse) SetBody(v *CreateTransitRouterRouteEntryResponseBody) *CreateTransitRouterRouteEntryResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
