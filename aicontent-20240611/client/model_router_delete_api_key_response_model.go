// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterDeleteApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterDeleteApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterDeleteApiKeyResponseBody) *ModelRouterDeleteApiKeyResponse
	GetBody() *ModelRouterDeleteApiKeyResponseBody
}

type ModelRouterDeleteApiKeyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterDeleteApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterDeleteApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterDeleteApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterDeleteApiKeyResponse) GetBody() *ModelRouterDeleteApiKeyResponseBody {
	return s.Body
}

func (s *ModelRouterDeleteApiKeyResponse) SetHeaders(v map[string]*string) *ModelRouterDeleteApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterDeleteApiKeyResponse) SetStatusCode(v int32) *ModelRouterDeleteApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterDeleteApiKeyResponse) SetBody(v *ModelRouterDeleteApiKeyResponseBody) *ModelRouterDeleteApiKeyResponse {
	s.Body = v
	return s
}

func (s *ModelRouterDeleteApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
