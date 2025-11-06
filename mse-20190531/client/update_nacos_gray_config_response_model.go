// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosGrayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNacosGrayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNacosGrayConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNacosGrayConfigResponseBody) *UpdateNacosGrayConfigResponse
	GetBody() *UpdateNacosGrayConfigResponseBody
}

type UpdateNacosGrayConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNacosGrayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNacosGrayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosGrayConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacosGrayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNacosGrayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNacosGrayConfigResponse) GetBody() *UpdateNacosGrayConfigResponseBody {
	return s.Body
}

func (s *UpdateNacosGrayConfigResponse) SetHeaders(v map[string]*string) *UpdateNacosGrayConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacosGrayConfigResponse) SetStatusCode(v int32) *UpdateNacosGrayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNacosGrayConfigResponse) SetBody(v *UpdateNacosGrayConfigResponseBody) *UpdateNacosGrayConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateNacosGrayConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
