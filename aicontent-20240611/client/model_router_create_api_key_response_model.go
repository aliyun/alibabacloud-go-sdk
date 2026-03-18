// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateApiKeyResponseBody) *ModelRouterCreateApiKeyResponse
	GetBody() *ModelRouterCreateApiKeyResponseBody
}

type ModelRouterCreateApiKeyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateApiKeyResponse) GetBody() *ModelRouterCreateApiKeyResponseBody {
	return s.Body
}

func (s *ModelRouterCreateApiKeyResponse) SetHeaders(v map[string]*string) *ModelRouterCreateApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateApiKeyResponse) SetStatusCode(v int32) *ModelRouterCreateApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateApiKeyResponse) SetBody(v *ModelRouterCreateApiKeyResponseBody) *ModelRouterCreateApiKeyResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
