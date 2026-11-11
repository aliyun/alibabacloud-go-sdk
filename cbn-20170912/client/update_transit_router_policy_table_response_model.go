// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPolicyTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterPolicyTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterPolicyTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterPolicyTableResponseBody) *UpdateTransitRouterPolicyTableResponse
	GetBody() *UpdateTransitRouterPolicyTableResponseBody
}

type UpdateTransitRouterPolicyTableResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterPolicyTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterPolicyTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPolicyTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPolicyTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterPolicyTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterPolicyTableResponse) GetBody() *UpdateTransitRouterPolicyTableResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterPolicyTableResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterPolicyTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterPolicyTableResponse) SetStatusCode(v int32) *UpdateTransitRouterPolicyTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterPolicyTableResponse) SetBody(v *UpdateTransitRouterPolicyTableResponseBody) *UpdateTransitRouterPolicyTableResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterPolicyTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
