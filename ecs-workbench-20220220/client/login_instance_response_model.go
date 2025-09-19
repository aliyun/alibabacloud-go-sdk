// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LoginInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LoginInstanceResponse
	GetStatusCode() *int32
	SetBody(v *LoginInstanceResponseBody) *LoginInstanceResponse
	GetBody() *LoginInstanceResponseBody
}

type LoginInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoginInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoginInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceResponse) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LoginInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LoginInstanceResponse) GetBody() *LoginInstanceResponseBody {
	return s.Body
}

func (s *LoginInstanceResponse) SetHeaders(v map[string]*string) *LoginInstanceResponse {
	s.Headers = v
	return s
}

func (s *LoginInstanceResponse) SetStatusCode(v int32) *LoginInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginInstanceResponse) SetBody(v *LoginInstanceResponseBody) *LoginInstanceResponse {
	s.Body = v
	return s
}

func (s *LoginInstanceResponse) Validate() error {
	return dara.Validate(s)
}
