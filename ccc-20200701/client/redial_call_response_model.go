// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedialCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RedialCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RedialCallResponse
	GetStatusCode() *int32
	SetBody(v *RedialCallResponseBody) *RedialCallResponse
	GetBody() *RedialCallResponseBody
}

type RedialCallResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedialCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedialCallResponse) String() string {
	return dara.Prettify(s)
}

func (s RedialCallResponse) GoString() string {
	return s.String()
}

func (s *RedialCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RedialCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RedialCallResponse) GetBody() *RedialCallResponseBody {
	return s.Body
}

func (s *RedialCallResponse) SetHeaders(v map[string]*string) *RedialCallResponse {
	s.Headers = v
	return s
}

func (s *RedialCallResponse) SetStatusCode(v int32) *RedialCallResponse {
	s.StatusCode = &v
	return s
}

func (s *RedialCallResponse) SetBody(v *RedialCallResponseBody) *RedialCallResponse {
	s.Body = v
	return s
}

func (s *RedialCallResponse) Validate() error {
	return dara.Validate(s)
}
