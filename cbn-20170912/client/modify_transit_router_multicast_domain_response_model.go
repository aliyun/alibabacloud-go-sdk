// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterMulticastDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTransitRouterMulticastDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTransitRouterMulticastDomainResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTransitRouterMulticastDomainResponseBody) *ModifyTransitRouterMulticastDomainResponse
	GetBody() *ModifyTransitRouterMulticastDomainResponseBody
}

type ModifyTransitRouterMulticastDomainResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTransitRouterMulticastDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTransitRouterMulticastDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterMulticastDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterMulticastDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTransitRouterMulticastDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTransitRouterMulticastDomainResponse) GetBody() *ModifyTransitRouterMulticastDomainResponseBody {
	return s.Body
}

func (s *ModifyTransitRouterMulticastDomainResponse) SetHeaders(v map[string]*string) *ModifyTransitRouterMulticastDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyTransitRouterMulticastDomainResponse) SetStatusCode(v int32) *ModifyTransitRouterMulticastDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainResponse) SetBody(v *ModifyTransitRouterMulticastDomainResponseBody) *ModifyTransitRouterMulticastDomainResponse {
	s.Body = v
	return s
}

func (s *ModifyTransitRouterMulticastDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
