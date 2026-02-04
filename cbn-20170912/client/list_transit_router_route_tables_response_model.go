// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterRouteTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterRouteTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterRouteTablesResponseBody) *ListTransitRouterRouteTablesResponse
	GetBody() *ListTransitRouterRouteTablesResponseBody
}

type ListTransitRouterRouteTablesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterRouteTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterRouteTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablesResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterRouteTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterRouteTablesResponse) GetBody() *ListTransitRouterRouteTablesResponseBody {
	return s.Body
}

func (s *ListTransitRouterRouteTablesResponse) SetHeaders(v map[string]*string) *ListTransitRouterRouteTablesResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterRouteTablesResponse) SetStatusCode(v int32) *ListTransitRouterRouteTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponse) SetBody(v *ListTransitRouterRouteTablesResponseBody) *ListTransitRouterRouteTablesResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterRouteTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
