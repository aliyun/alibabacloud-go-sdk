// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogstashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogstashResponse
	GetStatusCode() *int32
	SetBody(v *ListLogstashResponseBody) *ListLogstashResponse
	GetBody() *ListLogstashResponseBody
}

type ListLogstashResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogstashResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashResponse) GoString() string {
	return s.String()
}

func (s *ListLogstashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogstashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogstashResponse) GetBody() *ListLogstashResponseBody {
	return s.Body
}

func (s *ListLogstashResponse) SetHeaders(v map[string]*string) *ListLogstashResponse {
	s.Headers = v
	return s
}

func (s *ListLogstashResponse) SetStatusCode(v int32) *ListLogstashResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogstashResponse) SetBody(v *ListLogstashResponseBody) *ListLogstashResponse {
	s.Body = v
	return s
}

func (s *ListLogstashResponse) Validate() error {
	return dara.Validate(s)
}
