// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupsByApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupsByApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryModelGroupsByApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryModelGroupsByApiKeyResponseBody) *ModelRouterQueryModelGroupsByApiKeyResponse
	GetBody() *ModelRouterQueryModelGroupsByApiKeyResponseBody
}

type ModelRouterQueryModelGroupsByApiKeyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryModelGroupsByApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryModelGroupsByApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupsByApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponse) GetBody() *ModelRouterQueryModelGroupsByApiKeyResponseBody {
	return s.Body
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponse) SetHeaders(v map[string]*string) *ModelRouterQueryModelGroupsByApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponse) SetStatusCode(v int32) *ModelRouterQueryModelGroupsByApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponse) SetBody(v *ModelRouterQueryModelGroupsByApiKeyResponseBody) *ModelRouterQueryModelGroupsByApiKeyResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryModelGroupsByApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
