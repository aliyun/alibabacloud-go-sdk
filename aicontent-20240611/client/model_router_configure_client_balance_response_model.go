// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterConfigureClientBalanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterConfigureClientBalanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterConfigureClientBalanceResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterConfigureClientBalanceResponseBody) *ModelRouterConfigureClientBalanceResponse
	GetBody() *ModelRouterConfigureClientBalanceResponseBody
}

type ModelRouterConfigureClientBalanceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterConfigureClientBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterConfigureClientBalanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterConfigureClientBalanceResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterConfigureClientBalanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterConfigureClientBalanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterConfigureClientBalanceResponse) GetBody() *ModelRouterConfigureClientBalanceResponseBody {
	return s.Body
}

func (s *ModelRouterConfigureClientBalanceResponse) SetHeaders(v map[string]*string) *ModelRouterConfigureClientBalanceResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponse) SetStatusCode(v int32) *ModelRouterConfigureClientBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponse) SetBody(v *ModelRouterConfigureClientBalanceResponseBody) *ModelRouterConfigureClientBalanceResponse {
	s.Body = v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
