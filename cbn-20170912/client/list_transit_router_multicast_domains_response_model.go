// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterMulticastDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterMulticastDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterMulticastDomainsResponseBody) *ListTransitRouterMulticastDomainsResponse
	GetBody() *ListTransitRouterMulticastDomainsResponseBody
}

type ListTransitRouterMulticastDomainsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterMulticastDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterMulticastDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterMulticastDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterMulticastDomainsResponse) GetBody() *ListTransitRouterMulticastDomainsResponseBody {
	return s.Body
}

func (s *ListTransitRouterMulticastDomainsResponse) SetHeaders(v map[string]*string) *ListTransitRouterMulticastDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponse) SetStatusCode(v int32) *ListTransitRouterMulticastDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponse) SetBody(v *ListTransitRouterMulticastDomainsResponseBody) *ListTransitRouterMulticastDomainsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
