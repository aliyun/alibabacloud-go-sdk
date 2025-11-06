// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNacosServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNacosServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNacosServiceResponseBody) *DeleteNacosServiceResponse
	GetBody() *DeleteNacosServiceResponseBody
}

type DeleteNacosServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNacosServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNacosServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNacosServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNacosServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNacosServiceResponse) GetBody() *DeleteNacosServiceResponseBody {
	return s.Body
}

func (s *DeleteNacosServiceResponse) SetHeaders(v map[string]*string) *DeleteNacosServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNacosServiceResponse) SetStatusCode(v int32) *DeleteNacosServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNacosServiceResponse) SetBody(v *DeleteNacosServiceResponseBody) *DeleteNacosServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteNacosServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
