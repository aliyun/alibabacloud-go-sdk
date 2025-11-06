// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNacosServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNacosServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNacosServiceResponseBody) *UpdateNacosServiceResponse
	GetBody() *UpdateNacosServiceResponseBody
}

type UpdateNacosServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNacosServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNacosServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacosServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNacosServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNacosServiceResponse) GetBody() *UpdateNacosServiceResponseBody {
	return s.Body
}

func (s *UpdateNacosServiceResponse) SetHeaders(v map[string]*string) *UpdateNacosServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacosServiceResponse) SetStatusCode(v int32) *UpdateNacosServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNacosServiceResponse) SetBody(v *UpdateNacosServiceResponseBody) *UpdateNacosServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateNacosServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
