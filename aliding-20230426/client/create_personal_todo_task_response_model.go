// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTodoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalTodoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalTodoTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalTodoTaskResponseBody) *CreatePersonalTodoTaskResponse
	GetBody() *CreatePersonalTodoTaskResponseBody
}

type CreatePersonalTodoTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalTodoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalTodoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalTodoTaskResponse) GetBody() *CreatePersonalTodoTaskResponseBody {
	return s.Body
}

func (s *CreatePersonalTodoTaskResponse) SetHeaders(v map[string]*string) *CreatePersonalTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalTodoTaskResponse) SetStatusCode(v int32) *CreatePersonalTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalTodoTaskResponse) SetBody(v *CreatePersonalTodoTaskResponseBody) *CreatePersonalTodoTaskResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalTodoTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
