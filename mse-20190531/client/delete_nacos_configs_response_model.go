// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNacosConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNacosConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNacosConfigsResponseBody) *DeleteNacosConfigsResponse
	GetBody() *DeleteNacosConfigsResponseBody
}

type DeleteNacosConfigsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNacosConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNacosConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosConfigsResponse) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNacosConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNacosConfigsResponse) GetBody() *DeleteNacosConfigsResponseBody {
	return s.Body
}

func (s *DeleteNacosConfigsResponse) SetHeaders(v map[string]*string) *DeleteNacosConfigsResponse {
	s.Headers = v
	return s
}

func (s *DeleteNacosConfigsResponse) SetStatusCode(v int32) *DeleteNacosConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNacosConfigsResponse) SetBody(v *DeleteNacosConfigsResponseBody) *DeleteNacosConfigsResponse {
	s.Body = v
	return s
}

func (s *DeleteNacosConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
