// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTodoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTodoTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTodoTaskResponseBody) *UpdateTodoTaskResponse
	GetBody() *UpdateTodoTaskResponseBody
}

type UpdateTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTodoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTodoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTodoTaskResponse) GetBody() *UpdateTodoTaskResponseBody {
	return s.Body
}

func (s *UpdateTodoTaskResponse) SetHeaders(v map[string]*string) *UpdateTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTaskResponse) SetStatusCode(v int32) *UpdateTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTodoTaskResponse) SetBody(v *UpdateTodoTaskResponseBody) *UpdateTodoTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateTodoTaskResponse) Validate() error {
	return dara.Validate(s)
}
