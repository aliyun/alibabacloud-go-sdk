// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNacosConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNacosConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNacosConfigResponseBody) *UpdateNacosConfigResponse
	GetBody() *UpdateNacosConfigResponseBody
}

type UpdateNacosConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNacosConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacosConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNacosConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNacosConfigResponse) GetBody() *UpdateNacosConfigResponseBody {
	return s.Body
}

func (s *UpdateNacosConfigResponse) SetHeaders(v map[string]*string) *UpdateNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacosConfigResponse) SetStatusCode(v int32) *UpdateNacosConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNacosConfigResponse) SetBody(v *UpdateNacosConfigResponseBody) *UpdateNacosConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateNacosConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
