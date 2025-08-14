// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterMulticastDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterMulticastDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterMulticastDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterMulticastDomainResponseBody) *DeleteTransitRouterMulticastDomainResponse
	GetBody() *DeleteTransitRouterMulticastDomainResponseBody
}

type DeleteTransitRouterMulticastDomainResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterMulticastDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterMulticastDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterMulticastDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterMulticastDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterMulticastDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterMulticastDomainResponse) GetBody() *DeleteTransitRouterMulticastDomainResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterMulticastDomainResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterMulticastDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterMulticastDomainResponse) SetStatusCode(v int32) *DeleteTransitRouterMulticastDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainResponse) SetBody(v *DeleteTransitRouterMulticastDomainResponseBody) *DeleteTransitRouterMulticastDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterMulticastDomainResponse) Validate() error {
	return dara.Validate(s)
}
