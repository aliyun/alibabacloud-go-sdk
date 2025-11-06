// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNacosInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNacosInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateNacosInstanceResponseBody) *CreateNacosInstanceResponse
	GetBody() *CreateNacosInstanceResponseBody
}

type CreateNacosInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNacosInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNacosInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateNacosInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNacosInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNacosInstanceResponse) GetBody() *CreateNacosInstanceResponseBody {
	return s.Body
}

func (s *CreateNacosInstanceResponse) SetHeaders(v map[string]*string) *CreateNacosInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateNacosInstanceResponse) SetStatusCode(v int32) *CreateNacosInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNacosInstanceResponse) SetBody(v *CreateNacosInstanceResponseBody) *CreateNacosInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateNacosInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
