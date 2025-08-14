// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateTransitRouterMulticastDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateTransitRouterMulticastDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateTransitRouterMulticastDomainResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateTransitRouterMulticastDomainResponseBody) *DisassociateTransitRouterMulticastDomainResponse
	GetBody() *DisassociateTransitRouterMulticastDomainResponseBody
}

type DisassociateTransitRouterMulticastDomainResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateTransitRouterMulticastDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateTransitRouterMulticastDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateTransitRouterMulticastDomainResponse) GoString() string {
	return s.String()
}

func (s *DisassociateTransitRouterMulticastDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateTransitRouterMulticastDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateTransitRouterMulticastDomainResponse) GetBody() *DisassociateTransitRouterMulticastDomainResponseBody {
	return s.Body
}

func (s *DisassociateTransitRouterMulticastDomainResponse) SetHeaders(v map[string]*string) *DisassociateTransitRouterMulticastDomainResponse {
	s.Headers = v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainResponse) SetStatusCode(v int32) *DisassociateTransitRouterMulticastDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainResponse) SetBody(v *DisassociateTransitRouterMulticastDomainResponseBody) *DisassociateTransitRouterMulticastDomainResponse {
	s.Body = v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainResponse) Validate() error {
	return dara.Validate(s)
}
