// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetClientBalanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterGetClientBalanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterGetClientBalanceResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterGetClientBalanceResponseBody) *ModelRouterGetClientBalanceResponse
	GetBody() *ModelRouterGetClientBalanceResponseBody
}

type ModelRouterGetClientBalanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterGetClientBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterGetClientBalanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetClientBalanceResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterGetClientBalanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterGetClientBalanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterGetClientBalanceResponse) GetBody() *ModelRouterGetClientBalanceResponseBody {
	return s.Body
}

func (s *ModelRouterGetClientBalanceResponse) SetHeaders(v map[string]*string) *ModelRouterGetClientBalanceResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterGetClientBalanceResponse) SetStatusCode(v int32) *ModelRouterGetClientBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterGetClientBalanceResponse) SetBody(v *ModelRouterGetClientBalanceResponseBody) *ModelRouterGetClientBalanceResponse {
	s.Body = v
	return s
}

func (s *ModelRouterGetClientBalanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
