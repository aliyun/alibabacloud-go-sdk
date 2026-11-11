// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateTransitRouterPolicyTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateTransitRouterPolicyTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateTransitRouterPolicyTableResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateTransitRouterPolicyTableResponseBody) *DisassociateTransitRouterPolicyTableResponse
	GetBody() *DisassociateTransitRouterPolicyTableResponseBody
}

type DisassociateTransitRouterPolicyTableResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateTransitRouterPolicyTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateTransitRouterPolicyTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateTransitRouterPolicyTableResponse) GoString() string {
	return s.String()
}

func (s *DisassociateTransitRouterPolicyTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateTransitRouterPolicyTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateTransitRouterPolicyTableResponse) GetBody() *DisassociateTransitRouterPolicyTableResponseBody {
	return s.Body
}

func (s *DisassociateTransitRouterPolicyTableResponse) SetHeaders(v map[string]*string) *DisassociateTransitRouterPolicyTableResponse {
	s.Headers = v
	return s
}

func (s *DisassociateTransitRouterPolicyTableResponse) SetStatusCode(v int32) *DisassociateTransitRouterPolicyTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateTransitRouterPolicyTableResponse) SetBody(v *DisassociateTransitRouterPolicyTableResponseBody) *DisassociateTransitRouterPolicyTableResponse {
	s.Body = v
	return s
}

func (s *DisassociateTransitRouterPolicyTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
