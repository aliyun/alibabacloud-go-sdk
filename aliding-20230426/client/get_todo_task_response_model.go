// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTodoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTodoTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetTodoTaskResponseBody) *GetTodoTaskResponse
	GetBody() *GetTodoTaskResponseBody
}

type GetTodoTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTodoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTodoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTodoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTodoTaskResponse) GetBody() *GetTodoTaskResponseBody {
	return s.Body
}

func (s *GetTodoTaskResponse) SetHeaders(v map[string]*string) *GetTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTodoTaskResponse) SetStatusCode(v int32) *GetTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodoTaskResponse) SetBody(v *GetTodoTaskResponseBody) *GetTodoTaskResponse {
	s.Body = v
	return s
}

func (s *GetTodoTaskResponse) Validate() error {
	return dara.Validate(s)
}
