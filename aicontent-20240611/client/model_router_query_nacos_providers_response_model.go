// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryNacosProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryNacosProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryNacosProvidersResponseBody) *ModelRouterQueryNacosProvidersResponse
	GetBody() *ModelRouterQueryNacosProvidersResponseBody
}

type ModelRouterQueryNacosProvidersResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryNacosProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryNacosProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosProvidersResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryNacosProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryNacosProvidersResponse) GetBody() *ModelRouterQueryNacosProvidersResponseBody {
	return s.Body
}

func (s *ModelRouterQueryNacosProvidersResponse) SetHeaders(v map[string]*string) *ModelRouterQueryNacosProvidersResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponse) SetStatusCode(v int32) *ModelRouterQueryNacosProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponse) SetBody(v *ModelRouterQueryNacosProvidersResponseBody) *ModelRouterQueryNacosProvidersResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
