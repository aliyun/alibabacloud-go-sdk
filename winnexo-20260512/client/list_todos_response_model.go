// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTodosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTodosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTodosResponse
	GetStatusCode() *int32
	SetBody(v *ListTodosResponseBody) *ListTodosResponse
	GetBody() *ListTodosResponseBody
}

type ListTodosResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTodosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTodosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTodosResponse) GoString() string {
	return s.String()
}

func (s *ListTodosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTodosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTodosResponse) GetBody() *ListTodosResponseBody {
	return s.Body
}

func (s *ListTodosResponse) SetHeaders(v map[string]*string) *ListTodosResponse {
	s.Headers = v
	return s
}

func (s *ListTodosResponse) SetStatusCode(v int32) *ListTodosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTodosResponse) SetBody(v *ListTodosResponseBody) *ListTodosResponse {
	s.Body = v
	return s
}

func (s *ListTodosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
