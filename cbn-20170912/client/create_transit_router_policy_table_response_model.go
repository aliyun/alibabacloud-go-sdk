// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPolicyTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterPolicyTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterPolicyTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterPolicyTableResponseBody) *CreateTransitRouterPolicyTableResponse
	GetBody() *CreateTransitRouterPolicyTableResponseBody
}

type CreateTransitRouterPolicyTableResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterPolicyTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterPolicyTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPolicyTableResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPolicyTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterPolicyTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterPolicyTableResponse) GetBody() *CreateTransitRouterPolicyTableResponseBody {
	return s.Body
}

func (s *CreateTransitRouterPolicyTableResponse) SetHeaders(v map[string]*string) *CreateTransitRouterPolicyTableResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterPolicyTableResponse) SetStatusCode(v int32) *CreateTransitRouterPolicyTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterPolicyTableResponse) SetBody(v *CreateTransitRouterPolicyTableResponseBody) *CreateTransitRouterPolicyTableResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterPolicyTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
