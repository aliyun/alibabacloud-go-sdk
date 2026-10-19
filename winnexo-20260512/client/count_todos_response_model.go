// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountTodosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CountTodosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CountTodosResponse
	GetStatusCode() *int32
	SetBody(v *CountTodosResponseBody) *CountTodosResponse
	GetBody() *CountTodosResponseBody
}

type CountTodosResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountTodosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountTodosResponse) String() string {
	return dara.Prettify(s)
}

func (s CountTodosResponse) GoString() string {
	return s.String()
}

func (s *CountTodosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CountTodosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CountTodosResponse) GetBody() *CountTodosResponseBody {
	return s.Body
}

func (s *CountTodosResponse) SetHeaders(v map[string]*string) *CountTodosResponse {
	s.Headers = v
	return s
}

func (s *CountTodosResponse) SetStatusCode(v int32) *CountTodosResponse {
	s.StatusCode = &v
	return s
}

func (s *CountTodosResponse) SetBody(v *CountTodosResponseBody) *CountTodosResponse {
	s.Body = v
	return s
}

func (s *CountTodosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
