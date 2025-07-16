// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedirectTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RedirectTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RedirectTaskResponse
	GetStatusCode() *int32
	SetBody(v *RedirectTaskResponseBody) *RedirectTaskResponse
	GetBody() *RedirectTaskResponseBody
}

type RedirectTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedirectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedirectTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RedirectTaskResponse) GoString() string {
	return s.String()
}

func (s *RedirectTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RedirectTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RedirectTaskResponse) GetBody() *RedirectTaskResponseBody {
	return s.Body
}

func (s *RedirectTaskResponse) SetHeaders(v map[string]*string) *RedirectTaskResponse {
	s.Headers = v
	return s
}

func (s *RedirectTaskResponse) SetStatusCode(v int32) *RedirectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RedirectTaskResponse) SetBody(v *RedirectTaskResponseBody) *RedirectTaskResponse {
	s.Body = v
	return s
}

func (s *RedirectTaskResponse) Validate() error {
	return dara.Validate(s)
}
