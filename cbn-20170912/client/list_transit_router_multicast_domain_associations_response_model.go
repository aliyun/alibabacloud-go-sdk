// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterMulticastDomainAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterMulticastDomainAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterMulticastDomainAssociationsResponseBody) *ListTransitRouterMulticastDomainAssociationsResponse
	GetBody() *ListTransitRouterMulticastDomainAssociationsResponseBody
}

type ListTransitRouterMulticastDomainAssociationsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterMulticastDomainAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterMulticastDomainAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterMulticastDomainAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterMulticastDomainAssociationsResponse) GetBody() *ListTransitRouterMulticastDomainAssociationsResponseBody {
	return s.Body
}

func (s *ListTransitRouterMulticastDomainAssociationsResponse) SetHeaders(v map[string]*string) *ListTransitRouterMulticastDomainAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponse) SetStatusCode(v int32) *ListTransitRouterMulticastDomainAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponse) SetBody(v *ListTransitRouterMulticastDomainAssociationsResponseBody) *ListTransitRouterMulticastDomainAssociationsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponse) Validate() error {
	return dara.Validate(s)
}
