// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNacosInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNacosInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNacosInstanceResponseBody) *DeleteNacosInstanceResponse
	GetBody() *DeleteNacosInstanceResponseBody
}

type DeleteNacosInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNacosInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNacosInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNacosInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNacosInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNacosInstanceResponse) GetBody() *DeleteNacosInstanceResponseBody {
	return s.Body
}

func (s *DeleteNacosInstanceResponse) SetHeaders(v map[string]*string) *DeleteNacosInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNacosInstanceResponse) SetStatusCode(v int32) *DeleteNacosInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNacosInstanceResponse) SetBody(v *DeleteNacosInstanceResponseBody) *DeleteNacosInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteNacosInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
