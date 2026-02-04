// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterMulticastDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateTransitRouterMulticastDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateTransitRouterMulticastDomainResponse
	GetStatusCode() *int32
	SetBody(v *AssociateTransitRouterMulticastDomainResponseBody) *AssociateTransitRouterMulticastDomainResponse
	GetBody() *AssociateTransitRouterMulticastDomainResponseBody
}

type AssociateTransitRouterMulticastDomainResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateTransitRouterMulticastDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateTransitRouterMulticastDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterMulticastDomainResponse) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterMulticastDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateTransitRouterMulticastDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateTransitRouterMulticastDomainResponse) GetBody() *AssociateTransitRouterMulticastDomainResponseBody {
	return s.Body
}

func (s *AssociateTransitRouterMulticastDomainResponse) SetHeaders(v map[string]*string) *AssociateTransitRouterMulticastDomainResponse {
	s.Headers = v
	return s
}

func (s *AssociateTransitRouterMulticastDomainResponse) SetStatusCode(v int32) *AssociateTransitRouterMulticastDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainResponse) SetBody(v *AssociateTransitRouterMulticastDomainResponseBody) *AssociateTransitRouterMulticastDomainResponse {
	s.Body = v
	return s
}

func (s *AssociateTransitRouterMulticastDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
