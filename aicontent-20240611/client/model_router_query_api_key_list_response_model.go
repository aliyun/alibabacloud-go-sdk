// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryApiKeyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryApiKeyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryApiKeyListResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryApiKeyListResponseBody) *ModelRouterQueryApiKeyListResponse
	GetBody() *ModelRouterQueryApiKeyListResponseBody
}

type ModelRouterQueryApiKeyListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryApiKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryApiKeyListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryApiKeyListResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryApiKeyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryApiKeyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryApiKeyListResponse) GetBody() *ModelRouterQueryApiKeyListResponseBody {
	return s.Body
}

func (s *ModelRouterQueryApiKeyListResponse) SetHeaders(v map[string]*string) *ModelRouterQueryApiKeyListResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryApiKeyListResponse) SetStatusCode(v int32) *ModelRouterQueryApiKeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponse) SetBody(v *ModelRouterQueryApiKeyListResponseBody) *ModelRouterQueryApiKeyListResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryApiKeyListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
