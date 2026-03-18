// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryNacosTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryNacosTagsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryNacosTagsResponseBody) *ModelRouterQueryNacosTagsResponse
	GetBody() *ModelRouterQueryNacosTagsResponseBody
}

type ModelRouterQueryNacosTagsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryNacosTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryNacosTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosTagsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryNacosTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryNacosTagsResponse) GetBody() *ModelRouterQueryNacosTagsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryNacosTagsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryNacosTagsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryNacosTagsResponse) SetStatusCode(v int32) *ModelRouterQueryNacosTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponse) SetBody(v *ModelRouterQueryNacosTagsResponseBody) *ModelRouterQueryNacosTagsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryNacosTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
