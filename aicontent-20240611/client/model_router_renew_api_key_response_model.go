// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterRenewApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterRenewApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterRenewApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterRenewApiKeyResponseBody) *ModelRouterRenewApiKeyResponse
	GetBody() *ModelRouterRenewApiKeyResponseBody
}

type ModelRouterRenewApiKeyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterRenewApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterRenewApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterRenewApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterRenewApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterRenewApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterRenewApiKeyResponse) GetBody() *ModelRouterRenewApiKeyResponseBody {
	return s.Body
}

func (s *ModelRouterRenewApiKeyResponse) SetHeaders(v map[string]*string) *ModelRouterRenewApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterRenewApiKeyResponse) SetStatusCode(v int32) *ModelRouterRenewApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterRenewApiKeyResponse) SetBody(v *ModelRouterRenewApiKeyResponseBody) *ModelRouterRenewApiKeyResponse {
	s.Body = v
	return s
}

func (s *ModelRouterRenewApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
