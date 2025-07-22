// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableExpressConnectRouterRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableExpressConnectRouterRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableExpressConnectRouterRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DisableExpressConnectRouterRouteEntriesResponseBody) *DisableExpressConnectRouterRouteEntriesResponse
	GetBody() *DisableExpressConnectRouterRouteEntriesResponseBody
}

type DisableExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableExpressConnectRouterRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) GetBody() *DisableExpressConnectRouterRouteEntriesResponseBody {
	return s.Body
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DisableExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DisableExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetBody(v *DisableExpressConnectRouterRouteEntriesResponseBody) *DisableExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
