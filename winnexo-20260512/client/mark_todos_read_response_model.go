// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkTodosReadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MarkTodosReadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MarkTodosReadResponse
	GetStatusCode() *int32
	SetBody(v *MarkTodosReadResponseBody) *MarkTodosReadResponse
	GetBody() *MarkTodosReadResponseBody
}

type MarkTodosReadResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkTodosReadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MarkTodosReadResponse) String() string {
	return dara.Prettify(s)
}

func (s MarkTodosReadResponse) GoString() string {
	return s.String()
}

func (s *MarkTodosReadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MarkTodosReadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MarkTodosReadResponse) GetBody() *MarkTodosReadResponseBody {
	return s.Body
}

func (s *MarkTodosReadResponse) SetHeaders(v map[string]*string) *MarkTodosReadResponse {
	s.Headers = v
	return s
}

func (s *MarkTodosReadResponse) SetStatusCode(v int32) *MarkTodosReadResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkTodosReadResponse) SetBody(v *MarkTodosReadResponseBody) *MarkTodosReadResponse {
	s.Body = v
	return s
}

func (s *MarkTodosReadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
