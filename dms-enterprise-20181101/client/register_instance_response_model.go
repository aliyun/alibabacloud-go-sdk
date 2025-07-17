// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RegisterInstanceResponseBody) *RegisterInstanceResponse
	GetBody() *RegisterInstanceResponseBody
}

type RegisterInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterInstanceResponse) GoString() string {
	return s.String()
}

func (s *RegisterInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterInstanceResponse) GetBody() *RegisterInstanceResponseBody {
	return s.Body
}

func (s *RegisterInstanceResponse) SetHeaders(v map[string]*string) *RegisterInstanceResponse {
	s.Headers = v
	return s
}

func (s *RegisterInstanceResponse) SetStatusCode(v int32) *RegisterInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterInstanceResponse) SetBody(v *RegisterInstanceResponseBody) *RegisterInstanceResponse {
	s.Body = v
	return s
}

func (s *RegisterInstanceResponse) Validate() error {
	return dara.Validate(s)
}
