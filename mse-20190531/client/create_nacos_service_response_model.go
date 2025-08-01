// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNacosServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNacosServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateNacosServiceResponseBody) *CreateNacosServiceResponse
	GetBody() *CreateNacosServiceResponseBody
}

type CreateNacosServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNacosServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNacosServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateNacosServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNacosServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNacosServiceResponse) GetBody() *CreateNacosServiceResponseBody {
	return s.Body
}

func (s *CreateNacosServiceResponse) SetHeaders(v map[string]*string) *CreateNacosServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateNacosServiceResponse) SetStatusCode(v int32) *CreateNacosServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNacosServiceResponse) SetBody(v *CreateNacosServiceResponseBody) *CreateNacosServiceResponse {
	s.Body = v
	return s
}

func (s *CreateNacosServiceResponse) Validate() error {
	return dara.Validate(s)
}
