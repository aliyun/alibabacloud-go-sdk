// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterMulticastDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterMulticastDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterMulticastDomainResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterMulticastDomainResponseBody) *CreateTransitRouterMulticastDomainResponse
	GetBody() *CreateTransitRouterMulticastDomainResponseBody
}

type CreateTransitRouterMulticastDomainResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterMulticastDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterMulticastDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterMulticastDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterMulticastDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterMulticastDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterMulticastDomainResponse) GetBody() *CreateTransitRouterMulticastDomainResponseBody {
	return s.Body
}

func (s *CreateTransitRouterMulticastDomainResponse) SetHeaders(v map[string]*string) *CreateTransitRouterMulticastDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterMulticastDomainResponse) SetStatusCode(v int32) *CreateTransitRouterMulticastDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainResponse) SetBody(v *CreateTransitRouterMulticastDomainResponseBody) *CreateTransitRouterMulticastDomainResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterMulticastDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
