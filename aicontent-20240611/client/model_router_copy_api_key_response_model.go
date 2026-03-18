// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCopyApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCopyApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCopyApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCopyApiKeyResponseBody) *ModelRouterCopyApiKeyResponse
	GetBody() *ModelRouterCopyApiKeyResponseBody
}

type ModelRouterCopyApiKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCopyApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCopyApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCopyApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCopyApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCopyApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCopyApiKeyResponse) GetBody() *ModelRouterCopyApiKeyResponseBody {
	return s.Body
}

func (s *ModelRouterCopyApiKeyResponse) SetHeaders(v map[string]*string) *ModelRouterCopyApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCopyApiKeyResponse) SetStatusCode(v int32) *ModelRouterCopyApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCopyApiKeyResponse) SetBody(v *ModelRouterCopyApiKeyResponseBody) *ModelRouterCopyApiKeyResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCopyApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
