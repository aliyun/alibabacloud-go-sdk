// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainVSwitchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterMulticastDomainVSwitchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterMulticastDomainVSwitchesResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterMulticastDomainVSwitchesResponseBody) *ListTransitRouterMulticastDomainVSwitchesResponse
	GetBody() *ListTransitRouterMulticastDomainVSwitchesResponseBody
}

type ListTransitRouterMulticastDomainVSwitchesResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterMulticastDomainVSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterMulticastDomainVSwitchesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainVSwitchesResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponse) GetBody() *ListTransitRouterMulticastDomainVSwitchesResponseBody {
	return s.Body
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponse) SetHeaders(v map[string]*string) *ListTransitRouterMulticastDomainVSwitchesResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponse) SetStatusCode(v int32) *ListTransitRouterMulticastDomainVSwitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponse) SetBody(v *ListTransitRouterMulticastDomainVSwitchesResponseBody) *ListTransitRouterMulticastDomainVSwitchesResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponse) Validate() error {
	return dara.Validate(s)
}
