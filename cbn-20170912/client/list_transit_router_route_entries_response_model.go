// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterRouteEntriesResponseBody) *ListTransitRouterRouteEntriesResponse
	GetBody() *ListTransitRouterRouteEntriesResponseBody
}

type ListTransitRouterRouteEntriesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterRouteEntriesResponse) GetBody() *ListTransitRouterRouteEntriesResponseBody {
	return s.Body
}

func (s *ListTransitRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *ListTransitRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterRouteEntriesResponse) SetStatusCode(v int32) *ListTransitRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterRouteEntriesResponse) SetBody(v *ListTransitRouterRouteEntriesResponseBody) *ListTransitRouterRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
