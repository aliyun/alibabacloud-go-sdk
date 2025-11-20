// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTodoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTodoTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTodoTaskResponseBody) *CreateTodoTaskResponse
	GetBody() *CreateTodoTaskResponseBody
}

type CreateTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTodoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTodoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTodoTaskResponse) GetBody() *CreateTodoTaskResponseBody {
	return s.Body
}

func (s *CreateTodoTaskResponse) SetHeaders(v map[string]*string) *CreateTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTodoTaskResponse) SetStatusCode(v int32) *CreateTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTodoTaskResponse) SetBody(v *CreateTodoTaskResponseBody) *CreateTodoTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTodoTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
