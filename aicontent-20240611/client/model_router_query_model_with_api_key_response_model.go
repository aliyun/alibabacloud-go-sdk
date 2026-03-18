// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelWithApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelWithApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelWithApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelWithApiKeyResponseBody) *ModelRouterQueryModelWithApiKeyResponse
	GetBody() *ModelRouterQueryModelWithApiKeyResponseBody
}

type ModelRouterQueryModelWithApiKeyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelWithApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelWithApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelWithApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelWithApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelWithApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelWithApiKeyResponse) GetBody() *ModelRouterQueryModelWithApiKeyResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelWithApiKeyResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelWithApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponse) SetStatusCode(v int32) *ModelRouterQueryModelWithApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponse) SetBody(v *ModelRouterQueryModelWithApiKeyResponseBody) *ModelRouterQueryModelWithApiKeyResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
