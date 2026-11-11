// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPolicyTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterPolicyTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterPolicyTableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterPolicyTableResponseBody) *DeleteTransitRouterPolicyTableResponse
	GetBody() *DeleteTransitRouterPolicyTableResponseBody
}

type DeleteTransitRouterPolicyTableResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterPolicyTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterPolicyTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPolicyTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPolicyTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterPolicyTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterPolicyTableResponse) GetBody() *DeleteTransitRouterPolicyTableResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterPolicyTableResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterPolicyTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterPolicyTableResponse) SetStatusCode(v int32) *DeleteTransitRouterPolicyTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterPolicyTableResponse) SetBody(v *DeleteTransitRouterPolicyTableResponseBody) *DeleteTransitRouterPolicyTableResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterPolicyTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
