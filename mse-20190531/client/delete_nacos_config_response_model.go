// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNacosConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNacosConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNacosConfigResponseBody) *DeleteNacosConfigResponse
	GetBody() *DeleteNacosConfigResponseBody
}

type DeleteNacosConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNacosConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNacosConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNacosConfigResponse) GetBody() *DeleteNacosConfigResponseBody {
	return s.Body
}

func (s *DeleteNacosConfigResponse) SetHeaders(v map[string]*string) *DeleteNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteNacosConfigResponse) SetStatusCode(v int32) *DeleteNacosConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNacosConfigResponse) SetBody(v *DeleteNacosConfigResponseBody) *DeleteNacosConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteNacosConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
