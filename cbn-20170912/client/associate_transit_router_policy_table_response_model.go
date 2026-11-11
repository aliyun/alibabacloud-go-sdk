// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterPolicyTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateTransitRouterPolicyTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateTransitRouterPolicyTableResponse
	GetStatusCode() *int32
	SetBody(v *AssociateTransitRouterPolicyTableResponseBody) *AssociateTransitRouterPolicyTableResponse
	GetBody() *AssociateTransitRouterPolicyTableResponseBody
}

type AssociateTransitRouterPolicyTableResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateTransitRouterPolicyTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateTransitRouterPolicyTableResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterPolicyTableResponse) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterPolicyTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateTransitRouterPolicyTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateTransitRouterPolicyTableResponse) GetBody() *AssociateTransitRouterPolicyTableResponseBody {
	return s.Body
}

func (s *AssociateTransitRouterPolicyTableResponse) SetHeaders(v map[string]*string) *AssociateTransitRouterPolicyTableResponse {
	s.Headers = v
	return s
}

func (s *AssociateTransitRouterPolicyTableResponse) SetStatusCode(v int32) *AssociateTransitRouterPolicyTableResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateTransitRouterPolicyTableResponse) SetBody(v *AssociateTransitRouterPolicyTableResponseBody) *AssociateTransitRouterPolicyTableResponse {
	s.Body = v
	return s
}

func (s *AssociateTransitRouterPolicyTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
