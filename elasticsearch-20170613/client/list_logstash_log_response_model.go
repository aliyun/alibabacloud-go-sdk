// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstashLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogstashLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogstashLogResponse
	GetStatusCode() *int32
	SetBody(v *ListLogstashLogResponseBody) *ListLogstashLogResponse
	GetBody() *ListLogstashLogResponseBody
}

type ListLogstashLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogstashLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogstashLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogstashLogResponse) GoString() string {
	return s.String()
}

func (s *ListLogstashLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogstashLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogstashLogResponse) GetBody() *ListLogstashLogResponseBody {
	return s.Body
}

func (s *ListLogstashLogResponse) SetHeaders(v map[string]*string) *ListLogstashLogResponse {
	s.Headers = v
	return s
}

func (s *ListLogstashLogResponse) SetStatusCode(v int32) *ListLogstashLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogstashLogResponse) SetBody(v *ListLogstashLogResponseBody) *ListLogstashLogResponse {
	s.Body = v
	return s
}

func (s *ListLogstashLogResponse) Validate() error {
	return dara.Validate(s)
}
