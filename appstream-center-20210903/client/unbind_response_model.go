// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindResponse
	GetStatusCode() *int32
	SetBody(v *UnbindResponseBody) *UnbindResponse
	GetBody() *UnbindResponseBody
}

type UnbindResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindResponse) GoString() string {
	return s.String()
}

func (s *UnbindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindResponse) GetBody() *UnbindResponseBody {
	return s.Body
}

func (s *UnbindResponse) SetHeaders(v map[string]*string) *UnbindResponse {
	s.Headers = v
	return s
}

func (s *UnbindResponse) SetStatusCode(v int32) *UnbindResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindResponse) SetBody(v *UnbindResponseBody) *UnbindResponse {
	s.Body = v
	return s
}

func (s *UnbindResponse) Validate() error {
	return dara.Validate(s)
}
