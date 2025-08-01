// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNacosConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNacosConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateNacosConfigResponseBody) *CreateNacosConfigResponse
	GetBody() *CreateNacosConfigResponseBody
}

type CreateNacosConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNacosConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateNacosConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNacosConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNacosConfigResponse) GetBody() *CreateNacosConfigResponseBody {
	return s.Body
}

func (s *CreateNacosConfigResponse) SetHeaders(v map[string]*string) *CreateNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateNacosConfigResponse) SetStatusCode(v int32) *CreateNacosConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNacosConfigResponse) SetBody(v *CreateNacosConfigResponseBody) *CreateNacosConfigResponse {
	s.Body = v
	return s
}

func (s *CreateNacosConfigResponse) Validate() error {
	return dara.Validate(s)
}
