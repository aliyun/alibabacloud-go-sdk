// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryApiKeyResponseBody) *ModelRouterQueryApiKeyResponse
	GetBody() *ModelRouterQueryApiKeyResponseBody
}

type ModelRouterQueryApiKeyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryApiKeyResponse) GetBody() *ModelRouterQueryApiKeyResponseBody {
	return s.Body
}

func (s *ModelRouterQueryApiKeyResponse) SetHeaders(v map[string]*string) *ModelRouterQueryApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryApiKeyResponse) SetStatusCode(v int32) *ModelRouterQueryApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryApiKeyResponse) SetBody(v *ModelRouterQueryApiKeyResponseBody) *ModelRouterQueryApiKeyResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
