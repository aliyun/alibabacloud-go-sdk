// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRemindsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRemindsResponse
	GetStatusCode() *int32
	SetBody(v *ListRemindsResponseBody) *ListRemindsResponse
	GetBody() *ListRemindsResponseBody
}

type ListRemindsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRemindsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRemindsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRemindsResponse) GoString() string {
	return s.String()
}

func (s *ListRemindsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRemindsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRemindsResponse) GetBody() *ListRemindsResponseBody {
	return s.Body
}

func (s *ListRemindsResponse) SetHeaders(v map[string]*string) *ListRemindsResponse {
	s.Headers = v
	return s
}

func (s *ListRemindsResponse) SetStatusCode(v int32) *ListRemindsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemindsResponse) SetBody(v *ListRemindsResponseBody) *ListRemindsResponse {
	s.Body = v
	return s
}

func (s *ListRemindsResponse) Validate() error {
	return dara.Validate(s)
}
