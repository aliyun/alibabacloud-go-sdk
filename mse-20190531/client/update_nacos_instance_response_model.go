// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNacosInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNacosInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNacosInstanceResponseBody) *UpdateNacosInstanceResponse
	GetBody() *UpdateNacosInstanceResponseBody
}

type UpdateNacosInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNacosInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNacosInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacosInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNacosInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNacosInstanceResponse) GetBody() *UpdateNacosInstanceResponseBody {
	return s.Body
}

func (s *UpdateNacosInstanceResponse) SetHeaders(v map[string]*string) *UpdateNacosInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacosInstanceResponse) SetStatusCode(v int32) *UpdateNacosInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNacosInstanceResponse) SetBody(v *UpdateNacosInstanceResponseBody) *UpdateNacosInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateNacosInstanceResponse) Validate() error {
	return dara.Validate(s)
}
